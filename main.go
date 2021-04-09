package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {
	// p1 := person{
	// 	First: "Jenny",
	// }

	// p2 := person{
	// 	First: "James",
	// }

	// xp := []person{p1, p2}
	// bs err := json.Marshal(xp)
	// if err != il {
	// 	log.Panic(err)
	// }

	// fmt.Println("json", string(bs))

	// xp2 := []person{}
	// er = json.Unmarshal(bs, &xp2)
	// if err != il {
	// 	log.Panic(err)
	// }

	// fmt.Println("go struct", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("bad data", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p1 person
	err := json.NewDecoder(r.Body).Decode(&p1)
	if err != nil {
		log.Println("decode bad data", err)
	}

	log.Println(p1)

}
