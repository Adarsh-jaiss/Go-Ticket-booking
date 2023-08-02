package main

import (
	"fmt"
	"strings"
	"time"
)

// Define a global variable 'bookings' to store user data.
var bookings = make([]userData, 0)

// Define a struct 'userData' to store user information.
type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

// Function to greet users and show the conference details.
func greetUsers(conferenceName string, conferenceTicket int, remainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and out of which %v are currently available \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here!")
}

// Function to extract first names from the list of user data.
func GetfirstNames(bookings []userData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// Function to validate user input (first name, last name, email, and ticket count).
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName && isValidEmail && isValidTicketNumber
}

// Function to get user input for booking tickets
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of Tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

// Function to book tickets and update the remaining ticket count.
func bookTicket(remainingTickets *uint, userTickets uint, bookings *[]userData, firstName string, lastName string, email string) {
	*remainingTickets = *remainingTickets - userTickets

	// Create a new 'userData' struct to store user information.
	var userdata = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	// Add the new user data to the 'bookings' slice.
	*bookings = append(*bookings, userdata)

	// Print the list of booking data.
	fmt.Printf("list of booking list is %v\n", bookings)

	// Display a confirmation message for the user.
	fmt.Printf("Thank you %v %v for booking %v tickets. You will soon receive a confirmation mail at your email address %v\n", firstName, lastName, userTickets, email)

	// Display the remaining ticket count.
	fmt.Printf("%v tickets are still remaining for the go conference \n", *remainingTickets)
}

// Function to simulate sending tickets via email (using a 10-second sleep).
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var Ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("sending tickets:\n %v\n to email address %v\n", Ticket, email)
	fmt.Println("#####################")
}

func main() {
	var conferenceName string = "Go conference"
	var conferenceTickets uint = 50
	var remainingTickets uint = 50

	// Greet the users and show conference details.
	greetUsers(conferenceName, int(conferenceTickets), int(remainingTickets))

	// Continue taking user input until all tickets are sold out.
	for {
		// Get user input for booking tickets.
		firstName, lastName, email, userTickets := getUserInput()

		// Validate the user input.
		isValid := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// If the user input is valid, book the tickets and send a confirmation email.
		if isValid {
			bookTicket(&remainingTickets, userTickets, &bookings, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			// Get and print the list of first names from the booked users.
			firstNames := GetfirstNames(bookings)
			fmt.Printf("These are the first names of the bookings %v\n", firstNames)

			// If all tickets are sold out, break the loop.
			if remainingTickets == 0 {
				fmt.Println("Tickets are sold out!")
				break
			}
		} else {
			// If the user input is invalid, display error messages accordingly.
			if len(firstName) < 2 || len(lastName) < 2 {
				fmt.Println("Error: Firstname or lastname you entered is too short")
			}
			if !strings.Contains(email, "@") {
				fmt.Println("Error: Email address is invalid")
			}
			if userTickets > remainingTickets {
				fmt.Println("Error: The amount of ticket count you entered is more than the remaining slots")
			}
			fmt.Printf("Kindly fix the errors and try again!\n")
		}
	}
}
