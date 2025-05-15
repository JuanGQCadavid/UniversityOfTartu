package core

import (
	"log"
	"math"
	"math/rand"
)

func Swap(points []PointD) []PointD {
	var (
		poinsTmp = make([]PointD, len(points))
	)
	copy(poinsTmp, points)

	for {
		a, b := rand.Intn(len(points)), rand.Intn(len(points))
		if a == b {
			continue
		}
		poinsTmp[a], poinsTmp[b] = poinsTmp[b], poinsTmp[a]
		break
	}
	return poinsTmp
}

func SimulateAnnealing(point []PointD, initialTemp, finalTemp, alpha float64, maxIterations int) Paths {
	var (
		currentSolution = point
		currentTemp     = initialTemp
	)

	for currentTemp > finalTemp {
		log.Println("Epoch", currentTemp)
		for i := 0; i <= maxIterations; i++ {
			var (
				newSolution     = Swap(currentSolution)
				currentDistance = EvaluatePointsV2(currentSolution)
				newDistance     = EvaluatePointsV2(newSolution)
			)

			if newDistance < currentDistance || math.Exp((currentDistance-newDistance)/currentTemp) > rand.Float64() {
				currentSolution = newSolution
			}
		}
		currentTemp *= alpha
	}

	return EvaluatePoints(currentSolution)
}
