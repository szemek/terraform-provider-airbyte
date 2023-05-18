// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"time"
)

func (r *SourceBraintreeResourceModel) ToCreateSDKType() *shared.SourceBraintreeCreateRequest {
	environment := shared.SourceBraintreeEnvironmentEnum(r.Configuration.Environment.ValueString())
	merchantID := r.Configuration.MerchantID.ValueString()
	privateKey := r.Configuration.PrivateKey.ValueString()
	publicKey := r.Configuration.PublicKey.ValueString()
	sourceType := shared.SourceBraintreeBraintreeEnum(r.Configuration.SourceType.ValueString())
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceBraintree{
		Environment: environment,
		MerchantID:  merchantID,
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
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
	out := shared.SourceBraintreeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceBraintreeResourceModel) ToDeleteSDKType() *shared.SourceBraintreeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}