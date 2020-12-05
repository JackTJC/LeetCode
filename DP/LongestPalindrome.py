# 5. 最长回文子串
# 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
# python动态规划超时

class Solution:
    def longestPalindrome(self, s: str) -> str:
        dp = [[0] * len(s) for _ in range(len(s))]
        ans = ""
        for i in range(len(s)):
            for j in range(len(s) - i):
                if i == 0:
                    dp[j][i + j] = 1
                elif i == 1:
                    dp[j][i + j] = 1 if s[j] == s[i + j] else 0
                else:
                    dp[j][i + j] = 1 if (dp[j + 1][i + j - 1] == 1 and s[j] == s[i + j]) else 0
                if dp[j][i + j] == 1:
                    ans = s[j:i + j + 1] if len(s[j:i + j + 1]) > len(ans) else ans
        return ans

# Java通过
# class Solution {
#     public String longestPalindrome(String s) {
#         boolean [][] dp = new boolean[s.length()][s.length()];
#         String ans="";
#         for(int i=0;i<s.length();i++)
#         {
#             for(int j=0;j<s.length()-i;j++)
#             {
#                 if(i==0)
#                     dp[j][i+j]=true;
#                 else if (i==1)
#                     dp[j][i+j]=(s.charAt(j)==s.charAt(i+j));
#                 else
#                     dp[j][i+j]=(dp[j+1][i+j-1]&&s.charAt(j)==s.charAt(i+j));
#                 if(dp[j][i+j]&&(i+1>ans.length()))
#                     ans=s.substring(j,i+j+1);
#             }
#         }
#         return ans;
#     }
# }
