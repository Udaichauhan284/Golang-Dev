package main

import (
	"fmt"
	"maps"
)

//Maps -> hash, object, dict
func main(){
	//creating map

	m := make(map[string]string); //map[keytype]valueType

	//setting an element
	m["name"] = "golang";
	m["area"] = "backend";

	//getting an element
	fmt.Println(m["name"]);
	fmt.Println(m["area"]);
	fmt.Println("Getting the key wont there: ", m["phone"]); //If key does not exists in the map, it return the empty value

	fmt.Println("Printing whole map: ", m);
	fmt.Println("Printing the len of map: ", len(m));

	//deleting th element
	delete(m, "area"); //give map name and key name
	fmt.Println("After deletion: ", m);

	//empty the map
	clear(m);
	fmt.Println("After clearing: ", m);

	//Making map, using without make
	m1 := map[string]int{"price" : 40, "phone" : 50};
	fmt.Println("New Map: ", m1);

	//Checking the value and getting the value
	val, ok := m1["price"];
	fmt.Println("Value of Price: ", val);
	if ok{
		fmt.Println("all ok key is there");
	}else {
		fmt.Println("not okay");
	}

	//Checking maps if equal or not
	m2 := map[string]int{"price" : 40, "phones" : 30};
	m3 := map[string]int{"price" : 40, "phones" : 30};
	fmt.Println(maps.Equal(m2,m3));
}