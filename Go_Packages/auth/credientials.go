package auth

import "fmt"

func LoginWithCredentials(username string, password string){
	fmt.Println("Login user using: ", username, password);
}

/* if we declare the function name with small letters, like l,m,a etc, so we can
only use this function in same package,
but if we want to xport these function,
so that outer this package we can use it,
so just declare with capital Letter
*/