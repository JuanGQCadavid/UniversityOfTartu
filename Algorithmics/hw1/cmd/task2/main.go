package main

import (
	"fmt"
	"hw1/core"
	"math"
	"runtime"
	"time"
)

var (
	TestCases []int64 = []int64{
		int64(math.Pow10(3)),
		int64(math.Pow10(4)),
		int64(math.Pow10(5)),
		int64(math.Pow10(6)),
		int64(math.Pow10(7)),
		int64(math.Pow10(8)),

		// Well this one at lest finalized
		// int64(math.Pow10(9)),
		// int64(math.Pow10(10)), // This one is on my cp limits
		// int64(math.Pow10(11)),
		// int64(math.Pow10(12)),
	}
)

func main() {
	var results = make(map[int64]map[core.DataType]*core.RandSResult)
	for _, tt := range TestCases {
		includeNegatives := true
		fmt.Println(tt)
		runtime.GC() // We call GC before starting in order to clean the golang memory stack and have more free space
		results[tt] = runRand(tt, includeNegatives)
	}

	for result := range results {
		fmt.Println("Results for", result)
		for key, val := range results[result] {
			fmt.Printf("%s: %+v\n", key, val.Summary)
		}
	}

	core.GeneratePlots(results)
	core.GenerateNegativePositivePlots(results)
	core.GenerateSummaryText(results)
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

		start := time.Now()
		val := core.GenerateRandListG[V](n, includeNegatives)
		timing := time.Since(start).Seconds() // monotonic clock

		checks := core.CheckAvgDistanceG[V](val)
		positiveGap, negativeGap := core.CheckDistributionOfPositivesAndNegatiesG[V](val)

		// Adding Gaps
		checks.NegativePercentage = negativeGap
		checks.PositivesPercentage = positiveGap

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
