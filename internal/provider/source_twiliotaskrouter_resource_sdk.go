// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceTwilioTaskrouterResourceModel) ToCreateSDKType() *shared.SourceTwilioTaskrouterCreateRequest {
	accountSid := r.Configuration.AccountSid.ValueString()
	authToken := r.Configuration.AuthToken.ValueString()
	sourceType := shared.SourceTwilioTaskrouterTwilioTaskrouterEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceTwilioTaskrouter{
		AccountSid: accountSid,
		AuthToken:  authToken,
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
	out := shared.SourceTwilioTaskrouterCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTwilioTaskrouterResourceModel) ToDeleteSDKType() *shared.SourceTwilioTaskrouterCreateRequest {
	out := r.ToCreateSDKType()
	return out
}