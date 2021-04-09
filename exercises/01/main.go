package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	First string
}

func main() {
	http.HandleFunc("/encode", httpEncode)
	http.HandleFunc("/decode", httpDecode)
	http.ListenAndServe("localhost:8080", nil)
}

func httpEncode(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "Jennay",
	}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
	log.Println(p1)
}

func httpDecode(w http.ResponseWriter, r *http.Request) {
	var p1 Person

	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println(err)
	}
	log.Println(p1)
}
