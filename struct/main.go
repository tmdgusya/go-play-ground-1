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

type TestAssignment struct {
	A int8 // 1byte
	B int  // 8byte
	C int8 // 1byte
	D int  // 8byte
	E int8 // 1byte
} // 1 + 8 + 1 + 8 + 1 = 19 bytes, but padding makes it 24 bytes

type TestAssignment2 struct {
	A int8 // 1byte
	B int8 // 1byte
	C int8 // 1byte
	D int  // 8byte
	E int  // 8byte
} // 1 + 1 + 1 + 8 + 8 = 19 bytes, but padding makes it 5 bytes
// TestAssignment2 is a more efficient struct with less padding

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
	fmt.Printf("Student : %+v\n", student)                            // {Name:Roach Class:1 No:1 Score:99.5})
	fmt.Printf("Size of Student: %d bytes\n", unsafe.Sizeof(student)) // 40 bytes (output)

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

	testUser := TestAssignment{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
	}

	fmt.Printf("TestAssignment : %+v\n", testUser)
	fmt.Printf("Size of TestAssignment: %d bytes\n", unsafe.Sizeof(testUser)) // 40 bytes (output)

	testUser2 := TestAssignment2{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
	}

	fmt.Printf("TestAssignment2 : %+v\n", testUser2)
	fmt.Printf("Size of TestAssignment2: %d bytes\n", unsafe.Sizeof(testUser2)) // 24 bytes (output)
}
