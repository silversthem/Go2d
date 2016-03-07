package geometry
/*
 Basic operations in plan
 contains everything you could need in a plan :
  Points, Lines, Triangles (all shapes are just triangles next to each other)
  Scales (useful for sizing and scaling) and Transform (Transformations applied to an object) => resizing is done with a scale (Size), and rotating with Angle in rads
 Also contains everything needed for plan calculations
 ---- TODOS ----
 @TODO : Everything about triangles
 @TODO : Transform handling
 @TODO : More calculations :
  line intersections, alignment of points, translations using lines, triangles rotation and scaling
*/

import(
  "image"
  "math"
)

type Point image.Point // a point in plan

type Line struct { // a line is a link between two points
  Start,End Point
}

type Triangle struct { // 3 lines forms a triangle, useful when testing collisions or filling a shape
  Sides [3]Point
}

type Scale struct { // a scale for an object
  Width,Height float64
}

type Transform struct { // list a transformations a shape can undergo
  Angle float64 // angle, in radian
  Size Scale // Scaling
}

func NewPoint(x,y int) Point { // creates a point
  return Point{x,y,}
}

func NewLine(start,end Point) Line { // creates a line
  return Line{start,end,}
}

func NewTriangle() {
  // new triangle
}

func NewScale(w,h float64) Scale { // creates a scale structure
  return Scale{w,h,}
}

func NewTransform() Transform { // creates a Transform structure
  return Transform{0.0,NewScale(1.0,1.0),}
}

/* Line methods */

func (line *Line) distance() { // gets distance between two line points

}

/* Triangle methods */

// get the box surrounding the triangle

// is point in triangle

// does line go through triangle

/* Point methods */

func (point *Point) distance(point2 Point) float64 { // gets the distance between two points
  return math.Sqrt(math.Pow(float64(point.X - point2.X),2) + math.Pow(float64(point.Y - point2.Y),2)) // sqrt((x1-x2)^2 + (y1-y2)^2)
}

/* Transform methods */

func (transform *Transform) GetPoint(point Point) Point { // transforms a point
  point.X = int(float64(point.X)*transform.Size.Width) // Scaling
  point.Y = int(float64(point.Y)*transform.Size.Height)
  return Point{
    int(float64(point.X)*math.Cos(transform.Angle) + float64(point.Y)*math.Sin(transform.Angle)), // rotating
    int(float64((-1)*point.X)*math.Sin(transform.Angle) + float64(point.Y)*math.Cos(transform.Angle)),
  }
}

func (transform *Transform) getLine(line Line) Line { // transforms a line
  line.Start = transform.GetPoint(line.Start)
  line.End = transform.GetPoint(line.End)
  return line
}

// GetTriangle

// Zoom
// Rotation
