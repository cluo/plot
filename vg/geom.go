// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vg

// A Point is a location in 2d space.
//
// Points are used for drawing, not for data.  For
// data, see the XYer interface.
type Point struct {
	X, Y Length
}

// Dot returns the dot product of two points.
func (p Point) Dot(q Point) Length {
	return p.X*q.X + p.Y*q.Y
}

// Add returns the component-wise sum of two points.
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub returns the component-wise difference of two points.
func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

// Scale returns the component-wise product of a point and a scalar.
func (p Point) Scale(s Length) Point {
	return Point{p.X * s, p.Y * s}
}

// A Rectangle represents a rectangular region of 2d space.
type Rectangle struct {
	Min Point
	Max Point
}

// Size returns the width and height of a Rectangle.
func (r Rectangle) Size() Point {
	return Point{
		X: r.Max.X - r.Min.X,
		Y: r.Max.Y - r.Min.Y,
	}
}

// Path returns the path of a Rect specified by its
// upper left corner, width and height.
func (r Rectangle) Path() (p Path) {
	p.Move(r.Min.X, r.Min.Y)
	p.Line(r.Max.X, r.Min.Y)
	p.Line(r.Max.X, r.Max.Y)
	p.Line(r.Min.X, r.Max.Y)
	p.Close()
	return
}
