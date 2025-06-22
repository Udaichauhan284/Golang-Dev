package main

import "fmt"

//Enumerated types
/*Go doesn't have built.in enum support, but we can create enum-like constructs using the 
"iota" keyword and custom types.

iota is an identifier used with constants to simplify the definition of successive in integer constants start at 0 and increment by 1
for each subsequent constant in a constant block
*/
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)
func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to ", status);
}

func main(){
	changeOrderStatus(Received); //now this give 0, means Received
	changeOrderStatus(Delivered); //this give the 3, means Delivered
}