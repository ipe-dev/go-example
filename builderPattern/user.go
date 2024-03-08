package user

type User struct {
	param UserParam
}

type UserParam struct {
	name    string
	age     int
	petName string
	height  int
	weight  int
}

func NewBuilder(name string, age int) *UserParam {
	return &UserParam{
		name: name,
		age:  age,
	}
}

func (up *UserParam) PetName(petName string) *UserParam {
	up.petName = petName
	return up
}

func (up *UserParam) Height(height int) *UserParam {
	up.height = height
	return up
}

func (up *UserParam) Weight(weight int) *UserParam {
	up.weight = weight
	return up
}

func (up *UserParam) Build() *User {
	user := &User{
		param: *up,
	}

	return user
}
