package main

import "fmt"

//Iterating over data structures
func main(){
	nums := []int{6,7,8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println("Value of nums: ", nums[i]);
	// }

	//use of range, index, value := range nums
	for _, num := range nums {
		fmt.Println("Value of nums", num);
	}

	for ind, num := range nums {
		fmt.Println("The index: ",ind, "The value: ", num);
	}

	//now sum the nums value
	var sum int = 0;
	for _, num := range nums {
		sum += num;
	}
	fmt.Println("The sum of nums: ", sum);

	//same using this we can also iterate over the maps
	map1 := map[string]int{"age" : 25, "rollNum" : 77};
	for k,v := range map1 {
		fmt.Println("The key of map: ", k, "The value of map: ", v);
	}

	//Print only key
	for k := range map1 {
		fmt.Println("The only key of maps:", k);
	}

	//Iterating over the string
	for i,c := range "golang" {
		fmt.Println(i,c);
	}
	//The i and c will print the starting byte of rune and unicode code point of rune
	//Rune is data structure which used to store the unicode of g,o,l like wise
	
	//now if i want to print the char of string
	for i,c := range "golang" {
		fmt.Println(i, string(c));
	}
}