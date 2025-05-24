package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.0, 26.0, 27.0, 28.0}

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}

	days := [3]string{"monday", "tuesday", "wednesday"}
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	var s = [5]int{1: 10, 3: 30} // 인덱스 1에 10, 인덱스 3에 30을 할당하는 배열 선언
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for _, v := range days {
		fmt.Println(v)
	}
}
