package main

import (
	"fmt"
	"hw1/core"
)

func checkingAvgDistance() {
	testDemo := []int64{
		-10,
		1,
		2,
		3,
		4,
		5,
		23,
	}

	val := core.CheckAvgDistanceG[int64](testDemo)

	fmt.Println("Avg: ", val.Avg)
	fmt.Println("P50: ", val.P50)
}

func checkDistributionOfPositivesAndNegaties(values []int64) (float64, float64) {
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

func testingRamDistV2() {

	var capacity int64 = 1 << 8 // 62 Why is this the max ?
	fmt.Println(capacity)

	values := core.GenerateRandListG[int64](capacity)
	positiveGap, negativeGap := core.CheckDistributionOfPositivesAndNegatiesG[int64](values)
	fmt.Println("Positives: ", positiveGap)
	fmt.Println("Negatives: ", negativeGap)
}

func fullTets() {
	var capacity int64 = 1 << 32 // 62 Why is this the max ?
	fmt.Println(capacity)

	values := core.GenerateRandListG[int8](capacity)
	positiveGap, negativeGap := core.CheckDistributionOfPositivesAndNegatiesG[int8](values)
	fmt.Println("Positives: ", positiveGap)
	fmt.Println("Negatives: ", negativeGap)

	val := core.CheckAvgDistanceG[int8](values)

	fmt.Println("Avg: ", val.Avg)
	fmt.Println("P50: ", val.P50)
}
