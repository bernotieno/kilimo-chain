package main

import (
	"fmt"
	"net/http"
	"text/template"

	"kilimo-chain/asfuncss"
)

func SignUphandler(w http.ResponseWriter) {
	tmp, _ := template.ParseFiles("index.html")

	tmp.Execute(w, nil)
}

func Loginpageload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmp, _ := template.ParseFiles("login.html")
	tmp.Execute(w, nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SignUphandler(w)
	} else if r.URL.Path == "/reg" {
		asfuncss.Reg(w, r)
	} else if r.URL.Path == "/log" {
		Loginpageload(w, r)
	} else if r.URL.Path == "/login" {
		asfuncss.Login(w, r)
	} else if r.URL.Path == "/about" {
		asfuncss.AboutUs(w, r)
	}
}

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", router)
	fmt.Println("RUNNING SERVER")
	http.ListenAndServe("localhost:8078", mux)
}
