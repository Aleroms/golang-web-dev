package main

import (
	"log"
	"net/http"
	"text/template"
)

/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use only this:

``` Go
	fs := http.FileServer(http.Dir("public"))
```

Hint: look to see what type FileServer returns, then think it through.
*/
func main(){
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics/", fs)
	http.HandleFunc("/",home)
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
func home(w http.ResponseWriter, r *http.Request){
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		http.Error(w, "Internal Server Error: Unable to load template", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tpl.ExecuteTemplate(w,"index.gohtml",nil)
}