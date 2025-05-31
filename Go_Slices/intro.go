package main

import (
	"fmt"
	"slices"
)


func main(){
	/* Slices are similar to arrays, but are more powerful and flible.
	Like arrays, slice are also used to store muliplt values of the same type 
	in a single variable.
	However, unlike arrays, the length of slice can grow and shrink 
	as you seefit.

	In Go, there are several wasy to create a slice:
	1. using the []datatype{values} format
	2. create a slice from an array
	3. using the make() function
	*/

	mySlice := []int{1,2,3};
	//In Go there are two function to return the 
	//len and capacity of a slice, len() and cap()
	fmt.Println("The len of slice: ", len(mySlice));
	fmt.Println("The capacity of slice: ", cap(mySlice));

	//len(): function - return the length of the slice (the number of element in slice);
	//cap() : function - return the capacity of the slice (the number of elements the slice can grow or shrink to)

	//Create a slice from an array
	//myArray := [length]datatype{values}; //an array
	//mySlice := myArray[start:end] //an slice made from the array

	arr1 := [6]int{10,11,12,13,14,15};
	mySlice1 := arr1[2:4]; //start included : end not

	fmt.Println("The Slice: ", mySlice1); //[12, 13]
	fmt.Println("The length: ", len(mySlice1)); //2
	fmt.Println("The capacity: ", cap(mySlice1)); //4

	/* In this example mySlice1 is a slice with length 2. It is maade from arr1 which is an
	array with length 6.
	the slice starts from the third elem of an array which 
	has value 12 and goes till end-1 13, so slice is starting from index 2, and in array 
	the len is 6, so the capacity of slice can grow till 6, so now the cap of slice with 2
	elem inside is 4 remaining.
	If mySlice1 started from element 0, the slcie capacity would be 6
	*/

	//uninitialized sice is nil which is null
	var nums []int;
	fmt.Println("The nil slice: ", nums);
	fmt.Println("The ans of slice: ", nums == nil); //true


	//Create a Slice With the make() function, this also used to make a slice
	//Syntax: slice_name := make([]type,  length, capacity);
	//it capacity paramter is not defined, it will be equal to length
	
	mySlice2 := make([]int, 5, 10); //length, capacity
	fmt.Printf("mySlice2 = %v\n", mySlice2);
	fmt.Printf("length = %d\n", len(mySlice2));
	fmt.Printf("capacity = %d\n", cap(mySlice2));

	//with omitted capacity
	mySlice3 := make([]int, 5); //length
	fmt.Printf("mySlice3 = %v\n", mySlice3);
	fmt.Printf("length = %d\n", len(mySlice3));
	fmt.Printf("capacity = %d\n", cap(mySlice3));

	//using example of Video, using append
	//var nums1 = make([]int, 2, 5); //len, cap, this will give 2 zero in slice
	var nums1 = make([]int, 1, 5); //this is ideal way
	//nums1[0] = 1, so for achieving this we need to initialize slice slice with some len
	nums1 = append(nums1, 1); //[0,0,1]
	nums1 = append(nums1, 2);  //[0,0,1,2]
	nums1 = append(nums1, 3); //[0,0,1,2,3]
	nums1 = append(nums1, 4); //[0,0,1,2,3,4], cap 10
	fmt.Println(nums1);
	fmt.Println("The capacity now of nums1: ", cap(nums1));

	//Copy function
	var oldNums = make([]int, 0, 5);
	//now adding one element in oldone
	oldNums = append(oldNums, 2); //[2]
	var newNums = make([]int, len(oldNums));
	
	
	copy(newNums, oldNums); //[2], [2], it will copy
	fmt.Println("The old nums: ", oldNums , "The new nums ", newNums);

	//comparing
	var num1 = []int{1,2,3};
	var num2 = []int{1,2,3};
	fmt.Println("It will compare the slices: ", slices.Equal(num1, num2));

	//2d slices
	var twoD = [][]int{{1,2,3}, {4,5,6}};
	fmt.Println("The two d slices: ", twoD);
}