/*
	A struct(short for structure) is used to creaste a collection of memebers of differentata types

into a single variable.
While arrays are used to store multiple values of the same data type into a
single variable, stucts are used to store muliple values of ifferent data types into a
single variables.

# Syntax

type struct_name struct {
member1 datatype;
member2 datatype;
member2 datatype
}

type Person struct {
name string;
age int;
job string;
salary float;
}
*/
package main

import (
	"fmt"
	"time"
)

type order struct{
	id string;
	amount float32;
	status string;
	createdAt time.Time
}
func main(){
	// var order order 
	// order.id = "123456"
	// order.amount = 4567
	// order.status = "InProcess"

	myOrder := order{
		id : "1",
		amount : 4567,
		status : "done",
	}
	//now if i want to add later
	myOrder.createdAt = time.Now();

	//getting the field
	fmt.Println("The status of order: ", myOrder.status);

	fmt.Println("This is order: ", myOrder);
}