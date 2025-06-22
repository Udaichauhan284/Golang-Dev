package main

import "fmt"

//Now Start with Interface
type paymenter interface{
	pay(amount float32) 
}

type payment struct{
	//gateway stripe, so here dont provide any hardcoded stripe or razorpay
	gateway paymenter
}

//adding func of makePayment to struct
func (p payment) makePayment(amount float32){
	// razorpayPaymentGW := razorpay{};
	// razorpayPaymentGW.pay(amount);

	//here we are making changes in makePayment
	//So here we are violating the SOLID
	//stripePaymentGM := stripe{};
	//stripePaymentGM.pay(amount);

	//So calling here the struct gateway from payment
	p.gateway.pay(amount);
}

//adding the payment gateway
type razorpay struct{

}
//adding the func to rarozpay struct
func (r razorpay) pay(amount float32){
	//logic to make payment
	fmt.Println("making payment using razorpay", amount);
}


//now in future we want to add another payment gateway
type stripe struct{

}

//adding the func pay to stripe
func (s stripe) pay(amount float32){
	//logic to make payment
	fmt.Println("Making payement using the stripe: ", amount);
}
//So, In Go we dont need to write implement interface name, if in interface there is same 
//name of func which use in struct is impliciltly know to use the interface

//Suppose in future i need to use Paypal
type paypal struct{

}
func (p paypal) pay(amount float32){
	fmt.Println("making payment using the paypal: ", amount);
}
func main(){
	//so call the struct here
	//stripePaymentGM := stripe{};
	paypalPaymentGW := paypal{};
	newPayment := payment{
		gateway: paypalPaymentGW,
	};
	newPayment.makePayment(100);
}