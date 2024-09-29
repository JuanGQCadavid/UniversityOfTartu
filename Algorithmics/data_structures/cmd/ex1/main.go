package main

import (
	"log"
	"sort"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/search"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/stats"
	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/visualizer"
)

func main() {
	start := 51
	results := make([]int, 10)
	bst := binarysearch.NewBSTree[int, int]()
	for n := start; n > 1; n = nextCollatz(n) {
		results = append(results, n)
		log.Println(n)
		bst.Append(n, n)
	}
	// log.Println(results)
	visualizer.FromBSTToPNG(bst, "ex1")
	log.Println("Heigh: ", stats.GetDepth(bst.Root))
	wid := make(map[int]int)
	stats.GetWidth(bst.Root, 1, wid)

	keys := make([]int, 0, len(wid))
	for key := range wid {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	log.Println(keys)
	log.Println("Width: ")
	for i := range keys {
		log.Printf("%d : %d\n", keys[i], wid[keys[i]])
	}

	log.Println("BFS: ", search.BreadthFirstSearchAll(bst.Root, bst.NodeCounter(bst.Root)))

}

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}
