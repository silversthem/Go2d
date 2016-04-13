package geometry

/*
  Line methods and structure
*/

import(
  "math"
)

type Line struct { // a line is a link between two points
  Start,End Point
}

func NewLine(start,end Point) Line { // creates a line
  return Line{start,end,}
}

func NewLineFromPoint(startPoint Point,a,distance float64) (line Line) { // Creates a line from a point and an equation y = a*x+b , with end at distance from start
  line.Start = startPoint
  line.End.X = int(math.Acos(1/math.Hypot(1,a))*distance)
  line.End.Y = int(math.Asin(1/math.Hypot(1,a))*distance)
  return
}

/* Line methods */

func (line Line) IsVertical() bool { // if the line is vertical
  return line.Start.X == line.End.X
}

func (line Line) IsHorizontal() bool { // if the line is horizontal
  return line.Start.Y == line.End.Y
}

func (line Line) Angle() float64 { // Gets line oriented angle
  if line.IsVertical() {
    return math.Pi/2
  } else if line.IsHorizontal() {
    return 0
  } else {
    return math.Mod(math.Acos(1/math.Hypot(1,line.Coefficient())),math.Pi)
  }
}

func (line Line) GetPas() (pasx,pasy float64) { // Gets the amount of pixel the line moves at each iteration
  if !line.IsVertical() {
    if line.Coefficient() < 0 { // Lines goes down, reversing it
      pasx,pasy = NewLine(line.End,line.Start).GetPas()
      pasx = (-1)*pasx
      pasy = (-1)*pasy
    } else if line.Coefficient() > 1 { // Lines goes up faster than on the side
      if line.Start.X > line.End.X {
        pasx = -1
      } else {
        pasx = 1
      }
      pasy  = 1/line.Coefficient()
    } else if line.Coefficient() < 1 { // Lines goes on the side faster than up
      pasy = line.Coefficient()
      if line.Start.X > line.End.X {
        pasx = -1
      } else {
        pasx = 1
      }
    } else {
      pasx,pasy = 1,1
    }
  } else { // Vertical line
    pasx = 0
    if line.Start.Y > line.End.Y {
      pasy = -1
    } else {
      pasy = 1
    }
  }
  return
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

func (line Line) move(amount float64) (point Point) { // Returns the point we are at if we move by a certain amount along the line
  // ...
  return
}

func (line Line) GetPerpendicularLine() (l Line) { // Gets a line perpendicular to this one, of size 1
  // ...
  return
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
