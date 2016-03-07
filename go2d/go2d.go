package go2d
/*
  Basic functions of the go2d packages, contains :
  Shapes, which are bunches of points
  ---- TODOS ----
  @TODO : Use import C and SDL to be able to handle a window and events [new file]
  @TODO : Handle collisions [new file] using triangles and masks for textures [new file]
  @TODO : Default shapes and colors [new file] with NewX() function
  @TODO : Support Math graphing [new file]
  @TODO : Create Objects which are Drawable shapes [new file] with collisions ? with embedded event support ?
  @TODO : Change the way triangles are created in shapes
*/

import(
  "go2d/geometry"
)

type Shape struct { // a shape is a collection of points that can be transformed
  Points []geometry.Point // points, relatively to Position
  Transformations geometry.Transform // transformation on points
  Position geometry.Point // Position of shape
}

func NewShape(x,y int,points []geometry.Point) Shape { // creates a basic shape
  return Shape{points,geometry.NewTransform(),geometry.NewPoint(x,y),}
}

/* Shape methods */

func (shape *Shape) GetPointsCount() int { // returns a shape amount of points
  return len(shape.Points)
}

func (shape *Shape) GetRelativePoint(index int) geometry.Point { // returns a transformed point of a shape relative to shape origin (its position)
  if index >= shape.GetPointsCount() || index < 0 {
    panic("Trying to access non-existant point in shape")
  }
  return shape.Transformations.GetPoint(shape.Points[index])
}

func (shape *Shape) GetPoint(index int) (point geometry.Point) { // returns a point of shape in the plan
  point = shape.GetRelativePoint(index)
  point.X = point.X + shape.Position.X
  point.Y = point.Y + shape.Position.Y
  return
}

func (shape *Shape) GetLine(index int) (Line geometry.Line) { // returns the x-th line of the shape
  Line.Start = shape.GetPoint(index-1)
  Line.End = shape.GetPoint(index)
  return
}

func (shape *Shape) getTriangle(index int) (triangle geometry.Triangle) { // returns a triangle from 3 shape point, used in many things like filling shape or collisions
  triangle.Sides[0] = shape.GetPoint(index - 2)
  triangle.Sides[1] = shape.GetPoint(index - 1)
  triangle.Sides[2] = shape.GetPoint(index)
  return triangle
}
