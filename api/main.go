package main

import (
	"log"
	"net/http"

	bbql "dmh2000.xyz/bb-api/baseballGQL"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
        http.Handle("/query", &relay.Handler{Schema: bbql.RootSchema()})
        log.Fatal(http.ListenAndServe(":8080", nil))
}
