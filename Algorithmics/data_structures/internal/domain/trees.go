package domain

import "golang.org/x/exp/constraints"

type Node[K constraints.Ordered, D any] struct {
	Key           K
	SatalliteData D
	Color         string
	Parent        *Node[K, D]
	Left          *Node[K, D]
	Rigth         *Node[K, D]
}
