package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func generateQuad(quadType string, width, height int) string {
	if width <= 0 || height <= 0 {
		return ""
	}

	var result strings.Builder

	switch quadType {
	case "A":
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if (i == 0 || i == height-1) && (j == 0 || j == width-1) {
					result.WriteString("o")
				} else if i == 0 || i == height-1 {
					result.WriteString("-")
				} else if j == 0 || j == width-1 {
					result.WriteString("|")
				} else {
					result.WriteString(" ")
				}
			}
			result.WriteString("\n")
		}

	case "B":
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if (i == 0 && j == 0) || (i == height-1 && j == width-1) {
					result.WriteString("/")
				} else if (i == height-1 && j == 0) || (i == 0 && j == width-1) {
					result.WriteString("\\")
				} else if i == 0 || j == 0 || i == height-1 || j == width-1 {
					result.WriteString("*")
				} else {
					result.WriteString(" ")
				}
			}
			result.WriteString("\n")
		}

	case "C":
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i == 0 && (j == 0 || j == width-1) {
					result.WriteString("A")
				} else if i == height-1 && (j == 0 || j == width-1) {
					result.WriteString("C")
				} else if i == 0 || i == height-1 || j == 0 || j == width-1 {
					result.WriteString("B")
				} else {
					result.WriteString(" ")
				}
			}
			result.WriteString("\n")
		}

	case "D":
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i == 0 && j == 0 || i == height-1 && j == 0 {
					result.WriteString("A")
				} else if i == 0 && j == width-1 || i == height-1 && j == width-1 {
					result.WriteString("C")
				} else if i == 0 || i == height-1 || j == 0 || j == width-1 {
					result.WriteString("B")
				} else {
					result.WriteString(" ")
				}
			}
			result.WriteString("\n")
		}

	case "E":
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if (i == 0 && j == 0) || (i == height-1 && j == width-1) {
					result.WriteString("A")
				} else if (i == 0 && j == width-1) || (i == height-1 && j == 0) {
					result.WriteString("C")
				} else if i == 0 || j == 0 || i == height-1 || j == width-1 {
					result.WriteString("B")
				} else {
					result.WriteString(" ")
				}
			}
			result.WriteString("\n")
		}
	}

	return result.String()
}

func main() {
	// Read input from stdin
	var input strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}

	if scanner.Err() != nil {
		fmt.Println("Not a quad function")
		return
	}

	inputStr := input.String()
	if inputStr == "" {
		fmt.Println("Not a quad function")
		return
	}

	// Calculate dimensions from input
	lines := strings.Split(strings.TrimRight(inputStr, "\n"), "\n")
	if len(lines) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	height := len(lines)
	width := len(lines[0])

	// Check if all lines have the same width
	for _, line := range lines {
		if len(line) != width {
			fmt.Println("Not a quad function")
			return
		}
	}

	// Check against all quad types
	matches := []string{}
	for _, quad := range []string{"A", "B", "C", "D", "E"} {
		if generateQuad(quad, width, height) == inputStr {
			matches = append(matches, fmt.Sprintf("[quad%s] [%d] [%d]", quad, width, height))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// Sort and print matches
	sort.Strings(matches)
	fmt.Println(strings.Join(matches, " || "))
}
