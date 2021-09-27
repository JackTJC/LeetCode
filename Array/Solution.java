package array;

import java.util.HashMap;
import java.util.Map;

class Solution {
    public int longestArithSeqLength(int[] nums) {
        Map<Integer, Map<Integer, Integer>> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            map.put(i, new HashMap<>());
            for (int j = i + 1; j < nums.length; j++) {
                map.get(i).put(nums[j] - nums[i], j);
            }
        }
        return 0
    }
}