/* A mutex, short for Mutual Exclusion, is a synchronization mechanism used in Go to
ensure that only one Goroutine can access a
shared resource at a time.
This helps prevent race conditions, which occur when multiple goroutines access and
modify shared data concurrently, leading to
unpredictable results.

In Go, the sync.Mutex type provides mutual
exclusion. The Lock() method is used to acquire the lock, and the Unlock() method
is used to release it.
When a goroutine calls Lock(), it gains exclusive access to the shared resource. If
another goroutine tries to acquire the lock while it is held, it will block until the
lock is released.
*/

package main

import (
	"fmt"
	"sync"
)

// create the sturct
type post struct {
	views int
	//now here we use the mutex lock and unlock
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	//defer wg.Done()
	defer func(){
		wg.Done();
		p.mu.Unlock();
	}();
	//now here we are incrementing the views
	// so to avoid the race condition, we use the mutex Lock and then we will unlock it.
	p.mu.Lock(); //Try to lock only that, where some modiftication is happening
	p.views += 1;
	//p.mu.Unlock(); //best practice to put the Unlock in defer func, because while increment or any other operations there can be an error occur, so it will not Unlock it.
	//so move it to defer
}
func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	// myPost.inc();
	// fmt.Println("After increment: ", myPost.views);

	//Now i will use this in synchonously in for loop
	// for i := 0; i<100; i++{
	// 	myPost.inc(); //This is running in sync
	// }
	// fmt.Println("After the loop: ", myPost.views);

	//Now i want this to run concurrently
	//now we have to use the waitgroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println("After the concurrently: ", myPost.views) //So in this every time, i am getting the different ans, 100, 99, 98, 96
	//because these running the concurrenlty
	//and one change is overriding the other
	//now we can use the Mutex, to Lock the
	//views.

}
