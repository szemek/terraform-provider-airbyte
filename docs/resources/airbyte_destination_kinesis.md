---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_kinesis Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  DestinationKinesis Resource
---

# airbyte_destination_kinesis (Resource)

DestinationKinesis Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `access_key` (String)
- `buffer_size` (Number)
- `destination_type` (String)
- `endpoint` (String)
- `private_key` (String)
- `region` (String)
- `shard_count` (Number)

