package main

import (
	"fmt"
	"hw1/core"
	"math"
)

type RandSResult struct {
	NumberOfElements int64
	Epoch            int64
	StatsPerEpoch    []core.RandStats
	Summary          core.RandStats
}

var (
	testCases []int64 = []int64{
		int64(math.Pow10(3)),
		int64(math.Pow10(4)),
		int64(math.Pow10(5)),
		int64(math.Pow10(6)),
		int64(math.Pow10(7)),
		int64(math.Pow10(8)),
		int64(math.Pow10(9)),
		int64(math.Pow10(10)),
		int64(math.Pow10(11)),
		int64(math.Pow10(12)),
	}
)

func main() {
	for tt := range testCases {
		fmt.Println(testCases[tt])
	}
}
