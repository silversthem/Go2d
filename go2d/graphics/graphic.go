package graphics

/*
  basic graphic support
  Contains Surfaces and Drawable, Surfaces are where you draw and drawable is information about how to draw a shape
  ---- TODOS ----
  @TODO : Draw methods
  @TODO : Save in other format than png
  @TODO : Save animations in gif
  @TODO : Support texturing, by extracting triangle shaped parts of an image ? [new struct : Sprite]
*/

import(
  "go2d"
  "go2d/geometry"
  "image"
  "image/color"
  "image/png"
  "os"
  "math"
)

const(
  ABSOLUTE_FILL = 0 // Fills every possible triangle
  ORIGIN_FILL = 1 // Takes the origin, and draws triangles around it with all the points
  NEXT_FILL = 2 // Fills triangles one by one, by taking points 3 by 3
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
  drawable.FillType = ABSOLUTE_FILL
  drawable.ColorBorders = true
  drawable.Fill = true
  return
}

/* Surface methods */

func (surface *Surface) DrawPixel(x,y int,col color.Color) { // draws a pixel in the image
  surface.Image.Set(x,y,col)
}

func (surface *Surface) DrawPoint(point geometry.Point,col color.Color) { // draws a point in the image
  surface.DrawPixel(point.X,point.Y,col)
}

func (surface *Surface) DrawLine(line geometry.Line,thickness int,col color.Color) { // draws a line in the plan
  if thickness == 1 {
    if line.IsVertical() { // if the line is vertical
      if line.Start.Y > line.End.Y { // Always aligned from the shortest y to the longest
        line.Start,line.End = line.End,line.Start
      }
      for i := 0;i <= int(line.Length());i++ {
        surface.DrawPixel(line.Start.X, line.Start.Y+i, col)
      }
    } else { // Line is horizontal or different
      if line.Start.X > line.End.X { // Always aligned from the shortest x to the longest
        line.Start,line.End = line.End,line.Start
      }
      if line.IsHorizontal() { // If the line is horizontal
        for i := 0;i <= int(line.Length());i++ {
          surface.DrawPixel(line.Start.X+i, line.Start.Y, col)
        }
      } else { // Else
        a := line.Coefficient()
        pas := 1/math.Abs(a)
        if math.Abs(a) < pas {
          pas = math.Abs(a)
        }
        x,y := 0.0,0.0
        for ;x <= math.Abs(float64(line.Start.X) - float64(line.End.X));x = x + pas {
          y = x * a
          surface.DrawPixel(line.Start.X + int(x),line.Start.Y + int(y),col)
        }
      }
    }
  } else { // Line with bigger thickness

  }
}

func (surface *Surface) DrawTriangle(triangle geometry.Triangle,col color.Color) { // Draws a triangle
  surface.DrawLine(triangle.GetLine(0),1,col)
  surface.DrawLine(triangle.GetLine(1),1,col)
  surface.DrawLine(triangle.GetLine(2),1,col)
}

func (surface *Surface) DrawFillTriangle(triangle geometry.Triangle,col color.Color) { // fills a triangle section of the plan
  min,max := triangle.GetBox()
  if !triangle.IsTriangleFlat() {
    cpy := min.Y
    for ; min.X < max.X ; min.X++ {
      min.Y = cpy
      for ; min.Y < max.Y ; min.Y++ {
        if triangle.IsPointInTriangle(min) {
          surface.DrawPoint(min,col)
        }
      }
    }
  } else { // flat triangle
    surface.DrawLine(geometry.NewLine(min,max),1,col)
  }
}

func (surface *Surface) Draw(shape go2d.Shape,drawable Drawable) { // draws objects in plan
  // If object has 1 point -> point, use either fillColor or Border color
  // If object has 2 points -> line
  // else
    // if Fill is true, decompose object in triangles then fill each triangle
  // draw each lines with border
}

func (surface *Surface) Clear() { // clears the image
  // just writes over with an empty image
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
