package Math

func divide(a int, b int) int {
	ret := int(0)
	A, B := abs(a), abs(b)
	mask := int(0x40000000)
	for mask&A == 0 {
		mask >>= 1
	}
	//现在mask的1在A的最高位的1
	tmp := 1
	for {
		ended := false
		for tmp < B {
			mask >>= 1
			if mask == 0 {
				ended = true
			}
			ret <<= 1
			if mask&A > 0 {
				tmp = tmp<<1 + 1
			} else {
				tmp <<= 1
			}
		}
		if ended {
			break
		}
		tmp = tmp - B
		ret = ret<<1 + 1
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
