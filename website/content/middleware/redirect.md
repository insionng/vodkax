+++
title = "Redirect Middleware"
description = "Redirect middleware for Vodka"
[menu.side]
  name = "Redirect"
  parent = "middleware"
  weight = 5
+++

## HTTPSRedirect Middleware

HTTPSRedirect middleware redirects http requests to https.
For example, http://insionng.com will be redirected to https://insionng.com.

*Usage*

```go
e := vodka.New()
e.Pre(middleware.HTTPSRedirect())
```

## HTTPSWWWRedirect Middleware

HTTPSWWWRedirect redirects http requests to www https.
For example, http://insionng.com will be redirected to https://www.insionng.com.

*Usage*

```go
e := vodka.New()
e.Pre(middleware.HTTPSWWWRedirect())
```

## HTTPSNonWWWRedirect Middleware

HTTPSNonWWWRedirect redirects http requests to https non www.
For example, http://www.insionng.com will be redirect to https://insionng.com.

*Usage*

```go
e := vodka.New()
e.Pre(HTTPSNonWWWRedirect())
```

## WWWRedirect Middleware

WWWRedirect redirects non www requests to www.

For example, http://insionng.com will be redirected to http://www.insionng.com.

*Usage*

```go
e := vodka.New()
e.Pre(middleware.WWWRedirect())
```

## NonWWWRedirect Middleware

NonWWWRedirect redirects www requests to non www.
For example, http://www.insionng.com will be redirected to http://insionng.com.

*Usage*

```go
e := vodka.New()
e.Pre(middleware.NonWWWRedirect())
```

### Custom Configuration

*Usage*

```go
e := vodka.New()
e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
  Code: http.StatusTemporaryRedirect,
}))
```

Example above will redirect the request HTTP to HTTPS with status code `307 - StatusTemporaryRedirect`.

### Configuration

```go
RedirectConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Status code to be used when redirecting the request.
  // Optional. Default value http.StatusMovedPermanently.
  Code int `json:"code"`
}
```

*Default Configuration*

```go
DefaultRedirectConfig = RedirectConfig{
  Skipper: defaultSkipper,
  Code:    http.StatusMovedPermanently,
}
```
