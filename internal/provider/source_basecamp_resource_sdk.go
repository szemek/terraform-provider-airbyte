// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceBasecampResourceModel) ToSharedSourceBasecampCreateRequest() *shared.SourceBasecampCreateRequest {
	accountID, _ := r.Configuration.AccountID.ValueBigFloat().Float64()
	clientID := r.Configuration.ClientID.ValueString()
	clientRefreshToken2 := r.Configuration.ClientRefreshToken2.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceBasecamp{
		AccountID:           accountID,
		ClientID:            clientID,
		ClientRefreshToken2: clientRefreshToken2,
		ClientSecret:        clientSecret,
		StartDate:           startDate,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBasecampCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBasecampResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceBasecampResourceModel) ToSharedSourceBasecampPutRequest() *shared.SourceBasecampPutRequest {
	accountID, _ := r.Configuration.AccountID.ValueBigFloat().Float64()
	clientID := r.Configuration.ClientID.ValueString()
	clientRefreshToken2 := r.Configuration.ClientRefreshToken2.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceBasecampUpdate{
		AccountID:           accountID,
		ClientID:            clientID,
		ClientRefreshToken2: clientRefreshToken2,
		ClientSecret:        clientSecret,
		StartDate:           startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceBasecampPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
