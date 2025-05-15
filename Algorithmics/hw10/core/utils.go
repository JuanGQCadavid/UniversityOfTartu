package core

import (
	"fmt"
	"math"
)

func HeapsAlgo[T any](k int, a []T, solution *[][]T) {
	if k == 1 {
		newSolution := make([]T, len(a))
		// fmt.Printf(" %+v \n", a)
		copy(newSolution, a)
		*solution = append(*solution, newSolution)
		return
	}

	HeapsAlgo(k-1, a, solution)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			temp := a[i]
			a[i] = a[k-1]
			a[k-1] = temp
		} else {
			temp := a[0]
			a[0] = a[k-1]
			a[k-1] = temp
		}
		HeapsAlgo(k-1, a, solution)
	}
}

func EucledianDist(a PointD, b PointD) float64 {
	return math.Sqrt(
		math.Pow(float64(a.X)-float64(b.X), 2) +
			math.Pow(float64(a.Y)-float64(b.Y), 2))
}

func EvaluatePoints(points []PointD) Paths {
	dist := 0.0
	for i := 0; i < len(points)-1; i++ {
		dist += EucledianDist(points[i], points[i+1])
	}

	return Paths{
		Score:  dist,
		Points: points,
	}
}

func testingHeapAlgo() {
	A := []string{
		"A", "B", "C",
	}

	solutions := make([][]string, 0)
	HeapsAlgo(len(A), A, &solutions)

	for index, val := range solutions {
		fmt.Printf("%d %+v \n", index, val)
	}
}
