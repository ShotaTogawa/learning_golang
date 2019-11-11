package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// built int package以外は,go get でpackageを取得する必要がある
	s := `password`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(s)

	//loginPword1 := "password"
	loginPword2 := "password123"


	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword2))
	if err != nil {
		fmt.Println("You cant't login")
		return
	}
	fmt.Println("you logged in")
}
