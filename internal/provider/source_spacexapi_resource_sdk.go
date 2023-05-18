// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceSpacexAPIResourceModel) ToCreateSDKType() *shared.SourceSpacexAPICreateRequest {
	id := new(string)
	if !r.Configuration.ID.IsUnknown() && !r.Configuration.ID.IsNull() {
		*id = r.Configuration.ID.ValueString()
	} else {
		id = nil
	}
	options := new(string)
	if !r.Configuration.Options.IsUnknown() && !r.Configuration.Options.IsNull() {
		*options = r.Configuration.Options.ValueString()
	} else {
		options = nil
	}
	sourceType := shared.SourceSpacexAPISpacexAPIEnum(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSpacexAPI{
		ID:         id,
		Options:    options,
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
	out := shared.SourceSpacexAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSpacexAPIResourceModel) ToDeleteSDKType() *shared.SourceSpacexAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}