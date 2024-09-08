package core

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"gonum.org/v1/plot/plotter"
)

func GeneratePlots(results map[int64]map[DataType]*RandSResult) {

	keys := getSortedKeys(results)
	fmt.Printf("%v\n", keys)
	chunks := [][]int64{
		keys,
		keys[0:1],
		keys[len(keys)-1:],
	}

	for _, chunk := range chunks {
		index := 0
		toPlot := make(map[string]plotter.XYs)
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
		GenerateLinePlot(fmt.Sprintf("N %v vs Time per data type", chunk), "N", "Time in seconds", toPlot)

	}

}

func GenerateNegativePositivePlots(results map[int64]map[DataType]*RandSResult) {

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
		GenerateBarPlot(name, yLabel, xNames, dimensionNames, negativeDimension, positiveDimension)
	}
}

func GenerateSummaryText(results map[int64]map[DataType]*RandSResult) {
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

func getSortedKeys(results map[int64]map[DataType]*RandSResult) []int64 {
	keys := make([]int64, 0, len(results))

	for key := range results {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}
