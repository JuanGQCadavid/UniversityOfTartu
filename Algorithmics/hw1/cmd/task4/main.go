package main

import (
	"fmt"
	"hw1/core"
	"math"
	"runtime"
	"sort"
	"time"
)

var (
	TestCases []int64 = []int64{
		int64(math.Pow10(3)),
		int64(math.Pow10(4)),
		int64(math.Pow10(5)),
		int64(math.Pow10(6)),
		// int64(math.Pow10(7)),
		// int64(math.Pow10(8)),
		// int64(math.Pow10(9)),
	}
)

func main() {
	var results = make(map[int64]map[core.DataType]*core.RandSResult)
	for _, tt := range TestCases {
		includeNegatives := true
		runtime.GC()
		start := time.Now()
		results[tt] = runRand(tt, includeNegatives)
		timing := time.Since(start) // monotonic clock
		fmt.Println(tt, timing)
	}

	for result := range results {
		fmt.Println("Results for", result)
		for key, val := range results[result] {
			fmt.Printf("%s: %+v\n", key, val.Summary)
		}
	}

	// core.GeneratePlots(results)
	// core.GenerateSummaryText(results)
	core.GenerateFunction(results)
}

func runRand(n int64, includeNegatives bool) map[core.DataType]*core.RandSResult {
	result := make(map[core.DataType]*core.RandSResult)
	for _, dType := range core.DataCases {
		var randSResult core.RandSResult
		switch dType {
		case core.INT8:
			randSResult = runEpochsRand[int8](n, core.EPOCHS, includeNegatives)
		case core.INT16:
			randSResult = runEpochsRand[int16](n, core.EPOCHS, includeNegatives)
		case core.INT32:
			randSResult = runEpochsRand[int32](n, core.EPOCHS, includeNegatives)
		case core.INT64:
			randSResult = runEpochsRand[int64](n, core.EPOCHS, includeNegatives)
		}
		result[dType] = &randSResult
	}
	return result
}

// Here we should add the logic for the timer and the resumed of the randStats
func runEpochsRand[V int8 | int16 | int32 | int64](n int64, epochs int, includeNegatives bool) core.RandSResult {
	var randStats = make([]core.RandStats, epochs)

	for i := 0; i < epochs; i++ {

		val := core.GenerateRandListG[V](n, includeNegatives)

		start := time.Now()
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
		timing := time.Since(start).Seconds() // monotonic clock

		checks := core.CheckAvgDistanceG[V](val)

		// Adding timing
		checks.Timing = timing
		randStats[i] = checks
	}

	return core.RandSResult{
		NumberOfElements: n,
		Epoch:            epochs,
		StatsPerEpoch:    randStats,
		Summary:          core.ResumedEpochs(randStats),
	}
}
