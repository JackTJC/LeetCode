package Array;
//#532
import java.util.ArrayList;
import java.util.List;

public class KDiff {
    public static void main(String[] args) {
        new Solution().findPairs(new int[] {1,2,4,4,3,3,0,9,2,3},3);
    }
    static class Solution {
        public int findPairs(int[] nums, int k) {
            int sum=0;
            List<Integer> singleNums=new ArrayList<>();
            List<Integer> repeatNums=new ArrayList<>();
            for(int i:nums)
            {
                if(!singleNums.contains(i))
                    singleNums.add(i);
                else
                {
                    if(!repeatNums.contains(i))
                        repeatNums.add(i);
                }
            }
            for(int i=0;i<singleNums.size();i++)
            {
                for(int j=i+1;j<singleNums.size();j++)
                {
                    if(Math.abs(singleNums.get(i)-singleNums.get(j))==k)
                        sum++;
                }
            }
            if(k==0)
                return repeatNums.size();
            return sum;
        }
    }
}
