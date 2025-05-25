package main

import "fmt"

// pointer 는 메모리 주소를 값으로 가지는 타입이다.
func main() {
	var a int = 10
	var p *int = &a      // 포인터 변수 p는 a의 메모리 주소를 저장한다.
	fmt.Println("a:", a) // a: 10
	fmt.Println("p:", p) // p: 메모리 주소 (예: 0xc0000120a8)

	var array [3]int = [3]int{1, 2, 3}
	var pArray *[3]int = &array    // 포인터 변수 pArray는 array의 메모리 주소를 저장한다.
	fmt.Println("array:", array)   // array: [1 2 3]
	fmt.Println("pArray:", pArray) // pArray: 메모리 주소 (예: 0xc0000120b0)

	// print start address of array
	fmt.Printf("Start address of array: %p\n", &array[0]) // Start address of array: 0xc0000120b0
	// print value by referencing the pointer
	startAddress := &array[0]
	fmt.Printf("Value at start address: %d\n", *startAddress) // Value at start address: 1
	// dereference pointer to get value
	fmt.Printf("Value at pArray: %v\n", *pArray) // Value at pArray: [1 2 3]
	// add 8bytes to pointer to get next element
	nextElementAddress := (*pArray)[0] + 8                        // assuming 8 bytes for int
	fmt.Printf("Next element address: %p\n", &nextElementAddress) // Next element address: 0xc0000120b8
	// go prevent adding 8 bytes to pointer directly, we need to use index
	// So, if you want to access the next element by pointer arithmetic, you can do it like this:
	// using unsafe package to manipulate pointer arithmetic
	// arr := [3]int{1, 2, 3}
	// 1) arr[0]의 unsafe.Pointer
	// p0 := unsafe.Pointer(&arr[0])
	// 2) uintptr로 변환하여 바이트 오프셋을 더함
	// p1 := unsafe.Pointer(uintptr(p0) + unsafe.Sizeof(arr[0]))
}
