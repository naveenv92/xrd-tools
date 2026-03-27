package convert

import (
	"encoding/binary"
	"io"
	"log"
	"math"
	"os"
)

type decoder struct {
	size   int
	parser func([]byte, binary.ByteOrder) float64
}

var decoders = map[string]decoder{
	"uint16": {
		size: 2,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(bo.Uint16(b))
		},
	},
	"int16": {
		size: 2,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(int16(bo.Uint16(b)))
		},
	},
	"uint32": {
		size: 4,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(bo.Uint32(b))
		},
	},
	"int32": {
		size: 4,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(int32(bo.Uint32(b)))
		},
	},
	"float32": {
		size: 4,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(math.Float32frombits(bo.Uint32(b)))
		},
	},
	"uint64": {
		size: 8,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(bo.Uint64(b))
		},
	},
	"int64": {
		size: 8,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return float64(int64(bo.Uint64(b)))
		},
	},
	"float64": {
		size: 8,
		parser: func(b []byte, bo binary.ByteOrder) float64 {
			return math.Float64frombits(bo.Uint64(b))
		},
	},
}

func ConvertBinary(filename string, encoding string, littleEndian bool, width, height int32) []float64 {

	var byteOrder binary.ByteOrder
	if littleEndian {
		byteOrder = binary.LittleEndian
	} else {
		byteOrder = binary.BigEndian
	}

	d, ok := decoders[encoding]
	if !ok {
		log.Fatalf("unsupported encoding: %s", encoding)
	}

	buf := make([]byte, d.size)
	parseFunc := d.parser

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	res := []float64{}
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF || n == 0 {
				break
			}

			log.Fatalf("error reading file: %v", err)
		}

		if n < len(buf) {
			log.Fatalf("error: less than %d bytes read, check encoding", len(buf))
		}
		val := parseFunc(buf, byteOrder)
		res = append(res, val)
	}

	return res
	//var wg sync.WaitGroup
}
