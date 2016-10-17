+++
title = "MethodOverride Middleware"
description = "Method override middleware for Vodka"
[menu.side]
  name = "MethodOverride"
  parent = "middleware"
  weight = 5
+++

## MethodOverride Middleware

MethodOverride middleware checks for the overridden method from the request and
uses it instead of the original method.

For security reasons, only `POST` method can be overridden.

*Usage*

`e.Pre(middleware.MethodOverride())`

### Custom Configuration

*Usage*

```go
e := vodka.New()
e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
  Getter: middleware.MethodFromForm("_method"),
}))
```

### Configuration

```go
MethodOverrideConfig struct {
  // Skipper defines a function to skip middleware.
  Skipper Skipper

  // Getter is a function that gets overridden method from the request.
  // Optional. Default values MethodFromHeader(vodka.HeaderXHTTPMethodOverride).
  Getter MethodOverrideGetter
}
```

*Default Configuration*

```go
DefaultMethodOverrideConfig = MethodOverrideConfig{
  Skipper: defaultSkipper,
  Getter:  MethodFromHeader(vodka.HeaderXHTTPMethodOverride),
}
```
