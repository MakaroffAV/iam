package main

import (
	"fmt"
	"t1/pkg/tempconv"
)

func main() {
	zeroC := tempconv.Celsius(0)
	zeroK := tempconv.Kelvins(0)
	fmt.Println(tempconv.CToK(zeroC).String())
	fmt.Println(tempconv.KToC(zeroK).String())

}
