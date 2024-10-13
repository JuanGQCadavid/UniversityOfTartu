package generators

import "golang.org/x/exp/rand"

func GenRanInt(upperLimmit int32) int32 {
	return rand.Int31n(upperLimmit)
}

func GenerateRanList(n, upperLimmit int32) []int32 {
	arr := make([]int32, n)
	for index := range arr {
		arr[index] = GenRanInt(upperLimmit)
	}
	return arr
}
