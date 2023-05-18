---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_redshift Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  SourceRedshift Resource
---

# airbyte_source_redshift (Resource)

SourceRedshift Resource



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

- `database` (String)
- `host` (String)
- `password` (String)
- `port` (Number)
- `source_type` (String)
- `username` (String)

Optional:

- `jdbc_url_params` (String)
- `schemas` (List of String)

