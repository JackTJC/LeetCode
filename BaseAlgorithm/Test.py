class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        # 模拟手工乘法
        if num2=='0' or num1=='0':
            return 0
        resArray=[0]*(len(num1)+len(num2))
        m=len(num1)
        n=len(num2)
        Ci=0
        i=0
        for n1 in num1[::-1]:
            j=0
            Ci=0
            for n2 in num2[::-1]:
                bitSum=int(n1)*int(n2)+Ci+resArray[i+j]
                resArray[i+j]=bitSum%10
                if bitSum>=10:
                    Ci=bitSum//10
                else:
                    Ci=0
                j+=1
            i+=1
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
    print(Solution.multiply(None,"9999","9"))