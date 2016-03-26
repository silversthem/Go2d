package geometry

/*
  A point is the same as the point structure from the image package
  With a few helpers
*/

import(
  "image"
  "math"
)

type Point image.Point // a point in plan

func NewPoint(x,y int) Point { // creates a point
  return Point{x,y,}
}

/* Point methods */

func (point Point) Distance(point2 Point) float64 { // gets the distance between two points
  return math.Sqrt(math.Pow(float64(point.X - point2.X),2) + math.Pow(float64(point.Y - point2.Y),2)) // sqrt((x1-x2)^2 + (y1-y2)^2)
}

func (point Point) Substract(point2 Point) Point { // Substract point2 coordinates from point
  return Point{
    point.X - point2.X,
    point.Y - point2.Y,
  }
}
