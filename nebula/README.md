# Nebula

A Templ component library for the Sonr DWN (Decentralized Web Node) client.

## Overview

### 3rd Party

- [htmx](https://htmx.org/)
- [tailwindcss](https://tailwindcss.com/)
- [templ](https://templ.dev/)
- [alpinejs](https://alpinejs.dev/)

### Components

- Navbar
- Footer
- Login
- Register
- Profile
- Authorize

## Usage

```go
package main

import (
	"github.com/onsonr/sonr/nebula"
	"github.com/onsonr/sonr/nebula/components"
	"github.com/onsonr/sonr/nebula/pages"
)

func main() {
	e := echo.New()
	e.GET("/", pages.Home)
	e.GET("/login", pages.Login)
	e.GET("/register", pages.Register)
	e.GET("/profile", pages.Profile)
	e.GET("/authorize", pages.Authorize)
	e.GET("/components", components.Home)
	e.Logger.Fatal(e.Start(":1323"))
}
```

## License

[MIT](LICENSE)
