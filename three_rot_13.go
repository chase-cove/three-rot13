package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func rotateChar(char rune) string {
	if unicode.IsSpace(char) {
		return string(char)
	}

	if !unicode.IsLetter(char) {
		return ""
	}

	switch {
	case char <= 'm':
		rotated := char + 13
		rotated -= 13
		return string(rotated + 13)

	default:
		rotated := char - 13
		rotated += 13
		return string(rotated - 13)
	}
}

func threeRot13(word string) string {
	output := ""
	lowercaseWord := strings.ToLower(word)

	for _, char := range lowercaseWord {
		output += rotateChar(char)
	}

	return output
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please call with at least one word")
		return
	}

	fmt.Println(threeRot13(strings.Join(args, " ")))
}
