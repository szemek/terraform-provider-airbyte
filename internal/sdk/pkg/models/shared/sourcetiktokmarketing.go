// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType string

const (
	SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthTypeSandboxAccessToken SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType = "sandbox_access_token"
)

func (e SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType) ToPointer() *SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType {
	return &e
}

func (e *SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sandbox_access_token":
		*e = SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType: %v", v)
	}
}

// SourceTiktokMarketingAuthenticationMethodSandboxAccessToken - Authentication method
type SourceTiktokMarketingAuthenticationMethodSandboxAccessToken struct {
	// The long-term authorized access token.
	AccessToken string `json:"access_token"`
	// The Advertiser ID which generated for the developer's Sandbox application.
	AdvertiserID string                                                               `json:"advertiser_id"`
	AuthType     *SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType `json:"auth_type,omitempty"`
}

type SourceTiktokMarketingAuthenticationMethodOAuth20AuthType string

const (
	SourceTiktokMarketingAuthenticationMethodOAuth20AuthTypeOauth20 SourceTiktokMarketingAuthenticationMethodOAuth20AuthType = "oauth2.0"
)

func (e SourceTiktokMarketingAuthenticationMethodOAuth20AuthType) ToPointer() *SourceTiktokMarketingAuthenticationMethodOAuth20AuthType {
	return &e
}

func (e *SourceTiktokMarketingAuthenticationMethodOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceTiktokMarketingAuthenticationMethodOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTiktokMarketingAuthenticationMethodOAuth20AuthType: %v", v)
	}
}

// SourceTiktokMarketingAuthenticationMethodOAuth20 - Authentication method
type SourceTiktokMarketingAuthenticationMethodOAuth20 struct {
	// Long-term Authorized Access Token.
	AccessToken string `json:"access_token"`
	// The Advertiser ID to filter reports and streams. Let this empty to retrieve all.
	AdvertiserID *string `json:"advertiser_id,omitempty"`
	// The Developer Application App ID.
	AppID    string                                                    `json:"app_id"`
	AuthType *SourceTiktokMarketingAuthenticationMethodOAuth20AuthType `json:"auth_type,omitempty"`
	// The Developer Application Secret.
	Secret string `json:"secret"`
}

type SourceTiktokMarketingAuthenticationMethodType string

const (
	SourceTiktokMarketingAuthenticationMethodTypeSourceTiktokMarketingAuthenticationMethodOAuth20            SourceTiktokMarketingAuthenticationMethodType = "source-tiktok-marketing_Authentication Method_OAuth2.0"
	SourceTiktokMarketingAuthenticationMethodTypeSourceTiktokMarketingAuthenticationMethodSandboxAccessToken SourceTiktokMarketingAuthenticationMethodType = "source-tiktok-marketing_Authentication Method_Sandbox Access Token"
)

type SourceTiktokMarketingAuthenticationMethod struct {
	SourceTiktokMarketingAuthenticationMethodOAuth20            *SourceTiktokMarketingAuthenticationMethodOAuth20
	SourceTiktokMarketingAuthenticationMethodSandboxAccessToken *SourceTiktokMarketingAuthenticationMethodSandboxAccessToken

	Type SourceTiktokMarketingAuthenticationMethodType
}

func CreateSourceTiktokMarketingAuthenticationMethodSourceTiktokMarketingAuthenticationMethodOAuth20(sourceTiktokMarketingAuthenticationMethodOAuth20 SourceTiktokMarketingAuthenticationMethodOAuth20) SourceTiktokMarketingAuthenticationMethod {
	typ := SourceTiktokMarketingAuthenticationMethodTypeSourceTiktokMarketingAuthenticationMethodOAuth20

	return SourceTiktokMarketingAuthenticationMethod{
		SourceTiktokMarketingAuthenticationMethodOAuth20: &sourceTiktokMarketingAuthenticationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceTiktokMarketingAuthenticationMethodSourceTiktokMarketingAuthenticationMethodSandboxAccessToken(sourceTiktokMarketingAuthenticationMethodSandboxAccessToken SourceTiktokMarketingAuthenticationMethodSandboxAccessToken) SourceTiktokMarketingAuthenticationMethod {
	typ := SourceTiktokMarketingAuthenticationMethodTypeSourceTiktokMarketingAuthenticationMethodSandboxAccessToken

	return SourceTiktokMarketingAuthenticationMethod{
		SourceTiktokMarketingAuthenticationMethodSandboxAccessToken: &sourceTiktokMarketingAuthenticationMethodSandboxAccessToken,
		Type: typ,
	}
}

func (u *SourceTiktokMarketingAuthenticationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceTiktokMarketingAuthenticationMethodOAuth20 := new(SourceTiktokMarketingAuthenticationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceTiktokMarketingAuthenticationMethodOAuth20); err == nil {
		u.SourceTiktokMarketingAuthenticationMethodOAuth20 = sourceTiktokMarketingAuthenticationMethodOAuth20
		u.Type = SourceTiktokMarketingAuthenticationMethodTypeSourceTiktokMarketingAuthenticationMethodOAuth20
		return nil
	}

	sourceTiktokMarketingAuthenticationMethodSandboxAccessToken := new(SourceTiktokMarketingAuthenticationMethodSandboxAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceTiktokMarketingAuthenticationMethodSandboxAccessToken); err == nil {
		u.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken = sourceTiktokMarketingAuthenticationMethodSandboxAccessToken
		u.Type = SourceTiktokMarketingAuthenticationMethodTypeSourceTiktokMarketingAuthenticationMethodSandboxAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceTiktokMarketingAuthenticationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceTiktokMarketingAuthenticationMethodOAuth20 != nil {
		return json.Marshal(u.SourceTiktokMarketingAuthenticationMethodOAuth20)
	}

	if u.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken != nil {
		return json.Marshal(u.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken)
	}

	return nil, nil
}

type SourceTiktokMarketingTiktokMarketing string

const (
	SourceTiktokMarketingTiktokMarketingTiktokMarketing SourceTiktokMarketingTiktokMarketing = "tiktok-marketing"
)

func (e SourceTiktokMarketingTiktokMarketing) ToPointer() *SourceTiktokMarketingTiktokMarketing {
	return &e
}

func (e *SourceTiktokMarketingTiktokMarketing) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tiktok-marketing":
		*e = SourceTiktokMarketingTiktokMarketing(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTiktokMarketingTiktokMarketing: %v", v)
	}
}

type SourceTiktokMarketing struct {
	// The attribution window in days.
	AttributionWindow *int64 `json:"attribution_window,omitempty"`
	// Authentication method
	Credentials *SourceTiktokMarketingAuthenticationMethod `json:"credentials,omitempty"`
	// The date until which you'd like to replicate data for all incremental streams, in the format YYYY-MM-DD. All data generated between start_date and this date will be replicated. Not setting this option will result in always syncing the data till the current date.
	EndDate    *types.Date                           `json:"end_date,omitempty"`
	SourceType *SourceTiktokMarketingTiktokMarketing `json:"sourceType,omitempty"`
	// The Start Date in format: YYYY-MM-DD. Any data before this date will not be replicated. If this parameter is not set, all data will be replicated.
	StartDate *types.Date `json:"start_date,omitempty"`
}
