package heap

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	heapSize int
	data     []T
	k        int
}

func NewHeap[T constraints.Ordered](size int, k int) *Heap[T] {
	return &Heap[T]{
		heapSize: size - 1,
		data:     make([]T, size, size),
		k:        k,
	}
}

func NewHeapWithData[T constraints.Ordered](data []T, k int) *Heap[T] {
	return &Heap[T]{
		heapSize: len(data) - 1,
		data:     data,
		k:        k,
	}
}

func (h *Heap[T]) GetData() []T {
	return h.data
}

func (h *Heap[T]) BuildMaxHeap(n int) {
	h.heapSize = n
	log.Println("heap size: ", h.heapSize)
	log.Println("Data size: ", len(h.data))

	for i := h.heapSize / h.k; i >= 0; i-- {
		log.Println("Analyzing: ", i, h.data[i])
		h.MaxHeapify(i)
	}
}

func (h *Heap[T]) MaxHeapify(i int) {
	var (
		leftLimit  = h.left(i)
		rightLimit = h.right(i)
		largest    = i
	)

	log.Println("------ ", i, h.data[i], " ------")

	if leftLimit <= h.heapSize {
		log.Println("left limit: ", leftLimit, h.data[leftLimit])
	} else {
		log.Println("left limit is out of range")
		return
	}

	if rightLimit <= h.heapSize {
		log.Println("right limit: ", rightLimit, h.data[rightLimit])
	} else {
		log.Println("right limit is out of range, as the left is inside range, then setting to the upper limit heap size")
		rightLimit = h.heapSize
	}

	for index := leftLimit; index <= rightLimit; index++ {
		log.Println(index, h.data[index], "VS", h.data[largest])
		if h.data[index] > h.data[largest] {
			log.Println(h.data[index], ">", h.data[largest])
			largest = index
		}

	}

	// time.Sleep(1 * time.Second)
	if largest != i {
		log.Println("Swap")
		log.Println(h.data[i], h.data[largest])
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		log.Println("Done")
		h.MaxHeapify(largest)
	}
}

func (h *Heap[T]) ToString() string {
	return fmt.Sprintf("%+v", h.data)
}

func (h *Heap[T]) parent(i int) int {
	if i == 0 {
		return i
	}

	return (i - 1) / h.k
}

func (h *Heap[T]) left(i int) int {
	return h.k*i + 1
}

func (h *Heap[T]) right(i int) int {
	return h.k*i + h.k
}

func (h *Heap[T]) HeapSort(n int) {
	h.heapSize = n
	h.BuildMaxHeap(n)

	for i := h.heapSize; i > 0; i-- {
		log.Println("HeapSort: ", i, h.data[i])
		h.data[0], h.data[i] = h.data[i], h.data[0]
		h.heapSize = h.heapSize - 1
		h.MaxHeapify(0)
	}
}
