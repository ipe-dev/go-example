package main

import (
	"fmt"
	"go-example/user"
)

func main() {
	user := user.NewUser("taro", 20)
	user.petName = "goro"
	fmt.Println(*user)
}
