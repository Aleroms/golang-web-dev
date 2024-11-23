/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/
package main

import (
	"fmt"
	"net/http"
)

// not needed
// type santi int

// func(s santi) ServeHTTP(res http.ResponseWriter, req *http.Request){

// }

func main(){
	// HandleFunc needs to be called before ListenAndServe
	// or else a 404 page not found will occur
	
	http.HandleFunc("/", homepage)
	http.HandleFunc("/dog/", dogpage)
	http.HandleFunc("/me/", mepage)

	http.ListenAndServe(":8080", nil)
}

func homepage(res http.ResponseWriter, req *http.Request){
	fmt.Fprintln(res, "AYO this is da hommie page")
}
func dogpage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "dog dog dog woof!")
}
func mepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Santiago Morales 27 2024")
}