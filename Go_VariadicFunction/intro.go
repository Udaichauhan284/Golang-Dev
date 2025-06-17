package main

import "fmt"

//In Variadic Function, we can send multiple variables
func sum(nums ...int) int {
	var total int;
	for _, num := range nums {
		total += num;
	}
	return total;
}
func main(){
	//var nums = []int{1,2,3,4,5};
	//result := sum(nums...);

	var result = sum(1,2,3,4,5);
	fmt.Println("The sum is: ", result);
}