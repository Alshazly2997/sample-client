package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello World!"
	fmt.Fprintln(w, greeting)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("username")
		password := r.FormValue("password")

		jsonData := map[string]string{"name": name, "password": password}
		jsonValue, _ := json.Marshal(jsonData)

		response, _ := http.Post("http://localhost:8080/auth", "application/json", bytes.NewBuffer(jsonValue))

		tokenString, _ := ioutil.ReadAll(response.Body)

		http.SetCookie(w, &http.Cookie{Name: "jwt_token", Value: string(tokenString)})
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		template, _ := template.ParseFiles("login.html")
		template.Execute(w, nil)
	}

}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	req, _ := http.NewRequest("GET", "http://localhost:8080/query", nil)
	req.Header.Add("authorization", "Bearar "+cookie.Value)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The http request failed %v", err)
	} else {
		var data User
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, &data)
		fmt.Println(data.Name)
		template, _ := template.ParseFiles("dashboard.html")
		template.Execute(w, data)
	}

}
