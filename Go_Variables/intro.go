package main
import ("fmt")

// func main(){
// 	//simple var declaration, using var
// 	var name string = "Udai" //type string
// 	var age int = 25 //type int
// 	var sex = "male" //type inferred by complier
// 	job := "SDE" //type string, but inferred

// 	fmt.Println("Name of Person: ", name);
// 	fmt.Println("Age of Person: ", age);
// 	fmt.Println("Sex of Person: ", sex);
// 	fmt.Println("Job of Person: ", job);
// }

// Variable Declare Without Initial Value
// func main(){
// 	var a string;
// 	var b int;
// 	var c bool;
// 	var d float64;

// 	fmt.Println("value of a: ", a);
// 	fmt.Println("value of b: ", b);
// 	fmt.Println("value of c: ", c);
// 	fmt.Println("value of d: ", d);

// 	//initialize value after declaration
// 	var name string;
// 	name = "udai";
// 	fmt.Println("value of name: ", name);
// }

//Multiple variable declarations can also be group together
//use of const 
const name string = "UDAI"; //typed one
const age = 25;
func main(){
	var (
		a int
		b int = 1
		c string = "hello"
	);

	fmt.Println("value of a: ", a);
	fmt.Println("value of b: ", b);
	fmt.Println("value of c: ", c);


	const id int = 10112;

	fmt.Println("value of const id: ", id);
	fmt.Println("Value of const name: ", name);
	fmt.Println("value of const age: ", age);

	//use of Printf, %v is used to show the value 
	//%T is used to show the type of variable
	fmt.Printf("value of name %v and type %T\n", name, name);
	fmt.Printf("value of age %v and tyoe %T ", age, age);
}