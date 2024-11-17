package main

import (
	"fmt"
	"os"
	"text/template"
)

/*
* Create a data structure to pass to a template which
* contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
 */
type Menu struct {
    Breakfast []MenuItem
    Lunch     []MenuItem 
    Dinner    []MenuItem
}


type MenuItem struct {
    Name        string
    Description string 
    Price       float64
}

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("menu.txt"))
}
func main() {
    
    menu := Menu{
        Breakfast: []MenuItem{
            {"Pancakes", "Fluffy pancakes with syrup", 5.99},
            {"Omelette", "Cheese and vegetable omelette", 6.99},
        },
        Lunch: []MenuItem{
            {"Burger", "Juicy beef burger with fries", 10.99},
            {"Caesar Salad", "Fresh romaine lettuce with Caesar dressing", 8.99},
        },
        Dinner: []MenuItem{
            {"Steak", "Grilled steak with mashed potatoes", 15.99},
            {"Spaghetti", "Spaghetti with marinara sauce and garlic bread", 12.99},
        },
    }
	fmt.Println(menu)
	tpl.Execute(os.Stdout,nil)
}