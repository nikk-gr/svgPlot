// Package svgPlot Copyright 2023 Gryaznov Nikita
// Licensed under the Apache License, Version 2.0
package svgPlot

import (
	"errors"
	"fmt"
)

// DrawAngularFloat64 make the most compact plot with strait lines between points
// X and Y is a parallel arrays of point coordinates
func (s Style) DrawSmoothFloat64(X, Y []float64) (plot string, err error) {
	var (
		xMin, xMax, yMin, yMax float64
	)

	// get min & max values
	yMin, yMax, err = getMinMax(Y)
	if err != nil {
		return
	}
	xMin, xMax, err = getMinMax(X)
	if err != nil {
		return
	}
	plot, err = drawSmo(s.TotalHeight, s.TotalWidth, s.XDivisionsQty, s.YDivisionsQty, X, Y, xMin, xMax, yMin, yMax, s.NameOfX, s.NameOfY)
	return
}

// DrawAngularInt make the most compact plot with strait lines between points
// X and Y is a parallel arrays of point coordinates
func (s Style) DrawSmoothInt(X, Y []int) (plot string, err error) {
	var (
		xMin, xMax, yMin, yMax float64
		x1, y1                 []float64
	)
	// convert X, Y slices to float64
	x1, y1 = convertSliceToFloat64(X), convertSliceToFloat64(Y)

	// get min & max values
	yMin, yMax, err = getMinMax(y1)
	if err != nil {
		return
	}
	xMin, xMax, err = getMinMax(x1)
	if err != nil {
		return
	}
	plot, err = drawSmo(s.TotalHeight, s.TotalWidth, s.XDivisionsQty, s.YDivisionsQty, x1, y1, xMin, xMax, yMin, yMax, s.NameOfX, s.NameOfY)
	return
}

// DrawAngularInt64 make the most compact plot with strait lines between points
// X and Y is a parallel arrays of point coordinates
func (s Style) DrawSmoothInt64(X, Y []int64) (plot string, err error) {
	var (
		xMin, xMax, yMin, yMax float64
		x1, y1                 []float64
	)
	// convert X, Y slices to float64
	x1, y1 = convertSliceToFloat64(X), convertSliceToFloat64(Y)

	// get min & max values
	yMin, yMax, err = getMinMax(y1)
	if err != nil {
		return
	}
	xMin, xMax, err = getMinMax(x1)
	if err != nil {
		return
	}
	plot, err = drawSmo(s.TotalHeight, s.TotalWidth, s.XDivisionsQty, s.YDivisionsQty, x1, y1, xMin, xMax, yMin, yMax, s.NameOfX, s.NameOfY)
	return
}

// DrawAngularFloat64From0 make plot with strait lines between points
// Coordinate plane starts from (0, 0) point
// X and Y is a parallel arrays of point coordinates, only positive x and y is allowed
func (s Style) DrawSmoothFloat64From0(x, y []float64) (plot string, err error) {
	var (
		xMax, yMax float64
	)
	// convert x, y slices to float64
	err = checkPositive(x)
	if err != nil {
		return
	}
	err = checkPositive(y)
	if err != nil {
		return
	}
	// get min & max values
	_, yMax, err = getMinMax(y)
	if err != nil {
		return
	}
	_, xMax, err = getMinMax(x)
	if err != nil {
		return
	}
	plot, err = drawSmo(s.TotalHeight, s.TotalWidth, s.XDivisionsQty, s.YDivisionsQty, x, y, 0, xMax, 0, yMax, s.NameOfX, s.NameOfY)
	return
}

// DrawAngularIntFrom0 make plot with strait lines between points
// Coordinate plane starts from (0, 0) point
// X and Y is a parallel arrays of point coordinates, only positive x and y is allowed
func (s Style) DrawSmoothIntFrom0(x, y []int) (plot string, err error) {
	var (
		xMax, yMax float64
		x1, y1     []float64
	)
	// convert x, y slices to float64
	x1, y1 = convertSliceToFloat64(x), convertSliceToFloat64(y)
	err = checkPositive(x1)
	if err != nil {
		return
	}
	err = checkPositive(y1)
	if err != nil {
		return
	}
	// get min & max values
	_, yMax, err = getMinMax(y1)
	if err != nil {
		return
	}
	_, xMax, err = getMinMax(x1)
	if err != nil {
		return
	}
	plot, err = drawSmo(s.TotalHeight, s.TotalWidth, s.XDivisionsQty, s.YDivisionsQty, x1, y1, 0, xMax, 0, yMax, s.NameOfX, s.NameOfY)
	return
}

// DrawAngularInt64From0 make plot with strait lines between points
// Coordinate plane starts from (0, 0) point
// X and Y is a parallel arrays of point coordinates, only positive x and y is allowed
func (s Style) DrawSmoothInt64From0(x, y []int64) (plot string, err error) {
	var (
		xMax, yMax float64
		x1, y1     []float64
	)
	// convert x, y slices to float64
	x1, y1 = convertSliceToFloat64(x), convertSliceToFloat64(y)
	err = checkPositive(x1)
	if err != nil {
		return
	}
	err = checkPositive(y1)
	if err != nil {
		return
	}
	// get min & max values
	_, yMax, err = getMinMax(y1)
	if err != nil {
		return
	}
	_, xMax, err = getMinMax(x1)
	if err != nil {
		return
	}
	plot, err = drawSmo(s.TotalHeight, s.TotalWidth, s.XDivisionsQty, s.YDivisionsQty, x1, y1, 0, xMax, 0, yMax, s.NameOfX, s.NameOfY)
	return
}

// DrawAngular make the most compact plot with strait lines between points
// TotalHeight and TotalWidth defines the size of resulting picture
// X and Y is a parallel arrays of point coordinates
// NameOfX and NameOfY is axis labels. Max allowed length is 6 chars
// if NameOfX or NameOfX != "" it will replace the last number on the axis
func DrawSmooth[T numeric](TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty uint, X, Y []T, NameOfX, NameOfY string) (plot string, err error) {
	var (
		xMin, xMax, yMin, yMax float64
		x1, y1                 []float64
	)
	// convert X, Y slices to float64
	x1, y1 = convertSliceToFloat64(X), convertSliceToFloat64(Y)

	// get min & max values
	yMin, yMax, err = getMinMax(y1)
	if err != nil {
		return
	}
	xMin, xMax, err = getMinMax(x1)
	if err != nil {
		return
	}
	plot, err = drawSmo(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, x1, y1, xMin, xMax, yMin, yMax, NameOfX, NameOfY)
	return
}

// DrawAngularFrom0 make plot with strait lines between points
// Coordinate plane starts from (0, 0) point
// TotalHeight and TotalWidth defines the size of resulting picture
// X and Y is a parallel arrays of point coordinates, only positive x and y is allowed
// NameOfX and NameOfY is axis labels. Max allowed length is 6 chars
// if NameOfX or NameOfX != "" it will replace the last number on the axis
func DrawSmoothFrom0[T numeric](TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty uint, x, y []T, NameOfX, NameOfY string) (plot string, err error) {
	var (
		xMax, yMax float64
		x1, y1     []float64
	)
	// convert x, y slices to float64
	x1, y1 = convertSliceToFloat64(x), convertSliceToFloat64(y)
	err = checkPositive(x1)
	if err != nil {
		return
	}
	err = checkPositive(y1)
	if err != nil {
		return
	}
	// get min & max values
	_, yMax, err = getMinMax(y1)
	if err != nil {
		return
	}
	_, xMax, err = getMinMax(x1)
	if err != nil {
		return
	}
	plot, err = drawSmo(TotalHeight, TotalWidth, xDivisionsQty, yDivisionsQty, x1, y1, 0, xMax, 0, yMax, NameOfX, NameOfY)
	return
}

// draw returns a complete plot picture with points joined by straight lines
func drawSmo(height, width, xDivisionsQty, yDivisionsQty uint, x, y []float64, xMin, xMax, yMin, yMax float64, xName, yName string) (result string, err error) {
	var (
		greed, plot                      string
		x0, y0, gradX, gradY, xLen, yLen float64
		xZeroPos, yZeroPos               int
		xNums, yNums                     []string
	)

	if len(xName) > 6 {
		err = errors.New("xName max len is 6")
		return
	}
	if len(yName) > 6 {
		err = errors.New("yName max len is 6")
		return
	}
	xNums, xLen, xZeroPos, err = makeArr(xMin, xMax, xDivisionsQty)
	if err != nil {
		return
	}
	yNums, yLen, yZeroPos, err = makeArr(yMin, yMax, yDivisionsQty)
	if err != nil {
		return
	}
	if xName != "" {
		xNums[len(xNums)-1] = xName
	}
	if yName != "" {
		yNums[len(yNums)-1] = yName
	}
	greed, x0, y0, gradX, gradY, err = makeGreed(height, width, xNums, yNums, xLen, yLen, xZeroPos, yZeroPos)
	if err != nil {
		return
	}
	plot, err = makeSmoothCurve(x0, y0, gradX, gradY, x, y)
	if err != nil {
		return
	}
	result = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	result += fmt.Sprintf("<svg width=\"%d\" height=\"%d\" viewBox=\"0 0 %d %d\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\">\n\"", width, height, width, height)
	result += greed + "\n"
	result += plot + "\n"
	result += "</svg>"
	return
}