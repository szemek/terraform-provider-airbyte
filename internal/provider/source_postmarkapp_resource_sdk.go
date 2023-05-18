// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourcePostmarkappResourceModel) ToCreateSDKType() *shared.SourcePostmarkappCreateRequest {
	xPostmarkAccountToken := r.Configuration.XPostmarkAccountToken.ValueString()
	xPostmarkServerToken := r.Configuration.XPostmarkServerToken.ValueString()
	sourceType := shared.SourcePostmarkappPostmarkappEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourcePostmarkapp{
		XPostmarkAccountToken: xPostmarkAccountToken,
		XPostmarkServerToken:  xPostmarkServerToken,
		SourceType:            sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePostmarkappCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePostmarkappResourceModel) ToDeleteSDKType() *shared.SourcePostmarkappCreateRequest {
	out := r.ToCreateSDKType()
	return out
}