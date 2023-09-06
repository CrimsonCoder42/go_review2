package main

import "fmt"

type Room struct {
	name string
	cleaned bool 
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		if !rooms[i].cleaned {
			fmt.Println("The room", rooms[i].name, "is not clean")
		} else {
			fmt.Println("The room", rooms[i].name, "is clean")
		}
	}
}

func main() {
	rooms := [4]Room{
		{"kitchen", true},
		{"bathroom", false},
		{"bedroom", true},
		{"living room", false},
	}

	checkCleanliness(rooms)

 
}
