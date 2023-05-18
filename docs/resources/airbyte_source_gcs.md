---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_gcs Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  SourceGcs Resource
---

# airbyte_source_gcs (Resource)

SourceGcs Resource



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

- `gcs_bucket` (String)
- `gcs_path` (String)
- `service_account` (String)
- `source_type` (String)

