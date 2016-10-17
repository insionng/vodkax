+++
title = "Error Handling"
description = "Error handling in Vodka"
[menu.side]
  name = "Error Handling"
  parent = "guide"
  weight = 8
+++

## Error Handling

Vodka advocates centralized HTTP error handling by returning error from middleware
or handlers.

- Log errors from a unified location
- Send customized HTTP responses

For example, when basic auth middleware finds invalid credentials it returns
`401 - Unauthorized` error, aborting the current HTTP request.

```go
package main

import (
	"net/http"

	"github.com/insionng/vodka"
)

func main() {
	e := vodka.New()
	e.Use(func(c vodka.Context) error {
		// Extract the credentials from HTTP request header and perform a security
		// check

		// For invalid credentials
		return vodka.NewHTTPError(http.StatusUnauthorized)
	})
	e.GET("/welcome", welcome)
	e.Run(":1323")
}

func welcome(c vodka.Context) error {
	return c.String(http.StatusOK, "Welcome!")
}
```

See how [HTTPErrorHandler](/guide/customization#http-error-handler) handles it.
