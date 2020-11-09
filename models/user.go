package models



type UserStore interface{
	CreateUser(name string) (error)
	ReadUser(id int) (int, error)
	UpdateUser(User) (int, error)
	DeleteUser(name string) (int, error)
}


type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Employment string `json:"employment"`
}

func(u *User)CreateUser(newUser User)(error){
	u.Name = newUser.Name
	u.Age = newUser.Age
	u.Employment = newUser.Employment


	
	return nil
}

func(u *User)ReadUser(id int)(int,error){
	return 1,nil
}

func(u *User)UpdateUser(newUser User)(int,error){
	return 1,nil
}

func(u *User)DeleteUser()(int,error){
	return 1,nil
}
