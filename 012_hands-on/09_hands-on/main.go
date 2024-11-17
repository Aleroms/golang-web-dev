package main

import (
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

/*
1. Parse this CSV file, putting two fields from the contents of the CSV file into a data structure.
2. Parse an html template, then pass the data from step 1 into the CSV template;
   have the html template display the CSV data in a web page.
*/
type DataCsv struct {
	Headers []string
	Data [][]string
}
var tpl *template.Template
func init(){
	tpl = template.Must(template.ParseFiles("table.txt"))
}
func main(){
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header_csv, err := reader.Read()
	if err != nil {
		log.Fatalln(err)
	}
	
	xxcsv, err := reader.ReadAll()
	
	tpl.Execute(os.Stdout, DataCsv{header_csv,xxcsv})
}