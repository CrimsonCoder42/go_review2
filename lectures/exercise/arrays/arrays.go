//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import (
	"fmt"
)

type Products struct {
	Name  string
	Price float32
}

func priceTotal(items [4]Products) float32 {
	var total float32
	for i := 0; i < len(items); i++ {
		total += items[i].Price
	}
	return total

}

func main() {
	list := [4]Products{
		{"Milk", 4.99},
		{"bread", 3.99},
		{"eggs", 2.99},
	}

	fmt.Println(list[2])
	fmt.Println(len(list))
	fmt.Println(priceTotal(list))

	list[3] = Products{"Fish", 29.99}
	fmt.Println(list[len(list) - 1])
	fmt.Println(len(list))
	fmt.Println(priceTotal(list))

}
