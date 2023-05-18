// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceZendeskChatResourceModel) ToCreateSDKType() *shared.SourceZendeskChatCreateRequest {
	var credentials *shared.SourceZendeskChatAuthorizationMethod
	var sourceZendeskChatAuthorizationMethodOAuth20 *shared.SourceZendeskChatAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20 != nil {
		accessToken := new(string)
		if !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.AccessToken.IsNull() {
			*accessToken = r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.AccessToken.ValueString()
		} else {
			accessToken = nil
		}
		clientID := new(string)
		if !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		credentials1 := shared.SourceZendeskChatAuthorizationMethodOAuth20CredentialsEnum(r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.Credentials.ValueString())
		refreshToken := new(string)
		if !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.RefreshToken.IsUnknown() && !r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.RefreshToken.IsNull() {
			*refreshToken = r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodOAuth20.RefreshToken.ValueString()
		} else {
			refreshToken = nil
		}
		sourceZendeskChatAuthorizationMethodOAuth20 = &shared.SourceZendeskChatAuthorizationMethodOAuth20{
			AccessToken:  accessToken,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Credentials:  credentials1,
			RefreshToken: refreshToken,
		}
	}
	if sourceZendeskChatAuthorizationMethodOAuth20 != nil {
		credentials = &shared.SourceZendeskChatAuthorizationMethod{
			SourceZendeskChatAuthorizationMethodOAuth20: sourceZendeskChatAuthorizationMethodOAuth20,
		}
	}
	var sourceZendeskChatAuthorizationMethodAccessToken *shared.SourceZendeskChatAuthorizationMethodAccessToken
	if r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodAccessToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodAccessToken.AccessToken.ValueString()
		credentials2 := shared.SourceZendeskChatAuthorizationMethodAccessTokenCredentialsEnum(r.Configuration.Credentials.SourceZendeskChatAuthorizationMethodAccessToken.Credentials.ValueString())
		sourceZendeskChatAuthorizationMethodAccessToken = &shared.SourceZendeskChatAuthorizationMethodAccessToken{
			AccessToken: accessToken1,
			Credentials: credentials2,
		}
	}
	if sourceZendeskChatAuthorizationMethodAccessToken != nil {
		credentials = &shared.SourceZendeskChatAuthorizationMethod{
			SourceZendeskChatAuthorizationMethodAccessToken: sourceZendeskChatAuthorizationMethodAccessToken,
		}
	}
	sourceType := shared.SourceZendeskChatZendeskChatEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := new(string)
	if !r.Configuration.Subdomain.IsUnknown() && !r.Configuration.Subdomain.IsNull() {
		*subdomain = r.Configuration.Subdomain.ValueString()
	} else {
		subdomain = nil
	}
	configuration := shared.SourceZendeskChat{
		Credentials: credentials,
		SourceType:  sourceType,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskChatCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskChatResourceModel) ToDeleteSDKType() *shared.SourceZendeskChatCreateRequest {
	out := r.ToCreateSDKType()
	return out
}