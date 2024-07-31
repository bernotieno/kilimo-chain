package main

import (
	"fmt"
	"net/http"
	"os"
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
	} else if r.URL.Path == "/signin" {
		signin(w, r)
	} else if r.URL.Path == "/about" {
		asfuncss.DashboardHandler(w, r)
	} else if r.URL.Path == "/signup" {
		signup(w, r)
	} else if r.URL.Path == "/login" {
		asfuncss.Login(w, r)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmp, _ := template.ParseFiles("Signup.html")
	tmp.Execute(w, nil)
}

func signin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmp, _ := template.ParseFiles("Signin.html")
	tmp.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	// fs := http.FileServer(http.Dir("static"))
	// mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", router)
	mux.HandleFunc("/static/", serveStatic)
	fmt.Println("RUNNING SERVER")
	http.ListenAndServe("localhost:8078", mux)
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	info, err := os.Stat("." + path)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusBadGateway)
		return
	}
	if info.IsDir() {
		http.Error(w, "Permission denied", http.StatusBadGateway)
		return
	}
	http.ServeFile(w, r, "."+path)
}
