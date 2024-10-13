package main

import "hw6/internal/generators"

const (
	upperLimit     int32 = 999999999
	maxValuesSizes int32 = 1000000
	batchs         int32 = 1000
)

func main() {
	vals := generators.GenerateRanList(maxValuesSizes, upperLimit)

}
