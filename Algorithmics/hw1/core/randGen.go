package core

import (
	"fmt"
	mathV2 "math/rand/v2"
	"sort"
)

// Attempting Generics
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func GenerateRandListG[V int16 | int32 | int64](n int64) []V {
	values := make([]V, n)
	for i := range values {
		values[i] = V(mathV2.Int64() << 1)
	}
	return values
}

func CheckDistributionOfPositivesAndNegaties[V int16 | int32 | int64](values []V) (float64, float64) {
	positiveBalance := 0.0
	negativeBalance := 0.0

	for i := range values {
		if values[i] < 0 {
			negativeBalance++
			continue
		}
		positiveBalance++
	}

	fmt.Println("Positives: ", positiveBalance)
	fmt.Println("Negatives: ", negativeBalance)

	positiveGap := (positiveBalance / (positiveBalance + negativeBalance)) * 100
	negativeGap := (negativeBalance / (positiveBalance + negativeBalance)) * 100

	return positiveGap, negativeGap
}

// ORIGINAL
type RandStats struct {
	Avg float64
	P50 int64
}

func GenerateRandList(n int64) []int64 {
	values := make([]int64, n)
	for i := range values {
		values[i] = mathV2.Int64() << 1
	}
	return values
}

func sortSlice(slops []int64) {
	sort.Slice(slops, func(i, j int) bool {
		return slops[i] < slops[j]
	})
}

func absDiffInt(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}

func CheckAvgDistance(values []int64) RandStats {
	slops := make([]int64, len(values)-1) // the last number will not have a pair to check agains
	var accomulativeAvg int64 = 0
	for index := range values {
		if index == len(values)-1 {
			break
		}

		slops[index] = absDiffInt(values[index], values[index+1])
		accomulativeAvg += slops[index]
	}

	sortSlice(slops)
	return RandStats{
		Avg: float64(accomulativeAvg) / float64(len(slops)),
		P50: slops[len(slops)/2],
	}
}
