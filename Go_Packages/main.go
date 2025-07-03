/* Go packages are a fundamental concept in the Go langauage, allowing developers to organize and reuse code effectively.
A package in Go is a Collection of related Go scoure files that are compiled together into a single binary executable file. This organization helps in creating modular, reusable, and maintainable code.

packages in Go are named using a reverse domain name notation. For example, if your organization's domain name is example.com, you might name your package
com.example.mylib. It's also common to use shorted names for package that are frequently used in your code

*/

package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/udaichauhan284/podcast/Go_Packages/auth"
	"github.com/udaichauhan284/podcast/Go_Packages/user"
)

func main(){
	auth.LoginWithCredentials("udaichauhan", "1234");

	session := auth.GetSession();
	fmt.Println("Session ? ", session);

	//creating the user
	user := user.User{
		Email: "user@gmail.com",
		Name: "John Doe",
	}
	//fmt.Println(user.Email, user.Name);

	//now using the color package
	color.Red(user.Email, user.Name);

}