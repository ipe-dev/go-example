package user

type User struct {
	name    string
	age     int
	petName string
	height  int
	weight  int
}

type Option func(*User)

func WithHeight(height int) func(*User) {
	return func(u *User) {
		u.height = height
	}
}

func WithWeight(weight int) func(*User) {
	return func(u *User) {
		u.weight = weight
	}
}
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
