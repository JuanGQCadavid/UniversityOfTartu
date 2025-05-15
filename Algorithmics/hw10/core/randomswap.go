package core

import (
	"math"
	"math/rand"
)

func RandoSwap(points []PointD, times int) Paths {
	var (
		poinsTmp = make([]PointD, len(points))
	)

	copy(poinsTmp, points)

	for i := 0; i < times; i++ {
		prePath := EvaluatePointsV2(poinsTmp)
		// prePoints := make([]PointD, len(points))
		// copy(newPoints, prePath.Points)
		a, b := rand.Intn(len(points)), rand.Intn(len(points))

		if a == b {
			continue
		}

		poinsTmp[a], poinsTmp[b] = poinsTmp[b], poinsTmp[a]
		possiblePath := EvaluatePointsV2(poinsTmp)

		if possiblePath > prePath {
			poinsTmp[a], poinsTmp[b] = poinsTmp[b], poinsTmp[a]
		}
		// else {
		// 	log.Println(possiblePath)
		// }
	}

	return EvaluatePoints(poinsTmp)

}

func EvaluatePointsV2(points []PointD) float64 {
	dist := 0.0
	for i := 0; i < len(points)-1; i++ {
		dist += EucledianDist(points[i], points[i+1])
	}

	return dist
}

func EucledianDistV2(a PointD, b PointD) float64 {
	return math.Sqrt(
		math.Pow(float64(a.X)-float64(b.X), 2) +
			math.Pow(float64(a.Y)-float64(b.Y), 2))
}
