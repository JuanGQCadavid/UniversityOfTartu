package main

import (
	"encoding/csv"
	"fmt"
	"hw1/core"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var (
	testDataTypes = []core.DataType{
		core.INT8,
		core.INT64,
	}
	futureTestCases []int = []int{

		12,
		13,
		14,
	}
)

func main() {

	for _, dataType := range testDataTypes {
		fmt.Println(dataType)

		xs, ys := getXsAndYs("refression/data_plot.csv", dataType)

		fmt.Println(xs)
		fmt.Println(ys)

		w0, w1 := getFunction(xs, ys)

		fmt.Println(w0)
		fmt.Println(w1)

		y2 := make([]float64, len(xs))

		for index, value := range xs {
			y2[index] = getY2(w0, w1, value)
		}

		plotFunctions(xs, ys, y2, dataType)
		fmt.Println(calculateMSE(xs, ys, y2))

		for _, b := range futureTestCases {
			forcasting := getY2(w0, w1, float64(math.Pow10(b)))

			days := forcasting / (60 * 60 * 24)
			fmt.Printf("10^%d = %f\n", b, days)
		}
	}

}

func calculateMSE(xs, ys, y2s []float64) float64 {
	var mse float64
	for index := range xs {
		mse += math.Pow((ys[index] - y2s[index]), 2)
	}

	return mse / float64(len(xs))
}

func plotFunctions(xs, ys, y2s []float64, dataType core.DataType) {
	pointsOriginal := make(plotter.XYs, len(xs))
	pointsY2 := make(plotter.XYs, len(xs))

	for index := range xs {
		pointsOriginal[index] = plotter.XY{
			X: xs[index],
			Y: ys[index],
		}
		pointsY2[index] = plotter.XY{
			X: xs[index],
			Y: y2s[index],
		}
	}

	p := plot.New()
	p.Title.Text = "N Size VS Time sorting"
	p.X.Label.Text = "N Size"
	p.Y.Label.Text = "Time in secods"
	p.Add(plotter.NewGrid())

	err := plotutil.AddLinePoints(
		p,
		"Original", pointsOriginal,
		"Linear Regression", pointsY2,
	)

	checkErr(err)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("points-for-%s.png", string(dataType))); err != nil {
		panic(err)
	}

}

func getFunction(xs []float64, ys []float64) (float64, float64) {

	xsM := stat.Mean(xs, nil)
	ysM := stat.Mean(ys, nil)

	var x_less_xM_pow_2, x_less_xM_times_y_less_yM float64
	for index := range xs {
		x_less_xM := xs[index] - xsM
		y_less_yM := ys[index] - ysM

		x_less_xM_pow_2 += math.Pow(x_less_xM, 2)
		x_less_xM_times_y_less_yM += x_less_xM * y_less_yM
	}

	w1 := x_less_xM_times_y_less_yM / x_less_xM_pow_2
	w0 := ysM / (w1 * xsM)

	return w0, w1

}

func getY2(w0, w1, x float64) float64 {
	return w0 + w1*x
}

func getXsAndYs(filename string, dataType core.DataType) ([]float64, []float64) {

	file, err := os.Open(filename)
	checkErr(err)

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = rune(';')
	records, er := reader.ReadAll()
	checkErr(er)

	var (
		xs = make([]float64, 0, len(records))
		ys = make([]float64, 0, len(records))
	)

	for _, eachrecord := range records {
		if strings.Contains(eachrecord[1], string(dataType)) {
			x, err := strconv.ParseFloat(eachrecord[0], 64)
			checkErr(err)
			xs = append(xs, x)

			y, err := strconv.ParseFloat(strings.ReplaceAll(eachrecord[3], ",", "."), 64)
			checkErr(err)
			ys = append(ys, y)
		}

	}

	return xs, ys
}
