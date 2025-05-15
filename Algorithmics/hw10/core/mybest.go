package core

import (
	"log"
	"math/rand"
)

func MyBest(points []PointD, kIteration int, iterations int) Paths {
	var (
		tested     = make([]bool, len(points))
		pointsTemp = make([]PointD, len(points))
	)

	copy(pointsTemp, points)

	for i := 0; i < kIteration; i++ {
		log.Println("Iteration: ", i)
		index := rand.Intn(len(points))
		for {
			if !tested[index] {
				break
			}
			index = rand.Intn(len(points))
		}

		before := EvaluatePointsV2(pointsTemp)
		log.Println("Baseline: ", before)
		pointsTemp[0], pointsTemp[index] = pointsTemp[index], pointsTemp[0]
		newCandidate := KNN(points)
		newCandidate = RandoSwap(newCandidate, iterations).Points
		pointsTemp[0], pointsTemp[index] = pointsTemp[index], pointsTemp[0]

		if EvaluatePointsV2(newCandidate) < before {
			log.Println("Update:", EvaluatePointsV2(newCandidate))
			pointsTemp = newCandidate
		}
		tested[index] = true
	}
	return EvaluatePoints(pointsTemp)
}
