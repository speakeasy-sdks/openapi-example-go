# User

## Overview

Operations about user

### Available Operations

* [CreateUserForm](#createuserform) - Create user
* [CreateUserJSON](#createuserjson) - Create user
* [CreateUserRaw](#createuserraw) - Create user
* [CreateUsersWithListInput](#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](#deleteuser) - Delete user
* [GetUserByName](#getuserbyname) - Get user by user name
* [LoginUser](#loginuser) - Logs user into the system
* [LogoutUser](#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](#updateuserform) - Update user
* [UpdateUserJSON](#updateuserjson) - Update user
* [UpdateUserRaw](#updateuserraw) - Update user

## CreateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.CreateUserForm(ctx, shared.User{
        Email: openapiexample.String("john@email.com"),
        FirstName: openapiexample.String("John"),
        ID: openapiexample.Int64(10),
        LastName: openapiexample.String("James"),
        Password: openapiexample.String("12345"),
        Phone: openapiexample.String("12345"),
        UserStatus: openapiexample.Int(1),
        Username: openapiexample.String("theUser"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.CreateUserJSON(ctx, shared.User{
        Email: openapiexample.String("john@email.com"),
        FirstName: openapiexample.String("John"),
        ID: openapiexample.Int64(10),
        LastName: openapiexample.String("James"),
        Password: openapiexample.String("12345"),
        Phone: openapiexample.String("12345"),
        UserStatus: openapiexample.Int(1),
        Username: openapiexample.String("theUser"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.CreateUserRaw(ctx, []byte("quasi"))
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateUsersWithListInput

Creates list of users with given input array

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.CreateUsersWithListInput(ctx, []shared.User{
        shared.User{
            Email: openapiexample.String("john@email.com"),
            FirstName: openapiexample.String("John"),
            ID: openapiexample.Int64(10),
            LastName: openapiexample.String("James"),
            Password: openapiexample.String("12345"),
            Phone: openapiexample.String("12345"),
            UserStatus: openapiexample.Int(1),
            Username: openapiexample.String("theUser"),
        },
        shared.User{
            Email: openapiexample.String("john@email.com"),
            FirstName: openapiexample.String("John"),
            ID: openapiexample.Int64(10),
            LastName: openapiexample.String("James"),
            Password: openapiexample.String("12345"),
            Phone: openapiexample.String("12345"),
            UserStatus: openapiexample.Int(1),
            Username: openapiexample.String("theUser"),
        },
        shared.User{
            Email: openapiexample.String("john@email.com"),
            FirstName: openapiexample.String("John"),
            ID: openapiexample.Int64(10),
            LastName: openapiexample.String("James"),
            Password: openapiexample.String("12345"),
            Phone: openapiexample.String("12345"),
            UserStatus: openapiexample.Int(1),
            Username: openapiexample.String("theUser"),
        },
        shared.User{
            Email: openapiexample.String("john@email.com"),
            FirstName: openapiexample.String("John"),
            ID: openapiexample.Int64(10),
            LastName: openapiexample.String("James"),
            Password: openapiexample.String("12345"),
            Phone: openapiexample.String("12345"),
            UserStatus: openapiexample.Int(1),
            Username: openapiexample.String("theUser"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

## DeleteUser

This can only be done by the logged in user.

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
    res, err := s.User.DeleteUser(ctx, operations.DeleteUserRequest{
        Username: "Weston_Thiel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetUserByName

Get user by user name

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
    res, err := s.User.GetUserByName(ctx, operations.GetUserByNameRequest{
        Username: "Whitney.Bednar",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

## LoginUser

Logs user into the system

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
    res, err := s.User.LoginUser(ctx, operations.LoginUserRequest{
        Password: openapiexample.String("cum"),
        Username: openapiexample.String("Aiyana.Batz"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoginUser200ApplicationJSONString != nil {
        // handle response
    }
}
```

## LogoutUser

Logs out current logged in user session

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/operations"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserForm(ctx, operations.UpdateUserFormRequest{
        User: &shared.User{
            Email: openapiexample.String("john@email.com"),
            FirstName: openapiexample.String("John"),
            ID: openapiexample.Int64(10),
            LastName: openapiexample.String("James"),
            Password: openapiexample.String("12345"),
            Phone: openapiexample.String("12345"),
            UserStatus: openapiexample.Int(1),
            Username: openapiexample.String("theUser"),
        },
        Username: "Wilfrid.Carter",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/operations"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserJSON(ctx, operations.UpdateUserJSONRequest{
        User: &shared.User{
            Email: openapiexample.String("john@email.com"),
            FirstName: openapiexample.String("John"),
            ID: openapiexample.Int64(10),
            LastName: openapiexample.String("James"),
            Password: openapiexample.String("12345"),
            Phone: openapiexample.String("12345"),
            UserStatus: openapiexample.Int(1),
            Username: openapiexample.String("theUser"),
        },
        Username: "Jayden.Carter88",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi-example"
	"openapi-example/pkg/models/operations"
	"openapi-example/pkg/models/shared"
)

func main() {
    s := openapiexample.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserRaw(ctx, operations.UpdateUserRawRequest{
        RequestBody: []byte("commodi"),
        Username: "Terrill69",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
