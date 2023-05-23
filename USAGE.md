<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
	"openapi-example/pkg/models/operations"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: openapiexample.Int64(1),
            Name: openapiexample.String("Dogs"),
        },
        ID: openapiexample.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "provident",
            "distinctio",
            "quibusdam",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: openapiexample.Int64(544883),
                Name: openapiexample.String("Ben Mueller"),
            },
            shared.Tag{
                ID: openapiexample.Int64(437587),
                Name: openapiexample.String("Raquel Bednar"),
            },
            shared.Tag{
                ID: openapiexample.Int64(383441),
                Name: openapiexample.String("Alexandra Schulist"),
            },
            shared.Tag{
                ID: openapiexample.Int64(568045),
                Name: openapiexample.String("Mrs. Sophie Smith MD"),
            },
        },
    }, operations.AddPetFormSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->