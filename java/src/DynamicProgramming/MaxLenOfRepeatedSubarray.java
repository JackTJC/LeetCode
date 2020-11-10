package DynamicProgramming;
//#718
public class MaxLenOfRepeatedSubarray {
    public static void main(String[] args) {
        new Solution().findLength(new int[]{0,0,0,0,1},new int[]{1,0,0,0,0});
    }
    static class Solution {
        public int findLength(int[] A, int[] B) {
            int [][]dp=new int[A.length][B.length];
            int max=0;
            for(int i=A.length-1;i>=0;i--)
                for(int j=B.length-1;j>=0;j--)
                {
                    if(i==A.length-1)
                        dp[i][j]=(A[i]==B[j])?1:0;
                    else if(j==B.length-1)
                        dp[i][j]=(A[i]==B[j])?1:0;
                    else dp[i][j]=(A[i]==B[j])?(dp[i+1][j+1]+1):0;
                    max= Math.max(max, dp[i][j]);
                }
            return max;
        }
    }
}
