package main

import (
	"fmt"
	"net/http"
)

func init() {

	http.HandleFunc("/", handlerDefault)

}

func handlerDefault(responseWriter http.ResponseWriter, request *http.Request) {

	fmt.Fprint(responseWriter, "Welcome and will coming soon.")

}
