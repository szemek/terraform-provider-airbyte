// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceMailchimpResourceModel) ToCreateSDKType() *shared.SourceMailchimpCreateRequest {
	campaignID := new(string)
	if !r.Configuration.CampaignID.IsUnknown() && !r.Configuration.CampaignID.IsNull() {
		*campaignID = r.Configuration.CampaignID.ValueString()
	} else {
		campaignID = nil
	}
	var credentials *shared.SourceMailchimpAuthentication
	var sourceMailchimpAuthenticationOAuth20 *shared.SourceMailchimpAuthenticationOAuth20
	if r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.AccessToken.ValueString()
		authType := shared.SourceMailchimpAuthenticationOAuth20AuthTypeEnum(r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.AuthType.ValueString())
		clientID := new(string)
		if !r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.ClientID.IsNull() {
			*clientID = r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.ClientID.ValueString()
		} else {
			clientID = nil
		}
		clientSecret := new(string)
		if !r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.ClientSecret.IsNull() {
			*clientSecret = r.Configuration.Credentials.SourceMailchimpAuthenticationOAuth20.ClientSecret.ValueString()
		} else {
			clientSecret = nil
		}
		sourceMailchimpAuthenticationOAuth20 = &shared.SourceMailchimpAuthenticationOAuth20{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
		}
	}
	if sourceMailchimpAuthenticationOAuth20 != nil {
		credentials = &shared.SourceMailchimpAuthentication{
			SourceMailchimpAuthenticationOAuth20: sourceMailchimpAuthenticationOAuth20,
		}
	}
	var sourceMailchimpAuthenticationAPIKey *shared.SourceMailchimpAuthenticationAPIKey
	if r.Configuration.Credentials.SourceMailchimpAuthenticationAPIKey != nil {
		apikey := r.Configuration.Credentials.SourceMailchimpAuthenticationAPIKey.Apikey.ValueString()
		authType1 := shared.SourceMailchimpAuthenticationAPIKeyAuthTypeEnum(r.Configuration.Credentials.SourceMailchimpAuthenticationAPIKey.AuthType.ValueString())
		sourceMailchimpAuthenticationAPIKey = &shared.SourceMailchimpAuthenticationAPIKey{
			Apikey:   apikey,
			AuthType: authType1,
		}
	}
	if sourceMailchimpAuthenticationAPIKey != nil {
		credentials = &shared.SourceMailchimpAuthentication{
			SourceMailchimpAuthenticationAPIKey: sourceMailchimpAuthenticationAPIKey,
		}
	}
	sourceType := shared.SourceMailchimpMailchimpEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceMailchimp{
		CampaignID:  campaignID,
		Credentials: credentials,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMailchimpCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailchimpResourceModel) ToDeleteSDKType() *shared.SourceMailchimpCreateRequest {
	out := r.ToCreateSDKType()
	return out
}