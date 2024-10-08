---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_pipedrive Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePipedrive Resource
---

# airbyte_source_pipedrive (Resource)

SourcePipedrive Resource

## Example Usage

```terraform
resource "airbyte_source_pipedrive" "my_source_pipedrive" {
  configuration = {
    api_token              = "...my_api_token..."
    replication_start_date = "2017-01-25 00:00:00Z"
  }
  definition_id = "d9282ad1-9d25-4d52-93fa-02ef008f118d"
  name          = "Peter Hilll"
  secret_id     = "...my_secret_id..."
  workspace_id  = "f724d1e0-e7e7-408b-9f81-5bf9f1370c28"
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

- `api_token` (String, Sensitive) The Pipedrive API Token.
- `replication_start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. When specified and not None, then stream will behave as incremental

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_pipedrive.my_airbyte_source_pipedrive ""
```
