// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSmartsheetsResourceModel) ToCreateSDKType() *shared.SourceSmartsheetsCreateRequest {
	var credentials shared.SourceSmartsheetsAuthorizationMethod
	var sourceSmartsheetsAuthorizationMethodOAuth20 *shared.SourceSmartsheetsAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceSmartsheetsAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceSmartsheetsAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceSmartsheetsAuthorizationMethodOAuth20 = &shared.SourceSmartsheetsAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceSmartsheetsAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceSmartsheetsAuthorizationMethod{
			SourceSmartsheetsAuthorizationMethodOAuth20: sourceSmartsheetsAuthorizationMethodOAuth20,
		}
	}
	var sourceSmartsheetsAuthorizationMethodAPIAccessToken *shared.SourceSmartsheetsAuthorizationMethodAPIAccessToken
	if r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AccessToken.ValueString()
		authType1 := new(shared.SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthType)
		if !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AuthType.IsNull() {
			*authType1 = shared.SourceSmartsheetsAuthorizationMethodAPIAccessTokenAuthType(r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		sourceSmartsheetsAuthorizationMethodAPIAccessToken = &shared.SourceSmartsheetsAuthorizationMethodAPIAccessToken{
			AccessToken: accessToken1,
			AuthType:    authType1,
		}
	}
	if sourceSmartsheetsAuthorizationMethodAPIAccessToken != nil {
		credentials = shared.SourceSmartsheetsAuthorizationMethod{
			SourceSmartsheetsAuthorizationMethodAPIAccessToken: sourceSmartsheetsAuthorizationMethodAPIAccessToken,
		}
	}
	sourceType := shared.SourceSmartsheetsSmartsheets(r.Configuration.SourceType.ValueString())
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceSmartsheets{
		Credentials:   credentials,
		SourceType:    sourceType,
		SpreadsheetID: spreadsheetID,
		StartDatetime: startDatetime,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmartsheetsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartsheetsResourceModel) ToUpdateSDKType() *shared.SourceSmartsheetsPutRequest {
	var credentials shared.SourceSmartsheetsUpdateAuthorizationMethod
	var sourceSmartsheetsUpdateAuthorizationMethodOAuth20 *shared.SourceSmartsheetsUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceSmartsheetsUpdateAuthorizationMethodOAuth20 = &shared.SourceSmartsheetsUpdateAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceSmartsheetsUpdateAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceSmartsheetsUpdateAuthorizationMethod{
			SourceSmartsheetsUpdateAuthorizationMethodOAuth20: sourceSmartsheetsUpdateAuthorizationMethodOAuth20,
		}
	}
	var sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken *shared.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken
	if r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AccessToken.ValueString()
		authType1 := new(shared.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType)
		if !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AuthType.IsNull() {
			*authType1 = shared.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType(r.Configuration.Credentials.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken = &shared.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken{
			AccessToken: accessToken1,
			AuthType:    authType1,
		}
	}
	if sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken != nil {
		credentials = shared.SourceSmartsheetsUpdateAuthorizationMethod{
			SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken: sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken,
		}
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	startDatetime := new(time.Time)
	if !r.Configuration.StartDatetime.IsUnknown() && !r.Configuration.StartDatetime.IsNull() {
		*startDatetime, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDatetime.ValueString())
	} else {
		startDatetime = nil
	}
	configuration := shared.SourceSmartsheetsUpdate{
		Credentials:   credentials,
		SpreadsheetID: spreadsheetID,
		StartDatetime: startDatetime,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmartsheetsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartsheetsResourceModel) ToDeleteSDKType() *shared.SourceSmartsheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSmartsheetsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
