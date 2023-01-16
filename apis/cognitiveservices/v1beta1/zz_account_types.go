/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AccountObservation struct {

	// The endpoint used to connect to the Cognitive Service Account.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the Cognitive Service Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`
}

type AccountParameters struct {

	// If kind is TextAnalytics this specifies the ID of the Search service.
	// +kubebuilder:validation:Optional
	CustomQuestionAnsweringSearchServiceID *string `json:"customQuestionAnsweringSearchServiceId,omitempty" tf:"custom_question_answering_search_service_id,omitempty"`

	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CustomSubdomainName *string `json:"customSubdomainName,omitempty" tf:"custom_subdomain_name,omitempty"`

	// List of FQDNs allowed for the Cognitive Account.
	// +kubebuilder:validation:Optional
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the type of Cognitive Service Account that should be created. Possible values are Academic, AnomalyDetector, Bing.Autosuggest, Bing.Autosuggest.v7, Bing.CustomSearch, Bing.Search, Bing.Search.v7, Bing.Speech, Bing.SpellCheck, Bing.SpellCheck.v7, CognitiveServices, ComputerVision, ContentModerator, CustomSpeech, CustomVision.Prediction, CustomVision.Training, Emotion, Face,FormRecognizer, ImmersiveReader, LUIS, LUIS.Authoring, MetricsAdvisor, Personalizer, QnAMaker, Recommendations, SpeakerRecognition, Speech, SpeechServices, SpeechTranslation, TextAnalytics, TextTranslation and WebLM. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Kind *string `json:"kind" tf:"kind,omitempty"`

	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The Azure AD Client ID (Application ID). This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorAADClientID *string `json:"metricsAdvisorAadClientId,omitempty" tf:"metrics_advisor_aad_client_id,omitempty"`

	// The Azure AD Tenant ID. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorAADTenantID *string `json:"metricsAdvisorAadTenantId,omitempty" tf:"metrics_advisor_aad_tenant_id,omitempty"`

	// The super user of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorSuperUserName *string `json:"metricsAdvisorSuperUserName,omitempty" tf:"metrics_advisor_super_user_name,omitempty"`

	// The website name of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorWebsiteName *string `json:"metricsAdvisorWebsiteName,omitempty" tf:"metrics_advisor_website_name,omitempty"`

	// A network_acls block as defined below.
	// +kubebuilder:validation:Optional
	NetworkAcls []NetworkAclsParameters `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`

	// Whether outbound network access is restricted for the Cognitive Account. Defaults to false.
	// +kubebuilder:validation:Optional
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	// +kubebuilder:validation:Optional
	QnaRuntimeEndpoint *string `json:"qnaRuntimeEndpoint,omitempty" tf:"qna_runtime_endpoint,omitempty"`

	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the SKU Name for this Cognitive Service Account. Possible values are F0, F1, S, S0, S1, S2, S3, S4, S5, S6, P0, P1, and P2.
	// +kubebuilder:validation:Required
	SkuName *string `json:"skuName" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IdentityObservation struct {

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Cognitive Account.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Cognitive Account. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkAclsObservation struct {
}

type NetworkAclsParameters struct {

	// The Default Action to use when no rules match from ip_rules / virtual_network_rules. Possible values are Allow and Deny.
	// +kubebuilder:validation:Required
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
	// +kubebuilder:validation:Optional
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// A virtual_network_rules block as defined below.
	// +kubebuilder:validation:Optional
	VirtualNetworkRules []VirtualNetworkRulesParameters `json:"virtualNetworkRules,omitempty" tf:"virtual_network_rules,omitempty"`
}

type StorageObservation struct {
}

type StorageParameters struct {

	// The client ID of the managed identity associated with the storage resource.
	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Full resource id of a Microsoft.Storage resource.
	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type VirtualNetworkRulesObservation struct {
}

type VirtualNetworkRulesParameters struct {

	// Whether ignore missing vnet service endpoint or not. Default to false.
	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The ID of the subnet which should be able to access this Cognitive Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API. Manages a Cognitive Services Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccountSpec   `json:"spec"`
	Status            AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
