// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ConnectionScheduleCreate struct {
	CronExpression types.String `tfsdk:"cron_expression"`
	ScheduleType   types.String `tfsdk:"schedule_type"`
	BasicTiming    types.String `tfsdk:"basic_timing"`
}
