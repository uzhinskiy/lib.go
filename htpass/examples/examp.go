package main

import (
	"fmt"
	"github.com/uzhinskiy/htpass"
)

func main() {
	htp := make(htpass.HTPassFile)
	err := htp.OpenHTPASSFile("./.htpasswd")
	res, err := htp.Auth("user", "123")
	fmt.Printf("%v\t%v\t%v\n", htpass.IsAuth, res, err)
	if err != nil {
		fmt.Println(err)
	}
}
