// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *DestinationSftpJSONResourceModel) ToCreateSDKType() *shared.DestinationSftpJSONCreateRequest {
	destinationType := shared.DestinationSftpJSONSftpJSONEnum(r.Configuration.DestinationType.ValueString())
	destinationPath := r.Configuration.DestinationPath.ValueString()
	host := r.Configuration.Host.ValueString()
	password := r.Configuration.Password.ValueString()
	port := new(int64)
	if !r.Configuration.Port.IsUnknown() && !r.Configuration.Port.IsNull() {
		*port = r.Configuration.Port.ValueInt64()
	} else {
		port = nil
	}
	username := r.Configuration.Username.ValueString()
	configuration := shared.DestinationSftpJSON{
		DestinationType: destinationType,
		DestinationPath: destinationPath,
		Host:            host,
		Password:        password,
		Port:            port,
		Username:        username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationSftpJSONCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationSftpJSONResourceModel) ToDeleteSDKType() *shared.DestinationSftpJSONCreateRequest {
	out := r.ToCreateSDKType()
	return out
}