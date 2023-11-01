package main

import "net/http"

func responseHealth(w http.ResponseWriter, h *http.Request) {
	jsonmarshalhelper(w, 200, struct{}{})
}
