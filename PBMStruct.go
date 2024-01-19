package Netpbm

/*
 * Titouan Schotté
 * PBM -> Struct and Methods
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// STRUCT PBM
type PBM struct {
	data          [][]bool
	width, height int
	magicNumber   string
}

// Additionnal Method to see in terminal datas correctly formated
func (pbm *PBM) Show() {
	// Afficher le contenu de pbm.Data
	for _, row := range pbm.data {
		for _, pixel := range row {
			if pixel {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
}

// Size returns the width and height of the image.
func (pbm *PBM) Size() (int, int) {
	return pbm.width, pbm.height
}

// At returns the value of the pixel at (x, y).
func (pbm *PBM) At(x, y int) bool {
	// We check that the position exists in the matrix
	if x >= 0 && y >= 0 &&
		x < pbm.width && y < pbm.height {
		return pbm.data[y][x]
	}
	return false
}

// Set sets the value of the pixel at (x, y).
func (pbm *PBM) Set(x, y int, value bool) {
	// We check that the position exists in the matrix
	if x >= 0 && y >= 0 &&
		x < pbm.width && y < pbm.height {
		pbm.data[y][x] = value
	}
}

// Save saves the PBM image to a file and returns an error if there was a problem.
func (pbm *PBM) Save(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	// Write the magic number
	_, err = fmt.Fprintln(writer, pbm.magicNumber)
	if err != nil {
		return err
	}

	// Write the size
	_, err = fmt.Fprintf(writer, strconv.Itoa(pbm.width)+" "+strconv.Itoa(pbm.height)+"\n")
	if err != nil {
		return err
	}

	// Write image data (P1)
	if pbm.magicNumber == "P1" {
		for _, line := range pbm.data {
			for _, pixel := range line {
				if pixel {
					_, err = fmt.Fprint(writer, "1 ")
				} else {
					_, err = fmt.Fprint(writer, "0 ")
				}
				if err != nil {
					return err
				}
			}
			_, err = fmt.Fprintln(writer) // New line after each image line
			if err != nil {
				return err
			}
		}
	} else if pbm.magicNumber == "P4" { // Write image data (P4)
		for _, line := range pbm.data {
			buffOctet := 0
			P4BinaryIn := []int{}
			octetsNumber := 0

			// By respecting the padding method, we already read whole bytes
			for i := 0; i < len(line); i++ {
				if buffOctet == 8 {
					val, _ := BinaryToWindows1252(P4BinaryIn)
					_, err = fmt.Fprint(writer, val)
					P4BinaryIn = []int{}
					buffOctet = 0
					octetsNumber++
				}
				if line[i] {
					P4BinaryIn = append(P4BinaryIn, 1)
				} else {
					P4BinaryIn = append(P4BinaryIn, 0)
				}
				buffOctet++
			}

			// Padding
			if buffOctet != 0 {
				for i := 0; octetsNumber*8+len(P4BinaryIn)+i <= (octetsNumber+1)*8; i++ {
					P4BinaryIn = append(P4BinaryIn, 0)
				}
				val, _ := BinaryToWindows1252(P4BinaryIn)
				_, err = fmt.Fprint(writer, val)

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
func (pbm *PBM) Invert() {
	for i := range pbm.data {
		for j := range pbm.data[i] {
			pbm.data[i][j] = !pbm.data[i][j] // Allows us to return the inverse of the boolean corresponding to the position concerned.
		}
	}
}

// Flip flips the PBM image horizontally.
func (pgm *PBM) Flip() {
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
func (pbm *PBM) Flop() {
	for i := 0; i < len(pbm.data)/2; i++ {
		j := len(pbm.data) - 1 - i
		pbm.data[i], pbm.data[j] = pbm.data[j], pbm.data[i]
	}
}

// SetMagicNumber sets the magic number of the PBM image.
func (pgm *PBM) SetMagicNumber(magicNumber string) {
	pgm.magicNumber = magicNumber
}
