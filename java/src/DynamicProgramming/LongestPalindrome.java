package DynamicProgramming;
//#5 dp algorithm
public class LongestPalindrome {

    public static void main(String[] args) {
        System.out.println(new Solution().longestPalindrome("cbbd"));
    }

    static class Solution {
        public String longestPalindrome(String s) {
            boolean [][] dp=new boolean[s.length()][s.length()];
            String result="";
            /*
            [[x11,x12,x13,...x1n],
             [.  ,x22,x23,...x2n],
             [.  ,   ,x33,...,x3n],
             [.  ,   ,   ,...,...],
             [.  ,   ,   ,...,xnn]
            ]
            assign the value diagonally from left-bottom -> right-top
             */
            for(int index=0;index<s.length();index++)
                for(int i=0,j=index;j<s.length();i++,j++)
                {
                    boolean isEqual=s.charAt(i)==s.charAt(j);
                    if(i==j) dp[i][j]=isEqual;
                    else if(j==i+1)dp[i][j]=isEqual;
                    else dp[i][j]=dp[i+1][j-1]&&isEqual;
                    if(dp[i][j]&&(j-i+1>result.length()))
                        result=s.substring(i,j+1);
                }
            return result;
        }
    }
}

