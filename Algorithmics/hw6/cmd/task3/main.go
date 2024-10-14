package main

import (
	"fmt"
	"hw6/internal/heap"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

// ChatGPT made this function
func generateFibonacciMod(n int, mod int) []int {
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
		for i := 2; i < n; i++ {
			fib[i] = (fib[i-1] + fib[i-2]) % mod
		}
	}
	return fib
}

func generateRandomData(size int, mod int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(mod)
	}
	return data
}

type EpochResult struct {
	DataSize         int     `csv:"DataSize"`
	Mod              int     `csv:"Mod"`
	K                int     `csv:"K"`
	DataType         string  `csv:"DataType"`
	KthSmalles       int     `csv:"KthSmallest"`
	TimeTakenSeconds float64 `csv:"TimeTakenSeconds"`
}

func runExperiment(dataSize int, mod int, k int, dataType string) *EpochResult {
	var data []int
	if dataType == "random" {
		data = generateRandomData(dataSize, mod)
	} else if dataType == "fibonacci" {
		data = generateFibonacciMod(dataSize, mod)
	}

	heapA := heap.NewHeapWithData(data, k)
	startT := time.Now()
	heapA.HeapSort(heapA.GetHeapSize() - 1)
	duration := time.Since(startT)
	smallest := heapA.GetData()[k-1]

	fmt.Printf("Data Type: %s, Mod: %d, Data Size: %d, k: %d\n", dataType, mod, dataSize, k)
	fmt.Printf("k'th Smallest: %d, Time Taken: %v\n", smallest, duration)

	return &EpochResult{
		DataSize:         dataSize,
		Mod:              mod,
		K:                k,
		DataType:         dataType,
		KthSmalles:       smallest,
		TimeTakenSeconds: duration.Seconds(),
	}
}
func main() {
	dataSizes := []int{1000000, 100000000}   // 1M, 100M
	mods := []int{1000, 1000000, 1000000000} // 1k, 1M, 100M
	ks := []int{100, 10000, 1000000}         // 100, 10k, 1M
	// types := []string{"random", "fibonacci"}
	types := []string{"fibonacci"}
	results := make([]*EpochResult, 18)
	for _, dataType := range types {
		for _, size := range dataSizes {
			for _, mod := range mods {
				for _, k := range ks {
					results = append(results, runExperiment(size, mod, k, dataType))
					a := time.Now().Format(time.DateTime)
					b := strings.Replace(a, " ", "_", 1)
					fmt.Println(b)
					FromMatrixToFile(results, "Fib_"+b+".csv")
				}
			}
		}
	}

}

func FromMatrixToFile(matrix []*EpochResult, a string) {
	gFile, err := os.OpenFile(a, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	if err = gocsv.MarshalFile(matrix, gFile); err != nil {
		panic(err.Error())
	}
}
