---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_zoho_crm Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceZohoCrm Resource
---

# airbyte_source_zoho_crm (Resource)

SourceZohoCrm Resource

## Example Usage

```terraform
resource "airbyte_source_zoho_crm" "my_source_zohocrm" {
  configuration = {
    client_id      = "...my_client_id..."
    client_secret  = "...my_client_secret..."
    dc_region      = "US"
    edition        = "Professional"
    environment    = "Developer"
    refresh_token  = "...my_refresh_token..."
    start_datetime = "2000-01-01T13:00+00:00"
  }
  definition_id = "51f0c20e-4312-4d0c-bfe3-9df03e297d6f"
  name          = "Roxanne Yundt"
  secret_id     = "...my_secret_id..."
  workspace_id  = "b34f9589-f421-498f-b282-2b82a159ebc2"
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

- `client_id` (String) OAuth2.0 Client ID
- `client_secret` (String) OAuth2.0 Client Secret
- `dc_region` (String) Please choose the region of your Data Center location. More info by this <a href="https://www.zoho.com/crm/developer/docs/api/v2/multi-dc.html">Link</a>. must be one of ["US", "AU", "EU", "IN", "CN", "JP"]
- `environment` (String) Please choose the environment. must be one of ["Production", "Developer", "Sandbox"]
- `refresh_token` (String, Sensitive) OAuth2.0 Refresh Token

Optional:

- `edition` (String) Choose your Edition of Zoho CRM to determine API Concurrency Limits. must be one of ["Free", "Standard", "Professional", "Enterprise", "Ultimate"]; Default: "Free"
- `start_datetime` (String) ISO 8601, for instance: `YYYY-MM-DD`, `YYYY-MM-DD HH:MM:SS+HH:MM`

## Import

Import is supported using the following syntax:

```shell
terraform import airbyte_source_zoho_crm.my_airbyte_source_zoho_crm ""
```
