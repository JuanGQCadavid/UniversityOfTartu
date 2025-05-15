package main

import (
	"hw10/core"
	"log"
	"sync"
	"time"
)

const (
	POINTS_30   = "./data/hw10_tsp_points_data_0030.txt"
	POINTS_1000 = "./data/hw10_tsp_points_data_1000.txt"
)

var (
	FILES = []string{
		"./data/hw10_tsp_points_data_0030.txt",
		"./data/hw10_tsp_points_data_0100.txt",
		"./data/hw10_tsp_points_data_0200.txt",
		"./data/hw10_tsp_points_data_1000.txt",
	}
)

func main() {
	var (
		VERSION = "ve6"
	)
	// points := core.GetPointsFromFile(POINTS_1000)
	// log.Println(len(points))
	// myBest := core.MyBest(points, 90, 1e5)
	// core.Output("MY-BEST", VERSION, myBest)
	// log.Println("SWAP: ", myBest.Score, len(myBest.Points))

	points := core.GetPointsFromFile(POINTS_1000)
	log.Println(len(points))
	start := time.Now()
	myBest := core.MyBestBest(points, 1e6)
	end := time.Since(start)

	core.Output("MY-BEST", VERSION, myBest)
	log.Println("MY-BEST: ", myBest.Score, len(myBest.Points))

	log.Println("It tooks: ", end)

}

func TASK2() {

	var (
		VERSION = "v3"
		wg      = sync.WaitGroup{}
	)

	for _, file := range FILES {
		wg.Add(1)
		go func(fileToRun string) {
			defer wg.Done()

			points := core.GetPointsFromFile(file)
			log.Println(len(points))

			knnPoints := core.KNN(points)
			core.Output("KNN", VERSION, core.EvaluatePoints(knnPoints))

			sawp := core.RandoSwap(points, int(1e3))
			core.Output("KNN-SWAP", VERSION, sawp)

			initialTemp := 100.0
			finalTemp := 0.1
			alpha := 0.80
			maxIterations := int(1e6)
			burnThem := core.SimulateAnnealing(points, initialTemp, finalTemp, alpha, maxIterations)
			core.Output("KNN-BURN", VERSION, burnThem)

			log.Println("KNN: ", core.EvaluatePoints(knnPoints).Score)
			log.Println("SWAP: ", sawp.Score, len(sawp.Points))
			log.Println("BURN THEM: ", burnThem.Score, len(burnThem.Points))
		}(file)
	}

	wg.Wait()
}
