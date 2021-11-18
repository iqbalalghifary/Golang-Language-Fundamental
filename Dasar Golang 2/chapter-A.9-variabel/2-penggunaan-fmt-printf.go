package main

import "fmt"

func main() {
	var firstName string = "john"
	var lastName string
	lastName = "wick"

	var nama string = "aing"

	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s !\n", firstName)
	fmt.Println("halo", firstName, lastName+"!")

	fmt.Printf("hallo, nama %s iqbal", nama)
}
