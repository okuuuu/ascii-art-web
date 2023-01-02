package res

import (
	"strings"
)

// All this running around
// I can't fight it much longer
// Something's tryin' to get out
// And it's never been closer

func AsciiArt(wordstr string, banner string) string {

	// Set banner.

	banner = "banners/" + banner + ".txt"

	// Generate rune-[8]string map from ascii art file.

	mapChar := ReadAscii(banner)

	word := strings.Split(wordstr, "\r\n")
	wordAscii := WriteAscii(word, mapChar)

	// Print into different outputs.
	return PrintAscii(word, wordAscii)
}
