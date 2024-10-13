package domain

type BatchReport struct {
	HeapifyTimes    []float64
	BubbleSortTimes []float64
	NSizes          []int
	K               int
}
