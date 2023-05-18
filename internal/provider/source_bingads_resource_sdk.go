// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	customTypes "airbyte/internal/sdk/pkg/types"
)

func (r *SourceBingAdsResourceModel) ToCreateSDKType() *shared.SourceBingAdsCreateRequest {
	authMethod := new(shared.SourceBingAdsAuthMethodEnum)
	if !r.Configuration.AuthMethod.IsUnknown() && !r.Configuration.AuthMethod.IsNull() {
		*authMethod = shared.SourceBingAdsAuthMethodEnum(r.Configuration.AuthMethod.ValueString())
	} else {
		authMethod = nil
	}
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := new(string)
	if !r.Configuration.ClientSecret.IsUnknown() && !r.Configuration.ClientSecret.IsNull() {
		*clientSecret = r.Configuration.ClientSecret.ValueString()
	} else {
		clientSecret = nil
	}
	developerToken := r.Configuration.DeveloperToken.ValueString()
	lookbackWindow := new(int64)
	if !r.Configuration.LookbackWindow.IsUnknown() && !r.Configuration.LookbackWindow.IsNull() {
		*lookbackWindow = r.Configuration.LookbackWindow.ValueInt64()
	} else {
		lookbackWindow = nil
	}
	refreshToken := r.Configuration.RefreshToken.ValueString()
	reportsStartDate, _ := customTypes.NewDate(r.Configuration.ReportsStartDate.ValueString())
	sourceType := shared.SourceBingAdsBingAdsEnum(r.Configuration.SourceType.ValueString())
	tenantID := new(string)
	if !r.Configuration.TenantID.IsUnknown() && !r.Configuration.TenantID.IsNull() {
		*tenantID = r.Configuration.TenantID.ValueString()
	} else {
		tenantID = nil
	}
	configuration := shared.SourceBingAds{
		AuthMethod:       authMethod,
		ClientID:         clientID,
		ClientSecret:     clientSecret,
		DeveloperToken:   developerToken,
		LookbackWindow:   lookbackWindow,
		RefreshToken:     refreshToken,
		ReportsStartDate: reportsStartDate,
		SourceType:       sourceType,
		TenantID:         tenantID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBingAdsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBingAdsResourceModel) ToDeleteSDKType() *shared.SourceBingAdsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}