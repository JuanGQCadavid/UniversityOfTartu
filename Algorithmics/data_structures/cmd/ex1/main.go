package main

import (
	"log"

	"github.com/JuanGQCadavid/UniversityOfTartu/Algorithmics/data_structures/trees/binarysearch"
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
	// log.Println("Heigh: ", stats.GetDepth(bst.Root))
	// log.Println("Width: ", stats.GetWidth(bst.Root))

}

func nextCollatz(n int) int {
	if n%2 == 0 {
		return n / 2
	}
	return 3*n + 1
}
