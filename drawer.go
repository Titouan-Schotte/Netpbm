package Netpbm

/*
 * Titouan Schotté
 *  PPM Drawer
 */

import (
	"math"
	"math/rand"
	"sort"
)

// STRUCT POINT
type Point struct {
	X, Y int
}

// DrawLine draws a line between two points.
func (ppm *PPM) DrawLine(p1, p2 Point, color Pixel) {
	// To do this, you will need to use the Bresenham algorithm.: https://fr.wikipedia.org/wiki/Algorithme_de_trac%C3%A9_de_segment_de_Bresenham
	// Source algorithm : http://fredericgoset.ovh/mathematiques/courbes/fr/bresenham_line.html

	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	incX := Sgn(dx)
	incY := Sgn(dy)
	dx = Abs(dx)
	dy = Abs(dy)

	if dx == 0 {
		// Vertical line
		for y := p1.Y; y != p2.Y+incY; y += incY {
			ppm.Set(p1.X, y, color)
		}
	} else if dy == 0 {
		// Horizontal line
		for x := p1.X; x != p2.X+incX; x += incX {
			ppm.Set(x, p1.Y, color)
		}
	} else if dx >= dy {
		// More horizontal than vertical
		d := 2*dy - dx
		y := p1.Y
		for x := p1.X; x != p2.X+incX; x += incX {
			ppm.Set(x, y, color)
			if d > 0 {
				y += incY
				d -= 2 * dx
			}
			d += 2 * dy
		}
	} else {
		// More vertical than horizontal
		d := 2*dx - dy
		x := p1.X
		for y := p1.Y; y != p2.Y+incY; y += incY {
			ppm.Set(x, y, color)
			if d > 0 {
				x += incX
				d -= 2 * dy
			}
			d += 2 * dx
		}
	}
}

// DrawRectangle draws a rectangle.
func (ppm *PPM) DrawRectangle(p1 Point, width, height int, color Pixel) {
	ppm.DrawLine(p1, Point{X: p1.X + width, Y: p1.Y}, color)
	ppm.DrawLine(Point{X: p1.X + width, Y: p1.Y}, Point{X: p1.X + width, Y: p1.Y + height}, color)
	ppm.DrawLine(Point{X: p1.X + width, Y: p1.Y + height}, Point{X: p1.X, Y: p1.Y + height}, color)
	ppm.DrawLine(Point{X: p1.X, Y: p1.Y + height}, p1, color)
}

// DrawFilledRectangle draws a filled rectangle.
func (ppm *PPM) DrawFilledRectangle(p1 Point, width, height int, color Pixel) {
	// We traverse each point in Y to create a line that crosses the rectangle
	for y := p1.Y; y <= p1.Y+height; y++ {
		println(y)
		ppm.DrawLine(Point{X: p1.X, Y: y}, Point{X: p1.X + width, Y: y}, color)
	}
}

// DrawCircle draws a circle.
func (ppm *PPM) DrawCircle(center Point, radius int, color Pixel) {
	// Bresenham's Circle Drawing Algorithm (http://profmath.uqam.ca/~boileau/GRMS2014/cercles.html)

	// Initialize starting values for drawing the circle
	x := radius
	y := 0
	m := 0

	// Loop through the circle points using Bresenham's algorithm
	for x >= y {
		// Avoid drawing points on the circle and corners of the square
		if y > 0 && x > 0 {
			ppm.Set(center.X+x, center.Y-y, color)
			ppm.Set(center.X+y, center.Y-x, color)
			ppm.Set(center.X-y, center.Y-x, color)
			ppm.Set(center.X-x, center.Y-y, color)
			ppm.Set(center.X-x, center.Y+y, color)
			ppm.Set(center.X-y, center.Y+x, color)
			ppm.Set(center.X+y, center.Y+x, color)
			ppm.Set(center.X+x, center.Y+y, color)
		}

		// Increment y and update the decision parameter m
		y++
		if m <= 0 {
			m += 2*y + 1
		}

		// Decrement x and update the decision parameter m
		if m > 0 {
			x--
			m -= 2*x + 1
		}
	}

	// Draw the cardinal points to complete the circle
	ppm.Set(center.X+radius-1, center.Y, color)
	ppm.Set(center.X, center.Y+radius-1, color)
	ppm.Set(center.X-radius+1, center.Y, color)
	ppm.Set(center.X, center.Y-radius+1, color)
}

// DrawFilledCircle draws a filled circle.
func (ppm *PPM) DrawFilledCircle(center Point, radius int, color Pixel) {
	x := radius
	y := 0
	m := 0

	for x >= y {
		// Avoid drawing points on the circle and corners of the square
		if y > 0 && x > 0 {
			// Draw lines connecting points symmetrically to fill the circle
			ppm.DrawLine(Point{center.X + x, center.Y - y}, Point{center.X - x, center.Y + y}, color)
			ppm.DrawLine(Point{center.X + x, center.Y - y}, Point{center.X - y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X + x, center.Y - y}, Point{center.X + y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X + x, center.Y - y}, Point{center.X + x, center.Y + y}, color)

			ppm.DrawLine(Point{center.X + y, center.Y - x}, Point{center.X - x, center.Y + y}, color)
			ppm.DrawLine(Point{center.X + y, center.Y - x}, Point{center.X - y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X + y, center.Y - x}, Point{center.X + y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X + y, center.Y - x}, Point{center.X + x, center.Y + y}, color)

			ppm.DrawLine(Point{center.X - y, center.Y - x}, Point{center.X - x, center.Y + y}, color)
			ppm.DrawLine(Point{center.X - y, center.Y - x}, Point{center.X - y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X - y, center.Y - x}, Point{center.X + y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X - y, center.Y - x}, Point{center.X + x, center.Y + y}, color)

			ppm.DrawLine(Point{center.X - x, center.Y - y}, Point{center.X - x, center.Y + y}, color)
			ppm.DrawLine(Point{center.X - x, center.Y - y}, Point{center.X - y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X - x, center.Y - y}, Point{center.X + y, center.Y + x}, color)
			ppm.DrawLine(Point{center.X - x, center.Y - y}, Point{center.X + x, center.Y + y}, color)

			// Set pixels to fill the circle
			ppm.Set(center.X+x, center.Y-y, color)
			ppm.Set(center.X+y, center.Y-x, color)
			ppm.Set(center.X-y, center.Y-x, color)
			ppm.Set(center.X-x, center.Y-y, color)
			ppm.Set(center.X-x, center.Y+y, color)
			ppm.Set(center.X-y, center.Y+x, color)
			ppm.Set(center.X+y, center.Y+x, color)
			ppm.Set(center.X+x, center.Y+y, color)
		}

		// Update y and decision parameter m
		y++
		if m <= 0 {
			m += 2*y + 1
		}

		// Update x and decision parameter m
		if m > 0 {
			x--
			m -= 2*x + 1
		}
	}

	// Draw the cardinal points to complete the filled circle
	ppm.Set(center.X+radius-1, center.Y, color)
	ppm.Set(center.X, center.Y+radius-1, color)
	ppm.Set(center.X-radius+1, center.Y, color)
	ppm.Set(center.X, center.Y-radius+1, color)
}

// DrawTriangle draws a triangle.
func (ppm *PPM) DrawTriangle(p1, p2, p3 Point, color Pixel) {
	// Draw lines connecting the vertices to form the triangle
	ppm.DrawLine(p1, p2, color)
	ppm.DrawLine(p2, p3, color)
	ppm.DrawLine(p1, p3, color)
}

// DrawFilledTriangle draws a filled triangle.
func (ppm *PPM) DrawFilledTriangle(p1, p2, p3 Point, color Pixel) {
	//A triangle is a polygon so ...
	ppm.DrawFilledPolygon([]Point{p1, p2, p3}, color)
}

// DrawPolygon draws a polygon.
func (ppm *PPM) DrawPolygon(points []Point, color Pixel) {
	// Connect consecutive points with lines
	for point := 0; point < len(points)-1; point++ {
		ppm.DrawLine(points[point], points[point+1], color)
	}
	// Close the loop by connecting the last point to the first point
	ppm.DrawLine(points[0], points[len(points)-1], color)
}

// DrawFilledPolygon draws a filled polygon.
func (ppm *PPM) DrawFilledPolygon(points []Point, color Pixel) {
	// Draw the outline of the polygon
	ppm.DrawPolygon(points, color)

	// Find the bounding box of the polygon
	minX, minY, maxX, maxY := boundingBox(points)

	// Iterate over all pixels inside the bounding box
	for y := minY + 1; y < maxY; y++ {
		for x := minX + 1; x < maxX; x++ {
			// Check if the pixel is inside the polygon
			if isPointInsidePolygon(Point{X: x, Y: y}, points) {
				ppm.Set(x, y, color)
			}
		}
	}
}

// DrawKochSnowflake draws a Koch snowflake.
func (ppm *PPM) DrawKochSnowflake(n int, start Point, size int, color Pixel) {

	// Calculate the height of an equilateral triangle
	height := int(math.Sqrt(3) * float64(size) / 2)

	// Define the three initial points of the equilateral triangle
	p1 := start
	p2 := Point{X: start.X + size, Y: start.Y}
	p3 := Point{X: start.X + size/2, Y: start.Y + height}

	// Apply the Koch snowflake recursive formula to draw the snowflake
	ppm.KochSnowflake(n, p1, p2, color)
	ppm.KochSnowflake(n, p2, p3, color)
	ppm.KochSnowflake(n, p3, p1, color)
}

// KochSnowflake draws a Koch snowflake using the given recursive formula.
func (ppm *PPM) KochSnowflake(n int, p1, p2 Point, color Pixel) {
	if n == 0 {
		// Draw a line segment when recursion depth is 0
		ppm.DrawLine(p1, p2, color)
	} else {
		// Calculate 1/3 and 2/3 points along the line segment
		p1Third := Point{
			X: p1.X + (p2.X-p1.X)/3,
			Y: p1.Y + (p2.Y-p1.Y)/3,
		}
		p2Third := Point{
			X: p1.X + 2*(p2.X-p1.X)/3,
			Y: p1.Y + 2*(p2.Y-p1.Y)/3,
		}

		// Calculate the rotation angle for e**(iπ/3)
		angle := math.Pi / 3
		cosTheta := math.Cos(angle)
		sinTheta := math.Sin(angle)

		// Calculate the rotated point for the middle segment
		p3 := Point{
			X: int(float64(p1Third.X-p2Third.X)*cosTheta-float64(p1Third.Y-p2Third.Y)*sinTheta) + p2Third.X,
			Y: int(float64(p1Third.X-p2Third.X)*sinTheta+float64(p1Third.Y-p2Third.Y)*cosTheta) + p2Third.Y,
		}

		// Recursively draw segments for each third of the line
		ppm.KochSnowflake(n-1, p1, p1Third, color)
		ppm.KochSnowflake(n-1, p1Third, p3, color)
		ppm.KochSnowflake(n-1, p3, p2Third, color)
		ppm.KochSnowflake(n-1, p2Third, p2, color)
	}
}

// DrawSierpinskiTriangle draws a Sierpinski triangle.
func (ppm *PPM) DrawSierpinskiTriangle(n int, start Point, width int, color Pixel) {
	// Calculate the height of an equilateral triangle
	height := int(float64(width) / 2 * math.Sqrt(3))

	// Apply the Sierpinski triangle recursive formula to draw the triangle
	ppm.sierpinskiTriangle(n, start, Point{X: start.X + width, Y: start.Y}, Point{X: start.X + width/2, Y: start.Y + height}, color)
}

// sierpinskiTriangle draws a Sierpinski triangle using the given recursive formula.
func (ppm *PPM) sierpinskiTriangle(n int, p1, p2, p3 Point, color Pixel) {
	if n == 0 {
		// Draw a filled triangle when recursion depth is 0
		ppm.DrawFilledTriangle(p1, p2, p3, color)
	} else {
		// Calculate midpoints of the sides of the triangle
		mid1 := Point{X: (p1.X + p2.X) / 2, Y: (p1.Y + p2.Y) / 2}
		mid2 := Point{X: (p2.X + p3.X) / 2, Y: (p2.Y + p3.Y) / 2}
		mid3 := Point{X: (p3.X + p1.X) / 2, Y: (p3.Y + p1.Y) / 2}

		// Recursively draw Sierpinski triangles for the three sub-triangles
		ppm.sierpinskiTriangle(n-1, p3, mid2, mid3, color)
		ppm.sierpinskiTriangle(n-1, mid2, mid1, p2, color)
		ppm.sierpinskiTriangle(n-1, mid1, p1, mid3, color)
	}
}

// DrawPerlinNoise draws perlin noise.
// this function Draw a perlin noise of all the image.

func (ppm *PPM) DrawPerlinNoise(color1, color2 Pixel) {
	//Travel through each pixel
	for y := 0; y < len(ppm.data); y++ { // Go through each line
		for x := 0; x < len(ppm.data[0]); x++ { // Go through each column
			// Normalize coordinates to fit perlin noise function
			normalizedX := float64(x) / float64(len(ppm.data))
			normalizedY := float64(y) / float64(len(ppm.data[0]))

			// Map the noise value to a color between color1 and color2
			newColor := interpolateColor(color1, color2, PerlinNoise(5*normalizedX, 5*normalizedY)) // We zoom into the noise by multiplying the result by 5

			// Set the pixel color
			ppm.Set(x, y, newColor)
		}
	}
}
func PerlinNoise(x, y float64) float64 {
	// Generate a random gradient vector for each grid point
	gradX := rand.Float64()*2 - 1
	gradY := rand.Float64()*2 - 1

	// Calculate distance vectors from the grid points to the input point
	dx := x - math.Floor(x)
	dy := y - math.Floor(y)

	// Calculate dot products between the gradient vectors and the distance vectors
	lBottom := gradX*dx + gradY*(dy-1)
	rBottom := gradX*(dx-1) + gradY*(dy-1)
	lTop := gradX*dx + gradY*dy
	rTop := gradX*(dx-1) + gradY*dy

	// Interpolate along the x-axis
	inter_xTop := lTop + smoothing(dx)*(rTop-lTop)
	inter_xBottom := lBottom + smoothing(dx)*(rBottom-lBottom)

	return inter_xTop + smoothing(dy)*(inter_xBottom-inter_xTop)
}

// KNearestNeighbors resizes the PPM image using the k-nearest neighbors algorithm.
func (ppm *PPM) KNearestNeighbors(newWidth, newHeight, k int) {
	originalWidth := len(ppm.data[0])
	originalHeight := len(ppm.data)
	ppm.height = newHeight
	ppm.width = newWidth

	// Create a new image with the desired dimensions (new matrice)
	newData := make([][]Pixel, newHeight)
	for i := range newData {
		newData[i] = make([]Pixel, newWidth)
	}

	// Iterate over each pixel in the new image
	for newY := 0; newY < newHeight; newY++ {
		for newX := 0; newX < newWidth; newX++ {
			// Calculate the corresponding coordinates in the original image
			originalX := int(float64(originalWidth) * float64(newX) / float64(newWidth))
			originalY := int(float64(originalHeight) * float64(newY) / float64(newHeight))

			// Find k-nearest neighbors in the original image
			neighbors := ppm.findKNearestNeighbors(originalX, originalY, k)

			// Calculate the average color of the neighbors
			averageColor := ppm.calculateAverageColor(neighbors)

			// Set the color of the pixel in the new image
			newData[newY][newX] = averageColor
		}
	}

	// Update the PPM with the resized image
	ppm.data = newData
}

// findKNearestNeighbors finds the k-nearest neighbors of a given point in the original image.
func (ppm *PPM) findKNearestNeighbors(x, y, k int) []Point {
	points := make([]Point, 0)

	// Iterate over each pixel in the original image
	for i := range ppm.data {
		for j := range ppm.data[i] {
			points = append(points, Point{X: j, Y: i})
		}
	}

	// Sort points based on their distance to the target point
	sort.Slice(points, func(i, j int) bool {
		distI := Distance(Point{X: x, Y: y}, points[i])
		distJ := Distance(Point{X: x, Y: y}, points[j])
		return distI < distJ
	})

	// Return the first k points as the k-nearest neighbors
	return points[:k]
}
