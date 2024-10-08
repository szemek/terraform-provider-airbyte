// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/types"
)

type SourceAwsCloudtrailUpdate struct {
	// AWS CloudTrail Access Key ID. See the <a href="https://docs.airbyte.com/integrations/sources/aws-cloudtrail">docs</a> for more information on how to obtain this key.
	AwsKeyID string `json:"aws_key_id"`
	// The default AWS Region to use, for example, us-west-1 or us-west-2. When specifying a Region inline during client initialization, this property is named region_name.
	AwsRegionName string `json:"aws_region_name"`
	// AWS CloudTrail Access Key ID. See the <a href="https://docs.airbyte.com/integrations/sources/aws-cloudtrail">docs</a> for more information on how to obtain this key.
	AwsSecretKey string `json:"aws_secret_key"`
	// The date you would like to replicate data. Data in AWS CloudTrail is available for last 90 days only. Format: YYYY-MM-DD.
	StartDate *types.Date `default:"1970-01-01" json:"start_date"`
}

func (s SourceAwsCloudtrailUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceAwsCloudtrailUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceAwsCloudtrailUpdate) GetAwsKeyID() string {
	if o == nil {
		return ""
	}
	return o.AwsKeyID
}

func (o *SourceAwsCloudtrailUpdate) GetAwsRegionName() string {
	if o == nil {
		return ""
	}
	return o.AwsRegionName
}

func (o *SourceAwsCloudtrailUpdate) GetAwsSecretKey() string {
	if o == nil {
		return ""
	}
	return o.AwsSecretKey
}

func (o *SourceAwsCloudtrailUpdate) GetStartDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.StartDate
}
