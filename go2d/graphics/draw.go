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
  "math"
)

/* Drawing methods */

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
  } else { // Thicker lines
    // Get line perpendicular to the line
    // Translate using this line to two lines at half thickness distance each
    // Draws the rectangles created by the 2 lines
  }
}

func (surface *Surface) DrawTriangle(triangle geometry.Triangle,col color.Color) { // Draws a triangle
  surface.DrawLine(triangle.GetLine(0),1,col)
  surface.DrawLine(triangle.GetLine(1),1,col)
  surface.DrawLine(triangle.GetLine(2),1,col)
}

func (surface *Surface) DrawFillTriangle(triangle geometry.Triangle,col color.Color) { // fills a triangle section of the plan
  min,max := triangle.GetBox()
  if !triangle.IsTriangleFlat() { // Only drawing triangles, not lines
    cpy := min.Y
    for ; min.X < max.X ; min.X++ {
      min.Y = cpy
      for ; min.Y < max.Y ; min.Y++ {
        if triangle.IsPointInTriangle(min) {
          surface.DrawPoint(min,col)
        }
      }
    }
  }
}

func (surface *Surface) Draw(shape go2d.Shape,drawable Drawable) { // draws objects in plan
  if len(shape.Points) == 2 { // Shape is a line
    if drawable.ColorBorders { // Coloring borders
      surface.DrawLine(shape.GetLine(0),drawable.Thickness,drawable.BorderColor)
    }
  } else if len(shape.Points) > 2 { // Shape is a convex shape
    if drawable.Fill { // Filling object
      triangles := shape.GetTrianglesCount()
      switch drawable.FillType { // Drawing differently depending on shape
      case ABSOLUTE_FILL: // Drawing every possible triangle in the shape
        for i := 0;i < triangles;i++ {
          for j := 0;j < triangles;j++ {
            surface.DrawFillTriangle(shape.GetAbsoluteTriangle(i,j),drawable.FillColor)
          }
        }
      case ORIGIN_FILL: // Drawing every triangle from points two by two and origin
        for i := 0;i < triangles;i++ {
          surface.DrawFillTriangle(shape.GetTriangleFromOrigin(i),drawable.FillColor)
        }
      case NEXT_FILL: // Drawing triangles from points 3 by 3
        for i := 0;i < triangles;i++ {
          surface.DrawFillTriangle(shape.GetNextTriangle(i),drawable.FillColor)
        }
      case FIRST_FILL: // Drawing triangles from points two by two and the first point
        for i := 0;i < triangles;i++ {
          surface.DrawFillTriangle(shape.GetAbsoluteTriangle(i,0),drawable.FillColor)
        }
      }
    }
    if drawable.ColorBorders { // Coloring borders
      lines := shape.GetLinesCount()
      for i := 0;i < lines;i++ {
        surface.DrawLine(shape.GetLine(i),drawable.Thickness,drawable.BorderColor)
      }
    }
  }
}
