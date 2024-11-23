/*
* Take the previous program and change it so that:
* func main uses http.Handle instead of http.HandleFunc
 */
package main

import (
	"fmt"
	"net/http"
	"text/template"
)
 
 var tpl *template.Template
 
 var fc = template.FuncMap{
	 "loopThreeTimes": LoopThreeTimes,
 }
 
 func init(){
 
	 tpl = template.Must(template.New("test").Funcs(fc).ParseGlob("templates/*"))
 }
 func main(){

	http.Handle("/", http.HandlerFunc(homepage))
	 http.Handle("/dog/", http.HandlerFunc(dogpage))
	 http.Handle("/me/", http.HandlerFunc(mepage))
 
	 http.ListenAndServe(":8080", nil)
 }
 
 func homepage(res http.ResponseWriter, req *http.Request){
	 tpl.ExecuteTemplate(res,"test.html",nil)
 }
 func dogpage(w http.ResponseWriter, r *http.Request){
	 fmt.Println("reached the dog route")
	 data := struct {
		 Saying string
		 count int
	 }{
		 "woof", 3,
	 }
	 tpl.ExecuteTemplate(w,"dog.html",data)
 }
 func mepage(w http.ResponseWriter, r *http.Request){
	 fmt.Fprintln(w, "Santiago Morales 27 2024")
 }
 
 func LoopThreeTimes(s string) (result string){
	 return fmt.Sprintf("%v %v %v",s, s, s)
 }