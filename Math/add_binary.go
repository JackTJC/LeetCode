package Math

func addBinary(a string, b string) string {
	Cout := 0 //进位
	ret := make([]byte, 0)
	idxA, idxB := len(a)-1, len(b)-1
	for idxA >= 0 || idxB >= 0 {
		var Ci, Ai, Bi int
		if idxA < 0 {
			Ai, Bi = 0, toInt(b[idxB])
		} else if idxB < 0 {
			Ai, Bi = toInt(a[idxA]), 0
		} else {
			Ai, Bi = toInt(a[idxA]), toInt(b[idxB])
		}
		sum := Ai + Bi + Cout
		Ci, Cout = sum%2, sum/2
		ret = append([]byte{byte('0' + Ci)}, ret...)
		idxA--
		idxB--
	}
	if Cout == 1 {
		ret = append([]byte{'1'}, ret...)
	}
	return string(ret)
}

func toInt(x byte) int {
	if x == byte('0') {
		return 0
	}
	return 1
}
