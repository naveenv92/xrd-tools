package convert

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type PixelValue interface {
	uint32 | int32 | float32 | float64
}

func ConvertBinary(filename string, encoding string, littleEndian bool, width, height int32) []float64 {

	var byteOrder binary.ByteOrder
	if littleEndian {
		byteOrder = binary.LittleEndian
	} else {
		byteOrder = binary.BigEndian
	}

	var buf []byte
	var parseFunc func([]byte) float64
	switch encoding {
	case "uint32":
		buf = make([]byte, 4)
		parseFunc = func(b []byte) float64 {
			return float64(byteOrder.Uint32(b))
		}
	case "uint64":
		buf = make([]byte, 8)
		parseFunc = func(b []byte) float64 {
			return float64(byteOrder.Uint64(b))
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		return nil
	}
	defer file.Close()

	res := []float64{}
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF || n == 0 {
				break
			}
			fmt.Printf("error reading file: %v", err)
			return nil
		}

		val := parseFunc(buf)
		res = append(res, val)
	}

	return res
	//var wg sync.WaitGroup
}
