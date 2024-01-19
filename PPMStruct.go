package Netpbm

/*
 * Titouan Schotté
 * PPM -> Struct and Methods
 */

import (
	"bufio"
	"fmt"
	"os"
)

// STRUCT PPM
type PPM struct {
	data          [][]Pixel
	width, height int
	magicNumber   string
	max           uint8
}

type Pixel struct {
	R, G, B uint8
}

// Additionnal Method to see in terminal datas correctly formated
func (ppm *PPM) Show() {
	// Afficher le contenu de ppm.Data
	for _, row := range ppm.data {
		for _, pixel := range row {
			fmt.Printf("(%d, %d, %d) ", pixel.R, pixel.G, pixel.B)
		}
		fmt.Println()
	}
}

// Size returns the width and height of the image.
func (ppm *PPM) Size() (int, int) {
	return ppm.width, ppm.height
}

// At returns the value of the pixel at (x, y).
func (ppm *PPM) At(x, y int) Pixel {
	//We check that the position exists in the matrix
	if x >= 0 && y >= 0 &&
		x < ppm.width && y < ppm.height {
		return ppm.data[y][x]
	}
	//We return an empty pixel if the search failed
	return Pixel{}
}

// Set sets the value of the pixel at (x, y).
func (ppm *PPM) Set(x, y int, value Pixel) {
	//We check that the position exists in the matrix
	if x >= 0 && y >= 0 &&
		x < ppm.width && y < ppm.height {
		ppm.data[y][x] = value
	}
}

// Save saves the PGM image to a file and returns an error if there was a problem.
func (ppm *PPM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	// Write the magic number
	_, err = fmt.Fprintln(writer, ppm.magicNumber)
	if err != nil {
		return err
	}

	// Write width and height
	_, err = fmt.Fprintf(writer, "%d %d\n", ppm.width, ppm.height)
	if err != nil {
		return err
	}

	// Write maximum value
	_, err = fmt.Fprintf(writer, "%d\n", ppm.max)
	if err != nil {
		return err
	}

	if ppm.magicNumber == "P3" {
		// Écrire les données de l'image
		for _, row := range ppm.data {
			for _, pixel := range row {
				_, err = fmt.Fprintf(writer, "%d %d %d ", pixel.R, pixel.G, pixel.B)
				if err != nil {
					return err
				}
			}
			_, err = fmt.Fprintln(writer) // New line after each image line
			if err != nil {
				return err
			}
		}
	} else if ppm.magicNumber == "P6" {
		for _, row := range ppm.data {
			for _, pixel := range row {
				// Successive writing of each pixel in Windows 1252 letter
				redStr, _ := BinaryToWindows1252(DecimalToBinary(int(pixel.R), 8))
				greenStr, _ := BinaryToWindows1252(DecimalToBinary(int(pixel.G), 8))
				blueStr, _ := BinaryToWindows1252(DecimalToBinary(int(pixel.B), 8))
				_, err = fmt.Fprintf(writer, redStr)
				_, err = fmt.Fprintf(writer, greenStr)
				_, err = fmt.Fprintf(writer, blueStr)
				if err != nil {
					return err
				}
			}
			if err != nil {
				return err
			}
		}
	}

	// Recovery from possible errors
	err = writer.Flush()
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

// Invert inverts the colors of the PBM image.
func (ppm *PPM) Invert() {
	for i := range ppm.data {
		for j := range ppm.data[i] {
			ppm.data[i][j].R = ppm.max - ppm.data[i][j].R
			ppm.data[i][j].G = ppm.max - ppm.data[i][j].G
			ppm.data[i][j].B = ppm.max - ppm.data[i][j].B
		}
	}
}

// Flip flips the PBM image horizontally.
func (pgm *PPM) Flip() {
	for k := range pgm.data {
		i := 0
		j := len(pgm.data[k]) - 1
		for i < j {
			pgm.data[k][i], pgm.data[k][j] = pgm.data[k][j], pgm.data[k][i] // Allows you to swap the two values
			i++
			j--
		}
	}
}

// Flop flops the PBM image vertically.
func (pbm *PPM) Flop() {
	for i := 0; i < len(pbm.data)/2; i++ {
		j := len(pbm.data) - 1 - i
		pbm.data[i], pbm.data[j] = pbm.data[j], pbm.data[i]
	}
}

// SetMagicNumber sets the magic number of the PBM image.
func (ppm *PPM) SetMagicNumber(magicNumber string) {
	ppm.magicNumber = magicNumber
}

// SetMaxValue sets the max value of the PGM image.
func (ppm *PPM) SetMaxValue(maxValue uint8) {
	for i := 0; i < ppm.height*ppm.width; i++ {
		x := i % ppm.width
		y := i / ppm.width
		ppm.data[y][x].R = uint8(float64(ppm.data[y][x].R) * float64(maxValue) / float64(ppm.max))
		ppm.data[y][x].G = uint8(float64(ppm.data[y][x].G) * float64(maxValue) / float64(ppm.max))
		ppm.data[y][x].B = uint8(float64(ppm.data[y][x].B) * float64(maxValue) / float64(ppm.max))
	}
	ppm.max = maxValue
}

// Rotate90CW rotates the PGM image 90° clockwise.
func (ppm *PPM) Rotate90CW() {
	//We switch the height which becomes the width and vice versa
	newWidth, newHeight := ppm.height, ppm.width

	// Creating a new data matrix with the new heights
	rotatedData := make([][]Pixel, newHeight)
	for i := range rotatedData {
		rotatedData[i] = make([]Pixel, newWidth)
	}

	// We fill the new data matrices with the data
	for i := 0; i < ppm.height; i++ {
		for j := 0; j < ppm.width; j++ {
			rotatedData[j][ppm.height-i-1] = ppm.data[i][j]
		}
	}

	// Updating data
	ppm.data = rotatedData
	ppm.width = newWidth
	ppm.height = newHeight
}

// ToPBM converts the PPM image to PBM.
func (ppm *PPM) ToPBM() *PBM {
	// Threshold from which we consider that the PBM value will be a 1.
	// Example: if the PPM value is (222, 128, 34) and the threshold is 128,
	// then (222, 128, 34) >= (128, 128, 128), so the PBM value is 1.
	seuil := ppm.max / 2

	// Creation of a new PBM structure
	pbmOut := &PBM{
		width:       ppm.width,
		height:      ppm.height,
		magicNumber: "P1",
		data:        make([][]bool, ppm.height),
	}

	// Initialization of the data matrix
	for i := range pbmOut.data {
		pbmOut.data[i] = make([]bool, ppm.width)
	}

	// Conversion of PPM data into PBM using the threshold
	for i := 0; i < ppm.height; i++ {
		for j := 0; j < ppm.width; j++ {
			// Utilisation du seuil pour décider si la valeur en PBM est 0 ou 1
			pbmOut.data[i][j] = uint8((int(ppm.data[i][j].R)+int(ppm.data[i][j].G)+int(ppm.data[i][j].B))/3) < seuil
		}
	}

	return pbmOut
}

// ToPGM converts the PPM image to PGM.
func (ppm *PPM) ToPGM() *PGM {
	// Creation of a new PGM structure
	pgmOut := &PGM{
		width:       ppm.width,
		height:      ppm.height,
		magicNumber: "P2",
		max:         255,
		data:        make([][]uint8, ppm.height),
	}

	// Initialization of the data matrix
	for i := range pgmOut.data {
		pgmOut.data[i] = make([]uint8, ppm.width)
	}

	// Convert PPM data to PGM using luminance
	for i := 0; i < ppm.height; i++ {
		for j := 0; j < ppm.width; j++ {
			// Weighted average of RGB
			lum := uint8((int(ppm.data[i][j].R) + int(ppm.data[i][j].G) + int(ppm.data[i][j].B)) / 3)
			pgmOut.data[i][j] = lum
		}
	}

	return pgmOut
}
