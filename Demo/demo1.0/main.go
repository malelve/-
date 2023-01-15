package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServe := http.FileServer(http.Dir("CBA"))
	http.Handle("/", fileServe)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Listen on Port 8080", "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Fatal(err)
	}
	userName := req.FormValue("username")
	fmt.Fprintf(w, "Hello %s", userName)
}
