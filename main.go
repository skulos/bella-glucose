package main

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func data1() []opts.LineData {

	items := make([]opts.LineData, 0)

	//
	items = append(items, opts.LineData{Value: 22.1})
	items = append(items, opts.LineData{Value: 3.8})
	items = append(items, opts.LineData{Value: 3.3})
	items = append(items, opts.LineData{Value: 3.0})
	items = append(items, opts.LineData{Value: 5.4})

	//
	return items
}

func data2() []opts.LineData {

	items := make([]opts.LineData, 0)

	//
	items = append(items, opts.LineData{Value: 16.4})
	items = append(items, opts.LineData{Value: 8.4})
	items = append(items, opts.LineData{Value: 3.7})
	items = append(items, opts.LineData{Value: 5.4})
	items = append(items, opts.LineData{Value: 8.5})

	//
	return items
}

func data3() []opts.LineData {

	items := make([]opts.LineData, 0)

	//
	items = append(items, opts.LineData{Value: 18.4})
	items = append(items, opts.LineData{Value: 0})
	items = append(items, opts.LineData{Value: 0})
	items = append(items, opts.LineData{Value: 13.8})
	items = append(items, opts.LineData{Value: 10.4})

	//
	return items
}

func chartData() {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1600px",
			Height: "800px",
		}),
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeRoma}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Bella's Glucose Levels",
			Subtitle: "Line chart rendered for measurements on 12 & 13 August 2021",
		}),
	)

	line.SetXAxis([]string{"8:00am", "12:00pm", "04:00pm", "08:00pm", "12:00am"}).
		AddSeries("A", data1(),
			charts.WithLabelOpts(opts.Label{Show: true}),
		).
		AddSeries("B", data2(),
			charts.WithLabelOpts(opts.Label{Show: true}),
		).
		AddSeries("C",
			data3(),
			charts.WithLineStyleOpts(
				opts.LineStyle{
					Color: "#39D619",
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:  true,
					Color: "#39D619",
				},
			),
		).
		SetSeriesOptions(
			// charts.WithLabelOpts(opts.Label{
			// 	Show: true,
			// }),
			charts.WithLineChartOpts(opts.LineChart{Smooth: true}),
			// charts.WithMarkPointStyleOpts(
			// 	opts.MarkPointStyle{Label: &opts.Label{Show: true}}),
		)

	f, _ := os.Create("chart.html")

	line.Render(f)
}

func main() {
	chartData()
}
