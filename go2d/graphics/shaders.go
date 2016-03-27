package graphics

/*
  A shader is a structure containing a function (or multiple) and instructions on how to use it
  They're a way to draw cool things easily
  Shaders availables :
  - PixelShader, gives you the ability to draw pixel per pixel in  a rectangle area
  @TODO : Other types of shaders, moving differently or doing different things
  @TODO : Binary shaders -> Coloring a pixel in a color determined a boolean function
*/

import(
  "image"
  "image/color"
)

type PixelShader struct { // A shader that colors an area pixel per pixel
  Bounds image.Rectangle // Where to apply the shader
  Shader func(int,int) color.Color // Function
}

/* Creating shaders */

func NewPixelShader(min,max image.Point,shader func(int,int) color.Color) PixelShader { // Creates a pixel shader
  return PixelShader{image.Rectangle{min,max},shader}
}

/* Drawing shader in surfaces */

func (surface *Surface) DrawPixelShader(shader PixelShader) { // Drawing a pixel shader
  for i := shader.Bounds.Min.X;i <= shader.Bounds.Max.X;i++ { // Going through x
    for j := shader.Bounds.Min.Y;j <= shader.Bounds.Max.Y;j++ { // Going through y
      surface.DrawPixel(i,j,shader.Shader(i,j))
    }
  }
}
