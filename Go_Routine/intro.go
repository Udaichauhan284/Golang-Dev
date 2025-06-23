package main

import (
	"fmt"
	"sync"
	//"time"
)

/* Go Routine is a lightweight thread manages by the Go runtime.

go f(x,y,z)
starts a new goroutine running

The evaluation of f,x,y and z happens in the current goroutine and the execution of f happens in the new goroutine.

Go routines run in the same address space, so access to shared memory must be synchronized.
The sync package provides useful primitives,
although you wont need them much in Go as there are other primitives
*/
// func task(id int){
// 	fmt.Println("doing task", id);
// }

//now use of Inline function
//Now use of Wait group, instead use of time.Sleep

func task(id int, w *sync.WaitGroup){
	defer w.Done();
	fmt.Println("Done task from WaitGroup: ", id);
}
func main(){
	// for i:=0; i<=10; i++{
	// 	//inline function, amnoums function
	// 	//also closure.
	// 	go func(i int){
	// 		fmt.Println("doing task: ", i);
	// 	}(i);
	// }

	// time.Sleep(time.Second);

	var wg sync.WaitGroup;
	for i:=0; i<=10; i++ {
		wg.Add(1);
		go task(i, &wg);
	}
	wg.Wait();
}