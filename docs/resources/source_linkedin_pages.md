---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_linkedin_pages Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceLinkedinPages Resource
---

# airbyte_source_linkedin_pages (Resource)

SourceLinkedinPages Resource

## Example Usage

```terraform
resource "airbyte_source_linkedin_pages" "my_source_linkedinpages" {
  configuration = {
    credentials = {
      access_token = {
        access_token = "...my_access_token..."
      }
    }
    org_id                = "123456789"
    start_date            = "2022-01-06T04:41:29.233Z"
    time_granularity_type = "DAY"
  }
  definition_id = "ae1f1c37-b350-4ebb-b981-c89f963f1e61"
  name          = "Anita Ryan"
  secret_id     = "...my_secret_id..."
  workspace_id  = "788ff77a-5893-43f7-b38d-63dc7b7f8b16"
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

- `org_id` (String) Specify the Organization ID

Optional:

- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `start_date` (String) Start date for getting metrics per time period. Must be atmost 12 months before the request date (UTC) and atleast 2 days prior to the request date (UTC). See https://bit.ly/linkedin-pages-date-rules {{ "\n" }} {{ response.errorDetails }}. Default: "2023-01-01T00:00:00Z"
- `time_granularity_type` (String) Granularity of the statistics for metrics per time period. Must be either "DAY" or "MONTH". must be one of ["DAY", "MONTH"]; Default: "DAY"

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--access_token))
- `o_auth20` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--o_auth20))

<a id="nestedatt--configuration--credentials--access_token"></a>
### Nested Schema for `configuration.credentials.access_token`

Required:

- `access_token` (String, Sensitive) The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.


<a id="nestedatt--configuration--credentials--o_auth20"></a>
### Nested Schema for `configuration.credentials.o_auth20`

Required:

- `client_id` (String) The client ID of the LinkedIn developer application.
- `client_secret` (String) The client secret of the LinkedIn developer application.
- `refresh_token` (String, Sensitive) The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_linkedin_pages.my_airbyte_source_linkedin_pages ""
```
