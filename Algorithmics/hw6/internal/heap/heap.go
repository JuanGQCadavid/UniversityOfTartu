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

func NewHeap[T constraints.Ordered](size int) *Heap[T] {
	return &Heap[T]{
		heapSize: size - 1,
		data:     make([]T, size+1, size+1),
		k:        2,
	}
}

func NewHeapWithData[T constraints.Ordered](data []T) *Heap[T] {
	return &Heap[T]{
		heapSize: len(data) - 1,
		data:     data,
		k:        2,
	}
}

func (h *Heap[T]) GetData() []T {
	return h.data
}

func (h *Heap[T]) BuildMaxHeap(n int) {
	h.heapSize = n
	log.Println("heap size: ", h.heapSize)
	log.Println("Data size: ", len(h.data))

	for i := h.heapSize / 2; i > 0; i-- {
		log.Println("Analyzing: ", i, h.data[i])
		h.MaxHeapify(i)
	}
}

func (h *Heap[T]) MaxHeapify(i int) {
	var (
		l       = h.left(i)
		r       = h.right(i)
		largest = i
	)
	log.Println("------ ", i, h.data[i], " ------")

	if l <= h.heapSize {
		log.Println("left: ", l, h.data[l])
	}

	if r <= h.heapSize {
		log.Println("right: ", r, h.data[r])
	}

	if l <= h.heapSize && h.data[l] > h.data[i] {
		log.Println("left: ", h.data[l], ">", h.data[i])
		largest = l
	}

	if r <= h.heapSize && h.data[r] > h.data[largest] {
		log.Println("rigth: ", h.data[r], ">", h.data[largest])
		largest = r
	}

	if largest != i {
		log.Println("Swap")
		log.Println(h.data[i], h.data[largest])
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		log.Println(h.data[i], h.data[largest])
		h.MaxHeapify(largest)
	}
}

func (h *Heap[T]) ToString() string {
	return fmt.Sprintf("%+v", h.data)
}

func (h *Heap[T]) parent(i int) int {
	// (i-1) / h.k
	return i / h.k
}

func (h *Heap[T]) left(i int) int {
	return i * h.k
}

func (h *Heap[T]) right(i int) int {
	// K - i*k + 1
	return i*h.k + 1
}
