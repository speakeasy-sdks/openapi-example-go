# Store

## Overview

Access to Petstore orders

Find out more about our store
<http://swagger.io>
### Available Operations

* [DeleteOrder](#deleteorder) - Delete purchase order by ID
* [GetInventory](#getinventory) - Returns pet inventories by status
* [GetOrderByID](#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](#placeorderraw) - Place an order for a pet

## DeleteOrder

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/operations"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Store.DeleteOrder(ctx, operations.DeleteOrderRequest{
        OrderID: 662527,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetInventory

Returns a map of status codes to quantities

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/operations"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Store.GetInventory(ctx, operations.GetInventorySecurity{
        APIKey: "YOUR_API_KEY_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetInventory200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetOrderByID

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/operations"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Store.GetOrderByID(ctx, operations.GetOrderByIDRequest{
        OrderID: 820994,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

## PlaceOrderForm

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
	"openapi-example/pkg/types"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderForm(ctx, shared.Order{
        Complete: openapiexample.Bool(false),
        ID: openapiexample.Int64(10),
        PetID: openapiexample.Int64(198772),
        Quantity: openapiexample.Int(7),
        ShipDate: types.MustTimeFromString("2022-11-26T13:23:33.410Z"),
        Status: shared.OrderStatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

## PlaceOrderJSON

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
	"openapi-example/pkg/types"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderJSON(ctx, shared.Order{
        Complete: openapiexample.Bool(false),
        ID: openapiexample.Int64(10),
        PetID: openapiexample.Int64(198772),
        Quantity: openapiexample.Int(7),
        ShipDate: types.MustTimeFromString("2021-04-29T07:12:18.684Z"),
        Status: shared.OrderStatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

## PlaceOrderRaw

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
	"openapi-example/pkg/types"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderRaw(ctx, []byte("laborum"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```
