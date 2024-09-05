package main

import (
	"fmt"
	"hw1/core"
	"math"
)

type DataType string

const (
	INT8  DataType = "INT8"
	INT16 DataType = "INT16"
	INT32 DataType = "INT32"
	INT64 DataType = "INT64"
)

const (
	EPOCHS int = 10
)

type RandSResult struct {
	NumberOfElements int64
	Epoch            int
	StatsPerEpoch    []core.RandStats
	Summary          core.RandStats
}

var (
	testCases []int64 = []int64{
		int64(math.Pow10(3)),
		// int64(math.Pow10(4)),
		// int64(math.Pow10(5)),
		// int64(math.Pow10(6)),
		// int64(math.Pow10(7)),
		// int64(math.Pow10(8)),
		// int64(math.Pow10(9)),
		// int64(math.Pow10(10)),
		// int64(math.Pow10(11)),
		// int64(math.Pow10(12)),
	}
	dataCases []DataType = []DataType{
		INT8,
		INT16,
		INT32,
		INT64,
	}
)

func main() {
	var results = make(map[int64]map[DataType]*RandSResult)
	for _, tt := range testCases {
		fmt.Println(tt)
		results[tt] = runRand(tt)
	}

	for result := range results {
		fmt.Println("Results for", result)
		for key, val := range results[result] {
			fmt.Printf("%s: %+v\n", key, val)
		}
	}
}

func runRand(n int64) map[DataType]*RandSResult {
	result := make(map[DataType]*RandSResult)
	for _, dType := range dataCases {
		var randSResult RandSResult = RandSResult{
			NumberOfElements: n,
			Epoch:            EPOCHS,
		}
		switch dType {
		case INT8:
			randSResult.StatsPerEpoch = runEpochsRand[int8](n, EPOCHS)
		case INT16:
			randSResult.StatsPerEpoch = runEpochsRand[int16](n, EPOCHS)
		case INT32:
			randSResult.StatsPerEpoch = runEpochsRand[int32](n, EPOCHS)
		case INT64:
			randSResult.StatsPerEpoch = runEpochsRand[int64](n, EPOCHS)
		}
		result[dType] = &randSResult
	}
	return result
}

func runEpochsRand[V int8 | int16 | int32 | int64](n int64, epochs int) []core.RandStats {
	var randStats = make([]core.RandStats, epochs)
	for i := 0; i < epochs; i++ {
		val := core.GenerateRandListG[V](n)
		checks := core.CheckAvgDistanceG[V](val)
		positiveGap, negativeGap := core.CheckDistributionOfPositivesAndNegatiesG[V](val)

		// Adding Gaps
		checks.NegativePercentage = negativeGap
		checks.PositivesPercentage = positiveGap

		randStats[i] = checks
	}

	return randStats
}
