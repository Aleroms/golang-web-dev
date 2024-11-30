package main

import (
	"log"
	"net/http"
	"text/template"
)

/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use:

``` Go
	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler
```

Constraint: you are not allowed to change the route being used for images in the template file

*/
var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}
func main(){
	http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
func home(w http.ResponseWriter, r *http.Request){
	err := tpl.ExecuteTemplate(w,"index.gohtml", nil)
	if err != nil {
		http.Error(w, "Cannot find template", http.StatusInternalServerError)
	}
}