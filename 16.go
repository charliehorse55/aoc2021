package main

import (
	"fmt"
	"math"
	"math/big"
)

type bitsBuffer struct {
	pos uint
	data *big.Int
}

func (b *bitsBuffer)GetBits(n int) uint {
	bitLen := uint(b.data.BitLen())
	result := uint(0)
	for i := 0; i < n; i++ {
		result <<= 1
		currIndex := bitLen - (b.pos+1)
		result |= b.data.Bit(int(currIndex))
		b.pos++
	}
	return result
}

var operators = map[uint]func([]uint)uint {
	0: func(uints []uint) uint {
		sum := uint(0)
		for _,val := range uints {
			sum += val
		}
		return sum
	},
	1: func(uints []uint) uint {
		prod := uints[0]
		for _,val := range uints[1:] {
			prod *= val
		}
		return prod
	},
	2: func(uints []uint) uint {
		min := uint(math.MaxUint)
		for _,val := range uints {
			if val < min {
				min = val
			}
		}
		return min
	},
	3: func(uints []uint) uint {
		max := uint(0)
		for _,val := range uints {
			if val > max {
				max = val
			}
		}
		return max
	},
	5: func(uints []uint) uint {
		if uints[0] > uints[1] {
			return 1
		} else {
			return 0
		}
	},
	6: func(uints []uint) uint {
		if uints[0] < uints[1] {
			return 1
		} else {
			return 0
		}
	},
	7: func(uints []uint) uint {
		if uints[0] == uints[1] {
			return 1
		} else {
			return 0
		}
	},
}

func parsePacket(buf *bitsBuffer) (versionSum, value uint) {
	versionSum += buf.GetBits(3)
	kind := buf.GetBits(3)
	switch kind {
	case 0b100:
		literal := uint(0)
		for {
			contBit := buf.GetBits(1)
			data := buf.GetBits(4)
			literal <<= 4
			literal |= data
			if contBit == 0 {
				break
			}
		}
		value = literal
		//fmt.Printf("found packet version %d, type %d, literal value = %d\n", versionSum, kind, literal)
	default:
		lengthType := buf.GetBits(1)
		var packetValues []uint
		if lengthType == 0 {
			length := buf.GetBits(15)
			targetL := buf.pos + length
			//fmt.Printf("found packet version %d, type %d, length = %d bits\n", versionSum, kind, length)
			for buf.pos < targetL {
				vs, val := parsePacket(buf)
				versionSum += vs
				packetValues = append(packetValues, val)
			}
		} else {
			npackets := buf.GetBits(11)
			//fmt.Printf("found packet version %d, type %d, length = %d packets\n", versionSum, kind, npackets)
			for i := uint(0); i < npackets; i++ {
				vs, val := parsePacket(buf)
				versionSum += vs
				packetValues = append(packetValues, val)
			}
		}
		value = operators[kind](packetValues)
	}
	return
}


func adventDay16A(path string) {
	lines := readLines(path)
	buf := bitsBuffer{
		data: new(big.Int),
	}
	buf.data.SetString(lines[0], 16)

	vs, _ := parsePacket(&buf)
	fmt.Printf("version sum is %d\n", vs)
}


func adventDay16B(path string) {
	lines := readLines(path)
	buf := bitsBuffer{
		data: new(big.Int),
	}
	buf.data.SetString(lines[0], 16)

	_, val := parsePacket(&buf)
	fmt.Printf("val is %d\n", val)

}
