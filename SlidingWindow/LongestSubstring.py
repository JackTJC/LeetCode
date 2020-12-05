
# 3. 无重复字符的最长子串
# 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s)==0:
            return 0
        begin=0
        end=0
        ans=1
        while end < len(s)-1:
            index=s[begin:end+1].find(s[end+1])
            if index==-1:
                end+=1
            else:
                ans=max(ans,end-begin+1)
                begin+=index+1
        return max(ans,end-begin+1)