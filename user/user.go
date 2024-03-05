package user

type User struct {
	Name string
	Age int
	PetName string // ペットを飼っている人だけ追加したい
}

func New(name string, age int ) *User {
	return &User{
		Name: name,
		Age: age,
	}
}