package main

import (
	"fmt"
	"math"
)

func main() {
	var first, second int

	fmt.Scan(&first, &second)
	
	total := int64(first) * int64(second)
	fmt.Printf("Value %d saved in %s type data", total, printDataType(total))

}

func printDataType(total int64) (typeData string){

	switch {
	case total > 0:
		switch  {
		case total <= math.MaxUint8:
			typeData = "uint8"
		case total <= math.MaxUint16:
			typeData = "uint16"
		default:
			typeData = "uint32"
		}
	case total < 0:
		switch {
		case total >= math.MinInt8:
			typeData = "int8"
		case total >= math.MinInt16:
			typeData = "int16"
		case total >= math.MinInt32:
			typeData = "int32"
		default:
			typeData = "int64"
		}
	}
	return
}