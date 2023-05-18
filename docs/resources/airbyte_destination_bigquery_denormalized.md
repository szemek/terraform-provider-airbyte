---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_bigquery_denormalized Resource - terraform-provider-airbyte-new"
subcategory: ""
description: |-
  DestinationBigqueryDenormalized Resource
---

# airbyte_destination_bigquery_denormalized (Resource)

DestinationBigqueryDenormalized Resource



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

- `dataset_id` (String)
- `destination_type` (String)
- `project_id` (String)

Optional:

- `big_query_client_buffer_size_mb` (Number)
- `credentials_json` (String)
- `dataset_location` (String) The location of the dataset. Warning: Changes made after creation will not be applied. The default "US" value is used if not set explicitly. Read more <a href="https://cloud.google.com/bigquery/docs/locations">here</a>.
- `loading_method` (Attributes) (see [below for nested schema](#nestedatt--configuration--loading_method))

<a id="nestedatt--configuration--loading_method"></a>
### Nested Schema for `configuration.loading_method`

Optional:

- `destination_bigquery_denormalized_loading_method_gcs_staging` (Attributes) Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>. (see [below for nested schema](#nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_gcs_staging))
- `destination_bigquery_denormalized_loading_method_standard_inserts` (Attributes) Loading method used to send select the way data will be uploaded to BigQuery. <br/><b>Standard Inserts</b> - Direct uploading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In almost all cases, you should use staging. <br/><b>GCS Staging</b> - Writes large batches of records to a file, uploads the file to GCS, then uses <b>COPY INTO table</b> to upload the file. Recommended for most workloads for better speed and scalability. Read more about GCS Staging <a href="https://docs.airbyte.com/integrations/destinations/bigquery#gcs-staging">here</a>. (see [below for nested schema](#nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_standard_inserts))

<a id="nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_gcs_staging"></a>
### Nested Schema for `configuration.loading_method.destination_bigquery_denormalized_loading_method_gcs_staging`

Required:

- `credential` (Attributes) (see [below for nested schema](#nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_gcs_staging--credential))
- `gcs_bucket_name` (String)
- `gcs_bucket_path` (String)
- `method` (String)

Optional:

- `file_buffer_count` (Number)
- `keep_files_in_gcs_bucket` (String) This upload method is supposed to temporary store records in GCS bucket. By this select you can chose if these records should be removed from GCS when migration has finished. The default "Delete all tmp files from GCS" value is used if not set explicitly.

<a id="nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_gcs_staging--credential"></a>
### Nested Schema for `configuration.loading_method.destination_bigquery_denormalized_loading_method_gcs_staging.keep_files_in_gcs_bucket`

Optional:

- `destination_bigquery_denormalized_loading_method_gcs_staging_credential_hmac_key` (Attributes) An HMAC key is a type of credential and can be associated with a service account or a user account in Cloud Storage. Read more <a href="https://cloud.google.com/storage/docs/authentication/hmackeys">here</a>. (see [below for nested schema](#nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_gcs_staging--keep_files_in_gcs_bucket--destination_bigquery_denormalized_loading_method_gcs_staging_credential_hmac_key))

<a id="nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_gcs_staging--keep_files_in_gcs_bucket--destination_bigquery_denormalized_loading_method_gcs_staging_credential_hmac_key"></a>
### Nested Schema for `configuration.loading_method.destination_bigquery_denormalized_loading_method_gcs_staging.keep_files_in_gcs_bucket.destination_bigquery_denormalized_loading_method_gcs_staging_credential_hmac_key`

Required:

- `credential_type` (String)
- `hmac_key_access_id` (String)
- `hmac_key_secret` (String)




<a id="nestedatt--configuration--loading_method--destination_bigquery_denormalized_loading_method_standard_inserts"></a>
### Nested Schema for `configuration.loading_method.destination_bigquery_denormalized_loading_method_standard_inserts`

Required:

- `method` (String)

