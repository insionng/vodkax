+++
title = "Cookies"
description = "Handling cookie in Vodka"
[menu.side]
  name = "Cookies"
  parent = "guide"
  weight = 6
+++

## Cookies

Cookie is a small piece of data sent from a website and stored in the user's web
browser while the user is browsing. Every time the user loads the website, the browser
sends the cookie back to the server to notify the user's previous activity.
Cookies were designed to be a reliable mechanism for websites to remember stateful
information (such as items added in the shopping cart in an online store) or to
record the user's browsing activity (including clicking particular buttons, logging
in, or recording which pages were visited in the past). Cookies can also store
passwords and form content a user has previously entered, such as a credit card
number or an address.

### Cookie Attributes

Attribute | Optional
:--- | :---
`Name` | No
`Value` | No
`Path` | Yes
`Domain` | Yes
`Expires` | Yes
`Secure` | Yes
`HTTPOnly` | Yes

### Create a Cookie

```go
func writeCookie(c vodka.Context) error {
	cookie := new(vodka.Cookie)
	cookie.SetName("username")
	cookie.SetValue("jon")
	cookie.SetExpires(time.Now().Add(24 * time.Hour))
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}
```

- Cookie is created using `new(vodka.Cookie)`.
- Attributes for the cookie are set using `Setter` functions.
- Finally `c.SetCookie(cookies)` adds a `Set-Cookie` header in HTTP response.

### Read a Cookie

```go
func readCookie(c vodka.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name())
	fmt.Println(cookie.Value())
	return c.String(http.StatusOK, "read a cookie")
}
```

- Cookie is read by name using `c.Cookie("username")` from the HTTP request.
- Cookie attributes are accessed using `Getter` function.

### Read all Cookies

```go
func readAllCookies(c vodka.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name())
		fmt.Println(cookie.Value())
	}
	return c.String(http.StatusOK, "read all cookie")
}
```
