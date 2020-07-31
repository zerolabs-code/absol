# Absol

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![License: MIT](https://img.shields.io/badge/Conventional%20Commits-1.0.0-green.svg)](https://www.conventionalcommits.org)

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

    // users endpoints
    mux.Get("/users", api.GetUsers)
    mux.Post("/users", api.CreateUser)

    // user endpoints
    mux.Get("/user", api.GetUser)
    mux.Delete("/user", api.DeleteUser)

    log.Fatalln(http.ListenAndServe(":8080", mux))
}
```
