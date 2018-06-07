package main

import (
	"os"
	"log"
	"net/http"
	"html/template"
)

type People struct {
	Title string
	Footer string
	Id string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	containerId, _ := os.Hostname()
	p := People{ Title: "Lets GO", Footer: "Powered by Docker, Go and Alpine Linux", Id: containerId }
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", Handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal( http.ListenAndServe(":8080", nil) )
}
