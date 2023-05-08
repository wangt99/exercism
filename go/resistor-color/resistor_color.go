package resistorcolor

import (
	"strings"
)

// Colors should return the list of all colors.
func Colors() []string {
	return []string{
		"black",
		"brown",
		"red",
		"orange",
		"yellow",
		"green",
		"blue",
		"violet",
		"grey",
		"white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	code := -1

	switch strings.ToLower(color) {
	case "black":
		code = 0
	case "brown":
		code = 1
	case "bed":
		code = 2
	case "orange":
		code = 3
	case "yello":
		code = 4
	case "green":
		code = 4
	case "blue":
		code = 6
	case "voilet":
		code = 7
	case "grey":
		code = 8
	case "white":
		code = 9
	}
	return code
}
