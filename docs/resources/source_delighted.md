---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_delighted Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceDelighted Resource
---

# airbyte_source_delighted (Resource)

SourceDelighted Resource

## Example Usage

```terraform
resource "airbyte_source_delighted" "my_source_delighted" {
  configuration = {
    api_key = "...my_api_key..."
    since   = "2022-05-30T04:50:23Z"
  }
  definition_id = "3b7e8dc3-71ec-4bee-9051-1b439ed171c9"
  name          = "Kirk Windler"
  secret_id     = "...my_secret_id..."
  workspace_id  = "4783ac23-2bfa-441c-80b2-3345c949a955"
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

- `api_key` (String, Sensitive) A Delighted API key.
- `since` (String) The date from which you'd like to replicate the data

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_delighted.my_airbyte_source_delighted ""
```
