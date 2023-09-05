package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 123, false}
	fmt.Println(casey)

	var (
		bill = Passenger{"Bill", 456, false}
		ella = Passenger{"Ella", 789, false}

	)

	fmt.Println(bill, ella)
	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 101112
	heidi.Boarded = false
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	if bill.Boarded {
		fmt.Println(bill.Name, "has boarded")
	}

	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded")
	}


}
