// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi-example/pkg/models/shared"
)

type PlaceOrderJSONResponse struct {
	ContentType string
	// successful operation
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
}
