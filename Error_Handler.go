package main

import "net/http"

func responseError400(w http.ResponseWriter, h *http.Request) {
	errorhandleraction(w, 400, "Something went wrong with the request!")
}
