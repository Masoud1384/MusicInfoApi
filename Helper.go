package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonmarshalhelper(h http.ResponseWriter, code int, data interface{}) {
	mdata, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Something went wrong with casting data :", err)
		h.WriteHeader(500)
		return
	}
	h.Header().Add("Content-Type", "application/json")
	h.WriteHeader(code)
	h.Write(mdata)
}

func errorhandleraction(h http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("something went wrong with the server , status code error: ", code)
	}
	type customerror struct {
		Error string `json:"error"`
	}
	jsonmarshalhelper(h, code, customerror{
		msg,
	})
}
