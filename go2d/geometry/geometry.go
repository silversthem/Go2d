package geometry
/*
 Basic operations in plan
 contains everything you could need in a plan :
  Points, Lines, Triangles (all shapes are just triangles next to each other)
  Scales (useful for sizing and scaling) and Transform (Transformations applied to an object) => resizing is done with a scale (Size), and rotating with Angle in rads
 Also contains everything needed for plan calculations
 ---- TODOS ----
 @TODO : Everything about triangles
 @TODO : Transform handling
 @TODO : More calculations :
    line intersections, alignment of points, translations using lines, triangles rotation and scaling, line angles
*/

import(
  "image"
  "math"
)

type Point image.Point // a point in plan

type Line struct { // a line is a link between two points
  Start,End Point
}

type Triangle struct { // 3 lines forms a triangle, useful when testing collisions or filling a shape
  Points [3]Point
}

type Scale struct { // a scale for an object
  Width,Height float64
}

type Transform struct { // list a transformations a shape can undergo
  Angle float64 // angle, in radian
  Size Scale // Scaling
}

func NewPoint(x,y int) Point { // creates a point
  return Point{x,y,}
}

func NewLine(start,end Point) Line { // creates a line
  return Line{start,end,}
}

func NewTriangle(point,point2,point3 Point) Triangle { // creates a triangle
  return Triangle{
      [3]Point{point,point2,point3,},}
}

func NewScale(w,h float64) Scale { // creates a scale structure
  return Scale{w,h,}
}

func NewTransform() Transform { // creates a Transform structure
  return Transform{0.0,NewScale(1.0,1.0),}
}

/* Line methods */

func (line Line) Distance() float64 { // gets distance between two line points
  return line.Start.Distance(line.End)
}

func (line Line) ToVector() Point { // converts a line into a point -> a vector
  return line.End.Substract(line.Start)
}

func (line Line) Opposite() Line { // Returns the opposite line (swaps Start and End points)
  return Line{line.End,line.Start,}
}

func (line Line) IsPoint() bool { // if the line is a point
  // ...
  return true
}

/* Line methods with other lines */

func (line Line) DotProductLine(line2 Line) float64 { // calculates the dot product of two lines
  vect,vect2 := line.ToVector(),line2.ToVector()
  return float64(vect.X)*float64(vect2.X) + float64(vect.Y)*float64(vect2.Y)
}

func (line Line) AngleLine(line2 Line) float64 { // calculates the angle between two lines
  dotProduct := line.DotProductLine(line2)
  dist,dist2 := line.Distance(),line2.Distance()
  if dist == 0.0 || dist2 == 0.0 {
    panic("Line is actually a point") // Should AnglePoint be used directly ?
  } else {
    return math.Acos(dotProduct/(dist*dist2))
  }
}

func (line Line) Intersects(line2 Line) (point Point,test bool) { // if two line intersects
  // ...
  return
}

func (line) IsAlignedLine(line Line) bool { // if the line is aligned with another line
  // ...
  return true
}

/* Line methods with points */

func (line) IsAlignedPoint(point Point) bool { // if the line is aligned with a point
  // ...
  return true
}

func (line Line) DotProductPoint(point Point) float64 { // calculates the dot product of a line and a point
  return line.DotProductLine(NewLine(line.Start,point))
}

func (line Line) AnglePoint(point Point) float64 { // gets the oriented angle between two lines : (Line,Line2)
  return line.AngleLine(NewLine(line.Start,point))
}

/* Triangle methods */

func (triangle Triangle) GetLine(index int) Line { // returns one of the 3 lines of a triangle
  if index > 2 {
    panic("Trying to get non existant line in triangle")
  }
  return NewLine(triangle.Points[index],triangle.Points[index+1])
}

func IsTriangleFlat() bool { // if the triangle is float64
  return true
}

func (triangle Triangle) GetAcuteAngleLines() (line,line2 Line) { // gets two lines forming an acute angle
  return
}

func (triangle Triangle) GetBox() (min,max Point) { // returns the box in which the triangle is in
  min,max = triangle.Points[0],triangle.Points[0]
  for i := 1 ; i < 3 ; i++ { // Finding maximum and minimum
    max.X = int(math.Max(float64(max.X),float64(triangle.Points[i].X))) // Maximum
    max.Y = int(math.Max(float64(max.Y),float64(triangle.Points[i].Y)))
    min.X = int(math.Min(float64(min.X),float64(triangle.Points[i].X))) // Minimum
    min.Y = int(math.Min(float64(min.Y),float64(triangle.Points[i].Y)))
  }
  return
}

func (triangle Triangle) IsPointInTriangle(point Point) bool { // checks if a point is in the triangle
  min,max := triangle.GetBox()
  if IsPointInBox(min,max,point) { // if the point is in the triangle box first
    // we find two lines forming an acute angle in the triangle
    // Expressing point as a sum of the two lines
    // Checking if factors are both between [0 ; 1] -> true else false
  }
  return false
}

func (triangle Triangle) Intersects(line Line) bool { // if a line goes through a triangle
  // tests side by side
  return true
}

/* Box method */

func IsPointInBox(min,max,test Point) bool { // Returns whether a point is in a box, with min being the top left corner and max the bottom right
  return (min.X <= test.X && min.Y <= test.Y && test.X <= max.X && test.Y <= max.Y)
}

/* Point methods */

func (point Point) Distance(point2 Point) float64 { // gets the distance between two points
  return math.Sqrt(math.Pow(float64(point.X - point2.X),2) + math.Pow(float64(point.Y - point2.Y),2)) // sqrt((x1-x2)^2 + (y1-y2)^2)
}

func (point Point) Substract(point2 Point) Point { // Substract point2 coordinates from point
  return Point{
    point.X - point2.X,
    point.Y - point2.Y,
  }
}

/* Transform methods */

func (transform Transform) GetPoint(point Point) Point { // transforms a point
  point.X = int(float64(point.X)*transform.Size.Width) // Scaling
  point.Y = int(float64(point.Y)*transform.Size.Height)
  return Point{
    int(float64(point.X)*math.Cos(transform.Angle) + float64(point.Y)*math.Sin(transform.Angle)), // rotating
    int(float64((-1)*point.X)*math.Sin(transform.Angle) + float64(point.Y)*math.Cos(transform.Angle)),
  }
}

func (transform Transform) GetLine(line Line) Line { // transforms a line
  line.Start = transform.GetPoint(line.Start)
  line.End = transform.GetPoint(line.End)
  return line
}

func (transform Transform) GetTriangle(triangle Triangle) Triangle { // transforms a triangle
  return Triangle{[3]Point{
    transform.GetPoint(triangle.Points[0]),
    transform.GetPoint(triangle.Points[1]),
    transform.GetPoint(triangle.Points[2]),}}
}

func (transform *Transform) Rotate(angle float64) { // Adds rotation to the transformation
  // ...
}

func (transform *Transform) Zoom(scale float64) { // Zooms/Unzooms the transformation
  // ...
}

/* Angle calculations */

func GetAbsoluteAngleValue(angle float64) float64 { // gets angle absolute value : [0 ; 2Pi]
  return math.Mod(angle,2*math.Pi)
}

/* Angle testing */

func IsAngleFlat(angle float64) bool { // if the absolute angle value is = 0 or Pi or 2Pi
  return (angle == 0 || angle == math.Pi || angle == 2*math.Pi)
}

func IsAngleRight(angle float64) bool { // if the absolute angle value is = Pi/2 or 3Pi/2
  return (angle == math.Pi/2 || angle == (3*math.Pi)/2)
}

func IsAngleAcute(angle float64) bool { // if the absolute angle value is ]0 ; Pi/2[ or ]3Pi/2;2Pi[
  return ((0 < angle && angle < math.Pi/2) || ((3*math.Pi)/2 < angle && angle < 2*math.Pi))
}

func IsAngleObtuse(angle float64) bool { // if the angle absolute value ]Pi/2 ; Pi[ or ]Pi ; 3Pi/2[
  return ((math.Pi/2 < angle && angle < math.Pi) || (math.Pi < angle && angle < (3*math.Pi)/2))
}
