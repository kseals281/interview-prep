package main

import "strconv"

//Designer Mary wants to convert her CSS from statements like background-color: rgb(128, 128, 0); to:
//background-color: #808000;.
//
//Don't worry, you don't have to rewrite CSS, just complete the rgb_to_hex helper for her that takes a tuple of
//RGB (3 ints) and converts it to the corresponding hex value.
//
//For example for Maroon you would call it like rgb_to_hex{"name", [3]int{ 128, 0, 0)) and it would return the hex #800000.
//
//Make sure you check that each r, g and b int is within the valid range of 0 and 255 (both included), if not raise a
//ValueError, Explicit is better than implicit (Zen).
//
//Check the TESTS tab for more examples, there we check your code for 16 colors and 3 excepions. Have fun!

type ErrOutOfRange struct {
	err string
}

func NewErrOutOfRange(s string) *ErrOutOfRange {
	return &ErrOutOfRange{
		err: s,
	}
}

func (e *ErrOutOfRange) Error() string {
	return e.err
}

func rgbToHex(rgb [3]int) (hex string, err error) {
	//Receives (r, g, b)  array, checks if each rgb int is within RGB
	//boundaries (0, 255) and returns its converted hex, for example:
	//Silver: input tuple = (192,192,192) -> output hex str = #C0C0C0
	hex = "#"
	for _, i := range rgb {
		if i < 0 || i > 255 {
			return "", NewErrOutOfRange("Value %d not within range of 0-255")
		}
		hex += hexValue(int(i / 16))
		hex += hexValue(i % 16)
	}
	return hex, nil
}

func hexValue(num int) (hex string) {
	switch num {
	case 10:
		hex = "A"
	case 11:
		hex = "B"
	case 12:
		hex = "C"
	case 13:
		hex = "D"
	case 14:
		hex = "E"
	case 15:
		hex = "F"
	default:
		hex = strconv.Itoa(num)
	}
	return hex
}
