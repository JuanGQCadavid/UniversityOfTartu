package generators

import (
	"encoding/csv"
	"fmt"
	"hw6/internal/domain"
	"os"
	"strconv"

	"gonum.org/v1/gonum/stat"
)

type GenStats struct {
	K               int
	MeanHeapify     float64
	VarianceHeapify float64
	MeanBubble      float64
	VarianceBubble  float64
}

func GenerateStats(reports []domain.BatchReport) []GenStats {
	genStats := make([]GenStats, len(reports))

	for i, report := range reports {
		genStats[i] = GenStats{
			K:               report.K,
			MeanBubble:      stat.Mean(report.BubbleSortTimes, nil),
			VarianceBubble:  stat.Variance(report.BubbleSortTimes, nil),
			MeanHeapify:     stat.Mean(report.HeapifyTimes, nil),
			VarianceHeapify: stat.Variance(report.HeapifyTimes, nil),
		}
	}

	return genStats
}

func SaveGenStatsToCSV(stats []GenStats) error {
	if len(stats) == 0 {
		return fmt.Errorf("no data to save")
	}

	// Use the K value from the first entry to create the filename
	filename := fmt.Sprintf("gen_stats_k%d.csv", stats[0].K)

	// Create CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write to CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"K", "MeanHeapify", "VarianceHeapify", "MeanBubble", "VarianceBubble"})

	// Write data rows
	for _, stat := range stats {
		row := []string{
			strconv.Itoa(stat.K),
			strconv.FormatFloat(stat.MeanHeapify, 'f', 6, 64),
			strconv.FormatFloat(stat.VarianceHeapify, 'f', 6, 64),
			strconv.FormatFloat(stat.MeanBubble, 'f', 6, 64),
			strconv.FormatFloat(stat.VarianceBubble, 'f', 6, 64),
		}
		writer.Write(row)
	}

	fmt.Printf("Data saved to %s\n", filename)
	return nil
}
