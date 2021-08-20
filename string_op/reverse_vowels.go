package string_op

var (
	vowel = map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}
)

func reverseVowels(s string) string {
	ret := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < len(ret) {
			_, ok := vowel[ret[left]]
			if ok {
				break
			}
			left++
		}
		for right >= 0 {
			_, ok := vowel[ret[right]]
			if ok {
				break
			}
			right--
		}
		if left < right {
			ret[left], ret[right] = ret[right], ret[left]
			left++
			right--

		}
	}
	return string(ret)
}
