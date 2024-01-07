package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var length string
	var cipher string
	var key string
	fmt.Scanf("%s\n", &length)
	fmt.Scanf("%s\n", &cipher)
	fmt.Scanf("%s\n", &key)
	keyInt, err := strconv.Atoi(key)
	if err != nil {
		 panic(err)
	}
	fmt.Println("Cipher of", cipher, "is", cipherGet(cipher, keyInt))
}

func cipherGet(cipher string, key int) string {

	var sb strings.Builder
	for _, char := range cipher {
		newChar := int(char)
		if char >= 'a' && char <= 'z' {
			groundChar := char - 'a'
			newChar = (((int(groundChar)) + key) % ('z' - 'a' + 1)) + 'a'
			fmt.Println("Last", string(char), "New", string(newChar))	
		} else if char >= 'A' && char <= 'Z' {
			groundChar := char - 'A'
			newChar = (((int(groundChar)) + key) % ('Z' - 'A' + 1)) + 'A'
			fmt.Println("Last", string(char), "New", string(newChar))	
		}
		sb.WriteByte(byte(newChar))
	}

	return sb.String()
	
}
