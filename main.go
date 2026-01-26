package main

import (
	"fmt"
	"net/http"
)

var token string

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func main() {

	http.HandleFunc("/", greetingHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard", dashboardHandler)

	err := http.ListenAndServe("localhost:3000", http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
