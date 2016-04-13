package shapes

/*
  Defines a shape
  A shape is a collection of points, usually filled by triangles.
  Shapes include : square, rectangle, circle and any convex shape
*/

import(
  "go2d"
  "go2d/geometry"
)

/* Interfaces */

type Drawable struct { // Drawable instructions for Shapes
  Borders go2d.Drawable // instructions for borders
  Fill    go2d.Drawable // instructions for filling
}

type Shape interface { // A basic shape
  GetPoint(i int)  geometry.Point      // Returns a shape point
  GetPointsCount() int                 // Returns the amount of points the shape has
  GetDrawable()    Drawable            // Returns drawing context for shape
  GetTransform()   geometry.Transform  // Returns transformations on shape
  GetOrigin()      geometry.Point      // Returns shape's origin
}

func GetAbsolutePoint(point,origin geometry.Point,transform geometry.Transform) (final geometry.Point) { // Returns a transformed point in absolute coordinates
  final = transform.GetPoint(point)
  final.X += origin.X
  final.Y += origin.Y
  return
}

func DrawShape(shape Shape,surface go2d.Surface) { // Draws a shape in a surface
  for i := 0;i < shape.GetPointsCount();i++ { // Drawing borders
    surface.DrawLine(geometry.NewLine(GetAbsolutePoint(shape.GetPoint(i),
    shape.GetOrigin(),shape.GetTransform()),GetAbsolutePoint(shape.GetPoint(i+1),shape.GetOrigin(),shape.GetTransform())),shape.GetDrawable().Borders)
  }
  // @TODO : Filling the shape
}
