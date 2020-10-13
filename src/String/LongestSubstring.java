package String;

public class LongestSubstring {
    public static void main(String[] args) {
        System.out.println(new Solution().lengthOfLongestSubstring("ckilbkd"));
    }
}
class Solution {
    public int lengthOfLongestSubstring(String s) {
        if(s.length()==0)
            return 0;
        int max=1;
        int left=0;
        int right=0;
        String subString;
        while (right!=s.length()-1){
            subString=s.substring(left,right+1);
            if(subString.indexOf(s.charAt(right+1))==-1)
            {
                right++;
            }
            else{
                left+=subString.indexOf(s.charAt(right+1))+1;
                right++;
            }
            max= Math.max((right - left + 1), max);
        }
        return max;
    }
}