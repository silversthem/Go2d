package go2d
/*
  Basic functions of the go2d packages, contains :
  Shapes, which are bunches of points
  ---- TODOS ----
  @TODO : Text support /!\
  @TODO : Use import C and SDL to be able to handle a window and events [new file]
  @TODO : Handle collisions [new file] using triangles and masks for textures [new file]
  @TODO : Default shapes and colors [new file] with NewX() and ColorX() functions
  @TODO : Support Math graphing [new file] and diagrams [new file]
  @TODO : Create Objects which are Drawable shapes [new file] with collisions ? with embedded event support ?
  @TODO : Multithreading to update images while doing calculations, to make the library able to show cool, real time, rendering of another process [new file]
  @TODO : Shaders [new file]
  @TODO : Complex support in graphes [new file]
  @TODO : Nodes and stuff, to build cool diagrams [new file]
*/

import(
  "go2d/geometry"
  "math"
)

type Shape struct { // a shape is a collection of points that can be transformed
  Points []geometry.Point // points, relatively to Position
  Transformations geometry.Transform // transformation on points
  Position geometry.Point // Position of shape
}

func NewShape(x,y int,points... geometry.Point) Shape { // creates a basic shape
  return Shape{points,geometry.NewTransform(),geometry.NewPoint(x,y),}
}

/* Counting */

func (shape *Shape) GetPointsCount() int { // returns a shape amount of points
  return len(shape.Points)
}

func (shape *Shape) GetLinesCount() int { // returns the amount of lines in the shape
  return len(shape.Points)
}

func (shape *Shape) GetTrianglesCount() int { // returns the amount of triangles in the shape
  return len(shape.Points)
}

func (shape *Shape) GetRelativePoint(index int) geometry.Point { // returns a transformed point of a shape relative to shape origin (its position)
  if index < 0 {
    panic("Trying to access non-existant point in shape")
  }
  index = int(math.Mod(float64(index),float64(shape.GetPointsCount())))
  return shape.Transformations.GetPoint(shape.Points[index])
}

/* Getting */

func (shape *Shape) GetPoint(index int) (point geometry.Point) { // returns a point of shape in the plan
  point = shape.GetRelativePoint(index)
  point.X = point.X + shape.Position.X
  point.Y = point.Y + shape.Position.Y
  return
}

func (shape *Shape) GetLine(index int) (Line geometry.Line) { // returns the x-th line of the shape
  Line.Start = shape.GetPoint(index)
  Line.End = shape.GetPoint(index+1)
  return
}

func (shape *Shape) GetNextTriangle(index int) geometry.Triangle { // Returns a triangle composed of index,index+1,index+2 points of a shape
  line := shape.GetLine(index)
  return geometry.NewTriangle(line.Start,line.End,shape.GetPoint(index+2))
}

func (shape *Shape) GetAbsoluteTriangle(index,triangleIndex int) geometry.Triangle { // returns a triangle from 3 shape point, used in many things like filling shape or collisions
  line := shape.GetLine(index)
  return geometry.NewTriangle(line.Start,line.End,shape.GetPoint(triangleIndex))
}

func (shape *Shape) GetTriangleFromOrigin(index int) geometry.Triangle { // Returns a triangle formed by the origin and 2 points of the shape
  line := shape.GetLine(index)
  return geometry.NewTriangle(shape.Position,line.Start,line.End)
}
