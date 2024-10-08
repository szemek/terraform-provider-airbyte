---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_datascope Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceDatascope Resource
---

# airbyte_source_datascope (Resource)

SourceDatascope Resource

## Example Usage

```terraform
resource "airbyte_source_datascope" "my_source_datascope" {
  configuration = {
    api_key    = "...my_api_key..."
    start_date = "dd/mm/YYYY HH:MM"
  }
  definition_id = "8d140946-39cf-45dd-8a0c-05f536f6b9b8"
  name          = "Dave Willms"
  secret_id     = "...my_secret_id..."
  workspace_id  = "6afbf365-d687-4e08-be39-05b6a417faeb"
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

- `api_key` (String, Sensitive) API Key
- `start_date` (String) Start date for the data to be replicated

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_datascope.my_airbyte_source_datascope ""
```
