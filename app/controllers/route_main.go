package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/template/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, nil)
}
