package main

import (
	"bufio"
	"fmt"
	"hw1/core"
	"math"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot/plotter"
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
		int64(math.Pow10(4)),
		int64(math.Pow10(5)),
		// int64(math.Pow10(6)),
		// int64(math.Pow10(7)),
		// int64(math.Pow10(8)),

		// Well this one at lest finalized
		// int64(math.Pow10(9)),
		// int64(math.Pow10(10)), // This one is on my cp limits
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

	generatePlots(results)
	generateNegativePositivePlots(results)
	generateSummaryText(results)
}

func generateSummaryText(results map[int64]map[DataType]*RandSResult) {
	f, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	historicalFile, err := os.Create("historical.csv")
	if err != nil {
		panic(err)
	}
	defer historicalFile.Close()

	w := bufio.NewWriter(f)
	w.WriteString("N; Data type; Avg slope; Avg Negatives; Avg Positives; StdDe Time; Mean Time; Cov Time \n")

	historicalFileBuffer := bufio.NewWriter(historicalFile)
	historicalFileBuffer.WriteString("N; Data type; Epoch; Timing\n")

	for n, result := range results {
		for dataType, randResult := range result {

			newLine := fmt.Sprintf(
				"%d; %s; %f; %f; %f; %f; %f; %f\n",
				n,
				string(dataType),
				randResult.Summary.Avg,
				randResult.Summary.NegativePercentage,
				randResult.Summary.PositivesPercentage,
				randResult.Summary.TimingStdDev,
				randResult.Summary.Timing,
				randResult.Summary.TimingCoefficientOfVariation,
			)

			newLine = strings.ReplaceAll(newLine, ".", ",")
			w.WriteString(
				newLine,
			)

			for i, statsPerEpoch := range randResult.StatsPerEpoch {
				newLine := fmt.Sprintf(
					"%d; %s; %d; %f\n",
					n,
					string(dataType),
					i,
					statsPerEpoch.Timing,
				)

				newLine = strings.ReplaceAll(newLine, ".", ",")

				historicalFileBuffer.WriteString(
					newLine,
				)
			}

		}
	}

	w.Flush()
	historicalFileBuffer.Flush()

}

func generateNegativePositivePlots(results map[int64]map[DataType]*RandSResult) {

	for i, testCases := range results {
		var (
			name           string   = fmt.Sprintf("Negative and positives distribution for N=%d", i)
			xNames         []string = make([]string, 0, len(testCases))
			yLabel         string   = "Percentage"
			dimensionNames []string = []string{
				"Negative",
				"Positive",
			}
			negativeDimension []float64 = make([]float64, 0, len(testCases))
			positiveDimension []float64 = make([]float64, 0, len(testCases))
		)

		for ii, theCase := range testCases {
			xNames = append(xNames, string(ii))
			negativeDimension = append(negativeDimension, theCase.Summary.NegativePercentage)
			positiveDimension = append(positiveDimension, theCase.Summary.PositivesPercentage)
		}

		fmt.Println("Printing: ", i)
		core.GenerateBarPlot(name, yLabel, xNames, dimensionNames, negativeDimension, positiveDimension)
	}
}

func getSortedKeys(results map[int64]map[DataType]*RandSResult) []int64 {
	keys := make([]int64, 0, len(results))

	for key := range results {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return i < j
	})

	return keys
}

func generatePlots(results map[int64]map[DataType]*RandSResult) {

	keys := getSortedKeys(results)

	toPlot := make(map[string]plotter.XYs)
	index := 0

	fmt.Printf("%v\n", keys)

	chunks := [][]int64{
		// keys,
		keys[0 : len(keys)/2],
		keys[len(keys)/2:],
	}

	for _, chunk := range chunks {
		fmt.Printf("%v\n", chunk)
		for _, n := range chunk {
			result := results[n]
			for dataType, randResult := range result {
				if len(toPlot[string(dataType)]) == 0 {
					toPlot[string(dataType)] = make(plotter.XYs, len(results))
				}

				toPlot[string(dataType)][index].X = float64(n)
				toPlot[string(dataType)][index].Y = randResult.Summary.Timing

			}
			index++
		}
		core.GenerateLinePlot(fmt.Sprintf("N %v vs Time per data type", chunk), "N", "Time in seconds", toPlot)

	}

}

func runRand(n int64, includeNegatives bool) map[DataType]*RandSResult {
	result := make(map[DataType]*RandSResult)
	for _, dType := range dataCases {
		var randSResult RandSResult
		switch dType {
		case INT8:
			randSResult = runEpochsRand[int8](n, EPOCHS, includeNegatives)
		case INT16:
			randSResult = runEpochsRand[int16](n, EPOCHS, includeNegatives)
		case INT32:
			randSResult = runEpochsRand[int32](n, EPOCHS, includeNegatives)
		case INT64:
			randSResult = runEpochsRand[int64](n, EPOCHS, includeNegatives)
		}
		result[dType] = &randSResult
	}
	return result
}

// Here we should add the logic for the timer and the resumed of the randStats
func runEpochsRand[V int8 | int16 | int32 | int64](n int64, epochs int, includeNegatives bool) RandSResult {
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

	return RandSResult{
		NumberOfElements: n,
		Epoch:            epochs,
		StatsPerEpoch:    randStats,
		Summary:          resumedEpochs(randStats),
	}
}

func resumedEpochs(stats []core.RandStats) core.RandStats {

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

	return core.RandStats{
		Avg:                          avgSlope,
		Timing:                       mean,
		NegativePercentage:           avgNegatives,
		PositivesPercentage:          avgPositives,
		TimingStdDev:                 stdDev,
		TimingCoefficientOfVariation: cov,
	}

}
