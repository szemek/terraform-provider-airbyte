---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_smartengage Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSmartengage Resource
---

# airbyte_source_smartengage (Resource)

SourceSmartengage Resource

## Example Usage

```terraform
resource "airbyte_source_smartengage" "my_source_smartengage" {
  configuration = {
    api_key = "...my_api_key..."
  }
  definition_id = "65f64672-3901-4f87-89df-1af8f5013d5d"
  name          = "Robyn Weimann I"
  secret_id     = "...my_secret_id..."
  workspace_id  = "b2856e98-a695-40f0-807e-33047d95358a"
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

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_smartengage.my_airbyte_source_smartengage ""
```
