// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceDelightedResourceModel) ToCreateSDKType() *shared.SourceDelightedCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	since, _ := time.Parse(time.RFC3339Nano, r.Configuration.Since.ValueString())
	sourceType := shared.SourceDelightedDelightedEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceDelighted{
		APIKey:     apiKey,
		Since:      since,
		SourceType: sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDelightedCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDelightedResourceModel) ToDeleteSDKType() *shared.SourceDelightedCreateRequest {
	out := r.ToCreateSDKType()
	return out
}