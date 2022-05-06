package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("------------------------- SIGNED VALUES -------------------------\n")
	fmt.Printf("%30d %30d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("%30d %30d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("%30d %30d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("%30d %30d\n", math.MinInt64, math.MaxInt64)
	fmt.Println()

	fmt.Printf("------------------------- UNSIGNED VALUES ------------------------\n")
	fmt.Printf("%30d %30d\n", 0, math.MaxUint8)
	fmt.Printf("%30d %30d\n", 0, math.MaxUint16)
	fmt.Printf("%30d %30d\n", 0, math.MaxUint32)
	// fmt.Printf("%30d %30d\n", 0, math.MaxUint64) cannot be used here because of overflow value of MaxUint64
	fmt.Println()

	fmt.Printf("%30d %30d\n", 0, uint64(math.MaxUint64))
	fmt.Println()

	fmt.Printf("--------------------------- FLOAT VALUES -------------------------\n")
	fmt.Printf("%61v\n", math.MaxFloat32)
	fmt.Printf("%61v\n", math.MaxFloat64)

}
