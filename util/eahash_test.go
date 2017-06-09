package util

import (
	"log"
	"testing"
)

func TestEAHash(t *testing.T) {
	log.Println(EAHash("565656"))
}

func TestZeroFillRightShit(t *testing.T) {
	log.Println(zeroFillRightShit(123456789, 2))
}

func TestNumToHex(t *testing.T) {
	log.Println(numToHex(123456789))
}

func TestChunkMessage(t *testing.T) {
	log.Println(chunkMessage("123456789"))
}

func TestAdd(t *testing.T) {
	log.Println(add(12345, 12345))
}

func TestBitwiseRotate(t *testing.T) {
	log.Println(bitwiseRotate(12345, 10))
}

func TestCMN(t *testing.T) {
	log.Println(cmn(1, 2, 3, 4, 5, 6))
}

func TestMD5F(t *testing.T) {
	log.Println(md5_f(1, 2, 3, 4, 5, 6, 7))
}

func TestMD5G(t *testing.T) {
	log.Println(md5_g(1, 2, 3, 4, 5, 6, 7))
}

func TestMD5H(t *testing.T) {
	log.Println(md5_h(1, 2, 3, 4, 5, 6, 7))
}

func TestMD5I(t *testing.T) {
	log.Println(md5_i(1, 2, 3, 4, 5, 6, 7))
}
