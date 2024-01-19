package Netpbm

/*
 * Titouan Schotté
 * PGM -> Fonction Read PPM
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadPPM reads a PPM image from a file and returns a struct that represents the image.
func ReadPPM(filename string) (*PPM, error) {
	var ppmIn = &PPM{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur à l'ouverture du fichier:", err)
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
			ppmIn.magicNumber = line
			if ppmIn.magicNumber != "P3" && ppmIn.magicNumber != "P6" {
				return nil, fmt.Errorf("Format non pris en charge")
			}
		}
		if i == 1 { //We are currently reading height & width
			size := strings.Fields(scanner.Text())
			if len(size) != 2 {
				return nil, fmt.Errorf("Taille du format invalide")
			}
			ppmIn.width, err = strconv.Atoi(size[0])
			if err != nil {
				return nil, fmt.Errorf("largeur invalide")
			}
			ppmIn.height, err = strconv.Atoi(size[1])
			if err != nil {
				return nil, fmt.Errorf("hauteur invalide")
			}

			// Initialize the data matrix
			ppmIn.data = make([][]Pixel, ppmIn.height)
			for j := range ppmIn.data {
				ppmIn.data[j] = make([]Pixel, ppmIn.width)
			}
		}
		if i == 2 {
			// Read max value allowed
			maxValue, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("valeur maximale invalide")
			}
			ppmIn.max = uint8(maxValue)
		}
		if i > 2 {
			// Read the body (P3)
			if ppmIn.magicNumber == "P3" {
				lineData := strings.Fields(line)
				pixelCount := len(lineData)
				if pixelCount%3 != 0 || pixelCount/3 != ppmIn.width {
					return nil, fmt.Errorf("Largeur de la ligne du body invalide")
				}
				for j := 0; j < ppmIn.width; j++ {
					for k := 0; k < 3; k++ {
						val, err := strconv.Atoi(lineData[j*3+k])
						if err != nil {
							return nil, fmt.Errorf("Valeur de pixel invalide")
						}
						switch k {
						case 0:
							ppmIn.data[i-3][j].R = uint8(val)
						case 1:
							ppmIn.data[i-3][j].G = uint8(val)
						case 2:
							ppmIn.data[i-3][j].B = uint8(val)
						}
					}
				}
			} else if ppmIn.magicNumber == "P6" { // Read the body (P6)
				x, y := 0, 0
				for k := range line {

					if x == ppmIn.width {
						x = 0
						y++
					}
					//Reading in packs of 3 because there is RGB to fill
					if k%3 == 0 && x < ppmIn.width && y < ppmIn.height {
						ppmIn.data[y][x].R = line[k]
						if k+1 < len(line) {
							ppmIn.data[y][x].G = line[k+1]
						}
						if k+2 < len(line) {
							ppmIn.data[y][x].B = line[k+2]
						}
						x++
					}
				}
			} else {
				return nil, fmt.Errorf("Magic Number invalide") // On créé une erreur
			}

		}
		i++
	}

	// Error checking
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	file.Close()
	return ppmIn, nil
}
