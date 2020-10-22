package Greedy;
//870
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class AdvantageShuffle {
    public static void main(String[] args) {
        new Solution().advantageCount(new int[]{2,0,4,1,2},new int[]{1,3,0,0,2});
    }
    static class Solution {
        public int[] advantageCount(int[] A, int[] B) {
            Arrays.sort(A);//quick sort O(log(n))
            List<Integer> aList= new ArrayList<>();
            int [] result = new int[A.length];
            int index=0;
            for(int i:A)
                aList.add(i);
            for(int i:B)
            {
                //binary search
                int begin=0,end=aList.size()-1;
                while(end>=begin)
                {
                    int mid=(begin+end)/2;
                    if(i>=aList.get(mid))
                        begin=mid+1;
                    else if(i<aList.get(mid))
                        end=mid-1;
                    else
                        break;
                }
                result[index]=aList.get((begin==aList.size())?0:begin);
                aList.remove((begin==aList.size())?0:begin);
                index++;
            }
            return result;
        }
    }
}
