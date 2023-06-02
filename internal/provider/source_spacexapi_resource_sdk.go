// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	sourceType := new(shared.SourceSpacexAPISpacexAPI)
	if !r.Configuration.SourceType.IsUnknown() && !r.Configuration.SourceType.IsNull() {
		*sourceType = shared.SourceSpacexAPISpacexAPI(r.Configuration.SourceType.ValueString())
	} else {
		sourceType = nil
	}
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

func (r *SourceSpacexAPIResourceModel) ToUpdateSDKType() *shared.SourceSpacexAPIPutRequest {
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
	configuration := shared.SourceSpacexAPIUpdate{
		ID:      id,
		Options: options,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSpacexAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSpacexAPIResourceModel) ToDeleteSDKType() *shared.SourceSpacexAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSpacexAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
