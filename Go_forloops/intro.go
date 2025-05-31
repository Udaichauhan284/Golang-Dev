package main

import "fmt"

//in Go, there is only for loop
func main(){
	//there is no while loop in go, need to use the
	//for keyword
	//while loop
	//var i int = 0;
	i := 0;
	for i<=3{
		fmt.Println("The iteration: ", i);
		i++;
	}
	fmt.Println("---------");
	//classic for loop
	for i:= 0; i<=3; i++ {
		if i == 2{
			continue;
		}
		fmt.Println("the iteration: ", i);
		// if i == 2 {
		// 	break;
		// }
		
	}

	//in go version 1.22 range
	for i:= range 3 {
		fmt.Println(i); //0,1,2
	}
}