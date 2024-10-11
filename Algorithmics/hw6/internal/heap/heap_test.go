package heap

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMaxHeap(t *testing.T) {
	// Note: the heap data list behaves as a bfs on the data of a binary search tree

	var test = []struct {
		name   string
		k      int
		maxOn  int
		heap   []int
		result []int
	}{
		{
			name:  "No more childers",
			k:     3,
			maxOn: 1,
			heap: []int{
				16, 4, 10, 14, 7, 9, 3, 2, 8, 1,
			},
			result: []int{
				16, 9, 10, 14, 7, 4, 3, 2, 8, 1,
			},
		},
		{
			name:  "Two children instead of 3",
			k:     3,
			maxOn: 1,
			heap: []int{
				16, 4, 10, 14, 7, 9, 3, 2, 8, 1, 5, 5, 5, 5, 5, 5, 5, 5,
			},
			result: []int{
				16, 9, 10, 14, 7, 5, 3, 2, 8, 1, 5, 5, 5, 5, 5, 5, 4, 5,
			},
		},
		{
			name:  "Normal case",
			k:     2,
			maxOn: 1,
			heap: []int{
				16, 4, 10, 14, 7, 9, 3, 2, 8, 1,
			},
			result: []int{
				16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			heapy := NewHeapWithData[int](tt.heap, tt.k)
			heapy.MaxHeapify(tt.maxOn)
			log.Println(heapy.ToString())
			data := heapy.GetData()
			assert.Equal(t, tt.result, data)
		})
	}
}

func TestKBuildHeap(t *testing.T) {
	// Note: the heap data list behaves as a bfs on the data of a binary search tree

	var test = []struct {
		name   string
		k      int
		heap   []int
		result []int
	}{
		{
			name: "Normal case",
			k:    2,
			heap: []int{
				4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
			},
			result: []int{
				16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
			},
		},
		{
			name: "Normal case with three children",
			k:    3,
			heap: []int{
				4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
			},
			result: []int{
				16, 10, 14, 2, 1, 9, 4, 3, 8, 7,
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			heapy := NewHeapWithData[int](tt.heap, tt.k)
			heapy.BuildMaxHeap(heapy.heapSize)
			log.Println(heapy.ToString())
			data := heapy.GetData()
			assert.Equal(t, tt.result, data)
		})
	}
}

func TestKSortHeap(t *testing.T) {
	// Note: the heap data list behaves as a bfs on the data of a binary search tree

	var test = []struct {
		name   string
		k      int
		heap   []int
		result []int
	}{
		{
			name: "Normal case",
			k:    2,
			heap: []int{
				4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
			},
			result: []int{
				1, 2, 3, 4, 7, 8, 9, 10, 14, 16,
			},
		},
		{
			name: "Normal case with three children",
			k:    3,
			heap: []int{
				4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
			},
			result: []int{
				1, 2, 3, 4, 7, 8, 9, 10, 14, 16,
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			heapy := NewHeapWithData[int](tt.heap, tt.k)
			heapy.HeapSort(heapy.heapSize)
			log.Println(heapy.ToString())
			data := heapy.GetData()
			assert.Equal(t, tt.result, data)
		})
	}
}
