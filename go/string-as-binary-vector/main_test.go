package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {

	// a string printed as hex bytes
	s1 := "abcd"
	fmt.Printf("s1: % x\n", s1)

	// concat a "byte" cast to a string
	s2 := s1 + string([]byte{0}) + string([]byte{0x7f})
	fmt.Printf("s2: % x\n", s2)

	// concat a "byte" > 0x7f => encoding rules
	s3 := s1 + string([]byte{0}) + string([]byte{0x80}) // 2 bytes concatinated
	fmt.Printf("s3: % x\n", s3)

	// larger (>0xff) numbers are also "encoded" as UTF ruins
	s4 := s1 + string([]byte{0}) + string([]byte{0x0f, 0x7f, 0xff}) + string([]byte{0})
	fmt.Printf("s4: % x\n", s4)

	// Creating a buffer from string works as expected
	b1 := bytes.NewBufferString("ABCD")
	fmt.Printf("b1: % x\n", b1.String())

	// writing a byte does NOT trigger encoding
	b1.WriteByte(0x7f)
	b1.WriteByte(0xff) // this time 0xff is simply appended
	fmt.Printf("b1: % x\n", b1.String())

	// the binary encoding is variable length, and ..
	var num uint32 = 0xdeadbeef
	b2 := make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(b2, uint64(num))
	b1.WriteByte(0x00) // marker
	b1.Write(b2)
	fmt.Printf("b1: % x\n", b1.String())

	// //////////

	// AppendInt creates a base-n rendering of the value
	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	// //////////

	// Finally - direct encoding of an integer
	//
	//  binary encoding using the BigEndian ByteOrder
	//
	var val uint32 = 0xdeadbeef
	arr := make([]byte, 4)
	binary.BigEndian.PutUint32(arr, val) // 4 byte slice, expected encoding

	buf := bytes.NewBufferString("abcd")
	buf.WriteByte(0x00) // seperator
	buf.Write(arr)
	fmt.Printf("buf: % x\n", buf.String())
}
