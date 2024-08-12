// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type PutDestinationLangchainRequest struct {
	DestinationLangchainPutRequest *shared.DestinationLangchainPutRequest `request:"mediaType=application/json"`
	DestinationID                  string                                 `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *PutDestinationLangchainRequest) GetDestinationLangchainPutRequest() *shared.DestinationLangchainPutRequest {
	if o == nil {
		return nil
	}
	return o.DestinationLangchainPutRequest
}

func (o *PutDestinationLangchainRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

type PutDestinationLangchainResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PutDestinationLangchainResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutDestinationLangchainResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutDestinationLangchainResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
