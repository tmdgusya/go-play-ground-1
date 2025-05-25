package main

import (
	"fmt"
	"unsafe"
)

type Student struct {
	Name  string  // 16byte
	Class int     // 8byte
	No    int     // 8byte
	Score float64 // 8byte
}

type User struct {
	Name string // 16byte
	Age  int    // 8byte
}

type VIPUser struct {
	User      // embedded struct
	Level int // 8byte
}

func main() {
	var student Student = Student{
		Name:  "Roach",
		Class: 1,
		No:    1,
		Score: 99.5,
	}
	// 만약 컴퓨터 입장에서 64비트 시스템이라면
	// Student 구조체는 16 + 8 + 8 + 8 = 40바이트가 필요하다.
	// 만약 복사가 일어난다면 컴퓨터 입장에서는 메모리 복사가 일어나게 된다. (즉, 40바이트 복사)
	fmt.Printf("Student : %+v\n", student)
	fmt.Printf("Size of Student: %d bytes\n", unsafe.Sizeof(student))

	fmt.Printf("Student : %+v\n", student)
	fmt.Printf("Size of Student: %d bytes\n", unsafe.Sizeof(student))

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
