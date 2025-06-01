package main

import "fmt"

//creating the function, add
// func add(a int, b int) int { //this means it returning the number
// 	return a+b;
// }

// If in func parameters are same, we can also defined only one
func add(a, b int) int {
	return a + b
}

// now creating func which return 2 values
func addNMul(a int, b int) (int, int) {
	return a + b, a * b
}

func getLanguages() (string, string, string, bool) {
	return "golang", "javascripts", "java", true
}

//in function passing the func which take int and return int
// func processIt(fn func(a int) int){
// 	fn(1);
// }

// this func is returning the func which is returning the val
func processIt() func(a int) int {
	return func(a int) int {
		return 2
	}
}

func main() {
	//calling that function
	result := add(4, 5)
	fmt.Println("The result of add function: ", result)

	//now getting two values from function
	addVal, multipleVal := addNMul(3, 4)
	fmt.Println("The add: ", addVal, "The multiple: ", multipleVal)

	//getting multiples value from function
	// lang1, lang2, lang3, boolVal := getLanguages();
	// fmt.Println(lang1, lang2, lang3, boolVal);

	//suppose from getting value, i dont want use one value, we can use _
	lang1, lang2, lang3, _ := getLanguages()
	fmt.Println(lang1, lang2, lang3)

	// fn := func(a int) int {
	// 	return 2;
	// }
	// processIt(fn);

	fn := processIt()
	fn(2);
}
