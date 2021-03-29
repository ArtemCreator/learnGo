package main

import (
	"fmt"
	"math"
)

func main() {
	result := math.MaxUint32 / math.MaxUint8
	result2 := math.MaxUint32 / math.MaxUint16
	fmt.Println(result, result2)
}