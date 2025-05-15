package core

import "math"

func KNN(points []PointD) []PointD {
	var (
		result = make([]PointD, len(points))
		tmp    = make([]PointD, len(points))
	)

	copy(tmp, points)

	result[0] = tmp[0]
	tmp = tmp[1:]

	for idx := 1; idx < len(result); idx++ {
		closestIdx := -1
		closestDist := math.Inf(1)

		for jdx, point := range tmp {
			if EucledianDist(result[idx-1], point) < closestDist {
				closestDist = EucledianDist(result[idx-1], point)
				closestIdx = jdx
			}
		}

		tmp[0], tmp[closestIdx] = tmp[closestIdx], tmp[0]
		result[idx] = tmp[0]
		tmp = tmp[1:]
	}
	return result
}
