# Go2d
Simple packages to draw in your new favorite language
## How to use ?
Just install the go2d package by putting it in your $GOPATH and doing :
`go install go2d`

### What next ?

Well, you can create a surface to draw into and shapes to be drawn in that surface. Note however that the following are not
currently implemented (even though I want them to be, so eventually they'll make it) :
  - [ ] Using images as textures for shapes
  - [ ] Using shaders in shapes
  - [ ] Graphing mathematical functions
  - [ ] Write text
  - [ ] Gifs (No idea how it's going to look like)
  - [ ] Basic geometry operations, like telling when lines cross each other and stuff
  - [ ] Same, but with shapes
  - [ ] Drawing using functions (functional style)
  - [ ] Super cool shaders
  - [ ] Creating diagrams and stuff
  
#### What can it be used for ?

Obviously the project is still very young, but here's what we can hope for :
  - [x] Drawing simple images with shapes
  - [ ] ... And images
  - [ ] ... And text
  - [ ] Basic geometry/graphing operations
  - [ ] Basic diagrams features
  - [ ] Gifs
  - [ ] A visual log (by rendering while you're handling data or calculating stuff)
  - [ ] Generating images for something else (for example, creating images when called by a python script in a website)
  
##### And then ?

At some point, I would love to connect this to a G.U.I library, and build things such as an event machine and openGL
rendering, but these are far off in the future.

## Where do my drawings go ?

Drawings can now be saved as png files, but more will come.

## What can I draw ?

You can draw :
  - Pretty much any convex shape
  - Lines
  - Individual points
  - Shaders (Not real graphic card shaders (I wish they were)) but drawing using functions and coordinates

You can also draw using pixel shaders using a function like this one : 
```go
func myPixelShader(x,y int) color.Color {
  // Returning a color for the pixel in (x,y)
}
```
And then use a `image.Rectangle` to tell the square in which this function has to be applied.

#### Other shaders

More will come...

# How to use by examples

More will come...
