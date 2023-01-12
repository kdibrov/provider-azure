/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this ManagementLock.
func (mg *ManagementLock) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Scope),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ScopeRef,
		Selector:     mg.Spec.ForProvider.ScopeSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Scope")
	}
	mg.Spec.ForProvider.Scope = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ScopeRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourceGroupPolicyAssignment.
func (mg *ResourceGroupPolicyAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyDefinitionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyDefinitionIDRef,
		Selector:     mg.Spec.ForProvider.PolicyDefinitionIDSelector,
		To: reference.To{
			List:    &PolicyDefinitionList{},
			Managed: &PolicyDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyDefinitionID")
	}
	mg.Spec.ForProvider.PolicyDefinitionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyDefinitionIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ResourceGroupIDRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupIDSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupID")
	}
	mg.Spec.ForProvider.ResourceGroupID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourcePolicyAssignment.
func (mg *ResourcePolicyAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyDefinitionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyDefinitionIDRef,
		Selector:     mg.Spec.ForProvider.PolicyDefinitionIDSelector,
		To: reference.To{
			List:    &PolicyDefinitionList{},
			Managed: &PolicyDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyDefinitionID")
	}
	mg.Spec.ForProvider.PolicyDefinitionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyDefinitionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ResourcePolicyExemption.
func (mg *ResourcePolicyExemption) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyAssignmentID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyAssignmentIDRef,
		Selector:     mg.Spec.ForProvider.PolicyAssignmentIDSelector,
		To: reference.To{
			List:    &ResourcePolicyAssignmentList{},
			Managed: &ResourcePolicyAssignment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyAssignmentID")
	}
	mg.Spec.ForProvider.PolicyAssignmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyAssignmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      resource.ExtractParamPath("resource_id", false),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &ResourcePolicyAssignmentList{},
			Managed: &ResourcePolicyAssignment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RoleAssignment.
func (mg *RoleAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RoleDefinitionID),
		Extract:      resource.ExtractParamPath("role_definition_resource_id", true),
		Reference:    mg.Spec.ForProvider.RoleDefinitionIDRef,
		Selector:     mg.Spec.ForProvider.RoleDefinitionIDSelector,
		To: reference.To{
			List:    &RoleDefinitionList{},
			Managed: &RoleDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RoleDefinitionID")
	}
	mg.Spec.ForProvider.RoleDefinitionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RoleDefinitionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionPolicyAssignment.
func (mg *SubscriptionPolicyAssignment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyDefinitionID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyDefinitionIDRef,
		Selector:     mg.Spec.ForProvider.PolicyDefinitionIDSelector,
		To: reference.To{
			List:    &PolicyDefinitionList{},
			Managed: &PolicyDefinition{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyDefinitionID")
	}
	mg.Spec.ForProvider.PolicyDefinitionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyDefinitionIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SubscriptionPolicyExemption.
func (mg *SubscriptionPolicyExemption) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyAssignmentID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyAssignmentIDRef,
		Selector:     mg.Spec.ForProvider.PolicyAssignmentIDSelector,
		To: reference.To{
			List:    &SubscriptionPolicyAssignmentList{},
			Managed: &SubscriptionPolicyAssignment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyAssignmentID")
	}
	mg.Spec.ForProvider.PolicyAssignmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyAssignmentIDRef = rsp.ResolvedReference

	return nil
}
