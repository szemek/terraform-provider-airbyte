---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_amplitude Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAmplitude Resource
---

# airbyte_source_amplitude (Resource)

SourceAmplitude Resource

## Example Usage

```terraform
resource "airbyte_source_amplitude" "my_source_amplitude" {
  configuration = {
    api_key            = "...my_api_key..."
    data_region        = "EU Residency Server"
    request_time_range = 5
    secret_key         = "...my_secret_key..."
    start_date         = "2021-01-25T00:00:00Z"
  }
  definition_id = "21b517b1-6f1f-4884-abcd-5137451945c4"
  name          = "Mrs. Sylvia Howe"
  secret_id     = "...my_secret_id..."
  workspace_id  = "6ae8aa3c-4f28-4791-bb86-68105e1180fb"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String) Name of the source e.g. dev-mysql-instance.
- `workspace_id` (String)

### Optional

- `definition_id` (String) The UUID of the connector definition. One of configuration.sourceType or definitionId must be provided. Requires replacement if changed.
- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow. Requires replacement if changed.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `api_key` (String, Sensitive) Amplitude API Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
- `secret_key` (String, Sensitive) Amplitude Secret Key. See the <a href="https://docs.airbyte.com/integrations/sources/amplitude#setup-guide">setup guide</a> for more information on how to obtain this key.
- `start_date` (String) UTC date and time in the format 2021-01-25T00:00:00Z. Any data before this date will not be replicated.

Optional:

- `data_region` (String) Amplitude data region server. must be one of ["Standard Server", "EU Residency Server"]; Default: "Standard Server"
- `request_time_range` (Number) According to <a href="https://www.docs.developers.amplitude.com/analytics/apis/export-api/#considerations">Considerations</a> too big time range in request can cause a timeout error. In this case, set shorter time interval in hours. Default: 24

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_amplitude.my_airbyte_source_amplitude ""
```
