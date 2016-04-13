package colors

/*
  A collections of basic colors for your drawables
  Can also generate gradients, mesh colors and else
*/

import(
  "image/color"
  "go2d/geometry"
)

type Color struct { // Creates a color structure, to be applied as a drawable
  Col color.Color // The color
}

func (col Color) GetPointColor(point geometry.Point) color.Color { // Returns the color
  return col.Col
}
