package drawer

import (
"math/rand"
"gonum.org/v1/plot"
"gonum.org/v1/plot/plotter"
"gonum.org/v1/plot/plotutil"
"gonum.org/v1/plot/vg"
	"clusterdata-go/middle"
	"clusterdata-go/statistics"
)

func Test() {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "100 Times container resoure utilization"
	p.X.Label.Text = "times"
	p.Y.Label.Text = "Utilization"
	err = plotutil.AddLinePoints(p,
		"max", Select(1),
		"avg", Select(2),
		"min", Select(3))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(15*vg.Inch, 6*vg.Inch, middle.Prefix+"container100times.png"); err != nil {
		panic(err)
	}
}

// randomPoints returns some random x, y points.
func RandomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func Avg() plotter.XYs {
	pts := make(plotter.XYs, 12)
	var avg = [12]int{18,22,22,23,25,30,46,35,35,31,28,31}
	for i := range pts {
		pts[i] = plotter.XY{X:float64(i),Y:float64(avg[i])}
	}
	return pts

}

func Max() plotter.XYs {
	pts := make(plotter.XYs, 12)
	var avg = [12]int{96,97,95,95,95,98,96,95,99,96,96,94}
	for i := range pts {
		pts[i] = plotter.XY{X:float64(i),Y:float64(avg[i])}
	}
	return pts

}

func Min() plotter.XYs {
	pts := make(plotter.XYs, 12)
	var avg = [12]int{1,1,1,1,1,1,1,1,1,1,1,1}
	for i := range pts {
		pts[i] = plotter.XY{X:float64(i),Y:float64(avg[i])}
	}
	return pts

}

func Select(t int) plotter.XYs {
	pts := make(plotter.XYs, 101)
	for i := range pts {
		if t == 1{
			pts[i] = plotter.XY{X:float64(i),Y:float64(statistics.Clines[i].Max)}
		}
		if(t == 2){
			pts[i] = plotter.XY{X:float64(i),Y:float64(statistics.Clines[i].Avg)}
		}
		if(t == 3){
			pts[i] = plotter.XY{X:float64(i),Y:float64(statistics.Clines[i].Min)}
		}
	}
	return pts

}
