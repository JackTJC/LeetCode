package HashTable_;
//#954 time limit exceeded
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;


public class DoubledPairs {
    public static void main(String[] args) {
        new Solution().canReorderDoubled(new int[]{4,-2,2,-4});
    }
    static class Solution {
        public boolean canReorderDoubled(int[] A) {
            Arrays.sort(A);//O(log(n)),ascending
            List<Integer> container=new ArrayList<>();
            for(int i:A)
            {
                if((i%2==0)&&container.contains(i/2))
                    container.remove((Integer)(i/2));
                else if(container.contains(2*i))
                    container.remove((Integer)(2*i));
                else
                    container.add(i);
            }
            return container.isEmpty();
        }
    }
}
