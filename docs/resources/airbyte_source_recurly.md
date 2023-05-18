---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_recurly Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  SourceRecurly Resource
---

# airbyte_source_recurly (Resource)

SourceRecurly Resource



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

- `api_key` (String)
- `source_type` (String)

Optional:

- `begin_time` (String)
- `end_time` (String)

