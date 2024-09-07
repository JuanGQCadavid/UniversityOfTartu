package main

import "fmt"

func main() {

	fmt.Println(fiboV1(10))
	fmt.Println(fiboV2(10))

}

func fiboV1(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fiboV1(n-1) + fiboV1(n-2)
}

func fiboV2(n int) int {
	fibResults := make([]int, n)
	fibResults[0] = 0
	fibResults[1] = 0

	if n < 2 {
		fmt.Println("up")
		return fibResults[n-1]
	}

	for i := 2; i < n; i++ {
		fibResults[i] = fibResults[i-1] + fibResults[i-2]
	}

	return fibResults[n-1]
}
