// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceAsanaResourceModel) ToCreateSDKType() *shared.SourceAsanaCreateRequest {
	var credentials *shared.SourceAsanaAuthenticationMechanism
	var sourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth *shared.SourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth
	if r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth != nil {
		clientID := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.ClientSecret.ValueString()
		optionTitle := new(shared.SourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle)
		if !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.OptionTitle.IsNull() {
			*optionTitle = shared.SourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle(r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.OptionTitle.ValueString())
		} else {
			optionTitle = nil
		}
		refreshToken := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.RefreshToken.ValueString()
		sourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth = &shared.SourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			OptionTitle:  optionTitle,
			RefreshToken: refreshToken,
		}
	}
	if sourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth != nil {
		credentials = &shared.SourceAsanaAuthenticationMechanism{
			SourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth: sourceAsanaAuthenticationMechanismAuthenticateViaAsanaOauth,
		}
	}
	var sourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken *shared.SourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken
	if r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		optionTitle1 := new(shared.SourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle)
		if !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.OptionTitle.IsNull() {
			*optionTitle1 = shared.SourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle(r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.OptionTitle.ValueString())
		} else {
			optionTitle1 = nil
		}
		personalAccessToken := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.PersonalAccessToken.ValueString()
		sourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken = &shared.SourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken{
			OptionTitle:         optionTitle1,
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		credentials = &shared.SourceAsanaAuthenticationMechanism{
			SourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken: sourceAsanaAuthenticationMechanismAuthenticateWithPersonalAccessToken,
		}
	}
	sourceType := new(shared.SourceAsanaAsana)
	if !r.Configuration.SourceType.IsUnknown() && !r.Configuration.SourceType.IsNull() {
		*sourceType = shared.SourceAsanaAsana(r.Configuration.SourceType.ValueString())
	} else {
		sourceType = nil
	}
	configuration := shared.SourceAsana{
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
	out := shared.SourceAsanaCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAsanaResourceModel) ToUpdateSDKType() *shared.SourceAsanaPutRequest {
	var credentials *shared.SourceAsanaUpdateAuthenticationMechanism
	var sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth *shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth
	if r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth != nil {
		clientID := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.ClientSecret.ValueString()
		optionTitle := new(shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle)
		if !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.OptionTitle.IsNull() {
			*optionTitle = shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauthCredentialsTitle(r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.OptionTitle.ValueString())
		} else {
			optionTitle = nil
		}
		refreshToken := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth.RefreshToken.ValueString()
		sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth = &shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			OptionTitle:  optionTitle,
			RefreshToken: refreshToken,
		}
	}
	if sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth != nil {
		credentials = &shared.SourceAsanaUpdateAuthenticationMechanism{
			SourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth: sourceAsanaUpdateAuthenticationMechanismAuthenticateViaAsanaOauth,
		}
	}
	var sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken *shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken
	if r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		optionTitle1 := new(shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle)
		if !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.OptionTitle.IsNull() {
			*optionTitle1 = shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenCredentialsTitle(r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.OptionTitle.ValueString())
		} else {
			optionTitle1 = nil
		}
		personalAccessToken := r.Configuration.Credentials.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken.PersonalAccessToken.ValueString()
		sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken = &shared.SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken{
			OptionTitle:         optionTitle1,
			PersonalAccessToken: personalAccessToken,
		}
	}
	if sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		credentials = &shared.SourceAsanaUpdateAuthenticationMechanism{
			SourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken: sourceAsanaUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken,
		}
	}
	configuration := shared.SourceAsanaUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceAsanaPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceAsanaResourceModel) ToDeleteSDKType() *shared.SourceAsanaCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceAsanaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
