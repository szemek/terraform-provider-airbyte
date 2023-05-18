// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceGoogleDirectoryRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceGoogleDirectoryResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}