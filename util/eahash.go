package util

import "fmt"

var r1Shifts = []int32{7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22}
var r2Shifts = []int32{5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20}
var r3Shifts = []int32{4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23}
var r4Shifts = []int32{6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21}

var hexChar = "0123456789abcdef"

func zeroFillRightShit(data int64, bits byte) int64 {
	return (data & 0xffffffff) >> bits
}

func numToHex(num int32) string {
	temp := ""

	for i := byte(0); i < 4; i++ {
		x := hexChar[(num>>(i*8+4))&0x0F]
		y := hexChar[(num>>(i*8))&0x0F]
		temp += fmt.Sprintf("%c%c", x, y)
	}
	return temp
}

func chunkMessage(str string) []int32 {
	nblk := ((len(str) + 8) >> 6) + 1
	blks := make([]int32, nblk*16)

	for i := byte(0); i < byte(len(str)); i++ {
		blks[i>>2] |= int32(int64(str[i]) << ((i % 4) * 8))
		//println(i>>2, str[i], (i%4)*8, int32(int64(str[i])<<((i%4)*8)))
	}
	i := byte(len(str))

	blks[i>>2] |= 0x80 << ((i % 4) * 8)
	blks[nblk*16-2] = int32(len(str)) * 8
	return blks
}

func add(x, y int32) int32 {
	lsw := (x & 0xFFFF) + (y & 0xFFFF)
	msw := int32(x>>16) + int32(y>>16) + int32(lsw>>16)
	return int32(int32(msw<<16) | (lsw & 0xFFFF))
}

func bitwiseRotate(x, c int32) int32 {
	return int32(int32(x<<byte(c)) | int32(zeroFillRightShit(int64(x), byte(32-c))))
}

func cmn(q, a, b, x, s, t int32) int32 {
	z1 := add(a, q)
	z2 := add(x, t)
	a1 := add(z1, z2)

	x1 := bitwiseRotate(a1, s)
	return add(x1, b)
}

func md5_f(a, b, c, d, x, s, t int32) int32 {
	return cmn(int32((b&c)|((^b)&d)), a, b, x, s, t)
}

func md5_g(a, b, c, d, x, s, t int32) int32 {
	return cmn(int32((b&d)|(c&(^d))), a, b, x, s, t)
}

func md5_h(a, b, c, d, x, s, t int32) int32 {
	return cmn(int32(b^c^d), a, b, x, s, t)
}

func md5_i(a, b, c, d, x, s, t int32) int32 {
	return cmn(int32(c^(b|(^d))), a, b, x, s, t)
}

func EAHash(str string) string {
	x := chunkMessage(str)

	a := int32(1732584193)
	b := int32(-271733879)
	c := int32(-1732584194)
	d := int32(271733878)

	for i := 0; i < 16; i += 16 {

		tempA := a
		tempB := b
		tempC := c
		tempD := d

		// F
		a = md5_f(a, b, c, d, x[i+0], r1Shifts[0], -680876936)
		d = md5_f(d, a, b, c, x[i+1], r1Shifts[1], -389564586)
		c = md5_f(c, d, a, b, x[i+2], r1Shifts[2], 606105819)
		b = md5_f(b, c, d, a, x[i+3], r1Shifts[3], -1044525330)

		a = md5_f(a, b, c, d, x[i+4], r1Shifts[4], -176418897)
		d = md5_f(d, a, b, c, x[i+5], r1Shifts[5], 1200080426)
		c = md5_f(c, d, a, b, x[i+6], r1Shifts[6], -1473231341)
		b = md5_f(b, c, d, a, x[i+7], r1Shifts[7], -45705983)

		a = md5_f(a, b, c, d, x[i+8], r1Shifts[8], 1770035416)
		d = md5_f(d, a, b, c, x[i+9], r1Shifts[9], -1958414417)
		c = md5_f(c, d, a, b, x[i+10], r1Shifts[10], -42063)
		b = md5_f(b, c, d, a, x[i+11], r1Shifts[11], -1990404162)

		a = md5_f(a, b, c, d, x[i+12], r1Shifts[12], 1804603682)
		d = md5_f(d, a, b, c, x[i+13], r1Shifts[13], -40341101)
		c = md5_f(c, d, a, b, x[i+14], r1Shifts[14], -1502002290)
		b = md5_f(b, c, d, a, x[i+15], r1Shifts[15], 1236535329)

		// G
		a = md5_g(a, b, c, d, x[i+1], r2Shifts[0], -165796510)
		d = md5_g(d, a, b, c, x[i+6], r2Shifts[1], -1069501632)
		c = md5_g(c, d, a, b, x[i+11], r2Shifts[2], 643717713)
		b = md5_g(b, c, d, a, x[i+0], r2Shifts[3], -373897302)

		a = md5_g(a, b, c, d, x[i+5], r2Shifts[4], -701558691)
		d = md5_g(d, a, b, c, x[i+10], r2Shifts[5], 38016083)
		c = md5_g(c, d, a, b, x[i+15], r2Shifts[6], -660478335)
		b = md5_g(b, c, d, a, x[i+4], r2Shifts[7], -405537848)

		a = md5_g(a, b, c, d, x[i+9], r2Shifts[8], 568446438)
		d = md5_g(d, a, b, c, x[i+14], r2Shifts[9], -1019803690)
		c = md5_g(c, d, a, b, x[i+3], r2Shifts[10], -187363961)
		b = md5_g(b, c, d, a, x[i+8], r2Shifts[11], 1163531501)

		a = md5_g(a, b, c, d, x[i+13], r2Shifts[12], -1444681467)
		d = md5_g(d, a, b, c, x[i+2], r2Shifts[13], -51403784)
		c = md5_g(c, d, a, b, x[i+7], r2Shifts[14], 1735328473)
		b = md5_g(b, c, d, a, x[i+12], r2Shifts[15], -1926607734)

		// H
		a = md5_h(a, b, c, d, x[i+5], r3Shifts[0], -378558)
		d = md5_h(d, a, b, c, x[i+8], r3Shifts[1], -2022574463)
		// line below uses r2Shifts[2] where as MD5 would use r3Shifts[2]
		c = md5_h(c, d, a, b, x[i+11], r2Shifts[2], 1839030562)
		b = md5_h(b, c, d, a, x[i+14], r3Shifts[3], -35309556)

		a = md5_h(a, b, c, d, x[i+1], r3Shifts[4], -1530992060)
		d = md5_h(d, a, b, c, x[i+4], r3Shifts[5], 1272893353)
		c = md5_h(c, d, a, b, x[i+7], r3Shifts[6], -155497632)
		b = md5_h(b, c, d, a, x[i+10], r3Shifts[7], -1094730640)

		a = md5_h(a, b, c, d, x[i+13], r3Shifts[8], 681279174)
		d = md5_h(d, a, b, c, x[i+0], r3Shifts[9], -358537222)
		c = md5_h(c, d, a, b, x[i+3], r3Shifts[10], -722521979)
		b = md5_h(b, c, d, a, x[i+6], r3Shifts[11], 76029189)

		a = md5_h(a, b, c, d, x[i+9], r3Shifts[12], -640364487)
		d = md5_h(d, a, b, c, x[i+12], r3Shifts[13], -421815835)
		c = md5_h(c, d, a, b, x[i+15], r3Shifts[14], 530742520)
		b = md5_h(b, c, d, a, x[i+2], r3Shifts[15], -995338651)

		// I
		a = md5_i(a, b, c, d, x[i+0], r4Shifts[0], -198630844)
		d = md5_i(d, a, b, c, x[i+7], r4Shifts[1], 1126891415)
		c = md5_i(c, d, a, b, x[i+14], r4Shifts[2], -1416354905)
		b = md5_i(b, c, d, a, x[i+5], r4Shifts[3], -57434055)

		a = md5_i(a, b, c, d, x[i+12], r4Shifts[4], 1700485571)
		d = md5_i(d, a, b, c, x[i+3], r4Shifts[5], -1894986606)
		c = md5_i(c, d, a, b, x[i+10], r4Shifts[6], -1051523)
		b = md5_i(b, c, d, a, x[i+1], r4Shifts[7], -2054922799)

		a = md5_i(a, b, c, d, x[i+8], r4Shifts[8], 1873313359)
		d = md5_i(d, a, b, c, x[i+15], r4Shifts[9], -30611744)
		c = md5_i(c, d, a, b, x[i+6], r4Shifts[10], -1560198380)
		b = md5_i(b, c, d, a, x[i+13], r4Shifts[11], 1309151649)

		a = md5_i(a, b, c, d, x[i+4], r4Shifts[12], -145523070)
		d = md5_i(d, a, b, c, x[i+11], r4Shifts[13], -1120210379)
		c = md5_i(c, d, a, b, x[i+2], r4Shifts[14], 718787259)
		b = md5_i(b, c, d, a, x[i+9], r4Shifts[15], -343485551)
		// This line is doubled for some reason, line below is not in the MD5 version
		b = md5_i(b, c, d, a, x[i+9], r4Shifts[15], -343485551)

		a = add(a, tempA)
		b = add(b, tempB)
		c = add(c, tempC)
		d = add(d, tempD)
	}

	return numToHex(a) + numToHex(b) + numToHex(c) + numToHex(d)
}
