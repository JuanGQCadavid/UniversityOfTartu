package core

type PointD struct {
	X int16
	Y int16
}

type Paths struct {
	Points []PointD
	Score  float64
}
