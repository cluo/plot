// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot_test

import (
	"bytes"
	"fmt"
	"image/color"
	"reflect"
	"testing"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"
	"github.com/gonum/plot/vg/recorder"
)

func TestLegendAlignment(t *testing.T) {
	const fontSize = 10.189054726368159 // This font size gives an entry height of 10.
	font, err := vg.MakeFont(plot.DefaultFont, fontSize)
	if err != nil {
		t.Fatalf("failed to create font: %v", err)
	}
	l := plot.Legend{
		ThumbnailWidth: vg.Points(20),
		TextStyle:      draw.TextStyle{Font: font},
	}
	for _, n := range []string{"A", "B", "C", "D"} {
		b, err := plotter.NewBarChart(plotter.Values{0}, 1)
		if err != nil {
			t.Fatalf("failed to create bar chart %q: %v", n, err)
		}
		l.Add(n, b)
	}

	var r recorder.Canvas
	c := draw.NewCanvas(&r, 100, 100)
	l.Draw(draw.Canvas{
		Canvas: c.Canvas,
		Rectangle: vg.Rectangle{
			Min: vg.Point{0, 0},
			Max: vg.Point{100, 100},
		},
	})

	got := r.Actions

	// want is a snapshot of the actions for the code above when the
	// graphical output has been visually confirmed to be correct for
	// the bar charts example show in gonum/plot#25.
	want := []recorder.Action{
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.Fill{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 30},
				{Type: vg.LineComp, X: 80, Y: 40},
				{Type: vg.LineComp, X: 100, Y: 40},
				{Type: vg.LineComp, X: 100, Y: 30},
				{Type: vg.CloseComp},
			},
		},
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.SetLineWidth{
			Width: 1,
		},
		&recorder.SetLineDash{},
		&recorder.Stroke{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 30},
				{Type: vg.LineComp, X: 80, Y: 40},
				{Type: vg.LineComp, X: 100, Y: 40},
				{Type: vg.LineComp, X: 100, Y: 30},
				{Type: vg.LineComp, X: 80, Y: 30},
			},
		},
		&recorder.SetColor{},
		&recorder.FillString{
			Font:   string("Times-Roman"),
			Size:   fontSize,
			X:      70.09452736318407,
			Y:      30.18905472636816,
			String: "A",
		},
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.Fill{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 20},
				{Type: vg.LineComp, X: 80, Y: 30},
				{Type: vg.LineComp, X: 100, Y: 30},
				{Type: vg.LineComp, X: 100, Y: 20},
				{Type: vg.CloseComp},
			},
		},
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.SetLineWidth{
			Width: 1,
		},
		&recorder.SetLineDash{},
		&recorder.Stroke{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 20},
				{Type: vg.LineComp, X: 80, Y: 30},
				{Type: vg.LineComp, X: 100, Y: 30},
				{Type: vg.LineComp, X: 100, Y: 20},
				{Type: vg.LineComp, X: 80, Y: 20},
			},
		},
		&recorder.SetColor{},
		&recorder.FillString{
			Font:   string("Times-Roman"),
			Size:   fontSize,
			X:      70.65671641791045,
			Y:      20.18905472636816,
			String: "B",
		},
		&recorder.SetColor{
			Color: color.Gray16{
				Y: uint16(0),
			},
		},
		&recorder.Fill{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 10},
				{Type: vg.LineComp, X: 80, Y: 20},
				{Type: vg.LineComp, X: 100, Y: 20},
				{Type: vg.LineComp, X: 100, Y: 10},
				{Type: vg.CloseComp},
			},
		},
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.SetLineWidth{
			Width: 1,
		},
		&recorder.SetLineDash{},
		&recorder.Stroke{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 10},
				{Type: vg.LineComp, X: 80, Y: 20},
				{Type: vg.LineComp, X: 100, Y: 20},
				{Type: vg.LineComp, X: 100, Y: 10},
				{Type: vg.LineComp, X: 80, Y: 10},
			},
		},
		&recorder.SetColor{},
		&recorder.FillString{
			Font:   string("Times-Roman"),
			Size:   fontSize,
			X:      70.65671641791045,
			Y:      10.189054726368159,
			String: "C",
		},
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.Fill{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 0},
				{Type: vg.LineComp, X: 80, Y: 10},
				{Type: vg.LineComp, X: 100, Y: 10},
				{Type: vg.LineComp, X: 100, Y: 0},
				{Type: vg.CloseComp},
			},
		},
		&recorder.SetColor{
			Color: color.Gray16{},
		},
		&recorder.SetLineWidth{
			Width: 1,
		},
		&recorder.SetLineDash{},
		&recorder.Stroke{
			Path: vg.Path{
				{Type: vg.MoveComp, X: 80, Y: 0},
				{Type: vg.LineComp, X: 80, Y: 10},
				{Type: vg.LineComp, X: 100, Y: 10},
				{Type: vg.LineComp, X: 100, Y: 0},
				{Type: vg.LineComp, X: 80, Y: 0},
			},
		},
		&recorder.SetColor{},
		&recorder.FillString{
			Font:   string("Times-Roman"),
			Size:   fontSize,
			X:      70.09452736318407,
			Y:      0.189054726368159,
			String: "D",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected legend actions:\ngot:\n%s\nwant:\n%s", formatActions(got), formatActions(want))
	}
}

func formatActions(actions []recorder.Action) string {
	var buf bytes.Buffer
	for _, a := range actions {
		fmt.Fprintf(&buf, "\t%s\n", a.Call())
	}
	return buf.String()
}
