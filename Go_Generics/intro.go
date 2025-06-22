package main

import "fmt"

func printSlice(items []int) {
	for _, item := range items {
		fmt.Println("this is item val: ", item)
	}
}

//Now suppose i want to pass name in main func
//so i need to create the another func which accept the string slice.
func printStringSlice(items []string){
	for _, item := range items{
		fmt.Println("String val: ", item);
	}
}
//Here for printing the val, we are using the
//two func which have same logic, so in this case, we use the Generics

//In [T any], we can also use interface{}, on behalf of any
//func nameOfFunc[T any](items []T){}
//func nameOfFunc[T interface{}](items []T){}

//Now i want this func only take string and int
// func printGenericVal[T string | int | bool](items []T){
// 	for _, item := range items{
// 		fmt.Println("Generic val: ", item);
// 	}
// }

//we can also use the comparable in Generic
func printGenericVal[T comparable](items []T){
	for _, item := range items{
		fmt.Println("From Print func: ", item);
	}
}

//implementing Stack DS
type Stack[T any] struct{
	elements []T
}

func main() {

	stack := Stack[int]{
		elements: []int{1,2,3},
	}
	fmt.Println("Stack: ", stack);

	stack1 := Stack[string]{
		elements: []string{"udai", "tejswi"},
	}
	fmt.Println("Another Stack: ", stack1);

	nums := []int{1, 2, 3, 4, 5}
	//printSlice(nums);

	names := []string{"udai", "tanvi"};
	//printStringSlice(names);
	boolVal := []bool{true, false};
	printGenericVal(nums);
	printGenericVal(names);
	printGenericVal(boolVal);
}
