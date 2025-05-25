package main

import (
	"fmt"
)

func main() {
	hello := "Hello, 월드!"
	rune := []rune(hello) // rune 은 int32 의 별칭 타입(uytf-8 문자하나를 표현)

	fmt.Println("문자열 길이:", len(hello))

	for i := 0; i < len(hello); i++ {
		fmt.Printf("타입: %T 값:%d 문자값: %c\n", hello[i], hello[i], hello[i])
	}

	fmt.Println("문자열을 rune로 변환한 길이:", len(rune))

	for i := 0; i < len(rune); i++ {
		fmt.Printf("타입: %T 값:%d 문자값: %c\n", rune[i], rune[i], rune[i])
	}
}
