package main

import (
	"hw6/internal/domain"
	"hw6/internal/generators"
	"hw6/internal/heap"
	"log"
	"runtime"
	"time"
)

const (
	upperLimit int = 999999999
	// maxValuesSizes int = 1000000
	maxValuesSizes int = 5000 //100000
	batchs         int = 1000
)

func main() {
	var (
		vals = generators.GenerateRanList(int32(maxValuesSizes), int32(upperLimit))
		ks   = []int{
			2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		}

		results = make([]domain.BatchReport, 0, maxValuesSizes/batchs)
	)
	for _, k := range ks {
		var (
			heapifyTimes     = make([]float64, maxValuesSizes/batchs)
			batchIndex   int = 1
			bubbleTimes      = make([]float64, maxValuesSizes/batchs)
			nSizes           = make([]int, maxValuesSizes/batchs)
		)
		for ; batchIndex <= maxValuesSizes/batchs; batchIndex += 1 {
			// Call for garbage collector
			runtime.GC()
			var (
				low int = 0
				up  int = batchs * batchIndex
			)

			if up >= len(vals) {
				up = len(vals)
			}

			log.Println("K: ", k, "BatchId: ", batchIndex, "Upper: ", up, " Lower: ", low)
			heapifyA, bubbleSortA := make([]int32, up), make([]int32, up)

			_ = copy(heapifyA, vals[low:up])
			_ = copy(bubbleSortA, vals[low:up])

			heapA := heap.NewHeapWithData(heapifyA, k)
			startT := time.Now()
			heapA.HeapSort(heapA.GetHeapSize() - 1)
			endTime := time.Since(startT)
			heapifyTimes[batchIndex-1] = endTime.Seconds()

			startT = time.Now()
			buResult := bubbleSort(bubbleSortA)
			endTime = time.Since(startT)
			bubbleTimes[batchIndex-1] = endTime.Seconds()

			nSizes[batchIndex-1] = up

			for i, val := range heapifyA {
				if buResult[i] != val {
					panic("Shit, they are differents")
				}
			}
		}
		results = append(results, domain.BatchReport{
			HeapifyTimes:    heapifyTimes,
			BubbleSortTimes: bubbleTimes,
			NSizes:          nSizes,
			K:               k,
		})
		// hERE
	}
	generators.GeneratePlots(results)
	generators.FromGoToPythonImage(results)
	generators.SaveGenStatsToCSV(generators.GenerateStats(results))
}

func bubbleSort(arr []int32) []int32 {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
