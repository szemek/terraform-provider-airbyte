// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFacebookMarketingUpdateInsightConfig struct {
	ActionBreakdowns       []types.String `tfsdk:"action_breakdowns"`
	Breakdowns             []types.String `tfsdk:"breakdowns"`
	EndDate                types.String   `tfsdk:"end_date"`
	Fields                 []types.String `tfsdk:"fields"`
	InsightsLookbackWindow types.Int64    `tfsdk:"insights_lookback_window"`
	Level                  types.String   `tfsdk:"level"`
	Name                   types.String   `tfsdk:"name"`
	StartDate              types.String   `tfsdk:"start_date"`
	TimeIncrement          types.Int64    `tfsdk:"time_increment"`
}
