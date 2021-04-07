package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter two numbers")
	var first, second int16

	_, _ = fmt.Scan(&first, &second)

	total := int64(first) * int64(second)
	fmt.Printf("Value %d well be save in %s type data\n", total, printDataType(total))

}

func printDataType(total int64) (typeData string) {
	switch {
	case total > math.MaxUint32:
		typeData = "uint64"
	case total > math.MaxUint16:
		typeData = "uint32"
	case total > math.MaxUint8:
		typeData = "uint16"
	case total > 0:
		typeData = "uint8"
	case total > math.MinInt8:
		typeData = "int8"
	case total > math.MinInt16:
		typeData = "int16"
	case total > math.MinInt32:
		typeData = "int32"
	default:
		typeData = "int64"
	}
	return
}
