package main

import (
	"fmt"
	"hw1/core"
	mathV2 "math/rand/v2"
)

func testingRamDist() {
	fmt.Println(mathV2.Int32())

	var capacity int64 = 1 << 8 // 62 Why is this the max ?
	fmt.Println(capacity)

	values := core.GenerateRandList(capacity)

	positiveGap, negativeGap := checkDistributionOfPositivesAndNegaties(values)
	fmt.Println("Positives: ", positiveGap)
	fmt.Println("Negatives: ", negativeGap)
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

	values := core.GenerateRandListG[int16](capacity)

	positiveGap, negativeGap := core.CheckDistributionOfPositivesAndNegaties(values)
	fmt.Println("Positives: ", positiveGap)
	fmt.Println("Negatives: ", negativeGap)
}

func main() {
	testingRamDistV2()
}

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

	val := core.CheckAvgDistance(testDemo)

	fmt.Println("Avg: ", val.Avg)
	fmt.Println("P50: ", val.P50)
}
