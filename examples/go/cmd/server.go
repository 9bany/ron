package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("some"))
	})

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok api"))
	})

	fmt.Println("server start at port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err.Error())

	}

}

