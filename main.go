package main

import (
	"machine"
	"strings"
	"time"
)

var morseCode = map[rune][]byte{
	'A': {'.', '-'}, 'B': {'-', '.', '.', '.'}, 'C': {'-', '.', '-', '.'}, 'D': {'-', '.', '.'},
	'E': {'.'}, 'F': {'.', '.', '-', '.'}, 'G': {'-', '-', '.'}, 'H': {'.', '.', '.', '.'},
	'I': {'.', '.'}, 'J': {'.', '-', '-', '-'}, 'K': {'-', '.', '-'}, 'L': {'.', '-', '.', '.'},
	'M': {'-', '-'}, 'N': {'-', '.'}, 'O': {'-', '-', '-'}, 'P': {'.', '-', '-', '.'},
	'Q': {'-', '-', '.', '-'}, 'R': {'.', '-', '.'}, 'S': {'.', '.', '.'}, 'T': {'-'},
	'U': {'.', '.', '-'}, 'V': {'.', '.', '.', '-'}, 'W': {'.', '-', '-'}, 'X': {'-', '.', '.', '-'},
	'Y': {'-', '.', '-', '-'}, 'Z': {'-', '-', '.', '.'},
	'0': {'-', '-', '-', '-', '-'}, '1': {'.', '-', '-', '-', '-'}, '2': {'.', '.', '-', '-', '-'},
	'3': {'.', '.', '.', '-', '-'}, '4': {'.', '.', '.', '.', '-'}, '5': {'.', '.', '.', '.', '.'},
	'6': {'-', '.', '.', '.', '.'}, '7': {'-', '-', '.', '.', '.'}, '8': {'-', '-', '-', '.', '.'},
	'9': {'-', '-', '-', '-', '.'},
	'.': {'.', '-', '.', '-', '.', '-'}, ',': {'-', '-', '.', '.', '-', '-'},
	'?': {'.', '.', '-', '-', '.', '.'},
	'!': {'-', '.', '-', '.', '-', '-'},
}

const (
	DOT  = time.Millisecond * 100
	DASH = time.Millisecond * 300
	WORD = time.Millisecond * 700
)

var (
	led = machine.LED
)

func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func main() {
	word := strings.ToUpper("SOS Means I Need Help!")

	for i, char := range word {
		code := morseCode[char]

		for y, d := range code {
			var inType, endType time.Duration

			switch d {
			case '.':
				inType = DOT
			case '-':
				inType = DASH
			}

			if i+1 < len(word) && word[i+1] == ' ' {
				endType = WORD
			} else if y == len(code)-1 {
				endType = DASH
			} else {
				endType = DOT
			}

			flash(inType, endType)
		}

	}
}

func flash(inType time.Duration, endType time.Duration) {
	led.High()
	time.Sleep(inType)
	led.Low()
	time.Sleep(endType)
}
