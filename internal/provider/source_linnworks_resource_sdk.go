// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceLinnworksResourceModel) ToCreateSDKType() *shared.SourceLinnworksCreateRequest {
	applicationID := r.Configuration.ApplicationID.ValueString()
	applicationSecret := r.Configuration.ApplicationSecret.ValueString()
	sourceType := shared.SourceLinnworksLinnworksEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceLinnworks{
		ApplicationID:     applicationID,
		ApplicationSecret: applicationSecret,
		SourceType:        sourceType,
		StartDate:         startDate,
		Token:             token,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceLinnworksCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceLinnworksResourceModel) ToDeleteSDKType() *shared.SourceLinnworksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}