// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGitlabUpdate struct {
	APIURL      types.String                          `tfsdk:"api_url"`
	Credentials SourceGitlabUpdateAuthorizationMethod `tfsdk:"credentials"`
	Groups      types.String                          `tfsdk:"groups"`
	Projects    types.String                          `tfsdk:"projects"`
	StartDate   types.String                          `tfsdk:"start_date"`
	SourceType  types.String                          `tfsdk:"source_type"`
}
