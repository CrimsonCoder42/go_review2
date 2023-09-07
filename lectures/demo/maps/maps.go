package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 12
	shoppingList["milk"] = 10
	shoppingList["cheese"] += 1


	shoppingList["eggs"] += 1

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("No!")
	} else {
		fmt.Println("Yes, need", cereal)
	}

	totalItems := 0
	for _, quantity := range shoppingList {
		totalItems += quantity
	}

	fmt.Println("Total items:", totalItems)

}
