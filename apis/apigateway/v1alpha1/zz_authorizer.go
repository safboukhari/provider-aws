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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// AuthorizerParameters defines the desired state of Authorizer
type AuthorizerParameters struct {
	// Region is which region the Authorizer will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Optional customer-defined field, used in OpenAPI imports and exports without
	// functional impact.
	AuthType *string `json:"authType,omitempty"`
	// Specifies the required credentials as an IAM role for API Gateway to invoke
	// the authorizer. To specify an IAM role for API Gateway to assume, use the
	// role's Amazon Resource Name (ARN). To use resource-based permissions on the
	// Lambda function, specify null.
	AuthorizerCredentials *string `json:"authorizerCredentials,omitempty"`
	// The TTL in seconds of cached authorizer results. If it equals 0, authorization
	// caching is disabled. If it is greater than 0, API Gateway will cache authorizer
	// responses. If this field is not set, the default value is 300. The maximum
	// value is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *int64 `json:"authorizerResultTtlInSeconds,omitempty"`
	// Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or
	// REQUEST authorizers, this must be a well-formed Lambda function URI, for
	// example, arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form arn:aws:apigateway:{region}:lambda:path/{service_api},
	// where {region} is the same as the region hosting the Lambda function, path
	// indicates that the remaining substring in the URI should be treated as the
	// path to the resource, including the initial /. For Lambda functions, this
	// is usually of the form /2015-03-31/functions/[FunctionARN]/invocations.
	AuthorizerURI *string `json:"authorizerURI,omitempty"`
	// The identity source for which authorization is requested.
	//    * For a TOKEN or COGNITO_USER_POOLS authorizer, this is required and specifies
	//    the request header mapping expression for the custom header holding the
	//    authorization token submitted by the client. For example, if the token
	//    header name is Auth, the header mapping expression is method.request.header.Auth.
	//
	//    * For the REQUEST authorizer, this is required when authorization caching
	//    is enabled. The value is a comma-separated string of one or more mapping
	//    expressions of the specified request parameters. For example, if an Auth
	//    header, a Name query string parameter are defined as identity sources,
	//    this value is method.request.header.Auth, method.request.querystring.Name.
	//    These parameters will be used to derive the authorization caching key
	//    and to perform runtime validation of the REQUEST authorizer by verifying
	//    all of the identity-related request parameters are present, not null and
	//    non-empty. Only when this is true does the authorizer invoke the authorizer
	//    Lambda function, otherwise, it returns a 401 Unauthorized response without
	//    calling the Lambda function. The valid value is a string of comma-separated
	//    mapping expressions of the specified request parameters. When the authorization
	//    caching is not enabled, this property is optional.
	IdentitySource *string `json:"identitySource,omitempty"`
	// A validation expression for the incoming identity token. For TOKEN authorizers,
	// this value is a regular expression. For COGNITO_USER_POOLS authorizers, API
	// Gateway will match the aud field of the incoming token from the client against
	// the specified regular expression. It will invoke the authorizer's Lambda
	// function when there is a match. Otherwise, it will return a 401 Unauthorized
	// response without calling the Lambda function. The validation expression does
	// not apply to the REQUEST authorizer.
	IdentityValidationExpression *string `json:"identityValidationExpression,omitempty"`
	// [Required] The name of the authorizer.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// [Required] The authorizer type. Valid values are TOKEN for a Lambda function
	// using a single authorization token submitted in a custom header, REQUEST
	// for a Lambda function using incoming request parameters, and COGNITO_USER_POOLS
	// for using an Amazon Cognito user pool.
	// +kubebuilder:validation:Required
	Type                       *string `json:"type_"`
	CustomAuthorizerParameters `json:",inline"`
}

// AuthorizerSpec defines the desired state of Authorizer
type AuthorizerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AuthorizerParameters `json:"forProvider"`
}

// AuthorizerObservation defines the observed state of Authorizer
type AuthorizerObservation struct {
	// The identifier for the authorizer resource.
	ID *string `json:"id,omitempty"`
	// A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS authorizer.
	// Each element is of this format: arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}.
	// For a TOKEN or REQUEST authorizer, this is not defined.
	ProviderARNs []*string `json:"providerARNs,omitempty"`
}

// AuthorizerStatus defines the observed state of Authorizer.
type AuthorizerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AuthorizerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Authorizer is the Schema for the Authorizers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Authorizer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthorizerSpec   `json:"spec"`
	Status            AuthorizerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AuthorizerList contains a list of Authorizers
type AuthorizerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Authorizer `json:"items"`
}

// Repository type metadata.
var (
	AuthorizerKind             = "Authorizer"
	AuthorizerGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AuthorizerKind}.String()
	AuthorizerKindAPIVersion   = AuthorizerKind + "." + GroupVersion.String()
	AuthorizerGroupVersionKind = GroupVersion.WithKind(AuthorizerKind)
)

func init() {
	SchemeBuilder.Register(&Authorizer{}, &AuthorizerList{})
}
