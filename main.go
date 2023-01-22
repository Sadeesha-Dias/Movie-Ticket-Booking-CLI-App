package main

import "fmt"
import "math/rand"
import "strconv"

const (
	MOVIE1 = "Avatar : The way of water"
	MOVIE2 = "ANT MAN and the WASP : Quantumania"
	MOVIE3 = "Guardians of the Galaxy Vol.3"
	MOVIE4 = "Jhon Wick : Chapter 4"
)

const (
	SCR_TYPE1 = "IMAX 3D"
	SCR_TYPE2 = "3D"	
)

var bookedSeats int
var remainingSeatsIMAX3D int = 50
var remainingSeats3D int = 70

var bookings []string 

func main(){
	var MovieChoise uint

	fmt.Println("------ Welcome to EONMAX Cinema ------")
	fmt.Println("Now Screening")
	fmt.Printf("\t 1 - %v\n \t 2 - %v\n \t 3 - %v\n \t 4 - %v", MOVIE1, MOVIE2, MOVIE3, MOVIE4)
	fmt.Println()
	fmt.Println()
	fmt.Print("==> Please make your selection: ")
	fmt.Scan(&MovieChoise)

	if(MovieChoise == 1){
		Movie1_Booking()
	} else if(MovieChoise == 2){
		Movie2_Booking()
	} else if(MovieChoise == 3){
		Movie3_Booking()
	} else if(MovieChoise == 4){
		Movie4_Booking()
	} else {
		fmt.Println("Wrong movie selection")
	}

}

func Movie1_Booking(){
	fmt.Println("--- Movie Tite: ", MOVIE1)
	BookingDetails()
}

func Movie2_Booking(){
	fmt.Println("--- Movie Tite: ", MOVIE2)
	BookingDetails()
}

func Movie3_Booking(){
	fmt.Println("--- Movie Tite: ", MOVIE3)
	BookingDetails()
}

func Movie4_Booking(){
	fmt.Println("--- Movie Tite: ", MOVIE4)
	BookingDetails()
}

func BookingDetails(){

	var movieType int
	var min int = 10
	var max int = 990
	var moreBookings string

	var(
		firstName string
		lastName string
		email string
		numberOfTickets uint
	) 

	fmt.Println("1. ",SCR_TYPE1,"| 2. ",SCR_TYPE2)
	fmt.Println("Choose movie type: ")
	fmt.Scan(&movieType)

	if(movieType == 1){
		fmt.Println("Per person = Rs. 3200/=")

		for{
			fmt.Println("Avilable seats: ", remainingSeatsIMAX3D)

			fmt.Println("Enter first name: ")
			fmt.Scan(&firstName)

			fmt.Println("Enter last name: ")
			fmt.Scan(&lastName)
			
			fmt.Println("Enter email: ")
			fmt.Scan(&email)

			fmt.Println("Enter number of tickets: ")
			fmt.Scan(&numberOfTickets)

			var totalPrice = 3200 * numberOfTickets

			var secretCode string = strconv.Itoa((rand.Intn(max - min) + min))
			
			//adding bookings to the slice
			if(remainingSeatsIMAX3D == int(numberOfTickets) || remainingSeatsIMAX3D > int(numberOfTickets)){
				bookings = append(bookings, firstName + " " + lastName + " - Code: " + secretCode)

				//seat calculation part
				bookedSeats = bookedSeats + int(numberOfTickets)
				remainingSeatsIMAX3D = remainingSeatsIMAX3D - bookedSeats

				//printing out the tickets
				fmt.Println("Your booking was successful. Thank you for choosing EONMAX Cinema...!")
				fmt.Printf("Your bookings: %v\n", bookings)
				fmt.Println("Total Price : Rs. ", totalPrice,"/=")

				fmt.Println("Book more tickets? (yes/no): ")
				fmt.Scan(&moreBookings)

				if(moreBookings == "yes"){
					if(remainingSeatsIMAX3D == 0){
						fmt.Println("All tickects are sold out")
						break
					} else {
						continue
					}
		
				} else {
					break
				}

			} else {
				fmt.Printf("Sorry. Unable to make the booking. Only %v seat/seats available\n", remainingSeatsIMAX3D)
				break
			}
			
			
		}
		
	} else if (movieType == 2){
		fmt.Println("Per person = Rs. 2000/=")
		
		for{
			fmt.Println("Avilable seats: ", remainingSeats3D)

			fmt.Println("Enter first name: ")
			fmt.Scan(&firstName)

			fmt.Println("Enter last name: ")
			fmt.Scan(&lastName)
			
			fmt.Println("Enter email: ")
			fmt.Scan(&email)

			fmt.Println("Enter number of tickets: ")
			fmt.Scan(&numberOfTickets)

			var totalPrice = 2000 * numberOfTickets

			var secretCode string = strconv.Itoa((rand.Intn(max - min) + min))
			
			//adding to the slice
			if(remainingSeats3D == int(numberOfTickets) || remainingSeats3D > int(numberOfTickets)){
				bookings = append(bookings, firstName + " " + lastName + " - Code: " + secretCode)

				//seat calculation part
				bookedSeats = bookedSeats + int(numberOfTickets)
				remainingSeats3D = remainingSeats3D - bookedSeats

				//printing out the tickets
				fmt.Println("Your booking was successful. Thank you for choosing EONMAX Cinema...!")
				fmt.Printf("Your bookings: %v\n", bookings)
				fmt.Println("Total Price : Rs. ", totalPrice,"/=")

				fmt.Println("Book more tickets? (yes/no): ")
				fmt.Scan(&moreBookings)

				if(moreBookings == "yes"){
					if(remainingSeats3D == 0){
						fmt.Println("All tickects are sold out")
						break
					} else {
						continue
					}
		
				} else {
					break
				}

			} else {
				fmt.Printf("Sorry. Unable to make the booking. Only %v seat/seats available\n", remainingSeats3D)
				break
			}
			
			
		}
	}

}