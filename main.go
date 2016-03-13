package main

import(
  //"go2d"
  "image/color"
  "go2d/geometry"
  "go2d/graphics"
)

func main() {
  /* Colors */
  black := color.RGBA{0, 0, 0, 255}
  red := color.RGBA{255, 0, 0, 255}
  /* Creating a triangle */
  triangle := geometry.NewTriangle(geometry.NewPoint(3,2),geometry.NewPoint(235,90),geometry.NewPoint(51,490))
  /* Creating an image on which to draw */
  surface := graphics.NewSurface(500,500,"image.png")
  /* Drawing triangle */
  surface.DrawFillTriangle(triangle,black)
  /* Drawing the triangle outer */
  surface.DrawTriangle(triangle,red)
  /* Saving */
  surface.SaveAsPng()
}
