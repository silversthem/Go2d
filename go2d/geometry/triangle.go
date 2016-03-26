package geometry

/*
  A triangle is the basic shape for a 2D plan, used to fill shapes
*/

import(
  "math"
)

type Triangle struct { // 3 lines forms a triangle, useful when testing collisions or filling a shape
  Points [3]Point
}

func NewTriangle(point,point2,point3 Point) Triangle { // creates a triangle
  return Triangle{
      [3]Point{point,point2,point3,},}
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
