package math

//快速幂算法
//参考：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/solution/mian-shi-ti-16-shu-zhi-de-zheng-shu-ci-fang-kuai-s/
//核心在于将n表达为二进制形式，然后做累乘
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	ret := 1.0
	for n > 0 {
		if n&1 > 0 {
			ret *= x
		}
		x *= x
		n >>= 1
	}
	return ret
}
