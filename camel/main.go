package main

import (
	"fmt"
)

func main() {
	var camelString string
	fmt.Scanf("%s\n", &camelString)
	fmt.Println("The number of words in", camelString, "are", getCamelLength(camelString))	
}

func getCamelLength(camel string) int  {
	wordCount := 1
	for _, char := range camel {
		if char >= 'A' && char <= 'Z' {
			wordCount++			
		}
	}
	return wordCount	
}

