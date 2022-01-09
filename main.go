package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, " tickets and ", remainingTickets, "are still available")
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %q tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your booking tickets from here")

	for {
		var name string
		var userTickets uint
		var email string

		fmt.Println("Enter your name please")
		fmt.Scan(&name)
		fmt.Println("Enter your email")
		fmt.Scan(&email)
		fmt.Println("Number of tickets")
		fmt.Scan(&userTickets)
		isValid := len(name) > 2 && userTickets > 0
		isValidEmail := strings.Contains(email, "@")
		isValidNumber := userTickets > 0 && userTickets <= remainingTickets
		if isValid && isValidEmail {

			if remainingTickets >= userTickets {

				remainingTickets = remainingTickets - userTickets
				fmt.Printf("User %v booked %v tickets.\n", name, userTickets)
				fmt.Printf("%v tickets are remaining \n", remainingTickets)
				if remainingTickets == 0 {
					fmt.Printf("Our conference is Book out. Come next year \n")
					break
				}
			} else {
				fmt.Printf("You can not buy more tickets than  %v available. \n", remainingTickets)

			}

		} else {
			// validation messages
			if !isValid {
				fmt.Println("Invalid name of User")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email of User")
			}
			if !isValidNumber {
				fmt.Println("Invalid Ticket of User")
			}

		}

	}

}
