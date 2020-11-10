package String_;
//#5 middle expand algorithm
public class LongestPalindrome {

    public static void main(String[] args) {
        System.out.println(new Solution().longestPalindrome("xababax"));
    }
    static class Solution {
        public String longestPalindrome(String s) {
            int begin=0,end=0;
            int max=0;
            if(s==null||s.length()==0)
                return "";
            for(int i=0;i<s.length();i++)
            {
                int len1,len2;
                len1=expandFromCenter(s,i,i);
                len2=expandFromCenter(s,i,i+1);
                int maxLen=Math.max(len1,len2);
                if(maxLen>max)
                {
                    max=maxLen;
                    begin=i-(max-1)/2;
                    end=i+max/2;
                }
            }
            return s.substring(begin,end+1);
        }
        public int expandFromCenter(String s,int pos1,int pos2)
        {
            while(pos1>=0&&pos2<s.length()&&(s.charAt(pos1)==s.charAt(pos2)))
            {
                pos1--;
                pos2++;
            }
            return pos2-pos1-1;
        }
    }
}

