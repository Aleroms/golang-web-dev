package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
	fmt.Printf("mode %v parsename %v root %v tree %v",tpl.Mode, tpl.ParseName, tpl.Root, tpl.Tree)

}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
