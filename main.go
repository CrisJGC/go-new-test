package main

import (
	"fmt"
	"strings"
)

func main() {
	var storeName = "One Item Store"
	const totalItem = 100
	var stock uint = 100
	totalSells := []string{}

	fmt.Printf("Welcome to %v \n", storeName)
	fmt.Println("The store where you can only buy one item")
	fmt.Printf("Total item %v \n", totalItem)
	fmt.Printf("Remaining stock %v \n", stock)

	for {

		var firstName string
		var email string
		var userItems uint
		//ask info
		fmt.Println("Enter Your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter Your email address: ")
		fmt.Scan(&email)

		fmt.Println("Number of items you need: ")
		fmt.Scan(&userItems)

		if userItems > stock {
			fmt.Printf("We only have %v items, so we can't sell you %v items.\n", stock, userItems)
			continue
		}

		stock = stock - userItems
		totalSells = append(totalSells, firstName)

		// fmt.Printf("Array: %v\n", totalSells)
		// fmt.Printf("First value %v\n", totalSells[0])
		// fmt.Printf("Array type: %T\n", totalSells)
		// fmt.Printf("Array length: %v\n", len(totalSells))

		fmt.Printf("Thank you %v for buying %v items, a confirmation receipt will be sent to %v \n", firstName, userItems, email)
		fmt.Printf("%v items left \n", stock)

		firstNames := []string{}
		for _, nameList := range totalSells {
			var names = strings.Fields(nameList)
			firstNames = append(firstNames, names[0])
		}

		fmt.Println("Here our best buyers: ")
		fmt.Printf("%v\n", firstNames)

		if stock == 0 {
			fmt.Println("All our items are sold")
			break
		}
	}

}
