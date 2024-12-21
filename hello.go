package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		home().Render(r.Context(), w)
	})

	type MySignals struct {
		Name string
	}

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		var signals MySignals
		datastar.ReadSignals(r, &signals)
		sse := datastar.NewSSE(w, r)
		sse.MergeFragmentTempl(hello(signals.Name))
	})

	http.ListenAndServe(":3000", r)
}
