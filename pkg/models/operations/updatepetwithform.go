// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type UpdatePetWithFormSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type UpdatePetWithFormRequest struct {
	// Name of pet that needs to be updated
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// ID of pet that needs to be updated
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
	// Status of pet that needs to be updated
	Status *string `queryParam:"style=form,explode=true,name=status"`
}

type UpdatePetWithFormResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
