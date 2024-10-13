package generators

import (
	"encoding/csv"
	"fmt"
	"hw6/internal/domain"
	"os"
	"os/exec"
	"strconv"
)


const (
  PyFile string = "internal/generators/python_plotter.py"
)

func FromGoToPythonImage(reports []domain.BatchReport){
  for _, report := range(reports){
    // Save the report to CSV
    err := SaveToCSV(report)
    if err != nil {
      fmt.Println("Error saving CSV:", err)
      return
    }

    // Execute the Python script
    csvFile := fmt.Sprintf("batch_report_k%d.csv", report.K)
    err = executePythonScript(csvFile)
    if err != nil {
      fmt.Println("Error executing Python script:", err)
    }
  } 
}


func executePythonScript(csvFile string) error {
	// Command to run the Python script
	cmd := exec.Command("python3", PyFile, csvFile)

	// Execute the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}

	fmt.Println(string(output))
	return nil
}


func SaveToCSV(report domain.BatchReport) error {
	// Create a filename based on K
	filename := fmt.Sprintf("batch_report_k%d.csv", report.K)

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
	writer.Write([]string{"NSizes", "HeapifyTimes", "BubbleSortTimes"})

	// Write data rows
	for i := 0; i < len(report.NSizes); i++ {
		row := []string{
			strconv.Itoa(report.NSizes[i]),
			strconv.FormatFloat(report.HeapifyTimes[i], 'f', 6, 64),
			strconv.FormatFloat(report.BubbleSortTimes[i], 'f', 6, 64),
		}
		writer.Write(row)
	}

	fmt.Printf("Data saved to %s\n", filename)
	return nil
}
