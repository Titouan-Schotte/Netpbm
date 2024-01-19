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
			pbmIn.magicNumber = line
		}
		if i == 1 { //We are currently reading height & width
			size := strings.Fields(scanner.Text())
			if len(size) != 2 {
				return nil, fmt.Errorf("Taille du format invalide") //On créé une erreur
			}
			pbmIn.width, err = strconv.Atoi(size[0])
			if err != nil {
				return nil, fmt.Errorf("longueur invalide") //On créé une erreur
			}
			pbmIn.height, err = strconv.Atoi(size[1])
			if err != nil {
				return nil, fmt.Errorf("hauteur invalide") //On créé une erreur
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
					return nil, fmt.Errorf("Largeur de la ligne du body invalide") // On créé une erreur
				}
				for j, pixel := range lineData {
					val, err := strconv.Atoi(pixel)
					if err != nil {
						return nil, fmt.Errorf("Valeur de pixel invalide") // On créé une erreur
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
	return pbmIn, nil
}
