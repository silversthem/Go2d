package graphics

/*
  basic graphic support
  Contains Surfaces and Drawable, Surfaces are where you draw and drawable is information about how to draw a shape
  ---- TODOS ----
  @TODO : Draw methods
  @TODO : Save in other format than png
  @TODO : Save animations in gif
*/

import(
  "go2d"
  "go2d/geometry"
  "image"
  "image/color"
  "image/png"
  "os"
)

type Surface struct { // A surface is a structure containing a plan a filename and an actual image, to draw into
  Filename string
  Image *image.RGBA
}

type Drawable struct { // A structure to get data info on what to draw in a shape
  BorderColor color.Color // the border color
  FillColor color.Color // Fill color
  Thickness int // Border thickness
  ColorBorders,Fill bool // Should we color borders and fill the shape
}

func NewSurface(w,h int,Filename string) Surface { // creates a new surface
  return Surface{Filename,image.NewRGBA(image.Rect(0,0,w,h)),}
}

/* Surface methods */

func (surface *Surface) DrawPixel(x,y int,col color.Color) { // draws a pixel in the image

}

func (surface *Surface) DrawPoint(point geometry.Point) { // draws a point in the image

}

func (surface *Surface) DrawLine(line geometry.Line,thickness int,col color.Color) { // draws a line in the plan

}

func (surface *Surface) DrawFillTriangle(triangle geometry.Triangle,col color.Color) { // fills a triangle section of the plan

}

func (surface *Surface) Draw(shape go2d.Shape,drawable Drawable) { // draws objects in plan
  // If object has 1 point -> point, use either fillColor or Border color
  // If object has 2 points -> line
  // else
    // if Fill is true, decompose object in triangle then fill each triangle
  // draw each line with border
}

func (surface *Surface) Clear() { // clears the image

}

func (surface *Surface) setSize(size geometry.Scale) { // changes surface's size
  surface.Image.Rect.Max = image.Point{int(size.Width),int(size.Height),}
}

func (surface *Surface) SaveAsPng() { // saves the image in the png file
  file,err := os.Create(surface.Filename)
  if err != nil {
    panic(err)
  }
  png.Encode(file,surface.Image)
  defer file.Close()
}
