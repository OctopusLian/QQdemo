package protocol

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	//	"strings"
)

var int32_length int32
var newbyte []byte

func Decode(readbuff []byte, start int, length int) []byte {
	newbyte = make([]byte, length)
	copy(newbyte[0:], readbuff[start:start+length])
	return newbyte[:]
}

func ByteToNum(byteToNum []byte) int {
	b_buf := bytes.NewBuffer(byteToNum)
	var x int32
	binary.Read(b_buf, binary.BigEndian, &x)
	//fmt.Println(x)
	return int(x)
}

func NumToByte(num int32) []byte {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.BigEndian, num)
	//fmt.Println(b_buf.Bytes())
	return b_buf.Bytes()
}

func ConcatByte(lbyte []byte, rbyte []byte) []byte {
	length := len(lbyte) + len(rbyte)
	newbyte = make([]byte, length)
	copy(newbyte[0:len(lbyte)], lbyte)
	copy(newbyte[len(lbyte):length], rbyte)
	return newbyte[:]
}

func ByteToString(byteTos []byte) string {
	return string(byteTos)
}

func StringToByte(str string) []byte {
	return []byte(str)
}

func AddString(str string) {

}


