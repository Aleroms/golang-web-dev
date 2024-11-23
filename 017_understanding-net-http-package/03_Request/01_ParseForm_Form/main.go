package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog
	fmt.Fprintln(os.Stdout,"Listening on http://localhost:8080")
	http.ListenAndServe(":8080", d)
}
