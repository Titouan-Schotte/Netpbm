package Netpbm

/*
 * Titouan Schotté
 * PGM -> Fonction Read PGM
 */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadPGM reads a PGM image from a file and returns a struct that represents the image.
func ReadPGM(filename string) (*PGM, error) {
	var pgmIn = &PGM{}

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
			pgmIn.magicNumber = line
			if pgmIn.magicNumber != "P2" && pgmIn.magicNumber != "P5" {
				return nil, fmt.Errorf("Format non pris en charge")
			}
		}
		if i == 1 { //We are currently reading height & width
			size := strings.Fields(scanner.Text())
			if len(size) != 2 {
				return nil, fmt.Errorf("Taille du format invalide") // On créé une erreur
			}
			pgmIn.width, err = strconv.Atoi(size[0])
			if err != nil {
				return nil, fmt.Errorf("largeur invalide") // On créé une erreur
			}
			pgmIn.height, err = strconv.Atoi(size[1])
			if err != nil {
				return nil, fmt.Errorf("hauteur invalide") // On créé une erreur
			}

			// Initialize the data matrix
			pgmIn.data = make([][]uint8, pgmIn.height)
			for j := range pgmIn.data {
				pgmIn.data[j] = make([]uint8, pgmIn.width)
			}
		}
		if i == 2 {
			// Read max value allowed
			maxValue, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("valeur maximale invalide") // On créé une erreur
			}
			pgmIn.max = uint8(maxValue)
		}
		if i > 2 {
			// Read the body (P2)
			if pgmIn.magicNumber == "P2" {
				lineData := strings.Fields(line)
				if len(lineData) != pgmIn.width {
					return nil, fmt.Errorf("Largeur de la ligne du body invalide") // On créé une erreur
				}
				for j, pixel := range lineData {
					val, err := strconv.Atoi(pixel)
					if err != nil {
						return nil, fmt.Errorf("Valeur de pixel invalide") // On créé une erreur
					}
					pgmIn.data[i-3][j] = uint8(val)
				}
			} else if pgmIn.magicNumber == "P5" { // Read the body (P5)
				x, y := 0, 0

				for _, asciiCode := range line {

					if x == pgmIn.width {
						x = 0
						y++
					}
					pgmIn.data[y][x] = uint8(asciiCode)
					x++
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
	return pgmIn, nil
}
