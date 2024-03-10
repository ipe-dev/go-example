package main

import (
	"fmt"
	user "go-example/functional_option_pattern"
)

func main() {
	user := user.New("taro", 20, user.WithPetName("tom"))
	fmt.Println(*user)
}
