// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
)

func (r *SourceSnowflakeResourceModel) ToCreateSDKType() *shared.SourceSnowflakeCreateRequest {
	var credentials *shared.SourceSnowflakeAuthorizationMethod
	var sourceSnowflakeAuthorizationMethodOAuth20 *shared.SourceSnowflakeAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20 != nil {
		accessToken := new(string)
		if !r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.AccessToken.IsUnknown() && !r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.AccessToken.IsNull() {
			*accessToken = r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.AccessToken.ValueString()
		} else {
			accessToken = nil
		}
		authType := shared.SourceSnowflakeAuthorizationMethodOAuth20AuthTypeEnum(r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.AuthType.ValueString())
		clientID := r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := new(string)
		if !r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.RefreshToken.IsUnknown() && !r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.RefreshToken.IsNull() {
			*refreshToken = r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodOAuth20.RefreshToken.ValueString()
		} else {
			refreshToken = nil
		}
		sourceSnowflakeAuthorizationMethodOAuth20 = &shared.SourceSnowflakeAuthorizationMethodOAuth20{
			AccessToken:  accessToken,
			AuthType:     authType,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceSnowflakeAuthorizationMethodOAuth20 != nil {
		credentials = &shared.SourceSnowflakeAuthorizationMethod{
			SourceSnowflakeAuthorizationMethodOAuth20: sourceSnowflakeAuthorizationMethodOAuth20,
		}
	}
	var sourceSnowflakeAuthorizationMethodUsernameAndPassword *shared.SourceSnowflakeAuthorizationMethodUsernameAndPassword
	if r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		authType1 := shared.SourceSnowflakeAuthorizationMethodUsernameAndPasswordAuthTypeEnum(r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodUsernameAndPassword.AuthType.ValueString())
		password := r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodUsernameAndPassword.Password.ValueString()
		username := r.Configuration.Credentials.SourceSnowflakeAuthorizationMethodUsernameAndPassword.Username.ValueString()
		sourceSnowflakeAuthorizationMethodUsernameAndPassword = &shared.SourceSnowflakeAuthorizationMethodUsernameAndPassword{
			AuthType: authType1,
			Password: password,
			Username: username,
		}
	}
	if sourceSnowflakeAuthorizationMethodUsernameAndPassword != nil {
		credentials = &shared.SourceSnowflakeAuthorizationMethod{
			SourceSnowflakeAuthorizationMethodUsernameAndPassword: sourceSnowflakeAuthorizationMethodUsernameAndPassword,
		}
	}
	database := r.Configuration.Database.ValueString()
	host := r.Configuration.Host.ValueString()
	jdbcURLParams := new(string)
	if !r.Configuration.JdbcURLParams.IsUnknown() && !r.Configuration.JdbcURLParams.IsNull() {
		*jdbcURLParams = r.Configuration.JdbcURLParams.ValueString()
	} else {
		jdbcURLParams = nil
	}
	role := r.Configuration.Role.ValueString()
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	sourceType := shared.SourceSnowflakeSnowflakeEnum(r.Configuration.SourceType.ValueString())
	warehouse := r.Configuration.Warehouse.ValueString()
	configuration := shared.SourceSnowflake{
		Credentials:   credentials,
		Database:      database,
		Host:          host,
		JdbcURLParams: jdbcURLParams,
		Role:          role,
		Schema:        schema,
		SourceType:    sourceType,
		Warehouse:     warehouse,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSnowflakeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSnowflakeResourceModel) ToDeleteSDKType() *shared.SourceSnowflakeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}