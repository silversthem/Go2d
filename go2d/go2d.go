package go2d

/*
  go2d is a package to easily draw things into images.
  This file contains the basic interfaces for any drawing.
  ---- TODOS ----
  @TODO : Text support /!\
  @TODO : Use import C and SDL to be able to handle a window and events [new file]
  @TODO : Handle collisions [new file] using triangles and masks for textures [new file]
  @TODO : Default shapes and colors [new file] with NewX() and ColorX() functions
  @TODO : Support Math graphing [new file] and diagrams [new file]
  @TODO : Create Objects which are Drawable shapes [new file] with collisions ? with embedded event support ?
  @TODO : Multithreading to update images while doing calculations, to make the library able to show cool, real time, rendering of another process [new file]
  @TODO : Shaders [new file]
  @TODO : Complex support in graphes [new file]
  @TODO : Nodes and stuff, to build cool diagrams [new file]
*/

import(
  "go2d/geometry"
  "image/color"
)

/* Primary Interfaces */

type Surface interface { // A surface is something on which you can draw
  DrawPoint(geometry.Point,color.Color) // Drawing a colored point
  DrawLine(geometry.Line,Drawable)      // Draws a line
  Draw(Rendered)                        // Draws something in the surface
}

type Drawable interface { // Basic drawing instructions
  GetPointColor(geometry.Point) color.Color // Returns the color of a point
}

type Rendered interface { // Everything that can be drawn
  Draw(Surface) // Returns true while there's points to draw
}
