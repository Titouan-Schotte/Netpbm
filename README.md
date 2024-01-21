# Netpbm

## Aim

The aim of this project is to create a library for working with images (*.pbm, *.pgm, *.ppm). The library should be able to read and write images, and manipulate them.

This project is guided, so you will be given a list of functions to implement.

**WARNING** : you should use the same names for the functions and the types as the ones given in the instructions. If you don't, the tests will fail.

## tools

- [PBM/PGM/PPM Viewer VScode](https://marketplace.visualstudio.com/items?itemName=ngtystr.ppm-pgm-viewer-for-vscode)

## Create a go library

1. Create a new repository on GitHub named `Netpbm`.
> if you don't have a GitHub account, create one.
2. Clone the repository on your computer.
3. Create a new go module named `github.com/<username>/Netpbm`.
4. push the module to GitHub.

Well done, you are ready to start coding.

## PBM

[PBM](#file-pbm-md) (Portable BitMap) is a file format for images. It is a very simple format, and is easy to read and write.

[FileTest](#file-duck-pbm)

### Functions to implement

1. Take this struct and add it to your code.

```go
type PBM struct{
    data [][]bool
    width, height int
    magicNumber string
}
```

2. Write a function that takes a PBM image and returns a struct that represents the image.
> The function must be take all formats of PBM images into account (P1 and P4).

```go
// ReadPBM reads a PBM image from a file and returns a struct that represents the image.
func ReadPBM(filename string) (*PBM, error){
    // ...
}
```

3. Write a function attached to the PBM struct that returns the width and height.

```go
// Size returns the width and height of the image.
func (pbm *PBM) Size() (int,int){
    // ...
}
```

4. Write a function to get the value of a pixel.

```go
// At returns the value of the pixel at (x, y).
func (pbm *PBM) At(x, y int) bool{
    // ...
}
```

5. Write a function to set the value of a pixel.

```go
// Set sets the value of the pixel at (x, y).
func (pbm *PBM) Set(x, y int, value bool){
    // ...
}
```

6. Write a function to save the image to a file to the same format as the original image.

```go
// Save saves the PBM image to a file and returns an error if there was a problem.
func (pbm *PBM) Save(filename string) error{
    // ...
}
```

7. Write a function to invert the colors of the image.

```go
// Invert inverts the colors of the PBM image.
func (pbm *PBM) Invert(){
    // ...
}
```

8. Write a function to flip the image horizontally.

```go
// Flip flips the PBM image horizontally.
func (pbm *PBM) Flip(){
    // ...
}
```

9. Write a function to flop the image vertically.

```go
// Flop flops the PBM image vertically.
func (pbm *PBM) Flop(){
    // ...
}
```

10. Write a function to set the magic number of the image.

```go
// SetMagicNumber sets the magic number of the PBM image.
func (pbm *PBM) SetMagicNumber(magicNumber string){
    // ...
}
```

## PGM

[PGM](#file-pgm-md) (Portable GrayMap) is a file format for images. It is a very simple format, and is easy to read and write.

[FileTest](#file-duck-pgm)

### Functions to implement

1. Take this struct and add it to your code.

```go
type PGM struct{
    data [][]uint8
    width, height int
    magicNumber string
    max uint
}
```

2. Write a function that takes a PGM image and returns a struct that represents the image.
> The function must be take all formats of PGM images into account (P2 and P5).

```go
// ReadPGM reads a PGM image from a file and returns a struct that represents the image.
func ReadPGM(filename string) (*PGM, error){
    // ...
}
```

3. Write a function attached to the PGM struct that returns the width and height.

```go
// Size returns the width and height of the image.
func (pgm *PGM) Size() (int,int){
    // ...
}
```

4. Write a function to get the value of a pixel.

```go
// At returns the value of the pixel at (x, y).
func (pgm *PGM) At(x, y int) uint8{
    // ...
}
```

5. Write a function to set the value of a pixel.

```go
// Set sets the value of the pixel at (x, y).
func (pgm *PGM) Set(x, y int, value uint8){
    // ...
}
```

6. Write a function to save the image to a file to the same format as the original image.

```go
// Save saves the PGM image to a file and returns an error if there was a problem.
func (pgm *PGM) Save(filename string) error{
    // ...
}
```

7. Write a function to invert the colors of the image.

```go
// Invert inverts the colors of the PGM image.
func (pgm *PGM) Invert(){
    // ...
}
```

8. Write a function to flip the image horizontally.

```go
// Flip flips the PGM image horizontally.
func (pgm *PGM) Flip(){
    // ...
}
```

9. Write a function to flop the image vertically.

```go
// Flop flops the PGM image vertically.
func (pgm *PGM) Flop(){
    // ...
}
```

10. Write a function to set the magic number of the image.

```go
// SetMagicNumber sets the magic number of the PGM image.
func (pgm *PGM) SetMagicNumber(magicNumber string){
    // ...
}
```

11. Write a function to set the max value of the image.

```go
// SetMaxValue sets the max value of the PGM image.
func (pgm *PGM) SetMaxValue(maxValue uint8){
    // ...
}
```

12. Write a function to rotate the image 90° clockwise.

```go
// Rotate90CW rotates the PGM image 90° clockwise.
func (pgm *PGM) Rotate90CW(){
    // ...
}
```

13. Write a function to convert the image to PBM.

```go
// ToPBM converts the PGM image to PBM.
func (pgm *PGM) ToPBM() *PBM{
    // ...
}
```

## PPM

[PPM](#file-ppm-md) (Portable PixMap) is a file format for images. It is a very simple format, and is easy to read and write.

[FileTest](#file-duck-ppm)

### Functions to implement

1. Take this struct and add it to your code.

```go
type PPM struct{
    data [][]Pixel
    width, height int
    magicNumber string
    max uint
}

type Pixel struct{
    R, G, B uint8
}
```

2. Write a function that takes a PPM image and returns a struct that represents the image.
> The function must be take all formats of PPM images into account (P3 and P6).

```go
// ReadPPM reads a PPM image from a file and returns a struct that represents the image.
func ReadPPM(filename string) (*PPM, error){
    // ...
}
```

3. Write a function attached to the PPM struct that returns the width and height.

```go
// Size returns the width and height of the image.
func (ppm *PPM) Size() (int,int){
    // ...
}
```

4. Write a function to get the value of a pixel.

```go
// At returns the value of the pixel at (x, y).
func (ppm *PPM) At(x, y int) Pixel{
    // ...
}
```

5. Write a function to set the value of a pixel.

```go
// Set sets the value of the pixel at (x, y).
func (ppm *PPM) Set(x, y int, value Pixel){
    // ...
}
```

6. Write a function to save the image to a file to the same format as the original image.

```go
// Save saves the PPM image to a file and returns an error if there was a problem.
func (ppm *PPM) Save(filename string) error{
    // ...
}
```

7. Write a function to invert the colors of the image.

```go
// Invert inverts the colors of the PPM image.
func (ppm *PPM) Invert(){
    // ...
}
```

8. Write a function to flip the image horizontally.

```go
// Flip flips the PPM image horizontally.
func (ppm *PPM) Flip(){
    // ...
}
```

9. Write a function to flop the image vertically.

```go
// Flop flops the PPM image vertically.
func (ppm *PPM) Flop(){
    // ...
}
```

10. Write a function to set the magic number of the image.

```go
// SetMagicNumber sets the magic number of the PPM image.
func (ppm *PPM) SetMagicNumber(magicNumber string){
    // ...
}
```

11. Write a function to set the max value of the image.

```go
// SetMaxValue sets the max value of the PPM image.
func (ppm *PPM) SetMaxValue(maxValue uint8){
    // ...
}
```

12. Write a function to rotate the image 90° clockwise.

```go
// Rotate90CW rotates the PPM image 90° clockwise.
func (ppm *PPM) Rotate90CW(){
    // ...
}
```

13. Write a function to convert the image to PGM.

```go
// ToPGM converts the PPM image to PGM.
func (ppm *PPM) ToPGM() *PGM{
    // ...
}
```

14. Write a function to convert the image to PBM.

```go
// ToPBM converts the PPM image to PBM.
func (ppm *PPM) ToPBM() *PBM{
    // ...
}
```

15. Create a struct to represent a point in the image.

```go
type Point struct{
    X, Y int
}
```

16. Write a function to draw a line between two points.

```go
// DrawLine draws a line between two points.
func (ppm *PPM) DrawLine(p1, p2 Point, color Pixel){
    // ...
}
```

17. Write a function to draw a rectangle.

```go
// DrawRectangle draws a rectangle.
func (ppm *PPM) DrawRectangle(p1 Point, width , height int, color Pixel){
    // ...
}
```

18. Write a function to draw filled rectangle.

```go
// DrawFilledRectangle draws a filled rectangle.
func (ppm *PPM) DrawFilledRectangle(p1 Point, width , height int, color Pixel){
    // ...
}
```

19. Write a function to draw a circle.

```go
// DrawCircle draws a circle.
func (ppm *PPM) DrawCircle(center Point, radius int, color Pixel){
    // ...
}
```

20. Write a function to draw a filled circle.

```go
// DrawFilledCircle draws a filled circle.
func (ppm *PPM) DrawFilledCircle(center Point, radius int, color Pixel){
    // ...
}
```

21. Write a function to draw a triangle.

```go
// DrawTriangle draws a triangle.
func (ppm *PPM) DrawTriangle(p1, p2, p3 Point, color Pixel){
    // ...
}
```

22. Write a function to draw a filled triangle.

```go
// DrawFilledTriangle draws a filled triangle.
func (ppm *PPM) DrawFilledTriangle(p1, p2, p3 Point, color Pixel){
    // ...
}
```

23. Write a function to draw a polygon.

```go
// DrawPolygon draws a polygon.
func (ppm *PPM) DrawPolygon(points []Point, color Pixel){
    // ...
}
```

24. Write a function to draw a filled polygon.

```go
// DrawFilledPolygon draws a filled polygon.
func (ppm *PPM) DrawFilledPolygon(points []Point, color Pixel){
    // ...
}
```

25. Write a function to draw Koch snowflake.

```go
// DrawKochSnowflake draws a Koch snowflake.
func (ppm *PPM) DrawKochSnowflake(n int, start Point,width int,color Pixel){
    // N is the number of iterations.
    // Koch snowflake is a 3 times a Koch curve.
    // Start is the top point of the snowflake.
    // Width is the width all the lines.
    // Color is the color of the lines.
    // ...
}
```

26. Write a function to draw a Sierpinski triangle.

```go
// DrawSierpinskiTriangle draws a Sierpinski triangle.
func (ppm *PPM) DrawSierpinskiTriangle(n int, start Point,width int,color Pixel){
    // N is the number of iterations.
    // Start is the top point of the triangle.
    // Width is the width all the lines.
    // Color is the color of the lines.
    // ...
}
```

27. Write a function to draw perlin noise.

```go
// DrawPerlinNoise draws perlin noise.
// this function Draw a perlin noise of all the image.
func (ppm *PPM) DrawPerlinNoise(color1 Pixel , color2 Pixel){
    // Color1 is the color of 0.
    // Color2 is the color of 1.    
}
```

28. Write a function k-nearest neighbors to resize the image.

The k-nearest neighbors algorithm is a method for resizing images. It works by taking the k nearest pixels to the pixel that is being resized, and averaging their colors.

[k-nearest neighbors](https://en.wikipedia.org/wiki/K-nearest_neighbors_algorithm)

```go
// KNearestNeighbors resizes the PPM image using the k-nearest neighbors algorithm.
func (ppm *PPM) KNearestNeighbors(newWidth, newHeight int){
    // ...
}
```
