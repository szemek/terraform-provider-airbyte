// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFaunaUpdateCollection struct {
	Deletions SourceFaunaUpdateCollectionDeletionMode `tfsdk:"deletions"`
	PageSize  types.Int64                             `tfsdk:"page_size"`
}
