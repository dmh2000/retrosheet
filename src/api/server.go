package api

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go/relay"
)

func Run() {
	// non graphql for test
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.Handle("/query", &relay.Handler{Schema: RootSchema()})
	log.Fatal(http.ListenAndServe("172.24.46.88:8080", nil))
}
