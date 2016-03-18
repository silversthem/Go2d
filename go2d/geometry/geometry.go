package geometry
/*
 Basic operations in plan
 contains everything you could need in a plan :
  Points, Lines, Triangles (all shapes are just triangles next to each other)
  Scales (useful for sizing and scaling) and Transform (Transformations applied to an object) => resizing is done with a scale (Size), and rotating with Angle in rads
 Also contains everything needed for plan calculations
 ---- TODOS ----
 @TODO : Everything about triangles
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

func NewLineFromEquation(startPoint Point,a,b,distance float64) { // Creates a line from a point and an equation y = a*x+b , with end at distance from start
  
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

func (line Line) IsVertical() bool { // if the line is vertical
  return line.Start.X == line.End.X
}

func (line Line) IsHorizontal() bool { // if the line is horizontal
  return line.Start.Y == line.End.Y
}

func (line Line) Coefficient() float64 { // line coefficient the a in y = a*x + b
  return float64(line.Start.Y - line.End.Y)/float64(line.Start.X - line.End.X)
}

func (line Line) LinearConstant() float64 { // A line is y = a*x+b, this function returns b, Coefficient returns a
  return float64(line.Start.Y) - float64(line.Start.X)*line.Coefficient()
}

func (line Line) Length() float64 { // gets distance between two line points
  return line.Start.Distance(line.End)
}

func (line Line) ToVector() Point { // converts a line into a point -> a vector
  return line.End.Substract(line.Start)
}

func (line Line) Opposite() Line { // Returns the opposite line (swaps Start and End points)
  return Line{line.End,line.Start,}
}

func (line Line) IsPoint() bool { // if the line is a point
  return line.Start.Y == line.End.Y && line.Start.X == line.End.X
}

/* Line methods with other lines */

func (line Line) DotProductLine(line2 Line) float64 { // calculates the dot product of two lines
  vect,vect2 := line.ToVector(),line2.ToVector()
  return (float64(vect.X)*float64(vect2.X) + float64(vect.Y)*float64(vect2.Y))/math.Pow(line.Length(),2)
}

func (line Line) AngleLine(line2 Line) float64 { // calculates the angle between two lines
  dotProduct := line.DotProductLine(line2)
  dist,dist2 := line.Length(),line2.Length()
  if dist == 0.0 || dist2 == 0.0 {
    panic("Line is actually a point") // Should AnglePoint be used directly ?
  } else {
    return math.Acos(dotProduct/(dist*dist2))
  }
}

func (line Line) Intersects(line2 Line) (point Point,test bool) { // if two line intersects
  if line.IsVertical() {
    if !line2.IsVertical() {
      point.X = line.Start.X
      point.Y = int(float64(line.Start.X)*line.Coefficient() + line.LinearConstant())
    } else {
      if line.Start.X == line2.Start.X { // Overlapping
        test = true
      } else {
        test = false
      }
    }
  } else if line2.IsVertical() {
    point,test = line2.Intersects(line)
  } else {
    coeff,coeff2 := line.Coefficient(),line2.Coefficient()
    lnc,lnc2 := line.LinearConstant(),line2.LinearConstant()
    if coeff == coeff2 { // Parallel lines
      test = false
    } else { // solving x*a + b = x*a' + b' <=> x*(a - a') = b' - b <=> x = (b' - b)/(a - a')
      test = true // We know that non parallel line collides at some point
      point.X = int((lnc2 - lnc)/(coeff - coeff2))
      point.Y = int(float64(point.X)*coeff + lnc)
    }
  }
  return
}

func (line Line) IsAlignedLine(line2 Line) bool { // if the line is aligned with another line
  return line.IsAlignedPoint(line2.Start) && line.IsAlignedPoint(line2.End)
}

/* Line methods with points */

func (line Line) IsAlignedPoint(point Point) bool { // if the line is aligned with a point
  if line.IsVertical() { // vertical line
    return point.X == line.Start.X
  } else if line.IsHorizontal() { // horizontal line
    return point.Y == line.Start.Y
  } else { // a line
    test := float64(point.X)*line.Coefficient() + line.LinearConstant()
    return point.Y == int(test)
  }
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
  return NewLine(triangle.Points[index],triangle.Points[int(math.Mod(float64(index+1),3.0))])
}

func (triangle Triangle) IsTriangleFlat() bool { // if the triangle is flat
  return triangle.GetLine(0).IsAlignedPoint(triangle.Points[2])
}

func (triangle Triangle) GetAcuteAngleLines() (line,line2 Line,angleStart Point) { // gets two lines forming an acute angle and the angle point
  line = NewLine(triangle.Points[0],triangle.Points[1])
  line2 = NewLine(triangle.Points[1],triangle.Points[2])
  angleStart = triangle.Points[1]
  angle := line.AngleLine(line2)
  if IsAngleAcute(angle) {
    return
  } else if IsAngleFlat(angle) {
    panic("Triangle is flat")
  } else {
    line2 = NewLine(triangle.Points[0],triangle.Points[2])
    angleStart = triangle.Points[0]
    return
  }
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

func (triangle Triangle) GetBarycentric(point Point) (alpha,beta,gamma float64) { // Returns barycentric coordinates of a point the triangle
  p1,p2,p3 := triangle.Points[0],triangle.Points[1],triangle.Points[2]
  alpha = float64(((p2.Y - p3.Y)*(point.X - p3.X) + (p3.X - p2.X)*(point.Y - p3.Y)))/float64(((p2.Y - p3.Y)*(p1.X - p3.X) + (p3.X - p2.X)*(p1.Y - p3.Y)))
  beta = float64(((p3.Y - p1.Y)*(point.X - p3.X) + (p1.X - p3.X)*(point.Y - p3.Y)))/float64(((p2.Y - p3.Y)*(p1.X - p3.X) + (p3.X - p2.X)*(p1.Y - p3.Y)))
  gamma = 1 - alpha - beta
  return
}

func (triangle Triangle) IsPointInTriangle(point Point) bool { // checks if a point is in the triangle
  min,max := triangle.GetBox()
  if IsPointInBox(min,max,point) { // if the point is in the triangle box first
    alpha,beta,gamma := triangle.GetBarycentric(point)
    return 0 <= alpha && alpha <= 1 && 0 <= beta && beta <= 1 && 0 <= gamma && gamma <= 1
  }
  return false
}

func (triangle Triangle) IsLineInTriangle(line Line) bool { // If a line is inside the triangle
  return triangle.IsPointInTriangle(line.Start) && triangle.IsPointInTriangle(line.End)
}

func (triangle Triangle) Intersects(line Line) bool { // if a line goes through a triangle
  // Getting the three sides of the triangle
  s1,s2,s3 := NewLine(triangle.Points[0],triangle.Points[1]),NewLine(triangle.Points[1],triangle.Points[2]),NewLine(triangle.Points[2],triangle.Points[0])
  _,t1 := line.Intersects(s1)
  _,t2 := line.Intersects(s2)
  _,t3 := line.Intersects(s3)
  return t1 || t2 || t3
}

/* Box methods */

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
  transform.Angle = GetAbsoluteAngleValue(angle+transform.Angle);
}

func (transform *Transform) Zoom(scale float64) { // Zooms/Unzooms the transformation
  transform.Size.Width = scale*transform.Size.Width
  transform.Size.Height = scale*transform.Size.Height
}

func (transform *Transform) SetRotation(angle float64) { // Sets rotation
  transform.Angle = angle;
}

func (transform *Transform) SetZoom(scale float64) { // Sets zoom
  transform.Size.Width = scale
  transform.Size.Height = scale
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
