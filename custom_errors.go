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

func genericError(rw http.ResponseWriter, customError outJSON) (err error) {
	cj, err := json.Marshal(customError)
	if err != nil {
		return
	}
	if err = json.Indent(&out, cj, "", "\t"); err != nil {
		return
	}
	if _, err = out.WriteTo(rw); err != nil {
		return
	}
	return nil
}

func forgivenJSONPage(rw http.ResponseWriter, r *http.Request) {
	if err := genericError(rw, outJSON{Err: "The page you have requested is forbidden", Status: 403}); err != nil {
		panic(err) //Nunca debe pasar
	}
}

func notFoundJSONPage(rw http.ResponseWriter, r *http.Request) {
	if err := genericError(rw, outJSON{Err: "This entrypoint not exist", Status: 404}); err != nil {
		panic(err) //Nunca debe pasar
	}
}

func badRequestJSONPage(rw http.ResponseWriter, r *http.Request) {
	if err := genericError(rw, outJSON{Err: "bad syntax or bad request", Status: 400}); err != nil {
		panic(err) //Nunca debe pasar
	}
}

func notValidJSONPage(rw http.ResponseWriter, r *http.Request) {
	if err := genericError(rw, outJSON{Err: "This element is not validable", Status: 233}); err != nil {
		panic(err) //Nunca debe pasar
	}
}
