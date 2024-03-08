package user

import (
	"fmt"
)

func main() {
	user := New("taro", 20, WithPetName("tom"))
	fmt.Println(*user)
}
