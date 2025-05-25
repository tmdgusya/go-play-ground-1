package main

import (
	"bytes"
	"fmt"
)

/**
 * Go 언어에서 배열과 슬라이스의 차이점:
 * 1. 배열은 고정 크기이고, 슬라이스는 가변 크기입니다.
 * 2. 배열은 값 타입이며, 슬라이스는 참조 타입입니다.
 * 3. 배열은 메모리에서 연속된 공간을 차지하지만, 슬라이스는 배열의 일부를 참조합니다.

그래서 생각해보면 slice 는 그냥 아래 형태임

type sliceHeader struct {
	Length int
	ZerothElement *byte // 얘는 값이 배열의 첫번째 요소를 가리키기 때문에 call by value 가 아니라 call by reference 임
}
**/

func AddOneToEachElementOnArray(slice [5]byte) {
	for i := range slice {
		slice[i]++
	}
}

func AddOneToEachElementOnArray2(slice *[5]byte) {
	for i := range slice {
		slice[i]++
	}
}

func AddOneToEachElementOnSlice(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	print("Last index of slash: ", i, "\n")
	if i >= 0 {
		*p = (*p)[:i+1] // 마지막 슬래시 이후의 모든 바이트를 제거합니다.
	}
}

func main() {
	slice := []byte{1, 2, 3, 4, 5}
	fmt.Println("Before:", slice)
	AddOneToEachElementOnSlice(slice)
	fmt.Println("After:", slice)

	array := [5]byte{1, 2, 3, 4, 5}
	fmt.Println("Before:", array)
	AddOneToEachElementOnArray(array)
	fmt.Println("After:", array)

	array2 := [5]byte{1, 2, 3, 4, 5}
	fmt.Println("Before:", array2)
	AddOneToEachElementOnArray2(&array2)
	fmt.Println("After:", array2)

	pathName := path("/usr/local/bin")
	pathName.TruncateAtFinalSlash()
	fmt.Println("Truncated path:", string(pathName))
}
