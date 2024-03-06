package main

import (
	"fmt"
	"go-example/user"
)

func main() {
	user := user.New("taro", 20, user.WithPetName("tom"))
	fmt.Println(*user)
}
