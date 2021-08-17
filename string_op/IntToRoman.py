# 12. 整数转罗马数字
# 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
class Solution:
    def intToRoman(self, num: int) -> str:
        hashTable1=['I','II','III','IV','V','VI','VII','VIII','IX']
        hashTable2=['X','XX','XXX','XL','L','LX','LXX','LXXX','XC']
        hashTable3=['C','CC','CCC','CD','D','DC','DCC','DCCC','CM']
        hashTable4=['M','MM','MMM']
        allHash=[hashTable1,hashTable2,hashTable3,hashTable4]
        base=10
        i=0
        ans=[]
        while num!=0:
            res=num%base
            if res!=0:
                ans.append(allHash[i][res-1])
            # base*=10
            num=int(num/10)
            i+=1
        return ''.join(ans[::-1])

