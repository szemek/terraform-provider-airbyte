// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceBigcommerceResourceModel) ToCreateSDKType() *shared.SourceBigcommerceCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	sourceType := shared.SourceBigcommerceBigcommerceEnum(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	storeHash := r.Configuration.StoreHash.ValueString()
	configuration := shared.SourceBigcommerce{
		AccessToken: accessToken,
		SourceType:  sourceType,
		StartDate:   startDate,
		StoreHash:   storeHash,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBigcommerceCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBigcommerceResourceModel) ToDeleteSDKType() *shared.SourceBigcommerceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}