package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"os"
)

var secret = "CJ6V5YXGX3WSPJCPGHXVPXHAVRTIAKDR"
var passcode  = "177844"

func gen()  {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Example.com",
		AccountName: "alice@example.com",
	})
	if err != nil {
		fmt.Println("totp.Generate error!")
		return
	}

	fmt.Println(key.Secret())
}

func val()  {
	valid := totp.Validate(passcode, secret)
	if valid {
		println("Valid passcode!")
		os.Exit(0)
	} else {
		println("Invalid passocde!")
		os.Exit(1)
	}
}

func main() {
	gen()
	val()
}
