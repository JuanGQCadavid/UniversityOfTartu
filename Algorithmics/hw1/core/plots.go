package core

import (
	"fmt"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func GenerateLinePlot(title string, xLabel string, yLabel string, points map[string]plotter.XYs) {
	p := plot.New()
	p.Title.Text = title
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel

	p.Add(plotter.NewGrid())

	for label, dataPoints := range points {
		// l, err := plotter.NewLine(dataPoints)
		if err := plotutil.AddLinePoints(p, label, dataPoints); err != nil {
			panic(err.Error())
		}
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, fmt.Sprintf("%s.png", strings.ReplaceAll(title, " ", "_"))); err != nil {
		panic(err)
	}

}
