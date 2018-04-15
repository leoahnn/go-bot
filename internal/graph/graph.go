package graph

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/wcharczuk/go-chart"
)

// Render renders a bar chart given json
func Render() {
	bar := chart.BarChart{
		Title:      "Chart",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{Padding: chart.Box{Top: 40}},
		Height:     512,
		BarWidth:   60,
		XAxis:      chart.Style{Show: true},
		YAxis:      chart.YAxis{Style: chart.Style{Show: true}},
		Bars: []chart.Value{
			{Value: 5, Label: "name 1"},
			{Value: 2, Label: "name 2"},
			{Value: 1, Label: "name 3"},
			{Value: 4, Label: "name 4"},
			{Value: 5, Label: "name 5"},
		},
	}

	f, err := os.Create("test.png")
	err = bar.Render(chart.PNG, f)
	if err != nil {
		log.Error("error rendering chart", err)
	}
}
