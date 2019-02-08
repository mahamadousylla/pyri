package main

import (
	"net/http"

	"github.com/go-zoo/bone"
)

func main() {
	mux := bone.New()

	mux.GetFunc("/", nil)
	http.ListenAndServe(":8080", mux)
}
