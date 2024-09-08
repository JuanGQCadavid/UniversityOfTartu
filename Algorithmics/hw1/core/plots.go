package core

import (
	"fmt"
	"image/color"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	colors map[int][]color.RGBA = map[int][]color.RGBA{
		0: {
			color.RGBA{R: 255, G: 0, B: 0, A: 255},
			color.RGBA{R: 255, G: 128, B: 128, A: 255},
		},
		1: {
			color.RGBA{R: 0, G: 128, B: 0, A: 255},
			color.RGBA{R: 0, G: 255, B: 0, A: 255},
		},
		2: {
			color.RGBA{R: 0, G: 0, B: 255, A: 255},
			color.RGBA{R: 128, G: 128, B: 255, A: 255},
		},
		3: {
			color.RGBA{R: 255, G: 165, B: 0, A: 255},
			color.RGBA{R: 255, G: 223, B: 128, A: 255},
		},
		4: {
			color.RGBA{R: 128, G: 0, B: 128, A: 255},
			color.RGBA{R: 192, G: 128, B: 192, A: 255},
		},
		5: {
			color.RGBA{R: 0, G: 255, B: 255, A: 255},
			color.RGBA{R: 128, G: 255, B: 255, A: 255},
		},
		6: {
			color.RGBA{R: 255, G: 255, B: 0, A: 255},
			color.RGBA{R: 255, G: 255, B: 128, A: 255},
		},
		7: {
			color.RGBA{R: 165, G: 42, B: 42, A: 255},
			color.RGBA{R: 210, G: 105, B: 105, A: 255},
		},
		8: {
			color.RGBA{R: 75, G: 0, B: 130, A: 255},
			color.RGBA{R: 138, G: 43, B: 226, A: 255},
		},
		9: {
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 128, G: 128, B: 128, A: 255},
		},
	}
)

func GenerateLinePlot(title string, xLabel string, yLabel string, points map[string]plotter.XYs) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	p.Add(plotter.NewGrid())

	index := 0
	for label, dataPoints := range points {
		if index == len(colors) {
			index = 0
		}
		lpLine, lpPoints, err := plotter.NewLinePoints(dataPoints)
		if err != nil {
			panic(err.Error())
		}
		lpLine.Color = colors[index][0]
		lpPoints.Color = colors[index][1]

		p.Add(lpLine, lpPoints)
		p.Legend.Add(label, lpLine, lpPoints)

		index++

	}
	p.X.AutoRescale = true

	if err := p.Save(8*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s.png", strings.ReplaceAll(title, " ", "_"))); err != nil {
		panic(err)
	}

}

//
// There should be one item on dimensionNames per array added to dimensions
//

func GenerateBarPlot(title string, yLabel string, xNames []string, dimensionNames []string, dimensions ...[]float64) {
	p := plot.New()
	p.Title.Text = title
	p.Y.Label.Text = yLabel

	index := 0
	for i, dimension := range dimensions {
		group := make(plotter.Values, len(dimension))

		for i, val := range dimension {
			group[i] = val
		}

		if index == len(colors) {
			index = 0
		}
		w := vg.Points(20)
		bars, err := plotter.NewBarChart(group, w)
		bars.LineStyle.Width = vg.Length(0)
		bars.Color = colors[index][0]
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			bars.Offset = -w
		}
		p.Add(bars)
		p.Legend.Add(dimensionNames[i], bars)
		index++
	}
	p.NominalX(xNames...)
	p.Legend.Top = true

	if err := p.Save(6*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s.png", strings.ReplaceAll(title, " ", "_"))); err != nil {
		panic(err)
	}

}
