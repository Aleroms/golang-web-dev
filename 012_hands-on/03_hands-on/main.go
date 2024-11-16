package main

import (
	"log"
	"os"
	"text/template"
)

/* Create a data structure to pass to a template which
* contains information about California hotels including Name, Address, City, Zip, Region
* region can be: Southern, Central, Northern
* can hold an unlimited number of hotels
 */
type Hotel struct {
	Name, Address, City, Zip string
	Rating float32
}

type CaliforniaHotels struct {
	Region string
	Hotels []Hotel
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("californiaHotels.txt"))
}
func main(){
	ch := []CaliforniaHotels {
		{
			Region: "Northern California",
			Hotels: []Hotel{
				{"Golden Gate Inn", "123 Lombard Street", "San Francisco", "94133", 4.5},
				{"Redwood Retreat", "456 Redwood Drive", "Eureka", "95501", 4.2},
			},
		},
		{
			Region: "Central California",
			Hotels: []Hotel{
				{"Coastal Breeze Hotel", "789 Pacific Coast Highway", "Monterey", "93940", 4.8},
				{"Vineyard View Inn", "321 Napa Valley Way", "Napa", "94558", 4.7},
			},
		},
		{
			Region: "Southern California",
			Hotels: []Hotel{
				{"Sunset Boulevard Hotel", "654 Hollywood Blvd", "Los Angeles", "90028", 4.9},
				{"Palm Springs Oasis", "987 Palm Canyon Drive", "Palm Springs", "92262", 4.6},
			},
		},
		
	}
	
	err := tpl.Execute(os.Stdout, ch)
	if err != nil {
		log.Fatal(err)
	}
}