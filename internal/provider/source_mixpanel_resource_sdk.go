// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceMixpanelResourceModel) ToCreateSDKType() *shared.SourceMixpanelCreateRequest {
	attributionWindow := new(int64)
	if !r.Configuration.AttributionWindow.IsUnknown() && !r.Configuration.AttributionWindow.IsNull() {
		*attributionWindow = r.Configuration.AttributionWindow.ValueInt64()
	} else {
		attributionWindow = nil
	}
	var credentials *shared.SourceMixpanelAuthenticationWildcard
	var sourceMixpanelAuthenticationWildcardServiceAccount *shared.SourceMixpanelAuthenticationWildcardServiceAccount
	if r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount != nil {
		optionTitle := new(shared.SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle)
		if !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.OptionTitle.IsNull() {
			*optionTitle = shared.SourceMixpanelAuthenticationWildcardServiceAccountOptionTitle(r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.OptionTitle.ValueString())
		} else {
			optionTitle = nil
		}
		secret := r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.Secret.ValueString()
		username := r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.Username.ValueString()
		sourceMixpanelAuthenticationWildcardServiceAccount = &shared.SourceMixpanelAuthenticationWildcardServiceAccount{
			OptionTitle: optionTitle,
			Secret:      secret,
			Username:    username,
		}
	}
	if sourceMixpanelAuthenticationWildcardServiceAccount != nil {
		credentials = &shared.SourceMixpanelAuthenticationWildcard{
			SourceMixpanelAuthenticationWildcardServiceAccount: sourceMixpanelAuthenticationWildcardServiceAccount,
		}
	}
	var sourceMixpanelAuthenticationWildcardProjectSecret *shared.SourceMixpanelAuthenticationWildcardProjectSecret
	if r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret != nil {
		apiSecret := r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.APISecret.ValueString()
		optionTitle1 := new(shared.SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle)
		if !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.OptionTitle.IsNull() {
			*optionTitle1 = shared.SourceMixpanelAuthenticationWildcardProjectSecretOptionTitle(r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.OptionTitle.ValueString())
		} else {
			optionTitle1 = nil
		}
		sourceMixpanelAuthenticationWildcardProjectSecret = &shared.SourceMixpanelAuthenticationWildcardProjectSecret{
			APISecret:   apiSecret,
			OptionTitle: optionTitle1,
		}
	}
	if sourceMixpanelAuthenticationWildcardProjectSecret != nil {
		credentials = &shared.SourceMixpanelAuthenticationWildcard{
			SourceMixpanelAuthenticationWildcardProjectSecret: sourceMixpanelAuthenticationWildcardProjectSecret,
		}
	}
	dateWindowSize := new(int64)
	if !r.Configuration.DateWindowSize.IsUnknown() && !r.Configuration.DateWindowSize.IsNull() {
		*dateWindowSize = r.Configuration.DateWindowSize.ValueInt64()
	} else {
		dateWindowSize = nil
	}
	endDate := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	projectID := new(int64)
	if !r.Configuration.ProjectID.IsUnknown() && !r.Configuration.ProjectID.IsNull() {
		*projectID = r.Configuration.ProjectID.ValueInt64()
	} else {
		projectID = nil
	}
	projectTimezone := new(string)
	if !r.Configuration.ProjectTimezone.IsUnknown() && !r.Configuration.ProjectTimezone.IsNull() {
		*projectTimezone = r.Configuration.ProjectTimezone.ValueString()
	} else {
		projectTimezone = nil
	}
	region := new(shared.SourceMixpanelRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceMixpanelRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	selectPropertiesByDefault := new(bool)
	if !r.Configuration.SelectPropertiesByDefault.IsUnknown() && !r.Configuration.SelectPropertiesByDefault.IsNull() {
		*selectPropertiesByDefault = r.Configuration.SelectPropertiesByDefault.ValueBool()
	} else {
		selectPropertiesByDefault = nil
	}
	sourceType := new(shared.SourceMixpanelMixpanel)
	if !r.Configuration.SourceType.IsUnknown() && !r.Configuration.SourceType.IsNull() {
		*sourceType = shared.SourceMixpanelMixpanel(r.Configuration.SourceType.ValueString())
	} else {
		sourceType = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceMixpanel{
		AttributionWindow:         attributionWindow,
		Credentials:               credentials,
		DateWindowSize:            dateWindowSize,
		EndDate:                   endDate,
		ProjectID:                 projectID,
		ProjectTimezone:           projectTimezone,
		Region:                    region,
		SelectPropertiesByDefault: selectPropertiesByDefault,
		SourceType:                sourceType,
		StartDate:                 startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMixpanelCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMixpanelResourceModel) ToUpdateSDKType() *shared.SourceMixpanelPutRequest {
	attributionWindow := new(int64)
	if !r.Configuration.AttributionWindow.IsUnknown() && !r.Configuration.AttributionWindow.IsNull() {
		*attributionWindow = r.Configuration.AttributionWindow.ValueInt64()
	} else {
		attributionWindow = nil
	}
	var credentials *shared.SourceMixpanelUpdateAuthenticationWildcard
	var sourceMixpanelUpdateAuthenticationWildcardServiceAccount *shared.SourceMixpanelUpdateAuthenticationWildcardServiceAccount
	if r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount != nil {
		optionTitle := new(shared.SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle)
		if !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.OptionTitle.IsNull() {
			*optionTitle = shared.SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle(r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.OptionTitle.ValueString())
		} else {
			optionTitle = nil
		}
		secret := r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.Secret.ValueString()
		username := r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardServiceAccount.Username.ValueString()
		sourceMixpanelUpdateAuthenticationWildcardServiceAccount = &shared.SourceMixpanelUpdateAuthenticationWildcardServiceAccount{
			OptionTitle: optionTitle,
			Secret:      secret,
			Username:    username,
		}
	}
	if sourceMixpanelUpdateAuthenticationWildcardServiceAccount != nil {
		credentials = &shared.SourceMixpanelUpdateAuthenticationWildcard{
			SourceMixpanelUpdateAuthenticationWildcardServiceAccount: sourceMixpanelUpdateAuthenticationWildcardServiceAccount,
		}
	}
	var sourceMixpanelUpdateAuthenticationWildcardProjectSecret *shared.SourceMixpanelUpdateAuthenticationWildcardProjectSecret
	if r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret != nil {
		apiSecret := r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.APISecret.ValueString()
		optionTitle1 := new(shared.SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle)
		if !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.OptionTitle.IsUnknown() && !r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.OptionTitle.IsNull() {
			*optionTitle1 = shared.SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle(r.Configuration.Credentials.SourceMixpanelAuthenticationWildcardProjectSecret.OptionTitle.ValueString())
		} else {
			optionTitle1 = nil
		}
		sourceMixpanelUpdateAuthenticationWildcardProjectSecret = &shared.SourceMixpanelUpdateAuthenticationWildcardProjectSecret{
			APISecret:   apiSecret,
			OptionTitle: optionTitle1,
		}
	}
	if sourceMixpanelUpdateAuthenticationWildcardProjectSecret != nil {
		credentials = &shared.SourceMixpanelUpdateAuthenticationWildcard{
			SourceMixpanelUpdateAuthenticationWildcardProjectSecret: sourceMixpanelUpdateAuthenticationWildcardProjectSecret,
		}
	}
	dateWindowSize := new(int64)
	if !r.Configuration.DateWindowSize.IsUnknown() && !r.Configuration.DateWindowSize.IsNull() {
		*dateWindowSize = r.Configuration.DateWindowSize.ValueInt64()
	} else {
		dateWindowSize = nil
	}
	endDate := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	projectID := new(int64)
	if !r.Configuration.ProjectID.IsUnknown() && !r.Configuration.ProjectID.IsNull() {
		*projectID = r.Configuration.ProjectID.ValueInt64()
	} else {
		projectID = nil
	}
	projectTimezone := new(string)
	if !r.Configuration.ProjectTimezone.IsUnknown() && !r.Configuration.ProjectTimezone.IsNull() {
		*projectTimezone = r.Configuration.ProjectTimezone.ValueString()
	} else {
		projectTimezone = nil
	}
	region := new(shared.SourceMixpanelUpdateRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceMixpanelUpdateRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	selectPropertiesByDefault := new(bool)
	if !r.Configuration.SelectPropertiesByDefault.IsUnknown() && !r.Configuration.SelectPropertiesByDefault.IsNull() {
		*selectPropertiesByDefault = r.Configuration.SelectPropertiesByDefault.ValueBool()
	} else {
		selectPropertiesByDefault = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceMixpanelUpdate{
		AttributionWindow:         attributionWindow,
		Credentials:               credentials,
		DateWindowSize:            dateWindowSize,
		EndDate:                   endDate,
		ProjectID:                 projectID,
		ProjectTimezone:           projectTimezone,
		Region:                    region,
		SelectPropertiesByDefault: selectPropertiesByDefault,
		StartDate:                 startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMixpanelPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMixpanelResourceModel) ToDeleteSDKType() *shared.SourceMixpanelCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMixpanelResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}
