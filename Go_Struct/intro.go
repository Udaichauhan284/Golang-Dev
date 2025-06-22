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

//Now we can also use one struct inside another struct, which known as Struct Embedding
type customer struct {
	name string;
	age int
}

type order struct{
	id string;
	amount float32;
	status string;
	createdAt time.Time;
	customer //using another struct here
}

//Making a constructor in GO
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id : id,
		amount : amount,
		status: status,
	}

	return &myOrder; //returning the pointer of struct, so it can use the origanl struct, instead of creating new one
}

//using another func to get the order struct
//now using this way, this function attached to order struct
//use * when we want to change the value of struct
func (o *order) changeStatus(status string){
	o.status = status;
}

//another func to get the amount from struct
func (o order) getAmount() float32 {
	return o.amount;
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
	//so in myOrder a function is linked
	//we can use this function to chanage
	//the status
	myOrder.changeStatus("confirmed");
	//now if i want to add later
	//myOrder.createdAt = time.Now();

	//getting the field
	//fmt.Println("The status of order: ", myOrder.status);

	fmt.Println("This is order: ", myOrder);
	//get the amount from func
	fmt.Println("This is amount of order: ", myOrder.getAmount());

	// myOrder2 := order{
	// 	id : "2",
	// 	amount : 100.00,
	// 	createdAt : time.Now(),
	// }

	//changing the status of first order
	// myOrder.status = "paid"
	// fmt.Println("Order Struct after change: ", myOrder);

	// fmt.Println("New order: ", myOrder2);

	//Now calling the constructor
	myOrder3 := newOrder("1",30.50,"markDone");
	fmt.Println("This is by constructor: ", myOrder3);

	//creating no name struct
	language := struct {
		name string
		isGood bool
	}{"golang", true}

	fmt.Println("This is no name struct: ", language);

	//example of use of embedded struct
	// newCustomer := customer{
	// 	name : "john",
	// 	age : 25,
	// }
	myOrder4 := order{
		id : "2",
		amount : 304,
		status: "done",
		//customer: newCustomer,
		//or can we use as inline
		customer : customer{
			name : "udai",
			age : 23,
		},
	}

	//now want to change the customer name in myOrder4
	myOrder4.customer.name = "Robin";
	fmt.Println("Example of Embedded Struct:", myOrder4);
}