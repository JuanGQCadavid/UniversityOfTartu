package core

import (
	"log"
	"sync"
)

func MyBestBest(points []PointD, iterations int) Paths {
	var (
		wg         = sync.WaitGroup{}
		groups     = 8
		lowerLimit = 0
		jump       = len(points) / groups
		ch         = make(chan Paths, groups)
	)

	for i := 0; i < groups; i++ {
		wg.Add(1)
		upperLimit := lowerLimit + jump

		if i == groups-1 {
			upperLimit = len(points)
		}

		log.Println(lowerLimit, upperLimit)

		go func(start, end int, points []PointD, iterations int, ch chan Paths) {
			defer wg.Done()
			ch <- DoWork(start, end, points, iterations)
		}(lowerLimit, upperLimit, points, iterations, ch)
		lowerLimit += jump
	}
	wg.Wait()
	close(ch)

	var better Paths

	for possible := range ch {
		if len(better.Points) == 0 {
			better = possible
		}

		if better.Score > possible.Score {
			log.Println("A better onw has being founded")
			better = possible
		}
	}

	return better

}

func DoWork(start, end int, points []PointD, iterations int) Paths {
	pointsTemp := make([]PointD, len(points))
	copy(pointsTemp, points)

	for i := start; i < end; i++ {
		log.Println("Iteration: ", i)
		before := EvaluatePointsV2(pointsTemp)
		pointsTemp[0], pointsTemp[i] = pointsTemp[i], pointsTemp[0]
		newCandidate := KNN(points)
		newCandidate = RandoSwap(newCandidate, iterations).Points
		pointsTemp[0], pointsTemp[i] = pointsTemp[i], pointsTemp[0]

		if EvaluatePointsV2(newCandidate) < before {
			pointsTemp = newCandidate
		}
	}

	return EvaluatePoints(pointsTemp)
}
