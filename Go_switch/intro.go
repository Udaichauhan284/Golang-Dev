package main

import (
	"fmt"
	//"time"
)

func main(){
	//simple switch

	// i := 6;
	// switch i {
	// case 1:
	// 	fmt.Println("one");
	// case 2:
	// 	fmt.Println("two");
	// case 3:
	// 	fmt.Println("three");
	// case 4:
	// 	fmt.Println("four");
	// case 5:
	// 	fmt.Println("five");
	// default:
	// 	fmt.Println("Other number");
	// 	break;
	// }

	//Multiple condition switch
	//Here in this we can put multiple condition in swicth
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Enojot its weekend");
	// default:
	// 	fmt.Println("oh back to work");
	// 	break;
	// }

	//Type Switch
	whoAmI := func(i interface{}){
		switch i.(type){
		case int:
			fmt.Println("Its an integer");
		case string:
			fmt.Println("Its an string");
		case bool:
			fmt.Println("Its an boolean");
		default:
			fmt.Println("Other");
			break;
		}
	}

	//calling the function
	whoAmI(50); //Its an interger
	whoAmI("50.50"); //Its a string
}