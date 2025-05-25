package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

type User struct {
	Name string
	Age  int
}

type VIPUser struct {
	User  // embedded struct
	Level int
}

func main() {
	var student Student = Student{
		Name:  "Roach",
		Class: 1,
		No:    1,
		Score: 99.5,
	}

	fmt.Printf("Student : %+v\n", student)

	var user User = User{
		Name: "Roach",
		Age:  18,
	}

	fmt.Printf("User : %+v\n", user)

	var vipUser VIPUser = VIPUser{
		User:  user,
		Level: 1,
	}

	fmt.Printf("VIP User : %+v\n", vipUser)
	fmt.Printf("VIP User Name: %s, Age: %d, Level: %d\n", vipUser.Name, vipUser.Age, vipUser.Level)
}
