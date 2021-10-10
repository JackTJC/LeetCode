package Array;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public static void main(String[] args){
        Solution o = new Solution();
        int[] arr = new int[10];
        System.out.println(o.longestArithSeqLength(arr));
    }
    public int longestArithSeqLength(int[] nums) {
        Map<Integer, Map<Integer, Integer>> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            map.put(i, new HashMap<>());
            for (int j = i + 1; j < nums.length; j++) {
                map.get(i).put(nums[j] - nums[i], j);
            }
        }
        return 0;
    }
    public void testFunc(){
        List<Object> l = new ArrayList<>();
        l.add(1);
    }
}
