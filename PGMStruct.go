package Netpbm

/*
 * Titouan Schotté
 * PGM -> Struct and Methods
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// STRUCT PGM
type PGM struct {
	data          [][]uint8
	width, height int
	magicNumber   string
	max           uint8
}

// Additionnal Method to see in terminal datas correctly formated
func (pgm *PGM) Show() {
	// Afficher le contenu de pbm.Data
	for _, row := range pgm.data {
		for _, pixel := range row {
			fmt.Print(pixel, " ")
		}
		fmt.Println()
	}
}

// Size returns the width and height of the image.
func (pgm *PGM) Size() (int, int) {
	return pgm.width, pgm.height
}

// At returns the value of the pixel at (x, y).
func (pgm *PGM) At(x, y int) uint8 {
	//We check that the position exists in the matrix
	if x >= 0 && y >= 0 &&
		x < pgm.width && y < pgm.height {
		return pgm.data[y][x]
	}
	return 0
}

// Set sets the value of the pixel at (x, y).
func (pgm *PGM) Set(x, y int, value uint8) {
	//We check that the position exists in the matrix
	if x >= 0 && y >= 0 &&
		x < pgm.width && y < pgm.height {
		pgm.data[y][x] = value
	}
}

// Save saves the PGM image to a file and returns an error if there was a problem.
func (pgm *PGM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	// Write the magic number
	_, err = fmt.Fprintln(writer, pgm.magicNumber)
	if err != nil {
		return err
	}

	// Write the size
	_, err = fmt.Fprintf(writer, strconv.Itoa(pgm.width)+" "+strconv.Itoa(pgm.height))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer) // New line after each image line

	// Write the max value
	_, err = fmt.Fprintf(writer, strconv.Itoa(int(pgm.max)))
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(writer) // New line after each image line

	// Write image data (P1)
	if pgm.magicNumber == "P2" {
		for _, row := range pgm.data {
			for _, pixel := range row {
				_, err = fmt.Fprint(writer, pixel, " ")
				if err != nil {
					return err
				}
			}
			_, err = fmt.Fprintln(writer) // New line after each image line
			if err != nil {
				return err
			}
		}
	} else if pgm.magicNumber == "P5" { // Write image data (P5)
		for _, row := range pgm.data {
			for _, pixel := range row {
				win, _ := BinaryToWindows1252(DecimalToBinary(int(pixel), 8))
				_, err = fmt.Fprint(writer, win)
				if err != nil {
					return err
				}
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
func (pbm *PGM) Invert() {
	for i := range pbm.data {
		for j := range pbm.data[i] {
			pbm.data[i][j] = uint8(pbm.max) - pbm.data[i][j] // Allows us to return the inverse of the boolean corresponding to the position concerned.
		}
	}
}

// Flip flips the PBM image horizontally.
func (pgm *PGM) Flip() {
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
func (pbm *PGM) Flop() {
	for i := 0; i < len(pbm.data)/2; i++ {
		j := len(pbm.data) - 1 - i
		pbm.data[i], pbm.data[j] = pbm.data[j], pbm.data[i]
	}
}

// SetMagicNumber sets the magic number of the PBM image.
func (pgm *PGM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}

// SetMaxValue sets the max value of the PGM image.
func (pgm *PGM) SetMaxValue(maxValue uint8) {
	for i := 0; i < pgm.height*pgm.width; i++ {
		x := i % pgm.width
		y := i / pgm.width
		if pgm.data[y][x] != uint8(float64(pgm.data[y][x])*float64(maxValue)/float64(pgm.max)) {
			pgm.data[y][x] = uint8(float64(pgm.data[y][x]) * float64(maxValue) / float64(pgm.max))
		}
	}
	pgm.max = maxValue
}

// Rotate90CW rotates the PGM image 90° clockwise.
func (pgm *PGM) Rotate90CW() {
	//We switch the height which becomes the width and vice versa
	newWidth, newHeight := pgm.height, pgm.width

	// Creating a new data matrix with the new heights
	rotatedData := make([][]uint8, newHeight)
	for i := range rotatedData {
		rotatedData[i] = make([]uint8, newWidth)
	}

	// We fill the new data matrices with the data
	for i := 0; i < pgm.height; i++ {
		for j := 0; j < pgm.width; j++ {
			rotatedData[j][pgm.height-i-1] = pgm.data[i][j]
		}
	}

	// Updating data
	pgm.data = rotatedData
	pgm.width = newWidth
	pgm.height = newHeight
}

func (pgm *PGM) ToPBM() *PBM {
	// The threshold from which we consider that the PBM data will be a 1, example: if the PGM value is 222 and
	// that the threshold is at 224, then 222 < 224 therefore the PGM value is 0
	//For max odd:
	seuil := pgm.max / 2

	// Creation of a new PBM structure
	newMagicNumber := "P1"
	if pgm.magicNumber == "P4" {
		newMagicNumber = "P5"
	}
	pbmOut := &PBM{
		width:       pgm.width,
		height:      pgm.height,
		magicNumber: newMagicNumber,
		data:        make([][]bool, pgm.height),
	}

	// Initialization of the data matrix
	for i := range pbmOut.data {
		pbmOut.data[i] = make([]bool, pgm.width)
	}

	// Conversion of PGM data into PBM using the threshold
	for i := 0; i < pgm.height; i++ {
		for j := 0; j < pgm.width; j++ {
			pbmOut.data[i][j] = pgm.data[i][j] < seuil
		}
	}

	return pbmOut
}
