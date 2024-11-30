package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

/*
# ListenAndServe on port 8080 of localhost

For the default route "/"
Have a func called "foo"
which writes to the response "foo ran"

For the route "/dog/"
Have a func called "dog"
which parses a template called "dog.gohtml"
and writes to the response "<h1>This is from dog</h1>"
and also shows a picture of a dog when the template is executed.

Use "http.ServeFile"
to serve the file "dog.jpeg"
*/
func main() {
	http.HandleFunc("/",home)
	http.HandleFunc("/dog",dog)
	http.HandleFunc("/dog.jpg",dogPic)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/plain")
	io.WriteString(w,"<h1>foo ran</h1>")
}
func dog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","text/html")
	
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		http.Error(w, "Internal Server Error: Unable to load template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}
	tpl.ExecuteTemplate(w,"dog.gohtml",nil)
}
func dogPic(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"dog.jpg")
}