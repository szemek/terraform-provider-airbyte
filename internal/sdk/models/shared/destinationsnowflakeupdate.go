// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type DestinationSnowflakeUpdateSchemasAuthType string

const (
	DestinationSnowflakeUpdateSchemasAuthTypeOAuth20 DestinationSnowflakeUpdateSchemasAuthType = "OAuth2.0"
)

func (e DestinationSnowflakeUpdateSchemasAuthType) ToPointer() *DestinationSnowflakeUpdateSchemasAuthType {
	return &e
}
func (e *DestinationSnowflakeUpdateSchemasAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth2.0":
		*e = DestinationSnowflakeUpdateSchemasAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeUpdateSchemasAuthType: %v", v)
	}
}

type DestinationSnowflakeUpdateOAuth20 struct {
	// Enter you application's Access Token
	AccessToken string                                     `json:"access_token"`
	authType    *DestinationSnowflakeUpdateSchemasAuthType `const:"OAuth2.0" json:"auth_type"`
	// Enter your application's Client ID
	ClientID *string `json:"client_id,omitempty"`
	// Enter your application's Client secret
	ClientSecret *string `json:"client_secret,omitempty"`
	// Enter your application's Refresh Token
	RefreshToken string `json:"refresh_token"`
}

func (d DestinationSnowflakeUpdateOAuth20) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeUpdateOAuth20) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeUpdateOAuth20) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *DestinationSnowflakeUpdateOAuth20) GetAuthType() *DestinationSnowflakeUpdateSchemasAuthType {
	return DestinationSnowflakeUpdateSchemasAuthTypeOAuth20.ToPointer()
}

func (o *DestinationSnowflakeUpdateOAuth20) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *DestinationSnowflakeUpdateOAuth20) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *DestinationSnowflakeUpdateOAuth20) GetRefreshToken() string {
	if o == nil {
		return ""
	}
	return o.RefreshToken
}

type DestinationSnowflakeUpdateAuthType string

const (
	DestinationSnowflakeUpdateAuthTypeUsernameAndPassword DestinationSnowflakeUpdateAuthType = "Username and Password"
)

func (e DestinationSnowflakeUpdateAuthType) ToPointer() *DestinationSnowflakeUpdateAuthType {
	return &e
}
func (e *DestinationSnowflakeUpdateAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Username and Password":
		*e = DestinationSnowflakeUpdateAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeUpdateAuthType: %v", v)
	}
}

type UsernameAndPassword struct {
	authType *DestinationSnowflakeUpdateAuthType `const:"Username and Password" json:"auth_type"`
	// Enter the password associated with the username.
	Password string `json:"password"`
}

func (u UsernameAndPassword) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UsernameAndPassword) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *UsernameAndPassword) GetAuthType() *DestinationSnowflakeUpdateAuthType {
	return DestinationSnowflakeUpdateAuthTypeUsernameAndPassword.ToPointer()
}

func (o *UsernameAndPassword) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

type DestinationSnowflakeUpdateSchemasCredentialsAuthType string

const (
	DestinationSnowflakeUpdateSchemasCredentialsAuthTypeKeyPairAuthentication DestinationSnowflakeUpdateSchemasCredentialsAuthType = "Key Pair Authentication"
)

func (e DestinationSnowflakeUpdateSchemasCredentialsAuthType) ToPointer() *DestinationSnowflakeUpdateSchemasCredentialsAuthType {
	return &e
}
func (e *DestinationSnowflakeUpdateSchemasCredentialsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Key Pair Authentication":
		*e = DestinationSnowflakeUpdateSchemasCredentialsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationSnowflakeUpdateSchemasCredentialsAuthType: %v", v)
	}
}

type KeyPairAuthentication struct {
	authType *DestinationSnowflakeUpdateSchemasCredentialsAuthType `const:"Key Pair Authentication" json:"auth_type"`
	// RSA Private key to use for Snowflake connection. See the <a href="https://docs.airbyte.com/integrations/destinations/snowflake">docs</a> for more information on how to obtain this key.
	PrivateKey string `json:"private_key"`
	// Passphrase for private key
	PrivateKeyPassword *string `json:"private_key_password,omitempty"`
}

func (k KeyPairAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyPairAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *KeyPairAuthentication) GetAuthType() *DestinationSnowflakeUpdateSchemasCredentialsAuthType {
	return DestinationSnowflakeUpdateSchemasCredentialsAuthTypeKeyPairAuthentication.ToPointer()
}

func (o *KeyPairAuthentication) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}

func (o *KeyPairAuthentication) GetPrivateKeyPassword() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKeyPassword
}

type AuthorizationMethodType string

const (
	AuthorizationMethodTypeKeyPairAuthentication             AuthorizationMethodType = "Key Pair Authentication"
	AuthorizationMethodTypeUsernameAndPassword               AuthorizationMethodType = "Username and Password"
	AuthorizationMethodTypeDestinationSnowflakeUpdateOAuth20 AuthorizationMethodType = "destination-snowflake-update_OAuth2.0"
)

type AuthorizationMethod struct {
	KeyPairAuthentication             *KeyPairAuthentication
	UsernameAndPassword               *UsernameAndPassword
	DestinationSnowflakeUpdateOAuth20 *DestinationSnowflakeUpdateOAuth20

	Type AuthorizationMethodType
}

func CreateAuthorizationMethodKeyPairAuthentication(keyPairAuthentication KeyPairAuthentication) AuthorizationMethod {
	typ := AuthorizationMethodTypeKeyPairAuthentication

	return AuthorizationMethod{
		KeyPairAuthentication: &keyPairAuthentication,
		Type:                  typ,
	}
}

func CreateAuthorizationMethodUsernameAndPassword(usernameAndPassword UsernameAndPassword) AuthorizationMethod {
	typ := AuthorizationMethodTypeUsernameAndPassword

	return AuthorizationMethod{
		UsernameAndPassword: &usernameAndPassword,
		Type:                typ,
	}
}

func CreateAuthorizationMethodDestinationSnowflakeUpdateOAuth20(destinationSnowflakeUpdateOAuth20 DestinationSnowflakeUpdateOAuth20) AuthorizationMethod {
	typ := AuthorizationMethodTypeDestinationSnowflakeUpdateOAuth20

	return AuthorizationMethod{
		DestinationSnowflakeUpdateOAuth20: &destinationSnowflakeUpdateOAuth20,
		Type:                              typ,
	}
}

func (u *AuthorizationMethod) UnmarshalJSON(data []byte) error {

	var usernameAndPassword UsernameAndPassword = UsernameAndPassword{}
	if err := utils.UnmarshalJSON(data, &usernameAndPassword, "", true, true); err == nil {
		u.UsernameAndPassword = &usernameAndPassword
		u.Type = AuthorizationMethodTypeUsernameAndPassword
		return nil
	}

	var keyPairAuthentication KeyPairAuthentication = KeyPairAuthentication{}
	if err := utils.UnmarshalJSON(data, &keyPairAuthentication, "", true, true); err == nil {
		u.KeyPairAuthentication = &keyPairAuthentication
		u.Type = AuthorizationMethodTypeKeyPairAuthentication
		return nil
	}

	var destinationSnowflakeUpdateOAuth20 DestinationSnowflakeUpdateOAuth20 = DestinationSnowflakeUpdateOAuth20{}
	if err := utils.UnmarshalJSON(data, &destinationSnowflakeUpdateOAuth20, "", true, true); err == nil {
		u.DestinationSnowflakeUpdateOAuth20 = &destinationSnowflakeUpdateOAuth20
		u.Type = AuthorizationMethodTypeDestinationSnowflakeUpdateOAuth20
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for AuthorizationMethod", string(data))
}

func (u AuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.KeyPairAuthentication != nil {
		return utils.MarshalJSON(u.KeyPairAuthentication, "", true)
	}

	if u.UsernameAndPassword != nil {
		return utils.MarshalJSON(u.UsernameAndPassword, "", true)
	}

	if u.DestinationSnowflakeUpdateOAuth20 != nil {
		return utils.MarshalJSON(u.DestinationSnowflakeUpdateOAuth20, "", true)
	}

	return nil, errors.New("could not marshal union type AuthorizationMethod: all fields are null")
}

type DestinationSnowflakeUpdate struct {
	Credentials *AuthorizationMethod `json:"credentials,omitempty"`
	// Enter the name of the <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">database</a> you want to sync data into
	Database string `json:"database"`
	// Disable Writing Final Tables. WARNING! The data format in _airbyte_data is likely stable but there are no guarantees that other metadata columns will remain the same in future versions
	DisableTypeDedupe *bool `default:"false" json:"disable_type_dedupe"`
	// Enter your Snowflake account's <a href="https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#using-an-account-locator-as-an-identifier">locator</a> (in the format <account_locator>.<region>.<cloud>.snowflakecomputing.com)
	Host string `json:"host"`
	// Enter the additional properties to pass to the JDBC URL string when connecting to the database (formatted as key=value pairs separated by the symbol &). Example: key1=value1&key2=value2&key3=value3
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// The schema to write raw tables into (default: airbyte_internal)
	RawDataSchema *string `json:"raw_data_schema,omitempty"`
	// The number of days of Snowflake Time Travel to enable on the tables. See <a href="https://docs.snowflake.com/en/user-guide/data-time-travel#data-retention-period">Snowflake's documentation</a> for more information. Setting a nonzero value will incur increased storage costs in your Snowflake instance.
	RetentionPeriodDays *int64 `default:"1" json:"retention_period_days"`
	// Enter the <a href="https://docs.snowflake.com/en/user-guide/security-access-control-overview.html#roles">role</a> that you want to use to access Snowflake
	Role string `json:"role"`
	// Enter the name of the default <a href="https://docs.snowflake.com/en/sql-reference/ddl-database.html#database-schema-share-ddl">schema</a>
	Schema string `json:"schema"`
	// Enter the name of the user you want to use to access the database
	Username string `json:"username"`
	// Enter the name of the <a href="https://docs.snowflake.com/en/user-guide/warehouses-overview.html#overview-of-warehouses">warehouse</a> that you want to sync data into
	Warehouse string `json:"warehouse"`
}

func (d DestinationSnowflakeUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationSnowflakeUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationSnowflakeUpdate) GetCredentials() *AuthorizationMethod {
	if o == nil {
		return nil
	}
	return o.Credentials
}

func (o *DestinationSnowflakeUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationSnowflakeUpdate) GetDisableTypeDedupe() *bool {
	if o == nil {
		return nil
	}
	return o.DisableTypeDedupe
}

func (o *DestinationSnowflakeUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationSnowflakeUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationSnowflakeUpdate) GetRawDataSchema() *string {
	if o == nil {
		return nil
	}
	return o.RawDataSchema
}

func (o *DestinationSnowflakeUpdate) GetRetentionPeriodDays() *int64 {
	if o == nil {
		return nil
	}
	return o.RetentionPeriodDays
}

func (o *DestinationSnowflakeUpdate) GetRole() string {
	if o == nil {
		return ""
	}
	return o.Role
}

func (o *DestinationSnowflakeUpdate) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *DestinationSnowflakeUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *DestinationSnowflakeUpdate) GetWarehouse() string {
	if o == nil {
		return ""
	}
	return o.Warehouse
}
