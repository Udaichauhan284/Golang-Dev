package main

import "fmt"

const address string = "gurgram"; //okay
//job := "sde" this type of declaration not possible outside of function
func main(){
	const name string = "udai"; //valid
	const age = 25; // still valid
	// age = 26, cannot assign again to const

	fmt.Println("Address: ", address);

	//multiple value in const
	const (
		port = 5000
		host = "localhost"
	)
	//now if want to change the const, not possible
	// port = 5500, not possible
	//now print the const
	fmt.Println(port, host);
}