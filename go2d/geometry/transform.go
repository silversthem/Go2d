package geometry

/*
  A Transform structure is what you use to specify rotations and scaling transformations for a shape
*/

import(
  "math"
)

type Scale struct { // a scale for an object
  Width,Height float64
}

type Transform struct { // list a transformations a shape can undergo
  Angle float64 // angle, in radian
  Size Scale // Scaling
}

func NewTransform() Transform { // Creates a transformation for a shape
  return Transform{0,Scale{1,1}}
}

/* Transform methods */

func (transform Transform) GetPoint(point Point) Point { // transforms a point
  point.X = int(float64(point.X)*transform.Size.Width) // Scaling
  point.Y = int(float64(point.Y)*transform.Size.Height)
  return Point{
    int(float64(point.X)*math.Cos(transform.Angle) + float64(point.Y)*math.Sin(transform.Angle)), // rotating
    int(float64((-1)*point.X)*math.Sin(transform.Angle) + float64(point.Y)*math.Cos(transform.Angle)),
  }
}

func (transform Transform) GetLine(line Line) Line { // transforms a line
  line.Start = transform.GetPoint(line.Start)
  line.End = transform.GetPoint(line.End)
  return line
}

func (transform Transform) GetTriangle(triangle Triangle) Triangle { // transforms a triangle
  return Triangle{[3]Point{
    transform.GetPoint(triangle.Points[0]),
    transform.GetPoint(triangle.Points[1]),
    transform.GetPoint(triangle.Points[2]),}}
}

func (transform *Transform) Rotate(angle float64) { // Adds rotation to the transformation
  transform.Angle = GetAbsoluteAngleValue(angle+transform.Angle);
}

func (transform *Transform) Zoom(scale float64) { // Zooms/Unzooms the transformation
  transform.Size.Width = scale*transform.Size.Width
  transform.Size.Height = scale*transform.Size.Height
}

func (transform *Transform) SetRotation(angle float64) { // Sets rotation
  transform.Angle = angle;
}

func (transform *Transform) SetZoom(scale float64) { // Sets zoom
  transform.Size.Width = scale
  transform.Size.Height = scale
}
