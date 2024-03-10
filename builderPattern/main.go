package main

import (
	"fmt"
)

func main() {
	builder := NewBuilder("taro", 19).
		PetName("pochi").
		Height(170)
	user := builder.Build()
	fmt.Println(user)
}
