// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceSmailyResourceModel) ToCreateSDKType() *shared.SourceSmailyCreateRequest {
	apiPassword := r.Configuration.APIPassword.ValueString()
	apiSubdomain := r.Configuration.APISubdomain.ValueString()
	apiUsername := r.Configuration.APIUsername.ValueString()
	sourceType := shared.SourceSmailySmailyEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSmaily{
		APIPassword:  apiPassword,
		APISubdomain: apiSubdomain,
		APIUsername:  apiUsername,
		SourceType:   sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmailyCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmailyResourceModel) ToDeleteSDKType() *shared.SourceSmailyCreateRequest {
	out := r.ToCreateSDKType()
	return out
}