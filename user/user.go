package user

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	return &User{
		name: name,
		age:  age,
	}
}
