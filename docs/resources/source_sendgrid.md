---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_sendgrid Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSendgrid Resource
---

# airbyte_source_sendgrid (Resource)

SourceSendgrid Resource

## Example Usage

```terraform
resource "airbyte_source_sendgrid" "my_source_sendgrid" {
  configuration = {
    api_key    = "...my_api_key..."
    start_date = "2020-02-29T00:35:24.807Z"
  }
  definition_id = "57eb672b-8aa5-45d6-bfb2-a63da0917a61"
  name          = "Christine Weimann"
  secret_id     = "...my_secret_id..."
  workspace_id  = "3e8ec69b-abb3-4389-b4cd-0d539af2319a"
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

- `api_key` (String, Sensitive) Sendgrid API Key, use <a href=\"https://app.sendgrid.com/settings/api_keys/\">admin</a> to generate this key.
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_sendgrid.my_airbyte_source_sendgrid ""
```
