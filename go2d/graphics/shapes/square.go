package shapes

/*
  Defines a square shape
*/

import(
  "go2d"
  "go2d/geometry"
)

/* Structure */

type Square struct { // A struct
  Size            int                // Square's size
  Origin          geometry.Point     // Square's origin
  Transformations geometry.Transform // Transformations (rotation, scaling)
  Drawable        Drawable      // Drawing context
}

/* Methods */

func NewSquare(point geometry.Point,size int) (sq Square) { // Creates a square
  sq.Size            = size
  sq.Origin          = point
  sq.Transformations = geometry.NewTransform()
  return
}

func (square Square) Draw(surface go2d.Surface) { // Draws the square
  DrawShape(square,surface)
}

func (square Square) GetPoint(i int) (point geometry.Point) { // Returns one of the 4 edges of a square
  i = i % 4
  switch i {
  case 1:
    point.X = square.Size
  case 2:
    point.Y = square.Size
    point.X = square.Size
  case 3:
    point.Y = square.Size
  }
  return
}

func (square Square) GetPointsCount() int { // A square has 4 edges
  return 4;
}

func (square Square) GetDrawable() Drawable { // Returns drawing context for shape
  return square.Drawable
}

func (square Square) GetTransform() geometry.Transform { // Returns transformations on shape
  return square.Transformations
}

func (square Square) GetOrigin() geometry.Point { // Returns origin
  return square.Origin
}
