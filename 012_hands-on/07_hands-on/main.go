package main

import (
	"os"
	"text/template"
)

/*
1. Using the data structure created in the previous folder,
modify it to hold menu information for an unlimited number of restaurants.
*/
 
 
type MenuItem struct {
Name        string
Description string 
Price       float64
}

type Menu struct {
	Breakfast []MenuItem
	Lunch     []MenuItem 
	Dinner    []MenuItem
}
type Restaurant struct {
	Name, Address, City, Zip string
	Rating float32
	Menu Menu
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("menu.txt"))
}
func main() {
    
    restaurants := []Restaurant{
		{
			Name:    "The Breakfast Spot",
			Address: "123 Morning St",
			City:    "Pancake City",
			Zip:     "11111",
			Rating:  4.5,
			Menu: Menu{
				Breakfast: []MenuItem{
					{Name: "French Toast", Description: "Golden-brown French toast with syrup", Price: 7.99},
					{Name: "Scrambled Eggs", Description: "Freshly scrambled eggs with toast", Price: 5.49},
				},
				Lunch: []MenuItem{
					{Name: "Grilled Cheese", Description: "Classic grilled cheese sandwich", Price: 6.99},
				},
				Dinner: []MenuItem{
					{Name: "Meatloaf", Description: "Homestyle meatloaf with gravy", Price: 13.99},
				},
			},
		},
		{
			Name:    "Lunch Lovers",
			Address: "456 Noon Ave",
			City:    "Sandwich Town",
			Zip:     "22222",
			Rating:  4.8,
			Menu: Menu{
				Breakfast: []MenuItem{
					{Name: "Bagel", Description: "Freshly baked bagel with cream cheese", Price: 3.99},
				},
				Lunch: []MenuItem{
					{Name: "BLT Sandwich", Description: "Bacon, lettuce, and tomato on toasted bread", Price: 9.99},
					{Name: "Chicken Salad", Description: "Grilled chicken on a bed of greens", Price: 11.99},
				},
				Dinner: []MenuItem{
					{Name: "Pasta Primavera", Description: "Pasta with fresh vegetables in a light sauce", Price: 14.99},
				},
			},
		},
	}
	tpl.Execute(os.Stdout,restaurants)
}