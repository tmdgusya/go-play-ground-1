package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func toHex(s string) string {
	return hex.EncodeToString([]byte(s))
}

func main() {
	hello := "Hello, 월드!"
	// rune := []rune(hello) // rune 은 int32 의 별칭 타입(uytf-8 문자하나를 표현)

	// fmt.Println("문자열 길이:", len(hello))

	// for i := 0; i < len(hello); i++ {
	// 	fmt.Printf("타입: %T 값:%d 문자값: %c\n", hello[i], hello[i], hello[i])
	// }

	// fmt.Println("문자열을 rune로 변환한 길이:", len(rune))

	// for i := 0; i < len(rune); i++ {
	// 	fmt.Printf("타입: %T 값:%d 문자값: %c\n", rune[i], rune[i], rune[i])
	// }

	for _, v := range hello {
		fmt.Printf("타입: %T 값:%d 문자값: %c\n", v, v, v)
	}

	codePoint := rune(0xC6D4)

	binary16BitStr := fmt.Sprintf("%016b", codePoint)

	fmt.Println("16비트 이진수 표현:", binary16BitStr)

	r := '월' // 1100 0110 1101 0100
	c := string(r)
	bits := ""
	for i := 0; i < len(c); i++ {
		bits += fmt.Sprintf("%08b ", c[i])
	}
	fmt.Println("바이트: ", bits)

	standard3Bytes := "1110wwww 10xxxxyy 10yyzzzz"
	fmt.Println("표준 3바이트 인코딩:", standard3Bytes)
	wwww := (r >> 12) & 0b1111
	xxxxyy := (r >> 6) & 0b111111
	yyzzzz := r & 0b111111
	fmt.Println("첫 4비트:", fmt.Sprintf("%04b", wwww))     // 1100
	fmt.Println("다음 6비트:", fmt.Sprintf("%06b", xxxxyy))  // 011011
	fmt.Println("마지막 6비트:", fmt.Sprintf("%06b", yyzzzz)) // 010100

	standard3Bytes = strings.Replace(standard3Bytes, "wwww", fmt.Sprintf("%04b", wwww), 1)
	standard3Bytes = strings.Replace(standard3Bytes, "xxxxyy", fmt.Sprintf("%06b", xxxxyy), 1)
	standard3Bytes = strings.Replace(standard3Bytes, "yyzzzz", fmt.Sprintf("%06b", yyzzzz), 1)
	fmt.Println("최종 3바이트 인코딩:", standard3Bytes)
}
