package graphics

/*
  Contains basic drawing methods in a surface :
  - Points
  - Lines
  - Triangles
  - Shapes, which are basically drawn as a arrangement of triangles
*/

import(
  "go2d"
  "go2d/geometry"
  "image/color"
  "fmt"
)

/* Drawing methods */

func (surface Surface) DrawPixel(x,y int,col color.Color) { // draws a pixel in the image
  surface.Image.Set(x,y,col)
}

func (surface Surface) DrawPoint(point geometry.Point,col color.Color) { // draws a point in the image
  surface.DrawPixel(point.X,point.Y,col)
}

func (surface Surface) DrawLine(line geometry.Line,drawable go2d.Drawable) { // draws a line in the plan
  pasx,pasy := line.GetPas()
  fmt.Println(pasx,pasy)
  fmt.Println(line.Start,line.End)
  x,y := float64(line.Start.X),float64(line.Start.Y)
  for i := 0.0;i < line.Length();i++ {
    x += pasx
    y += pasy
    surface.DrawPoint(geometry.NewPoint(int(x),int(y)),drawable.GetPointColor(geometry.NewPoint(int(x),int(y))))
  }
}

func (surface Surface) Draw(rendered go2d.Rendered) { // Draws something in the image
  rendered.Draw(surface)
}
