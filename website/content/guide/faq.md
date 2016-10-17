+++
title = "FAQ"
description = "Frequently asked questions in Vodka"
[menu.side]
  name = "FAQ"
  parent = "guide"
  weight = 20
+++

## FAQ

Q: **How to retrieve `*http.Request` and `http.ResponseWriter` from `vodka.Context`?**

- `http.Request` > `c.Request().(*standard.Request).Request`
- `http.ResponseWriter` > `c.Response()`

>  Standard engine only

Q: **How to use standard handler `func(http.ResponseWriter, *http.Request)` with Vodka?**

```go
func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Handler!")
}

func main() {
	e := vodka.New()
	e.GET("/", standard.WrapHandler(http.HandlerFunc(handler)))
	e.Run(standard.New(":1323"))
}
```

Q: **How to use fasthttp handler `func(fasthttp.RequestCtx)` with Vodka?**

```go
func handler(c *fh.RequestCtx) {
	io.WriteString(c, "Handler!")
}

func main() {
	e := vodka.New()
	e.GET("/", fasthttp.WrapHandler(handler))
	e.Run(fasthttp.New(":1323"))
}
```

Q: **How to use standard middleware `func(http.Handler) http.Handler` with Vodka?**

```go
func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Middleware!")
		h.ServeHTTP(w, r)
	})
}

func main() {
	e := vodka.New()
	e.Use(standard.WrapMiddleware(middleware))
	e.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.Run(standard.New(":1323"))
}
```

Q: **How to use fasthttp middleware `func(http.Handler) http.Handler` with Vodka?**

```go
func middleware(h fh.RequestHandler) fh.RequestHandler {
	return func(ctx *fh.RequestCtx) {
		println("Middleware!")
		h(ctx)
	}
}

func main() {
	e := vodka.New()
	e.Use(fasthttp.WrapMiddleware(middleware))
	e.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.Run(fasthttp.New(":1323"))
}
```

<!-- ### Q: How to run Vodka on specific IP and port?

```go
``` -->
