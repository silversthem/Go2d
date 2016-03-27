package main

import(
  "go2d"
  "image/color"
  "go2d/geometry"
  "go2d/graphics"
)

func main() {
  /* Shape test */
  shape := go2d.NewShape(30,50,geometry.NewPoint(0,0),geometry.NewPoint(5,-5),
  geometry.NewPoint(10,0),geometry.NewPoint(10,10),geometry.NewPoint(0,10)) // Creating a pentagon

  shape.Transformations.Zoom(4); // Zooming it 4 times
  //shape.Transformations.Rotate(3.14/3) // Rotating it 60°

  drawable := graphics.NewDrawableShape(color.RGBA{255, 0, 0, 255},color.RGBA{0, 0, 0, 255}) // Creating a drawing context for our polygon, black with red border
  drawable.Thickness = 1 // 1px wide border, will soon be able to take different values

  surface := graphics.NewSurface(100,100,"image.png") // A surface in which to draw

  surface.Draw(shape,drawable) // Drawing the shape

  surface.SaveAsPng()

  /* Pixel shader test */
  surface2 := graphics.NewSurface(255,255,"image2.png") // Creating another surface, to test the stuff

  f := func(x,y int) color.Color { // Function used by the shader
    return color.RGBA{uint8(x),uint8(y),uint8(x),255}
  }
  shader := graphics.NewPixelShader(surface2.Image.Rect.Min,surface2.Image.Rect.Max,f) // Shader all along the image
  surface2.DrawPixelShader(shader)

  surface2.SaveAsPng()
}
