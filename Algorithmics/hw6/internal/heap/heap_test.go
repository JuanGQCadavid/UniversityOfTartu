package heap

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	// Note: the heap data list behaves as a bfs on the data of a binary search tree
	heapy := NewHeapWithData[int]([]int{
		0,
		16,
		4,
		10,
		14,
		7,
		9,
		3,
		2,
		8,
		1,
	})
	heapy.MaxHeapify(2)

	log.Println(heapy.ToString())

	maxHeap := []int{
		0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
	}
	data := heapy.GetData()
	assert.Equal(t, maxHeap, data)

}

func TestBuild(t *testing.T) {
	// Note: the heap data list behaves as a bfs on the data of a binary search tree
	heapy := NewHeapWithData[int]([]int{
		0, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	})

	heapy.BuildMaxHeap(len(heapy.data) - 1)

	maxHeap := []int{
		0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
	}

	log.Println(heapy.ToString())
	assert.Equal(t, maxHeap, heapy.data)
}

func TestSort(t *testing.T) {
	// Note: the heap data list behaves as a bfs on the data of a binary search tree
	heapy := NewHeapWithData[int]([]int{
		0, 4, 1, 3, 2, 16, 9, 10, 14, 8, 7,
	})

	heapy.BuildMaxHeap(heapy.heapSize)

	maxHeap := []int{
		0, 16, 14, 10, 8, 7, 9, 3, 2, 4, 1,
	}

	log.Println(heapy.ToString())
	assert.Equal(t, maxHeap, heapy.data)
	heapy.HeapSort(heapy.heapSize)
	log.Println(heapy.ToString())
}
