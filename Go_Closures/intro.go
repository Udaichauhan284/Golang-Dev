package main

import "fmt"

func counter() func() int {
	var count int = 0;

	return func() int {
		count += 1;
		return count;
	}
}
func main(){
	increment := counter();
	fmt.Println("The value of count: ", increment());
	fmt.Println("The value of count: ", increment());
}