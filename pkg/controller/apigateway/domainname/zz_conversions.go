/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package domainname

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigateway"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigateway/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetDomainNameInput returns input for read
// operation.
func GenerateGetDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.GetDomainNameInput {
	res := &svcsdk.GetDomainNameInput{}

	if cr.Spec.ForProvider.DomainName != nil {
		res.SetDomainName(*cr.Spec.ForProvider.DomainName)
	}

	return res
}

// GenerateDomainName returns the current state in the form of *svcapitypes.DomainName.
func GenerateDomainName(resp *svcsdk.DomainName) *svcapitypes.DomainName {
	cr := &svcapitypes.DomainName{}

	if resp.CertificateArn != nil {
		cr.Spec.ForProvider.CertificateARN = resp.CertificateArn
	} else {
		cr.Spec.ForProvider.CertificateARN = nil
	}
	if resp.CertificateName != nil {
		cr.Spec.ForProvider.CertificateName = resp.CertificateName
	} else {
		cr.Spec.ForProvider.CertificateName = nil
	}
	if resp.CertificateUploadDate != nil {
		cr.Status.AtProvider.CertificateUploadDate = &metav1.Time{*resp.CertificateUploadDate}
	} else {
		cr.Status.AtProvider.CertificateUploadDate = nil
	}
	if resp.DistributionDomainName != nil {
		cr.Status.AtProvider.DistributionDomainName = resp.DistributionDomainName
	} else {
		cr.Status.AtProvider.DistributionDomainName = nil
	}
	if resp.DistributionHostedZoneId != nil {
		cr.Status.AtProvider.DistributionHostedZoneID = resp.DistributionHostedZoneId
	} else {
		cr.Status.AtProvider.DistributionHostedZoneID = nil
	}
	if resp.DomainName != nil {
		cr.Spec.ForProvider.DomainName = resp.DomainName
	} else {
		cr.Spec.ForProvider.DomainName = nil
	}
	if resp.DomainNameStatus != nil {
		cr.Status.AtProvider.DomainNameStatus = resp.DomainNameStatus
	} else {
		cr.Status.AtProvider.DomainNameStatus = nil
	}
	if resp.DomainNameStatusMessage != nil {
		cr.Status.AtProvider.DomainNameStatusMessage = resp.DomainNameStatusMessage
	} else {
		cr.Status.AtProvider.DomainNameStatusMessage = nil
	}
	if resp.EndpointConfiguration != nil {
		f8 := &svcapitypes.EndpointConfiguration{}
		if resp.EndpointConfiguration.Types != nil {
			f8f0 := []*string{}
			for _, f8f0iter := range resp.EndpointConfiguration.Types {
				var f8f0elem string
				f8f0elem = *f8f0iter
				f8f0 = append(f8f0, &f8f0elem)
			}
			f8.Types = f8f0
		}
		if resp.EndpointConfiguration.VpcEndpointIds != nil {
			f8f1 := []*string{}
			for _, f8f1iter := range resp.EndpointConfiguration.VpcEndpointIds {
				var f8f1elem string
				f8f1elem = *f8f1iter
				f8f1 = append(f8f1, &f8f1elem)
			}
			f8.VPCEndpointIDs = f8f1
		}
		cr.Spec.ForProvider.EndpointConfiguration = f8
	} else {
		cr.Spec.ForProvider.EndpointConfiguration = nil
	}
	if resp.MutualTlsAuthentication != nil {
		f9 := &svcapitypes.MutualTLSAuthenticationInput{}
		if resp.MutualTlsAuthentication.TruststoreUri != nil {
			f9.TruststoreURI = resp.MutualTlsAuthentication.TruststoreUri
		}
		if resp.MutualTlsAuthentication.TruststoreVersion != nil {
			f9.TruststoreVersion = resp.MutualTlsAuthentication.TruststoreVersion
		}
		cr.Spec.ForProvider.MutualTLSAuthentication = f9
	} else {
		cr.Spec.ForProvider.MutualTLSAuthentication = nil
	}
	if resp.OwnershipVerificationCertificateArn != nil {
		cr.Spec.ForProvider.OwnershipVerificationCertificateARN = resp.OwnershipVerificationCertificateArn
	} else {
		cr.Spec.ForProvider.OwnershipVerificationCertificateARN = nil
	}
	if resp.RegionalCertificateArn != nil {
		cr.Spec.ForProvider.RegionalCertificateARN = resp.RegionalCertificateArn
	} else {
		cr.Spec.ForProvider.RegionalCertificateARN = nil
	}
	if resp.RegionalCertificateName != nil {
		cr.Spec.ForProvider.RegionalCertificateName = resp.RegionalCertificateName
	} else {
		cr.Spec.ForProvider.RegionalCertificateName = nil
	}
	if resp.RegionalDomainName != nil {
		cr.Status.AtProvider.RegionalDomainName = resp.RegionalDomainName
	} else {
		cr.Status.AtProvider.RegionalDomainName = nil
	}
	if resp.RegionalHostedZoneId != nil {
		cr.Status.AtProvider.RegionalHostedZoneID = resp.RegionalHostedZoneId
	} else {
		cr.Status.AtProvider.RegionalHostedZoneID = nil
	}
	if resp.SecurityPolicy != nil {
		cr.Spec.ForProvider.SecurityPolicy = resp.SecurityPolicy
	} else {
		cr.Spec.ForProvider.SecurityPolicy = nil
	}
	if resp.Tags != nil {
		f16 := map[string]*string{}
		for f16key, f16valiter := range resp.Tags {
			var f16val string
			f16val = *f16valiter
			f16[f16key] = &f16val
		}
		cr.Spec.ForProvider.Tags = f16
	} else {
		cr.Spec.ForProvider.Tags = nil
	}

	return cr
}

// GenerateCreateDomainNameInput returns a create input.
func GenerateCreateDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.CreateDomainNameInput {
	res := &svcsdk.CreateDomainNameInput{}

	if cr.Spec.ForProvider.CertificateARN != nil {
		res.SetCertificateArn(*cr.Spec.ForProvider.CertificateARN)
	}
	if cr.Spec.ForProvider.CertificateBody != nil {
		res.SetCertificateBody(*cr.Spec.ForProvider.CertificateBody)
	}
	if cr.Spec.ForProvider.CertificateChain != nil {
		res.SetCertificateChain(*cr.Spec.ForProvider.CertificateChain)
	}
	if cr.Spec.ForProvider.CertificateName != nil {
		res.SetCertificateName(*cr.Spec.ForProvider.CertificateName)
	}
	if cr.Spec.ForProvider.CertificatePrivateKey != nil {
		res.SetCertificatePrivateKey(*cr.Spec.ForProvider.CertificatePrivateKey)
	}
	if cr.Spec.ForProvider.DomainName != nil {
		res.SetDomainName(*cr.Spec.ForProvider.DomainName)
	}
	if cr.Spec.ForProvider.EndpointConfiguration != nil {
		f6 := &svcsdk.EndpointConfiguration{}
		if cr.Spec.ForProvider.EndpointConfiguration.Types != nil {
			f6f0 := []*string{}
			for _, f6f0iter := range cr.Spec.ForProvider.EndpointConfiguration.Types {
				var f6f0elem string
				f6f0elem = *f6f0iter
				f6f0 = append(f6f0, &f6f0elem)
			}
			f6.SetTypes(f6f0)
		}
		if cr.Spec.ForProvider.EndpointConfiguration.VPCEndpointIDs != nil {
			f6f1 := []*string{}
			for _, f6f1iter := range cr.Spec.ForProvider.EndpointConfiguration.VPCEndpointIDs {
				var f6f1elem string
				f6f1elem = *f6f1iter
				f6f1 = append(f6f1, &f6f1elem)
			}
			f6.SetVpcEndpointIds(f6f1)
		}
		res.SetEndpointConfiguration(f6)
	}
	if cr.Spec.ForProvider.MutualTLSAuthentication != nil {
		f7 := &svcsdk.MutualTlsAuthenticationInput{}
		if cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreURI != nil {
			f7.SetTruststoreUri(*cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreURI)
		}
		if cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreVersion != nil {
			f7.SetTruststoreVersion(*cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreVersion)
		}
		res.SetMutualTlsAuthentication(f7)
	}
	if cr.Spec.ForProvider.OwnershipVerificationCertificateARN != nil {
		res.SetOwnershipVerificationCertificateArn(*cr.Spec.ForProvider.OwnershipVerificationCertificateARN)
	}
	if cr.Spec.ForProvider.RegionalCertificateARN != nil {
		res.SetRegionalCertificateArn(*cr.Spec.ForProvider.RegionalCertificateARN)
	}
	if cr.Spec.ForProvider.RegionalCertificateName != nil {
		res.SetRegionalCertificateName(*cr.Spec.ForProvider.RegionalCertificateName)
	}
	if cr.Spec.ForProvider.SecurityPolicy != nil {
		res.SetSecurityPolicy(*cr.Spec.ForProvider.SecurityPolicy)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f12 := map[string]*string{}
		for f12key, f12valiter := range cr.Spec.ForProvider.Tags {
			var f12val string
			f12val = *f12valiter
			f12[f12key] = &f12val
		}
		res.SetTags(f12)
	}

	return res
}

// GenerateUpdateDomainNameInput returns an update input.
func GenerateUpdateDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.UpdateDomainNameInput {
	res := &svcsdk.UpdateDomainNameInput{}

	if cr.Spec.ForProvider.DomainName != nil {
		res.SetDomainName(*cr.Spec.ForProvider.DomainName)
	}

	return res
}

// GenerateDeleteDomainNameInput returns a deletion input.
func GenerateDeleteDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.DeleteDomainNameInput {
	res := &svcsdk.DeleteDomainNameInput{}

	if cr.Spec.ForProvider.DomainName != nil {
		res.SetDomainName(*cr.Spec.ForProvider.DomainName)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}
