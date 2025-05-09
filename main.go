package main

import (
	"html/template"
	"net/http"
)

func main(){
	fs:= http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/",fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp:= template.Must(template.ParseFiles("templates/index.html"))
		tmp.Execute(w,nil)
	})

	println("Server started and Listening at PORT 8080")
	http.ListenAndServe(":8080",nil)
}