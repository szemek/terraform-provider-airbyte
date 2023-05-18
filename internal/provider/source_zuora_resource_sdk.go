// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceZuoraResourceModel) ToCreateSDKType() *shared.SourceZuoraCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dataQuery := shared.SourceZuoraDataQueryTypeEnum(r.Configuration.DataQuery.ValueString())
	sourceType := shared.SourceZuoraZuoraEnum(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	tenantEndpoint := shared.SourceZuoraTenantEndpointLocationEnum(r.Configuration.TenantEndpoint.ValueString())
	windowInDays := new(string)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueString()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceZuora{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		DataQuery:      dataQuery,
		SourceType:     sourceType,
		StartDate:      startDate,
		TenantEndpoint: tenantEndpoint,
		WindowInDays:   windowInDays,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZuoraCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZuoraResourceModel) ToDeleteSDKType() *shared.SourceZuoraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}