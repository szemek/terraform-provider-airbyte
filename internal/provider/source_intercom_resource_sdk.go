// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceIntercomResourceModel) ToCreateSDKType() *shared.SourceIntercomCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	sourceType := shared.SourceIntercomIntercomEnum(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceIntercom{
		AccessToken: accessToken,
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
	out := shared.SourceIntercomCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceIntercomResourceModel) ToDeleteSDKType() *shared.SourceIntercomCreateRequest {
	out := r.ToCreateSDKType()
	return out
}