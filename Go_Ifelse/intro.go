package main

import "fmt"

func main(){
	//age := 20;

	// if age >= 18 {
	// 	fmt.Println("Person is an adult");
	// }else {
	// 	fmt.Println("Not an adult");
	// }

	// if age >= 18 {
	// 	fmt.Println("Person is an adult");
	// }else if age >= 12 {
	// 	fmt.Println("person is a teenager");
	// }else {
	// 	fmt.Println("Person is kid");
	// }

	// var role = "admin";
	// var hasPermission = false;

	// if role == "admin" && hasPermission {
	// 	fmt.Println("yes");
	// }else {
	// 	fmt.Println("You dont have permission");
	// }

	//we can declare a ariable inside if 
	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age);
	}else if age >= 12 {
		fmt.Println("Person is teenager", age);
	}

	//Go does not have ternary operator
}