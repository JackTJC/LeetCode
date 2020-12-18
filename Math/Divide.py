# 29. 两数相除
# 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
#
# 返回被除数 dividend 除以除数 divisor 得到的商。
#
# 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        def binary2decimal(binary):
            i=1
            sum=0
            for bit in binary[::-1]:
                sum+=i if bit=='1' else 0
                i<<=1
            return sum
        def binarySub(d1,d2):
            res=binary2decimal(d1)-binary2decimal(d2)
            return bin(res)[2:]
        # 二进制除法算法
        if abs(dividend)<abs(divisor):
            return 0
        sign1=True if dividend<0 else False
        sign2=True if divisor<0 else False
        b1=str(bin(abs(dividend)))[2:]
        b2=str(bin(abs(divisor)))[2:]
        headerBits=""
        res=[]
        for bit in b1:
            headerBits+=bit
            if int(headerBits)>=int(b2):
                res.append('1')
                headerBits=binarySub(headerBits,b2)
            else:
                res.append('0')
        strAns=''.join(res)
        sign=sign1^sign2
        ans=binary2decimal(strAns)
        ans = -ans if sign else ans
        if -2**31<=ans<=2**31-1:
            return ans
        else:
            return 2**31-1