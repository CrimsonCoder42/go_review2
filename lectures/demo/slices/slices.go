package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("----", title, "----")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(i, element)
	}
}

func main() {
	route := []string{"Grocery", "Deparment Store", "Salon"}
	printSlice("route 1", route)

	route = append(route, "School")
	printSlice("route 2", route)

	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printSlice("route 3", route)



}
