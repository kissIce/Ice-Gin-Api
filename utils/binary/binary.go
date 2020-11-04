package binary

func SetBitOne(num int, step int) int {
	return num | (1 << step)
}

func SetBitZero(num int, step int) int {
	mask := ^(1 << step)
	return num & mask
}

func GetBit(num int, step int) bool {
	return (num & (1 << step)) != 0
}

func CountBitOne(num int) (count int) {
	for count = 0; num > 0; count++ {
		num &= num - 1
	}
	return
}

func CountBitOne1(num int) (tmp int) {
	tmp = num - ((num >> 1) & 033333333333) - ((num >> 2) & 011111111111)
	return ((tmp + (tmp >> 3)) & 030707070707) % 63
}

func CountSuccessiveOne(num int) (k int) {
	for k = 0; num != 0; k++ {
		num = num & (num << 1)
	}
	return
}