// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi-example/pkg/models/shared"
)

type UpdateUserFormRequest struct {
	// Update an existent user in the store
	User *shared.User `request:"mediaType=application/x-www-form-urlencoded"`
	// name that needs to be updated
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type UpdateUserFormResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
