package main

import(
  "go2d"
  "image/color"
  "go2d/geometry"
  "go2d/graphics"
)

func main() {
  /* Shape */
  shape := go2d.NewShape(30,50,geometry.NewPoint(0,0),geometry.NewPoint(5,-5),
  geometry.NewPoint(10,0),geometry.NewPoint(10,10),geometry.NewPoint(0,10)) // Creating a pentagon
  shape.Transformations.Zoom(4); // Zooming it 4 times
  shape.Transformations.Rotate(3.14/3) // Rotating it 60°
  /* Drawable */
  drawable := graphics.NewDrawableShape(color.RGBA{255, 0, 0, 255},color.RGBA{0, 0, 0, 255}) // Creating a drawing context for our polygon, black with red border
  // drawable.Thickness = 3 // 3px wide border, not implemented yet
  /* Surface */
  surface := graphics.NewSurface(100,100,"image.png")
  /* Drawing the shape */
  surface.Draw(shape,drawable)
  /* Saving */
  surface.SaveAsPng()
}
