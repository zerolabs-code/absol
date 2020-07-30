# Absol

A simple HTTP request handler that works without variable parameters in request paths.

### Install

```
go get github.com/zerolabs-code/absol
```

### Usage

```go
package main

import (
    "log"
    "net/http"

    "github.com/zerolabs-code/absol"
)

func main() {
    mux := absol.NewMux()
    mux.GET("/users", api.GetUsers)
    mux.POST("/users", api.CreateUser)
    mux.GET("/user", api.GetUser)
    mux.DELETE("/user", api.DeleteUser)
    log.Fatalln(http.ListenAndServe(":8080", mux))
}
```
