package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("We have total of", conferenceTickets, " tickets and ", remainingTickets, "are still available")
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %q tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your booking tickets from here")
	var name string
	var userTickets uint
	var email string
	fmt.Println("Enter your name please")
	fmt.Scan(&name)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Number of tickets")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("User %v booked %v tickets.\n", name, userTickets)
	fmt.Printf("%v tickets are remaining", remainingTickets)

}
