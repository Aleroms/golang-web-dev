package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseGlob("starting-files/templates/*"))
}
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		err := tpl.ExecuteTemplate(w,"index.gohtml", nil)
		if err != nil {
			http.Error(w,"Cannot open template", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		err := tpl.ExecuteTemplate(w,"about.gohtml", nil)
		if err != nil {
			http.Error(w,"Cannot open template", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
		err := tpl.ExecuteTemplate(w,"contact.gohtml", nil)
		if err != nil {
			http.Error(w,"Cannot open template", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/apply", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost {
			err := tpl.ExecuteTemplate(w,"applyProcess.gohtml", nil)
			if err != nil {
				http.Error(w,"Cannot open template", http.StatusInternalServerError)
			}
		} else {
			err := tpl.ExecuteTemplate(w,"apply.gohtml", nil)
			if err != nil {
				http.Error(w,"Cannot open template", http.StatusInternalServerError)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}