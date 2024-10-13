package generators

import (
	"fmt"
	"hw6/internal/domain"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)
func GeneratePlots(reports []domain.BatchReport) {
	for _, report := range reports {
		if len(report.HeapifyTimes) != len(report.NSizes) || len(report.BubbleSortTimes) != len(report.NSizes) {
			log.Printf("Mismatch in lengths for report with K=%d, skipping...", report.K)
			continue
		}

		// Plot both HeapifyTimes and BubbleSortTimes on the same chart
		err := PlotCombined(report)
		if err != nil {
			log.Printf("Error generating combined plot for K=%d: %v", report.K, err)
		}

		// Plot HeapifyTimes separately
		err = PlotSingle(report.NSizes, report.HeapifyTimes, fmt.Sprintf("HeapifyTimes_K%d.png", report.K), "Heapify Time")
		if err != nil {
			log.Printf("Error generating Heapify plot for K=%d: %v", report.K, err)
		}

		// Plot BubbleSortTimes separately
		err = PlotSingle(report.NSizes, report.BubbleSortTimes, fmt.Sprintf("BubbleSortTimes_K%d.png", report.K), "Bubble Sort Time")
		if err != nil {
			log.Printf("Error generating Bubble Sort plot for K=%d: %v", report.K, err)
		}
	}
}

func PlotCombined(report domain.BatchReport) error {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Heapify vs bubble sort (k=%d)", report.K)
	p.X.Label.Text = "N size"
	p.Y.Label.Text = "Time in seconds"

	heapifyPoints := makeXYs(report.NSizes, report.HeapifyTimes)
	bubbleSortPoints := makeXYs(report.NSizes, report.BubbleSortTimes)

  err := plotutil.AddLinePoints(p,
		"Heapify", heapifyPoints,
		"Bubble", bubbleSortPoints)
	if err != nil {
		return err
	}

	// Save the plot
	return p.Save(8*vg.Inch, 4*vg.Inch, fmt.Sprintf("Combined_K%d.png", report.K))
}

func PlotSingle(x []int, y []float64, filename, title string) error {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = "N size"
	p.Y.Label.Text = "Time in seconds"

	points := makeXYs(x, y)

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	p.Add(line)

	return p.Save(8*vg.Inch, 4*vg.Inch, filename)
}

func makeXYs(x []int, y []float64) plotter.XYs {
	points := make(plotter.XYs, len(x))
	for i := range x {
		points[i].X = float64(x[i])
		points[i].Y = y[i]
	}
	return points
}

// func main() {
// 	// Example usage
// 	reports := []domain.BatchReport{
// 		{
// 			HeapifyTimes:    []float64{0.1, 0.3, 0.5, 0.8},
// 			BubbleSortTimes: []float64{0.2, 0.4, 0.6, 1.0},
// 			NSizes:          []int{10, 50, 100, 200},
// 			K:               1,
// 		},
// 		// Add more reports as needed
// 	}
//
// 	generatePlots(reports)
// }
