---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_aha Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAha Resource
---

# airbyte_source_aha (Resource)

SourceAha Resource

## Example Usage

```terraform
resource "airbyte_source_aha" "my_source_aha" {
  configuration = {
    api_key = "...my_api_key..."
    url     = "...my_url..."
  }
  definition_id = "6eef0475-7630-4ddb-82db-f188dfabd571"
  name          = "Joyce Pagac"
  secret_id     = "...my_secret_id..."
  workspace_id  = "1dffa69f-e714-43a3-a9a2-44d7bc1a5a6e"
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
- `url` (String) URL

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_aha.my_airbyte_source_aha ""
```
