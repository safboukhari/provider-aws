/*
Copyright 2019 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcerecordset

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"

	"github.com/crossplane/provider-aws/apis/route53/v1alpha1"
	awsclients "github.com/crossplane/provider-aws/pkg/clients"
)

const (
	errResourceRecordSetNotFound = "ResourceRecordSet.NotFound"
)

// Client defines ResourceRecordSet operations
type Client interface {
	ChangeResourceRecordSetsRequest(input *route53.ChangeResourceRecordSetsInput) route53.ChangeResourceRecordSetsRequest
	ListResourceRecordSetsRequest(input *route53.ListResourceRecordSetsInput) route53.ListResourceRecordSetsRequest
}

// NotFoundError will be raised when there is no ResourceRecordSet
type NotFoundError struct{}

// Error satisfies the Error interface for NotFoundError.
// We need to implement our own error for this because AWS SDK doesn't have
// a predefined error for Resource Record not found.
func (r *NotFoundError) Error() string {
	return errResourceRecordSetNotFound
}

// IsNotFound returns true if the error code indicates that the requested Resource Record was not found
func IsNotFound(err error) bool {
	_, ok := err.(*NotFoundError)
	return ok
}

// NewClient creates new AWS client with provided AWS Configuration/Credentials
func NewClient(cfg aws.Config) Client {
	return route53.New(cfg)
}

// GetResourceRecordSet returns recordSet if present or err
func GetResourceRecordSet(ctx context.Context, name string, params v1alpha1.ResourceRecordSetParameters, c Client) (*route53.ResourceRecordSet, error) {
	res, err := c.ListResourceRecordSetsRequest(&route53.ListResourceRecordSetsInput{
		HostedZoneId:    params.ZoneID,
		StartRecordName: &name,
		StartRecordType: route53.RRType(params.Type),
	}).Send(ctx)
	if err != nil {
		return nil, err
	}
	appendDot := func(s string) string {
		if !strings.HasSuffix(s, ".") {
			return fmt.Sprintf("%s.", s)
		}
		return s
	}
	for _, rr := range res.ResourceRecordSets {
		if appendDot(aws.StringValue(rr.Name)) == appendDot(name) &&
			string(rr.Type) == params.Type &&
			aws.StringValue(rr.SetIdentifier) == aws.StringValue(params.SetIdentifier) {
			return &rr, nil
		}
	}
	return nil, &NotFoundError{}
}

// GenerateChangeResourceRecordSetsInput prepares input for a ChangeResourceRecordSetsInput
func GenerateChangeResourceRecordSetsInput(name string, p v1alpha1.ResourceRecordSetParameters, action route53.ChangeAction) *route53.ChangeResourceRecordSetsInput {
	r := &route53.ResourceRecordSet{
		Name:                    aws.String(name),
		Type:                    route53.RRType(p.Type),
		TTL:                     p.TTL,
		SetIdentifier:           p.SetIdentifier,
		Weight:                  p.Weight,
		Failover:                route53.ResourceRecordSetFailover(p.Failover),
		HealthCheckId:           p.HealthCheckID,
		MultiValueAnswer:        p.MultiValueAnswer,
		Region:                  route53.ResourceRecordSetRegion(p.Region),
		TrafficPolicyInstanceId: p.TrafficPolicyInstanceID,
	}
	for _, v := range p.ResourceRecords {
		r.ResourceRecords = append(r.ResourceRecords, route53.ResourceRecord{
			Value: aws.String(v.Value),
		})
	}
	if p.AliasTarget != nil {
		r.AliasTarget = &route53.AliasTarget{
			DNSName:              aws.String(p.AliasTarget.DNSName),
			EvaluateTargetHealth: aws.Bool(p.AliasTarget.EvaluateTargetHealth),
			HostedZoneId:         aws.String(p.AliasTarget.HostedZoneID),
		}
	}
	if p.GeoLocation != nil {
		r.GeoLocation = &route53.GeoLocation{
			ContinentCode:   p.GeoLocation.ContinentCode,
			CountryCode:     p.GeoLocation.CountryCode,
			SubdivisionCode: p.GeoLocation.SubdivisionCode,
		}
	}

	return &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: p.ZoneID,
		ChangeBatch: &route53.ChangeBatch{
			Changes: []route53.Change{
				{
					Action:            action,
					ResourceRecordSet: r,
				},
			},
		},
	}
}

// IsUpToDate checks if object is up to date
func IsUpToDate(p v1alpha1.ResourceRecordSetParameters, rrset route53.ResourceRecordSet) (bool, error) {
	patch, err := CreatePatch(&rrset, &p)
	if err != nil {
		return false, err
	}
	return cmp.Equal(&v1alpha1.ResourceRecordSetParameters{}, patch,
		cmpopts.IgnoreTypes(&runtimev1alpha1.Reference{}, &runtimev1alpha1.Selector{}),
		cmpopts.IgnoreFields(v1alpha1.ResourceRecordSetParameters{}, "Region")), nil
}

// LateInitialize fills the empty fields in *v1alpha1.ResourceRecordSetParameters with
// the values seen in route53.ResourceRecordSet.
func LateInitialize(in *v1alpha1.ResourceRecordSetParameters, rrSet *route53.ResourceRecordSet) {
	if rrSet == nil || in == nil {
		return
	}
	rrType := string(rrSet.Type)
	in.Type = awsclients.LateInitializeString(in.Type, &rrType)
	in.TTL = awsclients.LateInitializeInt64Ptr(in.TTL, rrSet.TTL)
	if len(in.ResourceRecords) == 0 && len(rrSet.ResourceRecords) != 0 {
		in.ResourceRecords = make([]v1alpha1.ResourceRecord, len(rrSet.ResourceRecords))
		for i, val := range rrSet.ResourceRecords {
			in.ResourceRecords[i] = v1alpha1.ResourceRecord{
				Value: awsclients.StringValue(val.Value),
			}
		}
	}
}

// CreatePatch creates a *v1beta1.ResourceRecordSetParameters that has only the changed
// values between the target *v1beta1.ResourceRecordSetParameters and the current
// *route53.ResourceRecordSet
func CreatePatch(in *route53.ResourceRecordSet, target *v1alpha1.ResourceRecordSetParameters) (*v1alpha1.ResourceRecordSetParameters, error) {
	currentParams := &v1alpha1.ResourceRecordSetParameters{}
	LateInitialize(currentParams, in)

	// ZoneID doesn't exist in *route53.ResourceRecordSet object, so, we have to
	// skip its comparison.
	currentParams.ZoneID = target.ZoneID

	jsonPatch, err := awsclients.CreateJSONPatch(currentParams, target)
	if err != nil {
		return nil, err
	}
	patch := &v1alpha1.ResourceRecordSetParameters{}
	if err := json.Unmarshal(jsonPatch, patch); err != nil {
		return nil, err
	}
	return patch, nil
}
