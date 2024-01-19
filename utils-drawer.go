package Netpbm

/*
 * Titouan Schotté
 * UTILS For PPM Drawer
 */

import "math"

func Sgn(x int) int {
	//SGN returns the sign of a value: either -1, 1 or 0 if the number is negative, positive or zero.
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	return 0
}

func Abs(x int) int {
	//ABS returns the absolute value of a number.
	if x < 0 {
		return -x
	}
	return x
}

// **** DRAW FILLED POLYGON ****

// Function to check if a point is inside a polygon using the ray-casting algorithm.
func boundingBox(points []Point) (minX, minY, maxX, maxY int) {
	// Initialize the minimum and maximum coordinates with the first point
	minX, minY = points[0].X, points[0].Y
	maxX, maxY = points[0].X, points[0].Y

	// Iterate through each point and update the bounding box coordinates
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// Return the calculated bounding box coordinates
	return minX, minY, maxX, maxY
}

// Function to check if a point is inside a polygon using the ray-casting algorithm.
func isPointInsidePolygon(point Point, polygon []Point) bool {
	// Use the ray-casting algorithm to determine if the point is inside the polygon
	intersections := 0

	// Iterate through each edge of the polygon
	for i := 0; i < len(polygon); i++ {
		p1, p2 := polygon[i], polygon[(i+1)%len(polygon)]

		// Check if the point is on an edge of the polygon
		if (p1.Y == point.Y && p2.Y == point.Y) &&
			((p1.X <= point.X && point.X <= p2.X) || (p2.X <= point.X && point.X <= p1.X)) {
			return true
		}

		// Check if the horizontal ray from the point intersects the edge
		if (p1.Y > point.Y) != (p2.Y > point.Y) &&
			point.X < int(float64(p1.X+(point.Y-p1.Y)*(p2.X-p1.X)/(p2.Y-p1.Y))) {
			intersections++
		}
	}

	// If the number of intersections is odd, the point is inside the polygon
	return intersections%2 == 1
}

// ** PERLIN NOISE **

// smoothing applies the smooth step function to a given input value t.
func smoothing(t float64) float64 {
	// The formula t * t * (3 - 2*t) is a smoothing function often used in perlin noise generation and other graphics applications
	return t * t * (3 - 2*t)
}

// EXPLAIN : Linear interpolation is a mathematical method that allows you to estimate an intermediate value between two known values.
// interpolateColor interpolates between two colors based on a given parameter t.
// It returns a new color that is a linear blend between color1 and color2.
// The parameter t should be in the range [0, 1], where 0 corresponds to color1,
// 1 corresponds to color2, and values in between yield interpolated colors.
func interpolateColor(color1, color2 Pixel, t float64) Pixel {
	// Interpolate the red component of the colors
	r := uint8(float64(color1.R)*(1-t) + float64(color2.R)*t)
	// Interpolate the green component of the colors
	g := uint8(float64(color1.G)*(1-t) + float64(color2.G)*t)
	// Interpolate the blue component of the colors
	b := uint8(float64(color1.B)*(1-t) + float64(color2.B)*t)
	// Return the new interpolated color
	return Pixel{R: r, G: g, B: b}
}

// ** KNearestNeighbors **
// Distance calculates the Euclidean distance between two points.
func Distance(p1, p2 Point) float64 {
	// Calculate the differences in x and y coordinates
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y

	return math.Sqrt(float64(dx*dx + dy*dy))
}

// calculateAverageColor calculates the average color of a list of points in the original image.
// It takes a list of points (the neighbors) and computes the average color by summing up the color
// components of all neighbors and then dividing by the number of neighbors.
// It return a PPM Pixel representing the average color
func (ppm *PPM) calculateAverageColor(neighbors []Point) Pixel {
	// Variables to accumulate the total color components
	var totalR, totalG, totalB uint32

	// Sum up the color components of all neighbors
	for _, neighbor := range neighbors {
		totalR += uint32(ppm.data[neighbor.Y][neighbor.X].R)
		totalG += uint32(ppm.data[neighbor.Y][neighbor.X].G)
		totalB += uint32(ppm.data[neighbor.Y][neighbor.X].B)
	}

	// Calculate the average color components
	averageR := uint8(totalR / uint32(len(neighbors)))
	averageG := uint8(totalG / uint32(len(neighbors)))
	averageB := uint8(totalB / uint32(len(neighbors)))

	// Return a Pixel representing the average color
	return Pixel{R: averageR, G: averageG, B: averageB}
}
