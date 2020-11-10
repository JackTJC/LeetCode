package Array;
//240
//binary search example
public class Search2D {
    public static void main(String[] args) {
        int [][] arr ={{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}};
        new Solution().searchMatrix(arr,5);
    }
    /*
    *binary search Note:
    * 1. when after a comparison, we will set the begin or end to mid-1 or mid+1
    * 2. we set the while ending condition is end >= begin
    * 3. we can judge find it whether or not in while body
     */
    static class Solution {
        public boolean searchMatrix(int[][] matrix, int target) {
            if(matrix.length==0)return false;//process empty matrix
            int rowBegin=0,rowEnd=matrix.length-1;
            if(matrix[0].length==0)return false;//process empty matrix
            //row binary search
            while(rowEnd>rowBegin)
            {
                int mid=(rowBegin+rowEnd)/2;
                if(target>matrix[mid][0])
                    rowBegin=mid+1;
                else if(matrix[mid][0]==target)
                    return true;
                else
                    rowEnd=mid-1;
            }
            for(int i=rowBegin;i>=0;i--)
            {
                //row binary search
                int colBegin=0,colEnd=matrix[0].length-1;
                while (colEnd>=colBegin)
                {
                    int mid=(colBegin+colEnd)/2;
                    if(target>matrix[i][mid])
                        colBegin=mid+1;
                    else if(target<matrix[i][mid])
                        colEnd=mid-1;
                    else
                        return true;
                }
            }
            return false;
        }
    }
}
