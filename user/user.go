package user

type User struct {
	name    string
	age     int
	petName string // ペットを飼っている人だけ追加したい
}

type Option func(*User)

func WithPetName(petName string) func(*User) {
	return func(u *User) {
		u.petName = petName
	}
}

func New(name string, age int, options ...Option) *User {
	user := &User{
		name: name,
		age:  age,
	}
	for _, option := range options {
		option(user) // ここでoptionで指定した値をuserにセット
	}
	return user
}
