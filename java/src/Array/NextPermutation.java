package Array;
//#31
import java.util.Arrays;

public class NextPermutation {
    public static void main(String[] args) {
        new Solution().nextPermutation(new int[]{6,8,7,4,3,2});
    }
    static class Solution {
        public void nextPermutation(int[] nums) {
            int firstNotAscendingPos=0;
            int firstMorePos=0;
            for(int i=nums.length-1;i>=0;i--)
            {
                if(i==0)
                {
                    Arrays.sort(nums);
                    return;
                }
                else if(nums[i-1]<nums[i])
                {
                    firstNotAscendingPos=i-1;
                    break;
                }
            }
            for(int i=nums.length-1;i>=0;i--)
            {
                if(nums[i]>nums[firstNotAscendingPos]) {
                    firstMorePos = i;
                    break;
                }
            }
            //exchange
            int t=nums[firstNotAscendingPos];
            nums[firstNotAscendingPos]=nums[firstMorePos];
            nums[firstMorePos]=t;
            //reverse
            reverse(nums,firstNotAscendingPos);
        }
        public void reverse(int [] nums,int pos)
        {
            for(int begin=pos+1,end=nums.length-1;begin<end;end--,begin++)
            {
                int t=nums[begin];
                nums[begin]=nums[end];
                nums[end]=t;
            }
        }
    }
}
