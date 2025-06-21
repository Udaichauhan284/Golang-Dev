package main

import "fmt"

//In Go, if we sending num in function, so it send by value
//so in another function it have dynamically new copy num
// By value
// func changeNum(num int){
// 	num = 5;
// 	fmt.Println("In changeNum func: ", num);
// }

//By reference
func changeNum(num *int){
	//num = 5 this will give the error, because we are getting
	//the pointer, means a memory address, so we need to dereference it
	*num = 5; //change the value in that memory address
	fmt.Println("In changeNum func: ", *num); //use of agian * means dereferrence
}
func main(){
	num := 1
	//changeNum(num); //so it sending 1 to changeNum function
	//so changeNum func should change the num to 5, but 
	//when function move to println num is still 1

	changeNum(&num); //sending the memory address
	fmt.Println("After the changeNum in main: ", num);
}