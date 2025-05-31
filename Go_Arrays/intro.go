package main

import "fmt"

func main(){
	//Array is a numbered squence of specific length, of a specific type
	var nums [4]int //var variableName [space]type

	fmt.Println(len(nums)); //length of nums array

	//now i want to add element in array
	nums[0] = 1;
	fmt.Println("The nums array: ", nums);

	//mow in for loop i want to add the elem in array
	for i := 0; i < len(nums); i++ {
		nums[i] = i;
	}

	fmt.Println("The nums array after foor loop: ", nums);

	var vals [4]bool;
	vals[2] = true;
	fmt.Println("The vals boolean array: ", vals);

	var name [3]string;
	name[1] = "udai"; //adding the string at 1 index
	fmt.Println("The name string array: ", name);

	//Assigning the value at declaring time
	values := [3]int{1,2,3};
	fmt.Println("The inline declare values array: ", values);

	//2d arrays
	twoDNums := [2][2]int{{1,2},{3,4}};
	fmt.Println("The two d nums values array: ", twoDNums);

	//Now if you dont know the length
	arr1 := [...]int{1,2,3,4,5};
	arr2 := [...]int{4,5,90};
	fmt.Println("The len not know at declare time: ", arr1);
	fmt.Println("The len not know at declare time for arr2: ", arr2);

	//fixed in size, that is predictable, also if we dont know the len
	//we can defined using [...]
	//Memory optimization
	//Constant time access

	//We can also initialize only specifix elements
	arr3 := [5]int{1:10,2:50}
	fmt.Println("this is partially initialize: ",arr3);
}