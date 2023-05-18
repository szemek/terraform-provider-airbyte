---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_faker Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  SourceFaker Resource
---

# airbyte_source_faker (Resource)

SourceFaker Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String)

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `count` (Number)
- `source_type` (String)

Optional:

- `parallelism` (Number)
- `records_per_slice` (Number)
- `records_per_sync` (Number)
- `seed` (Number)

