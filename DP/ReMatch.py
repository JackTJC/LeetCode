# 10. 正则表达式匹配
# 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
#
# '.' 匹配任意单个字符
# '*' 匹配零个或多个前面的那一个元素
# 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        m, n = len(s), len(p)
        dp = [[False] * (n + 1) for _ in range(m + 1)]
        dp[0][0] = True

        def matchChr(i, j):
            if i == 0:
                return False
            if p[j - 1] == '.':
                return True
            return s[i - 1] == p[j - 1]

        for i in range(m + 1):
            for j in range(1, n + 1):
                if p[j - 1] == '*':
                    if matchChr(i, j - 1):
                        dp[i][j] = dp[i][j - 2] or dp[i - 1][j]
                    else:
                        dp[i][j] = dp[i][j - 2]
                else:
                    dp[i][j] = dp[i - 1][j - 1] and matchChr(i, j)
        return dp[m][n]


"""
以一个例子详解动态规划转移方程：
S = abbbbc
P = ab*d*c
1. 当 i, j 指向的字符均为字母（或 '.' 可以看成一个特殊的字母）时，
   只需判断对应位置的字符即可，
   若相等，只需判断 i,j 之前的字符串是否匹配即可，转化为子问题 f[i-1][j-1].
   若不等，则当前的 i,j 肯定不能匹配，为 false.
   
       f[i-1][j-1]   i
            |        |
   S [a  b  b  b  b][c] 
   
   P [a  b  *  d  *][c]
                     |
                     j
   

2. 如果当前 j 指向的字符为 '*'，则不妨把类似 'a*', 'b*' 等的当成整体看待。
   看下面的例子

            i
            |
   S  a  b [b] b  b  c  
   
   P  a [b  *] d  *  c
            |
            j
   
   注意到当 'b*' 匹配完 'b' 之后，它仍然可以继续发挥作用。
   因此可以只把 i 前移一位，而不丢弃 'b*', 转化为子问题 f[i-1][j]:
   
         i
         | <--
   S  a [b] b  b  b  c  
   
   P  a [b  *] d  *  c
            |
            j
   
   另外，也可以选择让 'b*' 不再进行匹配，把 'b*' 丢弃。
   转化为子问题 f[i][j-2]:

            i
            |
   S  a  b [b] b  b  c  
    
   P [a] b  *  d  *  c
      |
      j <--

3. 冗余的状态转移不会影响答案，
   因为当 j 指向 'b*' 中的 'b' 时, 这个状态对于答案是没有用的,
   原因参见评论区 稳中求胜 的解释, 当 j 指向 '*' 时,
   dp[i][j]只与dp[i][j-2]有关, 跳过了 dp[i][j-1].
"""
