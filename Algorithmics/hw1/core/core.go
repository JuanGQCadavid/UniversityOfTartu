package core

import (
	"gonum.org/v1/gonum/stat"
)

const (
	EPOCHS int = 10
)

type DataType string

const (
	INT8  DataType = "INT8"
	INT16 DataType = "INT16"
	INT32 DataType = "INT32"
	INT64 DataType = "INT64"
)

type RandSResult struct {
	NumberOfElements int64
	Epoch            int
	StatsPerEpoch    []RandStats
	Summary          RandStats
}

var (
	DataCases []DataType = []DataType{
		INT8,
		INT16,
		INT32,
		INT64,
	}
)

func ResumedEpochs(stats []RandStats) RandStats {

	var avgSlope, avgNegatives, avgPositives float64
	var timingClokcs []float64 = make([]float64, 0, len(stats))

	for _, stat := range stats {
		avgSlope += stat.Avg
		avgNegatives += stat.NegativePercentage
		avgPositives += stat.PositivesPercentage
		timingClokcs = append(timingClokcs, stat.Timing)
	}
	avgSlope = avgSlope / float64(len(stats))
	avgNegatives = avgNegatives / float64(len(stats))
	avgPositives = avgPositives / float64(len(stats))

	mean := stat.Mean(timingClokcs, nil)
	stdDev := stat.StdDev(timingClokcs, nil)
	cov := (stdDev / mean) * 100

	return RandStats{
		Avg:                          avgSlope,
		Timing:                       mean,
		NegativePercentage:           avgNegatives,
		PositivesPercentage:          avgPositives,
		TimingStdDev:                 stdDev,
		TimingCoefficientOfVariation: cov,
	}

}
