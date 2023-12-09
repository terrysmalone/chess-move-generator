package lookuptables

// Hardcoded since these will be referenced A LOT
const (
	A1 = uint64(1)
	B1 = uint64(2)
	C1 = uint64(4)
	D1 = uint64(8)
	E1 = uint64(16)
	F1 = uint64(32)
	G1 = uint64(64)
	H1 = uint64(128)

	A2 = uint64(256)
	B2 = uint64(512)
	C2 = uint64(1024)
	D2 = uint64(2048)
	E2 = uint64(4096)
	F2 = uint64(8192)
	G2 = uint64(16384)
	H2 = uint64(32768)

	A3 = uint64(65536)
	B3 = uint64(131072)
	C3 = uint64(262144)
	D3 = uint64(524288)
	E3 = uint64(1048576)
	F3 = uint64(2097152)
	G3 = uint64(4194304)
	H3 = uint64(8388608)

	A4 = uint64(16777216)
	B4 = uint64(33554432)
	C4 = uint64(67108864)
	D4 = uint64(134217728)
	E4 = uint64(268435456)
	F4 = uint64(536870912)
	G4 = uint64(1073741824)
	H4 = uint64(2147483648)

	A5 = uint64(4294967296)
	B5 = uint64(8589934592)
	C5 = uint64(17179869184)
	D5 = uint64(34359738368)
	E5 = uint64(68719476736)
	F5 = uint64(137438953472)
	G5 = uint64(274877906944)
	H5 = uint64(549755813888)

	A6 = uint64(1099511627776)
	B6 = uint64(2199023255552)
	C6 = uint64(4398046511104)
	D6 = uint64(8796093022208)
	E6 = uint64(17592186044416)
	F6 = uint64(35184372088832)
	G6 = uint64(70368744177664)
	H6 = uint64(140737488355328)

	A7 = uint64(281474976710656)
	B7 = uint64(562949953421312)
	C7 = uint64(1125899906842624)
	D7 = uint64(2251799813685248)
	E7 = uint64(4503599627370496)
	F7 = uint64(9007199254740992)
	G7 = uint64(18014398509481984)
	H7 = uint64(36028797018963968)

	A8 = uint64(72057594037927936)
	B8 = uint64(144115188075855872)
	C8 = uint64(288230376151711744)
	D8 = uint64(576460752303423488)
	E8 = uint64(1152921504606846976)
	F8 = uint64(2305843009213693952)
	G8 = uint64(4611686018427387904)
	H8 = uint64(9223372036854775808)
)

var BitboardValueFromPosition [8][8]uint64
var BitboardValueFromIndex [64]uint64

func init() {
	initialiseBitboardValueFromPosition()
	initialiseBitboardValueFromIndex()
}

func initialiseBitboardValueFromPosition() {
	squareValue := uint64(1)

	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			BitboardValueFromPosition[column][row] = squareValue
			squareValue <<= 1
		}
	}
}

func initialiseBitboardValueFromIndex() {
	BitboardValueFromIndex[0] = 1

	for i := 1; i < len(BitboardValueFromIndex); i++ {
		//Left shift gives the same result as multiplying by two but is faster
		BitboardValueFromIndex[i] = BitboardValueFromIndex[i-1] << 1
	}
}
