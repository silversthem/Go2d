package main

import(
  "go2d/graphics"
  "go2d/graphics/shapes"
  "go2d/graphics/colors"
  "go2d/geometry"
  "image/color"
)

func main() {
  /* Square Shape test */
  image := graphics.NewSurface(200,200,"image3.png") // An image
  square := shapes.NewSquare(geometry.NewPoint(30,30),100) // Creates a square
  square.Drawable.Fill = colors.Color{color.RGBA{255, 0, 0, 255}} // Red square
  square.Drawable.Borders = colors.Color{color.RGBA{0, 255, 0, 255}} // Green borders
  // @TODO : square.Transformations.Rotate(40) // Rotating the square
  image.Draw(square)
  image.SaveAsPng()
}
