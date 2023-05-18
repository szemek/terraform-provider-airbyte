// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceIp2whoisResourceModel) ToCreateSDKType() *shared.SourceIp2whoisCreateRequest {
	apiKey := new(string)
	if !r.Configuration.APIKey.IsUnknown() && !r.Configuration.APIKey.IsNull() {
		*apiKey = r.Configuration.APIKey.ValueString()
	} else {
		apiKey = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	sourceType := shared.SourceIp2whoisIp2whoisEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceIp2whois{
		APIKey:     apiKey,
		Domain:     domain,
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
	out := shared.SourceIp2whoisCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceIp2whoisResourceModel) ToDeleteSDKType() *shared.SourceIp2whoisCreateRequest {
	out := r.ToCreateSDKType()
	return out
}