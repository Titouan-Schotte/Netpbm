package Netpbm

/*
 * Titouan Schotté
 * Utils
 */

import (
	"fmt"
	"math"
)

func DecimalToBinary(decimal int, fixedLength int) []int {
	//Allow to convert a decimal number to binary (useful for conversions to Windows-1252)
	binaryArray := []int{}

	for decimal > 0 {
		remainder := decimal % 2
		binaryArray = append([]int{remainder}, binaryArray...)
		decimal = decimal / 2
	}

	// PADDING : Pad with leading zeros to reach fixed length
	for len(binaryArray) < fixedLength {
		binaryArray = append([]int{0}, binaryArray...)
	}

	return binaryArray
}

func BinaryToDecimal(binaryArray []int) int {
	// Allow to convert binary to a decimal number (useful for conversions to Windows-1252)
	decimal := 0
	power := len(binaryArray) - 1

	for _, bit := range binaryArray {
		decimal += bit * int(math.Pow(2, float64(power)))
		power--
	}

	return decimal
}

func BinaryToWindows1252(binaryArray []int) (string, error) {
	// Allow to convert binary to windows 1252 string
	// Check that the length of the binary array is a multiple of 8
	if len(binaryArray)%8 != 0 {
		return "", fmt.Errorf("la longueur du tableau binaire doit être un multiple de 8")
	}

	// Convert binary sequence to a byte array
	var byteArray []byte
	for i := 0; i < len(binaryArray); i += 8 {
		byteValue := BinaryToDecimal(binaryArray[i : i+8])
		byteArray = append(byteArray, byte(byteValue))
	}

	return string(byteArray), nil
}
