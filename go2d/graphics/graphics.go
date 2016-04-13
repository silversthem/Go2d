package graphics

/*
  basic graphic support
  Contains Surface, a structure to draw in png files
  ---- TODOS ----
  @TODO : Save in other format than png
  @TODO : Save animations in gif
  @TODO : Support texturing, by extracting triangle shaped parts of an image ? [new struct : Sprite]
  @TODO : Support erasing with a clearArea() method and else
*/

import(
  "go2d/geometry"
  "image"
  "image/png"
  "os"
)

type Surface struct { // A surface is a structure containing a plan a filename and an actual image, to draw into
  Filename string
  Image *image.RGBA
}

func NewSurface(w,h int,Filename string) Surface { // creates a new surface
  return Surface{Filename,image.NewRGBA(image.Rect(0,0,w,h)),}
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
