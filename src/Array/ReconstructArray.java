package Array;
//#1253
import java.util.ArrayList;
import java.util.List;

public class ReconstructArray {
    public static void main(String[] args) {
        new Solution().reconstructMatrix(1,4,new int[]{2,1,2,0,0,2});
    }
    static class Solution {
        public List<List<Integer>> reconstructMatrix(int upper, int lower, int[] colsum) {
            List<List<Integer>> result=new ArrayList<List<Integer>>();
            int count=0;
            List<Integer> up=new ArrayList<Integer>();
            for(int i:colsum)
            {

                //value is 2, so this position must be 1
                if(i==2)
                {
                    up.add(1);
                    count++;
                }
                else
                    up.add(0);
            }
            if(count>upper)
                return new ArrayList<List<Integer>>();
            for(int i=0;i<up.size();i++)
            {
                //use greedy algorithm, until this row's sum reach to upper value
                if(colsum[i]==1)
                {
                    if(count>=upper)
                        break;
                    up.set(i,1);
                    count++;
                }
            }
            result.add(up);
            count=0;
            List<Integer> low=new ArrayList<>();
            for(int i=0;i<up.size();i++)
            {
                low.add(colsum[i]-up.get(i));
                if(colsum[i]-up.get(i)==1)
                    count++;
            }
            if(count!=lower)
                return new ArrayList<List<Integer>>();
            result.add(low);
            return result;
        }
    }
}
