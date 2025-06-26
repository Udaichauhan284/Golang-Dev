package main

import (
	//"math/rand"
	"fmt"
	"time"
)

/* func processChan(takeChan chan int){
	// fmt.Println("Processing num: ", <-takeChan);

	for num := range takeChan{
		fmt.Println("processing number: ", num);
		//making little bit slow
		time.Sleep(time.Second);
	}
} */

//This is how data is Recived
/* func sum(takeChan chan int, num1 int, num2 int){
	// numResult := num1 + num2;
	// takeChan <- numResult

	takeChan <- num1 + num2;
}*/

//Use of goroutine synchronizer, using channel
// func task(done chan bool){
// 	//defer func is cleaning func, this will run one time
// 	defer func() { done <- true}();
// 	fmt.Println("processing...");
// }

//Now we will create the queue like
// func emailSender(emailChain chan string, done chan bool){
// 	defer func() {done <- true}();
// 	for email := range emailChain{
// 		fmt.Println("sending email to ", email);
// 		time.Sleep(time.Second);
// 	}
// }

//Now making the channel, which can only recieve the data 
//<-chan means recieving
//chan<- means sending
func emailSender(emailChan <-chan string, done chan<- bool){
	defer func() {done <- true}();
	
	//now if i want to send something to emailChain recieve channel, throw the error
	//emailChan <- "hello@gmail.com"
	

	//<-done this will thorw the error, here we are recieving, but in parameter we make the type safty of sending in channel

	for email := range emailChan {
		fmt.Println("sending the mail to ", email);
		time.Sleep(time.Second);
	}
}
func main(){

	//now make the two channel
	chan1 := make(chan int);
	chan2 := make(chan string);

	go func(){
		chan1 <- 10
	}()

	go func(){
		chan2 <- "pong"
	}()

	//Now to recieve from multile channel
	// for i := 0; i<2; i++ {
	// 	select {
	// 	case chan1Val := <- chan1:
	// 		fmt.Println("received data from chan1 ", chan1Val);
	// 	case chan2Val := <- chan2:
	// 		fmt.Println("received data from chan2 ", chan2Val);
	// 	}
	
	// }

	//now here we will make emailChain channel
	// emailChan := make(chan string, 100);
	// done := make(chan bool);

	// go emailSender(emailChan, done);

	// for i := 0; i<5; i++{
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i);
	// }

	// fmt.Println("done sending");
	// close(emailChan); //this is important to close the channel
	// <- done

	// emailChain := make(chan string, 100);

	// emailChain <- "1@example.com"
	// emailChain <- "2@example.com"
	// fmt.Println(<-emailChain);
	// fmt.Println(<-emailChain);
	//done := make(chan bool);
	//go task(done);

	//<- done //only reciving, blocking


	/* result := make(chan int);
	go sum(result, 4, 5);
	res := <- result;
	fmt.Println("The result: ", res);
	*/

	/* numChan := make(chan int);
	go processChan(numChan);
	//numChan <- 5;

	//sending the muliple val to chan in infinte loop
	for {
		numChan <- rand.Intn(100);
	} */
	//time.Sleep(time.Second);

	/* messageChain := make(chan string); //making of channel

	messageChain <- "ping" //sending data to channel //Blocking

	msg := <-messageChain; //receving the data into variable
	fmt.Println("The message from channel: ", msg); //This giving the deadlock
	*/
}