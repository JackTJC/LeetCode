//#1329
package Array;

public class SortDiagonally {
    public static void main(String[] args) {
        int [][] mat={{3,3,1,1},{2,2,1,2},{1,1,1,2}};
        mat=new Solution().diagonalSort(mat);
        System.out.println("done");
    }
    static class Solution {
        public int[][] diagonalSort(int[][] mat) {
            int n=mat[0].length;//column
            int m=mat.length;//row
            //mat[m-2][0]->mat[0,0]
            for(int i=m-2;i>=0;i--)
            {
                int len=Math.min(m-i,n);//count of every diagonal
                for(int j=1;j<len;j++){
                    for(int k=0;k<len-j;k++){
                        if(mat[i+k][k]>mat[i+k+1][k+1])
                        {
                            int temp=mat[i+k][k];
                            mat[i+k][k]=mat[i+k+1][k+1];
                            mat[i+k+1][k+1]=temp;
                        }
                    }
                }
            }
            //mat[0][1]->mat[n-1][0]
            for (int i=1;i<=n-2;i++)
            {
                int len=Math.min(m,n-i);
                for(int j=1;j<len;j++){
                    for(int k=0;k<len-j;k++){
                        if(mat[k][i+k]>mat[k+1][i+k+1])
                        {
                            int temp=mat[k][i+k];
                            mat[k][i+k]=mat[k+1][i+k+1];
                            mat[k+1][i+k+1]=temp;
                        }
                    }
                }
            }
            return mat;
        }
    }
}
