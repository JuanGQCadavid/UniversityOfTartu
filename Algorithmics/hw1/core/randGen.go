package core

import (
	mathV2 "math/rand/v2"
	"sort"
)

// Attempting Generics

type RandStats struct {
	Avg                          float64
	P50                          int64
	PositivesPercentage          float64
	NegativePercentage           float64
	Timing                       float64
	TimingStdDev                 float64
	TimingCoefficientOfVariation float64
}

func GenerateRandListG[V int8 | int16 | int32 | int64](n int64) []V {
	values := make([]V, n)
	for i := range values {
		values[i] = V(mathV2.Int64() << 1)
	}
	return values
}

func CheckDistributionOfPositivesAndNegatiesG[V int8 | int16 | int32 | int64](values []V) (float64, float64) {
	positiveBalance := 0.0
	negativeBalance := 0.0

	for i := range values {
		if values[i] < 0 {
			negativeBalance++
			continue
		}
		positiveBalance++
	}

	positiveGap := (positiveBalance / (positiveBalance + negativeBalance)) * 100
	negativeGap := (negativeBalance / (positiveBalance + negativeBalance)) * 100

	return positiveGap, negativeGap
}

func CheckAvgDistanceG[V int8 | int16 | int32 | int64](values []V) RandStats {
	slops := make([]int64, len(values)-1) // the last number will not have a pair to check agains
	var accomulativeAvg int64 = 0
	for index := range values {
		if index == len(values)-1 {
			break
		}

		slops[index] = int64(absDiffIntG[V](values[index], values[index+1]))
		accomulativeAvg += slops[index]
	}

	sortSliceG(slops)
	return RandStats{
		Avg: float64(accomulativeAvg) / float64(len(slops)),
		P50: slops[len(slops)/2],
	}
}

func sortSliceG[V int8 | int16 | int32 | int64](slops []V) {
	sort.Slice(slops, func(i, j int) bool {
		return slops[i] < slops[j]
	})
}

func absDiffIntG[V int8 | int16 | int32 | int64](x, y V) V {
	if x < y {
		return y - x
	}
	return x - y
}
