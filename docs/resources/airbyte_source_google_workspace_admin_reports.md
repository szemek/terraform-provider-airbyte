---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_workspace_admin_reports Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  SourceGoogleWorkspaceAdminReports Resource
---

# airbyte_source_google_workspace_admin_reports (Resource)

SourceGoogleWorkspaceAdminReports Resource



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

- `credentials_json` (String)
- `email` (String)
- `source_type` (String)

Optional:

- `lookback` (Number)

