package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type outJSON struct {
	Err    string `json:"error"`
	Status int    `json:"status"`
}

var out bytes.Buffer

func genericError(rw http.ResponseWriter, customError outJSON) error {
	cj, err := json.Marshal(customError)
	if err != nil {
		return err
	}
	json.Indent(&out, cj, "", "\t")
	out.WriteTo(rw)
	return nil
}

func forgivenJsonPage(rw http.ResponseWriter, r *http.Request) {
	genericError(rw, outJSON{Err: "The page you have requested is forbidden", Status: 403})
}

func notFoundJsonPage(rw http.ResponseWriter, r *http.Request) {
	genericError(rw, outJSON{Err: "This entrypoint not exist", Status: 404})
}

func BadRequestJsonPage(rw http.ResponseWriter, r *http.Request) {
	genericError(rw, outJSON{Err: "bad syntax or bad request", Status: 400})
}

func notValidJsonPage(rw http.ResponseWriter, r *http.Request) {
	genericError(rw, outJSON{Err: "This element is not validable", Status: 233})
}
