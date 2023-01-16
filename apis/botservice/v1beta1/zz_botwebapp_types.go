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

type BotWebAppObservation struct {

	// The ID of the Bot Web App.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BotWebAppParameters struct {

	// The Application Insights API Key to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsAPIKeySecretRef *v1.SecretKeySelector `json:"developerAppInsightsApiKeySecretRef,omitempty" tf:"-"`

	// The Application Insights Application ID to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Web App Bot will be displayed as. This defaults to name if not specified.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Web App Bot endpoint.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A list of LUIS App IDs to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

	// The LUIS key to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	LuisKeySecretRef *v1.SecretKeySelector `json:"luisKeySecretRef,omitempty" tf:"-"`

	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	MicrosoftAppID *string `json:"microsoftAppId" tf:"microsoft_app_id,omitempty"`

	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Web App Bot. Valid values include F0 or S1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BotWebAppSpec defines the desired state of BotWebApp
type BotWebAppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotWebAppParameters `json:"forProvider"`
}

// BotWebAppStatus defines the observed state of BotWebApp.
type BotWebAppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotWebAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotWebApp is the Schema for the BotWebApps API. Manages a Web App Bot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotWebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotWebAppSpec   `json:"spec"`
	Status            BotWebAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotWebAppList contains a list of BotWebApps
type BotWebAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotWebApp `json:"items"`
}

// Repository type metadata.
var (
	BotWebApp_Kind             = "BotWebApp"
	BotWebApp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotWebApp_Kind}.String()
	BotWebApp_KindAPIVersion   = BotWebApp_Kind + "." + CRDGroupVersion.String()
	BotWebApp_GroupVersionKind = CRDGroupVersion.WithKind(BotWebApp_Kind)
)

func init() {
	SchemeBuilder.Register(&BotWebApp{}, &BotWebAppList{})
}
