package main

func isFIFO(takeOut, dineIn, served []int) bool {
	for i := 0; i < len(served); i++ {
		if len(takeOut) != 0 && served[i] == takeOut[0] {
			takeOut = takeOut[1:]
		} else if len(dineIn) != 0 && served[i] == dineIn[0] {
			dineIn = dineIn[1:]
		} else {
			return false
		}

	}
	return true
}

// Not Valid
// Take Out Orders: [1, 3, 5]
// Dine In Orders: [2, 4, 6]
// Served Orders: [1, 2, 4, 6, 5, 3]

// Valid
// Take Out Orders: [17, 8, 24]
// Dine In Orders: [12, 19, 2]
// Served Orders: [17, 8, 12, 19, 24, 2]
