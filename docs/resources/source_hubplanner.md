---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_hubplanner Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceHubplanner Resource
---

# airbyte_source_hubplanner (Resource)

SourceHubplanner Resource

## Example Usage

```terraform
resource "airbyte_source_hubplanner" "my_source_hubplanner" {
  configuration = {
    api_key = "...my_api_key..."
  }
  definition_id = "85c78fa7-d86b-4df5-bf91-acb121083728"
  name          = "Johnnie Maggio"
  secret_id     = "...my_secret_id..."
  workspace_id  = "51e868df-1f2c-45ad-84a4-6153eb240d62"
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

- `api_key` (String, Sensitive) Hubplanner API key. See https://github.com/hubplanner/API#authentication for more details.

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_hubplanner.my_airbyte_source_hubplanner ""
```
