package geometry

/*
  This package contains basic representations in plan :
    - Points
    - Lines, defined as two points connected to each other
    - Triangles, which are the basic shapes for drawing stuff
 ---- TODOS ----
 @TODO : Vectors, matrices
 @TODO : Everything about triangles
 @TODO : More calculations :
    line intersections, alignment of points, translations using lines, triangles rotation and scaling, line angles
*/

import(
  "math"
)

/* Box methods */

func IsPointInBox(min,max,test Point) bool { // Returns whether a point is in a box, with min being the top left corner and max the bottom right
  return (min.X <= test.X && min.Y <= test.Y && test.X <= max.X && test.Y <= max.Y)
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
