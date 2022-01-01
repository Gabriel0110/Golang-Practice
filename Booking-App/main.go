package main

import (
	"fmt"
	"sync"
	"time"
	//utilities "github.com/Gabriel0110/go-utilities"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50 // set the type to 'uint' to ensure it cannot be a negative value
//var bookings = make([]map[string]string, 0) // when making a slice of maps, you need to give it a starter size even though slices are dynamic
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//var s = []int{4}

	// if utilities.ContainsInt(s, 4) {
	//  fmt.Println("INT FOUND")
	// }

	greetUsers()

	for {

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year!")
			break
		}

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)                                              // sets the number of goroutines to wait for (increases the counter by the provided number)
			go sendTicket(userTickets, firstName, lastName, email) // 'go' starts a goroutine (another thread) - that's it
			printFirstNames()

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered does not contain the @ symbol.")
			}

			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
		}
	}
	wg.Wait() // tells the application to wait for all threads to complete before it can end (until the WaitGroup counter is 0)
}

func greetUsers() {
	fmt.Println("\nWelcome to the", conferenceName, "booking application!")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend.")
	fmt.Println()
}

func printFirstNames() {
	var firstNames []string
	for _, booking := range bookings { // _ is used to identify unused variables (like in python)
		//var names = strings.Fields(booking) // strings.Fields() splits a string with white space as separator, like Python's split(' ')
		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Println("These are all of our bookings:")
	for _, booking := range firstNames {
		fmt.Printf("- %v\n", booking)
	}
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of desired tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) // Sprintf returns the resulting string to the variable
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("################")

	wg.Done() // removes the thread from the WaitGroup to decrement the counter
}
