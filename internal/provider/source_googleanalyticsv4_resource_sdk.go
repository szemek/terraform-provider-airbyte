// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceGoogleAnalyticsV4ResourceModel) ToCreateSDKType() *shared.SourceGoogleAnalyticsV4CreateRequest {
	var credentials *shared.SourceGoogleAnalyticsV4Credentials
	var sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth *shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth
	if r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth != nil {
		accessToken := new(string)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken.IsNull() {
			*accessToken = r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AccessToken.ValueString()
		} else {
			accessToken = nil
		}
		authType := new(shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthAuthTypeEnum)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType.IsNull() {
			*authType = shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauthAuthTypeEnum(r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth.RefreshToken.ValueString()
		sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth = &shared.SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth != nil {
		credentials = &shared.SourceGoogleAnalyticsV4Credentials{
			SourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth: sourceGoogleAnalyticsV4CredentialsAuthenticateViaGoogleOauth,
		}
	}
	var sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication *shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication
	if r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication != nil {
		authType1 := new(shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthenticationAuthTypeEnum)
		if !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType.IsNull() {
			*authType1 = shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthenticationAuthTypeEnum(r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		credentialsJSON := r.Configuration.Credentials.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication.CredentialsJSON.ValueString()
		sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication = &shared.SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication{
			AuthType:        authType1,
			CredentialsJSON: credentialsJSON,
		}
	}
	if sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication != nil {
		credentials = &shared.SourceGoogleAnalyticsV4Credentials{
			SourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication: sourceGoogleAnalyticsV4CredentialsServiceAccountKeyAuthentication,
		}
	}
	customReports := new(string)
	if !r.Configuration.CustomReports.IsUnknown() && !r.Configuration.CustomReports.IsNull() {
		*customReports = r.Configuration.CustomReports.ValueString()
	} else {
		customReports = nil
	}
	sourceType := shared.SourceGoogleAnalyticsV4GoogleAnalyticsV4Enum(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	viewID := r.Configuration.ViewID.ValueString()
	windowInDays := new(int64)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueInt64()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceGoogleAnalyticsV4{
		Credentials:   credentials,
		CustomReports: customReports,
		SourceType:    sourceType,
		StartDate:     startDate,
		ViewID:        viewID,
		WindowInDays:  windowInDays,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleAnalyticsV4CreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleAnalyticsV4ResourceModel) ToDeleteSDKType() *shared.SourceGoogleAnalyticsV4CreateRequest {
	out := r.ToCreateSDKType()
	return out
}