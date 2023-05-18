// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceQuickbooksResourceModel) ToCreateSDKType() *shared.SourceQuickbooksCreateRequest {
	var credentials shared.SourceQuickbooksAuthorizationMethod
	var sourceQuickbooksAuthorizationMethodOAuth20 *shared.SourceQuickbooksAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceQuickbooksAuthorizationMethodOAuth20AuthTypeEnum)
		if !r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceQuickbooksAuthorizationMethodOAuth20AuthTypeEnum(r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.ClientSecret.ValueString()
		realmID := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.RealmID.ValueString()
		refreshToken := r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceQuickbooksAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceQuickbooksAuthorizationMethodOAuth20 = &shared.SourceQuickbooksAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RealmID:         realmID,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceQuickbooksAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceQuickbooksAuthorizationMethod{
			SourceQuickbooksAuthorizationMethodOAuth20: sourceQuickbooksAuthorizationMethodOAuth20,
		}
	}
	sandbox := r.Configuration.Sandbox.ValueBool()
	sourceType := shared.SourceQuickbooksQuickbooksEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceQuickbooks{
		Credentials: credentials,
		Sandbox:     sandbox,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceQuickbooksCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQuickbooksResourceModel) ToDeleteSDKType() *shared.SourceQuickbooksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}