# 43. 字符串相乘
# 给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积
# 它们的乘积也表示为字符串形式。

class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        # 模拟手工乘法
        if num2=='0' or num1=='0':
            return 0
        resArray=[0]*(len(num1)+len(num2))
        m=len(num1)
        n=len(num2)
        Ci=0
        for i in range(1,m+1):
            Ci=0
            for j in range(1,n+1):
                bitSum=(int(num1[m-i]))*(int(num2[n-j]))+Ci+resArray[i+j-2]
                resArray[i+j-2]=bitSum%10
                if bitSum>=10:
                    Ci=bitSum//10
                else:
                    Ci=0
        if Ci!=0:
            resArray[-1]=Ci
            resArray=resArray[::-1]
        else:#去除首部0
            resArray=resArray[::-1]
            idx=0
            while resArray[idx]==0:
                idx+=1
            resArray=resArray[idx:]
        return ''.join([str(i) for i in resArray])

if __name__ == '__main__':
    Solution.multiply(None,"999","999")