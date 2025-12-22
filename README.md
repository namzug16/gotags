[![GoDoc](https://godoc.org/github.com/namzug16/gotags?status.svg)](http://godoc.org/github.com/namzug16/gotags)
[![Go Report Card](https://goreportcard.com/badge/github.com/namzug16/gotags)](https://goreportcard.com/report/github.com/namzug16/gotags)

```
   __________  _________   ___________
  / ____/ __ \/_  __/   | / ____/ ___/
 / / __/ / / / / / / /| |/ / __ \__ \ 
/ /_/ / /_/ / / / / ___ / /_/ /___/ / 
\____/\____/ /_/ /_/  |_\____//____/  

```

Template-less. Type-Safe. Simple. Familiar. Pure Go.

## Features

* Compose reusable HTML with pure functions
```go
func Card(content ...HTML) HTML {
  return Div(
    X.Class("card bg-base-100 w-96"),
    content,
  )
}
```
* Tags accept `any`: strings are automatically wrapped as text nodes, while other Go values panic immediately so problems surface early instead of silently failing. If you need raw HTML, use `Raw`;
plain strings (or `Text` components) are automatically escaped

* Attributes are accessed through the `X` value to make them immediately recognizable and easy to skim 

* Compose multiple root components with `Fragment` when you need to return more than one component without introducing a wrapper tag or a full `Doc`.
`Fragment` concatenates its child components in order. It is especially useful for HTTP responses (for example with htmx) where a single response must include multiple independent elements
```go
return Fragment(
  H1("Settings"),
  Form(
    Input(X.Type("text"), X.Name("email")),
    Button("Save"),
  ),
)
```
* Conditional composition with `If` (for components) and `X.If` (for attributes)
```go
button := Button(
  X.Class(
    "btn",
    X.If(isPrimary, "btn-primary"),
    X.If(!isPrimary, "btn-secondary"),
  ),
  If(isPrimary, customIcon()),
  "Save changes",
)
```
* List of components via `Range`
* Mutate tags in place with `AddToTag` (Helpful when using htmx with swap-oob)
* Build **custom tags** directly with `NewTagComponent` when you need a tag outside the built-in set
```go
func XMLFeed(content ...HTML) HTML {
  // An XML-like root for a feed export.
  return NewTagComponent("feed", false, content...)
}
```
* Build **custom attributes**
```go
func HxGet(value string) HTML {
  return X.Attr("hx-get", value)
}
```
* Built-in **htmx integration** via the `htmx` package, which provides strongly-typed helpers for all common htmx attributes and headers
```go
import hx "github.com/namzug16/gotags/htmx"

button := Button(
  hx.Post("/save"),
  hx.On("click", "console.log('saving')"),
  "Save",
)
```

## Usage

```shell
go get github.com/namzug16/gotags
```

```go
package main

import (
  "fmt"

  . "github.com/namzug16/gotags"
)

func main() {
  isSignedIn := true
  users := []string{"Federica", "Mateo", "Ciro"}

  page := Doc(
    Head(
      Title("gotags example"),
      Link(X.Rel("stylesheet"), X.Href("/styles.css")),
    ),
    Body(
      Header(
        H1("Team dashboard"),
        Nav(
          Ul(
            Li(A(X.Href("/"), "Home")),
            Li(A(X.Href("/projects"), "Projects")),
            Li(A(X.Href("/profile"), If(isSignedIn, Text("Profile")))),
          ),
        ),
      ),
      Main(
        Section(
          H2("People"),
          Ul(
            Range(users, func(_ int, user string) HTML {
              return Li(
                X.Class("person", X.If(user == "Ciro", "highlight")),
                user,
              )
            }),
          ),
        ),
      ),
    ),
  )

  fmt.Println(page.String())
}
```
