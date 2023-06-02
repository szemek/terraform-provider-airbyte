// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZendeskSunshine struct {
	Credentials *SourceZendeskSunshineAuthorizationMethod `tfsdk:"credentials"`
	SourceType  types.String                              `tfsdk:"source_type"`
	StartDate   types.String                              `tfsdk:"start_date"`
	Subdomain   types.String                              `tfsdk:"subdomain"`
}
