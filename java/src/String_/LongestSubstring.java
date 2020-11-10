//#3
package String_;

//sliding-window algorithm
public class LongestSubstring {
    public static void main(String[] args) {
        System.out.println(new Solution().lengthOfLongestSubstring("ckilbkd"));
    }
    static class Solution {
        public int lengthOfLongestSubstring(String s) {
            if(s.length()==0)
                return 0;
            int max=1;
            int left=0;
            int right=0;//[...,left,...,right,...],left->right is the window
            String subString;
            while (right!=s.length()-1){
                subString=s.substring(left,right+1);//get the substring according to the right and the left value
                if(subString.indexOf(s.charAt(right+1))==-1)
                {
                    //when the next character is not in current window, we increase the window by 1
                    right++;
                }
                else{
                    //or we update the left value to make the next window
                    left+=subString.indexOf(s.charAt(right+1))+1;
                    right++;
                }
                max= Math.max((right - left + 1), max);
            }
            return max;
        }
    }
}
