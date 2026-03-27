package main

import (
	"fmt"

	"github.com/naveenv92/xrd-tools/pkg/convert"
)

func main() {
	res := convert.ConvertBinary("/Users/naveenvenkatesan/Desktop/data.bin", "uint16", true, 0, 0)
	fmt.Println(res)
}
