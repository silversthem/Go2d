package graphics

/*
  basic graphic support
  Contains Surfaces and Drawable, Surfaces are where you draw and drawable is information about how to draw a shape
  ---- TODOS ----
  @TODO : Save in other format than png
  @TODO : Save animations in gif
  @TODO : Support texturing, by extracting triangle shaped parts of an image ? [new struct : Sprite]
  @TODO : Support erasing with a clearArea() method and else
*/

import(
  "go2d/geometry"
  "image"
  "image/color"
  "image/png"
  "os"
)

const( // Differents ways of filling a shape
  ABSOLUTE_FILL = 0 // Fills every possible triangle
  ORIGIN_FILL = 1 // Takes the origin, and draws triangles around it with all the points
  NEXT_FILL = 2 // Fills triangles one by one, by taking points 3 by 3
  FIRST_FILL = 3 // Takes the first points, creates triangles with every other point 2 by 2 and draws that
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
  FillType int // How the shape should be filled, if it should
}

func NewSurface(w,h int,Filename string) Surface { // creates a new surface
  return Surface{Filename,image.NewRGBA(image.Rect(0,0,w,h)),}
}

func NewDrawableShape(border,fill color.Color) (drawable Drawable) { // Creates a drawable for a shape
  drawable.BorderColor = border
  drawable.FillColor = fill
  drawable.Thickness = 1
  drawable.FillType = FIRST_FILL
  drawable.ColorBorders = true
  drawable.Fill = true
  return
}

/* Surface methods */

func (surface *Surface) Clear() { // clears the image
  surface.Image = image.NewRGBA(surface.Image.Rect) // Creates a new image of the same size over the precedent
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
