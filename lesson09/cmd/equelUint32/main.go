package main

import (
	"fmt"
	"math"
)

func main() {
	result := math.MaxUint32 / (1<<8 - 1)
	result2 := math.MaxUint32 / (1<<16 - 1)

	fmt.Printf("the number of overflows of types uint8: %d, uint16: %d in the range 0 to uint32\n",
		result, result2)

	var counterUint8, counterUint16 uint

	for i := 1; i < math.MaxUint32; i++ {
		if uint16(i) == 0 {
			counterUint16++
		}
		if uint8(i) == 0 {
		counterUint8++
		}
	}

	fmt.Printf("the number of overflows of types uint8: %d, uint16: %d in the range 0 to uint32",
		counterUint8, counterUint16)
}
