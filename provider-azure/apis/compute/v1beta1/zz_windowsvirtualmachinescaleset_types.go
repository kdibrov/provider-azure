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

type IPConfigurationPublicIPAddressIPTagObservation struct {
}

type IPConfigurationPublicIPAddressIPTagParameters struct {

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkInterfaceIPConfigurationPublicIPAddressObservation struct {
}

type NetworkInterfaceIPConfigurationPublicIPAddressParameters struct {

	// +kubebuilder:validation:Optional
	DomainNameLabel *string `json:"domainNameLabel,omitempty" tf:"domain_name_label,omitempty"`

	// +kubebuilder:validation:Optional
	IPTag []IPConfigurationPublicIPAddressIPTagParameters `json:"ipTag,omitempty" tf:"ip_tag,omitempty"`

	// +kubebuilder:validation:Optional
	IdleTimeoutInMinutes *float64 `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPPrefixID *string `json:"publicIpPrefixId,omitempty" tf:"public_ip_prefix_id,omitempty"`
}

type WindowsVirtualMachineScaleSetAdditionalCapabilitiesObservation struct {
}

type WindowsVirtualMachineScaleSetAdditionalCapabilitiesParameters struct {

	// +kubebuilder:validation:Optional
	UltraSsdEnabled *bool `json:"ultraSsdEnabled,omitempty" tf:"ultra_ssd_enabled,omitempty"`
}

type WindowsVirtualMachineScaleSetAdditionalUnattendContentObservation struct {
}

type WindowsVirtualMachineScaleSetAdditionalUnattendContentParameters struct {

	// +kubebuilder:validation:Required
	ContentSecretRef v1.SecretKeySelector `json:"contentSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	Setting *string `json:"setting" tf:"setting,omitempty"`
}

type WindowsVirtualMachineScaleSetAutomaticInstanceRepairObservation struct {
}

type WindowsVirtualMachineScaleSetAutomaticInstanceRepairParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	GracePeriod *string `json:"gracePeriod,omitempty" tf:"grace_period,omitempty"`
}

type WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyObservation struct {
}

type WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyParameters struct {

	// +kubebuilder:validation:Required
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback" tf:"disable_automatic_rollback,omitempty"`

	// +kubebuilder:validation:Required
	EnableAutomaticOsUpgrade *bool `json:"enableAutomaticOsUpgrade" tf:"enable_automatic_os_upgrade,omitempty"`
}

type WindowsVirtualMachineScaleSetBootDiagnosticsObservation struct {
}

type WindowsVirtualMachineScaleSetBootDiagnosticsParameters struct {

	// +kubebuilder:validation:Optional
	StorageAccountURI *string `json:"storageAccountUri,omitempty" tf:"storage_account_uri,omitempty"`
}

type WindowsVirtualMachineScaleSetDataDiskObservation struct {
}

type WindowsVirtualMachineScaleSetDataDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// +kubebuilder:validation:Optional
	CreateOption *string `json:"createOption,omitempty" tf:"create_option,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Required
	Lun *float64 `json:"lun" tf:"lun,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	UltraSsdDiskIopsReadWrite *float64 `json:"ultraSsdDiskIopsReadWrite,omitempty" tf:"ultra_ssd_disk_iops_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	UltraSsdDiskMbpsReadWrite *float64 `json:"ultraSsdDiskMbpsReadWrite,omitempty" tf:"ultra_ssd_disk_mbps_read_write,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type WindowsVirtualMachineScaleSetExtensionObservation struct {
}

type WindowsVirtualMachineScaleSetExtensionParameters struct {

	// +kubebuilder:validation:Optional
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticUpgradeEnabled *bool `json:"automaticUpgradeEnabled,omitempty" tf:"automatic_upgrade_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	ForceUpdateTag *string `json:"forceUpdateTag,omitempty" tf:"force_update_tag,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProtectedSettingsSecretRef *v1.SecretKeySelector `json:"protectedSettingsSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ProvisionAfterExtensions []*string `json:"provisionAfterExtensions,omitempty" tf:"provision_after_extensions,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Optional
	Settings *string `json:"settings,omitempty" tf:"settings,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	TypeHandlerVersion *string `json:"typeHandlerVersion" tf:"type_handler_version,omitempty"`
}

type WindowsVirtualMachineScaleSetIdentityObservation struct {
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type WindowsVirtualMachineScaleSetIdentityParameters struct {

	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIPConfigurationObservation struct {
}

type WindowsVirtualMachineScaleSetNetworkInterfaceIPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	ApplicationGatewayBackendAddressPoolIds []*string `json:"applicationGatewayBackendAddressPoolIds,omitempty" tf:"application_gateway_backend_address_pool_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ApplicationSecurityGroupIds []*string `json:"applicationSecurityGroupIds,omitempty" tf:"application_security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerBackendAddressPoolIds []*string `json:"loadBalancerBackendAddressPoolIds,omitempty" tf:"load_balancer_backend_address_pool_ids,omitempty"`

	// +kubebuilder:validation:Optional
	LoadBalancerInboundNATRulesIds []*string `json:"loadBalancerInboundNatRulesIds,omitempty" tf:"load_balancer_inbound_nat_rules_ids,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// +kubebuilder:validation:Optional
	PublicIPAddress []NetworkInterfaceIPConfigurationPublicIPAddressParameters `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type WindowsVirtualMachineScaleSetNetworkInterfaceObservation struct {
}

type WindowsVirtualMachineScaleSetNetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`

	// +kubebuilder:validation:Optional
	EnableIPForwarding *bool `json:"enableIpForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`

	// +kubebuilder:validation:Required
	IPConfiguration []WindowsVirtualMachineScaleSetNetworkInterfaceIPConfigurationParameters `json:"ipConfiguration" tf:"ip_configuration,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`
}

type WindowsVirtualMachineScaleSetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Identity []WindowsVirtualMachineScaleSetIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	UniqueID *string `json:"uniqueId,omitempty" tf:"unique_id,omitempty"`
}

type WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsObservation struct {
}

type WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsParameters struct {

	// +kubebuilder:validation:Required
	Option *string `json:"option" tf:"option,omitempty"`
}

type WindowsVirtualMachineScaleSetOsDiskObservation struct {
}

type WindowsVirtualMachineScaleSetOsDiskParameters struct {

	// +kubebuilder:validation:Required
	Caching *string `json:"caching" tf:"caching,omitempty"`

	// +kubebuilder:validation:Optional
	DiffDiskSettings []WindowsVirtualMachineScaleSetOsDiskDiffDiskSettingsParameters `json:"diffDiskSettings,omitempty" tf:"diff_disk_settings,omitempty"`

	// +kubebuilder:validation:Optional
	DiskEncryptionSetID *string `json:"diskEncryptionSetId,omitempty" tf:"disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	DiskSizeGb *float64 `json:"diskSizeGb,omitempty" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	SecureVMDiskEncryptionSetID *string `json:"secureVmDiskEncryptionSetId,omitempty" tf:"secure_vm_disk_encryption_set_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityEncryptionType *string `json:"securityEncryptionType,omitempty" tf:"security_encryption_type,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountType *string `json:"storageAccountType" tf:"storage_account_type,omitempty"`

	// +kubebuilder:validation:Optional
	WriteAcceleratorEnabled *bool `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
}

type WindowsVirtualMachineScaleSetParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalCapabilities []WindowsVirtualMachineScaleSetAdditionalCapabilitiesParameters `json:"additionalCapabilities,omitempty" tf:"additional_capabilities,omitempty"`

	// +kubebuilder:validation:Optional
	AdditionalUnattendContent []WindowsVirtualMachineScaleSetAdditionalUnattendContentParameters `json:"additionalUnattendContent,omitempty" tf:"additional_unattend_content,omitempty"`

	// +kubebuilder:validation:Required
	AdminPasswordSecretRef v1.SecretKeySelector `json:"adminPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	AdminUsername *string `json:"adminUsername" tf:"admin_username,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticInstanceRepair []WindowsVirtualMachineScaleSetAutomaticInstanceRepairParameters `json:"automaticInstanceRepair,omitempty" tf:"automatic_instance_repair,omitempty"`

	// +kubebuilder:validation:Optional
	AutomaticOsUpgradePolicy []WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyParameters `json:"automaticOsUpgradePolicy,omitempty" tf:"automatic_os_upgrade_policy,omitempty"`

	// +kubebuilder:validation:Optional
	BootDiagnostics []WindowsVirtualMachineScaleSetBootDiagnosticsParameters `json:"bootDiagnostics,omitempty" tf:"boot_diagnostics,omitempty"`

	// +kubebuilder:validation:Optional
	ComputerNamePrefix *string `json:"computerNamePrefix,omitempty" tf:"computer_name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	CustomDataSecretRef *v1.SecretKeySelector `json:"customDataSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DataDisk []WindowsVirtualMachineScaleSetDataDiskParameters `json:"dataDisk,omitempty" tf:"data_disk,omitempty"`

	// +kubebuilder:validation:Optional
	DoNotRunExtensionsOnOverprovisionedMachines *bool `json:"doNotRunExtensionsOnOverprovisionedMachines,omitempty" tf:"do_not_run_extensions_on_overprovisioned_machines,omitempty"`

	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// +kubebuilder:validation:Optional
	EnableAutomaticUpdates *bool `json:"enableAutomaticUpdates,omitempty" tf:"enable_automatic_updates,omitempty"`

	// +kubebuilder:validation:Optional
	EncryptionAtHostEnabled *bool `json:"encryptionAtHostEnabled,omitempty" tf:"encryption_at_host_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	EvictionPolicy *string `json:"evictionPolicy,omitempty" tf:"eviction_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Extension []WindowsVirtualMachineScaleSetExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// +kubebuilder:validation:Optional
	ExtensionsTimeBudget *string `json:"extensionsTimeBudget,omitempty" tf:"extensions_time_budget,omitempty"`

	// +kubebuilder:validation:Optional
	HealthProbeID *string `json:"healthProbeId,omitempty" tf:"health_probe_id,omitempty"`

	// +kubebuilder:validation:Optional
	Identity []WindowsVirtualMachineScaleSetIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// +kubebuilder:validation:Required
	Instances *float64 `json:"instances" tf:"instances,omitempty"`

	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	MaxBidPrice *float64 `json:"maxBidPrice,omitempty" tf:"max_bid_price,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterface []WindowsVirtualMachineScaleSetNetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Required
	OsDisk []WindowsVirtualMachineScaleSetOsDiskParameters `json:"osDisk" tf:"os_disk,omitempty"`

	// +kubebuilder:validation:Optional
	Overprovision *bool `json:"overprovision,omitempty" tf:"overprovision,omitempty"`

	// +kubebuilder:validation:Optional
	Plan []WindowsVirtualMachineScaleSetPlanParameters `json:"plan,omitempty" tf:"plan,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformFaultDomainCount *float64 `json:"platformFaultDomainCount,omitempty" tf:"platform_fault_domain_count,omitempty"`

	// +kubebuilder:validation:Optional
	Priority *string `json:"priority,omitempty" tf:"priority,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisionVMAgent *bool `json:"provisionVmAgent,omitempty" tf:"provision_vm_agent,omitempty"`

	// +kubebuilder:validation:Optional
	ProximityPlacementGroupID *string `json:"proximityPlacementGroupId,omitempty" tf:"proximity_placement_group_id,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RollingUpgradePolicy []WindowsVirtualMachineScaleSetRollingUpgradePolicyParameters `json:"rollingUpgradePolicy,omitempty" tf:"rolling_upgrade_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ScaleInPolicy *string `json:"scaleInPolicy,omitempty" tf:"scale_in_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Secret []WindowsVirtualMachineScaleSetSecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// +kubebuilder:validation:Optional
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty" tf:"secure_boot_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	SinglePlacementGroup *bool `json:"singlePlacementGroup,omitempty" tf:"single_placement_group,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImageID *string `json:"sourceImageId,omitempty" tf:"source_image_id,omitempty"`

	// +kubebuilder:validation:Optional
	SourceImageReference []WindowsVirtualMachineScaleSetSourceImageReferenceParameters `json:"sourceImageReference,omitempty" tf:"source_image_reference,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	TerminateNotification []WindowsVirtualMachineScaleSetTerminateNotificationParameters `json:"terminateNotification,omitempty" tf:"terminate_notification,omitempty"`

	// +kubebuilder:validation:Optional
	TerminationNotification []WindowsVirtualMachineScaleSetTerminationNotificationParameters `json:"terminationNotification,omitempty" tf:"termination_notification,omitempty"`

	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`

	// +kubebuilder:validation:Optional
	UpgradeMode *string `json:"upgradeMode,omitempty" tf:"upgrade_mode,omitempty"`

	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`

	// +kubebuilder:validation:Optional
	VtpmEnabled *bool `json:"vtpmEnabled,omitempty" tf:"vtpm_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	WinrmListener []WindowsVirtualMachineScaleSetWinrmListenerParameters `json:"winrmListener,omitempty" tf:"winrm_listener,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneBalance *bool `json:"zoneBalance,omitempty" tf:"zone_balance,omitempty"`

	// +kubebuilder:validation:Optional
	Zones []*string `json:"zones,omitempty" tf:"zones,omitempty"`
}

type WindowsVirtualMachineScaleSetPlanObservation struct {
}

type WindowsVirtualMachineScaleSetPlanParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Product *string `json:"product" tf:"product,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`
}

type WindowsVirtualMachineScaleSetRollingUpgradePolicyObservation struct {
}

type WindowsVirtualMachineScaleSetRollingUpgradePolicyParameters struct {

	// +kubebuilder:validation:Required
	MaxBatchInstancePercent *float64 `json:"maxBatchInstancePercent" tf:"max_batch_instance_percent,omitempty"`

	// +kubebuilder:validation:Required
	MaxUnhealthyInstancePercent *float64 `json:"maxUnhealthyInstancePercent" tf:"max_unhealthy_instance_percent,omitempty"`

	// +kubebuilder:validation:Required
	MaxUnhealthyUpgradedInstancePercent *float64 `json:"maxUnhealthyUpgradedInstancePercent" tf:"max_unhealthy_upgraded_instance_percent,omitempty"`

	// +kubebuilder:validation:Required
	PauseTimeBetweenBatches *string `json:"pauseTimeBetweenBatches" tf:"pause_time_between_batches,omitempty"`
}

type WindowsVirtualMachineScaleSetSecretCertificateObservation struct {
}

type WindowsVirtualMachineScaleSetSecretCertificateParameters struct {

	// +kubebuilder:validation:Required
	Store *string `json:"store" tf:"store,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type WindowsVirtualMachineScaleSetSecretObservation struct {
}

type WindowsVirtualMachineScaleSetSecretParameters struct {

	// +kubebuilder:validation:Required
	Certificate []WindowsVirtualMachineScaleSetSecretCertificateParameters `json:"certificate" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`
}

type WindowsVirtualMachineScaleSetSourceImageReferenceObservation struct {
}

type WindowsVirtualMachineScaleSetSourceImageReferenceParameters struct {

	// +kubebuilder:validation:Required
	Offer *string `json:"offer" tf:"offer,omitempty"`

	// +kubebuilder:validation:Required
	Publisher *string `json:"publisher" tf:"publisher,omitempty"`

	// +kubebuilder:validation:Required
	Sku *string `json:"sku" tf:"sku,omitempty"`

	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type WindowsVirtualMachineScaleSetTerminateNotificationObservation struct {
}

type WindowsVirtualMachineScaleSetTerminateNotificationParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type WindowsVirtualMachineScaleSetTerminationNotificationObservation struct {
}

type WindowsVirtualMachineScaleSetTerminationNotificationParameters struct {

	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Timeout *string `json:"timeout,omitempty" tf:"timeout,omitempty"`
}

type WindowsVirtualMachineScaleSetWinrmListenerObservation struct {
}

type WindowsVirtualMachineScaleSetWinrmListenerParameters struct {

	// +kubebuilder:validation:Optional
	CertificateURL *string `json:"certificateUrl,omitempty" tf:"certificate_url,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// WindowsVirtualMachineScaleSetSpec defines the desired state of WindowsVirtualMachineScaleSet
type WindowsVirtualMachineScaleSetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     WindowsVirtualMachineScaleSetParameters `json:"forProvider"`
}

// WindowsVirtualMachineScaleSetStatus defines the observed state of WindowsVirtualMachineScaleSet.
type WindowsVirtualMachineScaleSetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        WindowsVirtualMachineScaleSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsVirtualMachineScaleSet is the Schema for the WindowsVirtualMachineScaleSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type WindowsVirtualMachineScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WindowsVirtualMachineScaleSetSpec   `json:"spec"`
	Status            WindowsVirtualMachineScaleSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WindowsVirtualMachineScaleSetList contains a list of WindowsVirtualMachineScaleSets
type WindowsVirtualMachineScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WindowsVirtualMachineScaleSet `json:"items"`
}

// Repository type metadata.
var (
	WindowsVirtualMachineScaleSet_Kind             = "WindowsVirtualMachineScaleSet"
	WindowsVirtualMachineScaleSet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WindowsVirtualMachineScaleSet_Kind}.String()
	WindowsVirtualMachineScaleSet_KindAPIVersion   = WindowsVirtualMachineScaleSet_Kind + "." + CRDGroupVersion.String()
	WindowsVirtualMachineScaleSet_GroupVersionKind = CRDGroupVersion.WithKind(WindowsVirtualMachineScaleSet_Kind)
)

func init() {
	SchemeBuilder.Register(&WindowsVirtualMachineScaleSet{}, &WindowsVirtualMachineScaleSetList{})
}
