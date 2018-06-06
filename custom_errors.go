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

func forgivenJsonPage(rw http.ResponseWriter, r *http.Request) {
	customJSON := outJSON{Err: "The page you have requested is forbidden", Status: 403}
	cj, _ := json.Marshal(customJSON)
	json.Indent(&out, cj, "", "\t")
	out.WriteTo(rw)
}

func notFoundJsonPage(rw http.ResponseWriter, r *http.Request) {
	customJSON := outJSON{Err: "This entrypoint not exist", Status: 404}
	cj, _ := json.Marshal(customJSON)
	json.Indent(&out, cj, "", "\t")
	out.WriteTo(rw)
}

func BadRequestJsonPage(rw http.ResponseWriter, r *http.Request) {
	customJSON := outJSON{Err: "bad syntax or bad request", Status: 400}
	cj, _ := json.Marshal(customJSON)
	json.Indent(&out, cj, "", "\t")
	out.WriteTo(rw)
}
