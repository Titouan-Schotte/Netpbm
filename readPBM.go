package Netpbm

/*
 * Titouan Schotté
 * PBM -> Fonction Read PBM
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadPBM reads a PBM image from a file and returns a struct that represents the image.
func ReadPBM(filename string) (*PBM, error) {

	var pbmIn = &PBM{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR : Can't open file:", err)
		return nil, err
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "#") { //We ignore comments
			continue
		}
		if i == 0 { //We are currently reading the magic number
			pbmIn.magicNumber = line
			if pbmIn.magicNumber != "P1" && pbmIn.magicNumber != "P4" {
				return nil, fmt.Errorf("ERROR : file format is not PBM format")
			}
		}
		if i == 1 { //We are currently reading height & width
			size := strings.Fields(scanner.Text())
			if len(size) != 2 {
				return nil, fmt.Errorf("ERROR : File lenght not valid")
			}
			pbmIn.width, err = strconv.Atoi(size[0])
			if err != nil {
				return nil, fmt.Errorf("ERROR : Width not valid")
			}
			pbmIn.height, err = strconv.Atoi(size[1])
			if err != nil {
				return nil, fmt.Errorf("ERROR : Height not valid")
			}

			if pbmIn.magicNumber == "P1" {
				pbmIn.data = make([][]bool, pbmIn.height)
				for j := range pbmIn.data {
					pbmIn.data[j] = make([]bool, pbmIn.width)
				}
			}

		}
		if i > 1 {
			// Read the body
			if pbmIn.magicNumber == "P1" {
				// Initialize the data matrix

				lineData := strings.Fields(line)
				if len(lineData) != pbmIn.width {
					return nil, fmt.Errorf("ERROR : width of the body line not valid") // On créé une erreur
				}
				for j, pixel := range lineData {
					val, err := strconv.Atoi(pixel)
					if err != nil {
						return nil, fmt.Errorf("ERROR : pixel value not valid") // On créé une erreur
					}
					pbmIn.data[i-2][j] = val == 1
				}
			} else if pbmIn.magicNumber == "P4" && pbmIn.data == nil {
				// Initialize the data matrix
				pbmIn.data = make([][]bool, pbmIn.height)
				for j := range pbmIn.data {
					pbmIn.data[j] = make([]bool, pbmIn.width)
				}
				p4Buff := 0
				p4LineIn := 0
				for _, asciiCode := range []byte(line) {
					binaryCode := DecimalToBinary(int(asciiCode), 8)
					for b := 0; b < len(binaryCode); b++ {
						if p4Buff >= pbmIn.width {
							p4Buff = 0
							p4LineIn++
							continue
						}
						if p4LineIn >= pbmIn.height {
							break
						}

						pbmIn.data[p4LineIn][p4Buff] = binaryCode[b] == 1
						p4Buff++
					}
				}
			} else {
				return nil, fmt.Errorf("ERROR : magic number unknown") // On créé une erreur
			}
		}
		i++
	}
	// Error checking
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	file.Close()
	return pbmIn, nil
}
