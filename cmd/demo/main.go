package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/namzug16/gotags"
	hx "github.com/namzug16/gotags/htmx"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, homePage().String())
	})

	http.HandleFunc("/get-datetime", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		now := time.Now().UTC().Format(time.RFC3339)
		fmt.Fprint(w,
			Fragment(
				AddToTag(
					dateContainer(now),
					hx.SwapOob("true"),
				),
				successToast(),
			).String(),
		)
	})

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage() HTML {
	return Doc(
		Head(
			Link(X.Rel("stylesheet"), X.Type("text/css"), X.Href("https://cdn.jsdelivr.net/npm/daisyui@5")),
			Script(X.Src("https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4")),
			Script(X.Src("https://cdn.jsdelivr.net/npm/htmx.org@2.0.8/dist/htmx.min.js")),
			Script(X.Src("https://unpkg.com/hyperscript.org@0.9.14")),
		),
		Body(
			X.Class("p-4 flex flex-col gap-4"),
			Button(
				X.Class("btn btn-primary"),
				hx.Get("/get-datetime"),
				hx.Swap("none"),
				"hx GET datetime",
			),
			dateContainer("---"),
		),
	)
}

func dateContainer(date string) HTML {
	return Div(
		X.Class("alert alert-info alert-outline"),
		X.Role("alert"),
		X.Id("date-container"),
		Span(date),
	)
}

func successToast() HTML {
	return Div(
		hx.SwapOob("beforeend:body"),
		Div(
			X.Attr("_", "init wait 3 seconds then transition my opacity to 0 then remove me"),
			X.Class("toast toast-top toast-end"),
			Div(
				X.Class("alert alert-success"),
				X.Role("alert"),
				"Success",
			),
		),
	)
}
