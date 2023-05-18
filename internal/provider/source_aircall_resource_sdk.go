// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceAircallResourceModel) ToCreateSDKType() *shared.SourceAircallCreateRequest {
	apiID := r.Configuration.APIID.ValueString()
	apiToken := r.Configuration.APIToken.ValueString()
	sourceType := shared.SourceAircallAircallEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceAircall{
		APIID:      apiID,
		APIToken:   apiToken,
		SourceType: sourceType,
		StartDate:  startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAircallCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAircallResourceModel) ToDeleteSDKType() *shared.SourceAircallCreateRequest {
	out := r.ToCreateSDKType()
	return out
}