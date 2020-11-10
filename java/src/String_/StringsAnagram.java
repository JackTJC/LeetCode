package String_;

public class StringsAnagram {
    public static void main(String[] args) {
        System.out.println(new Solution().minSteps("leetcode","practice"));
    }
    static class Solution {
        public int minSteps(String s, String t) {
            int [] hashMapS=new int[26];
            int [] hashMapT=new int[26];
            int sum=0;
            //construct the hash table
            for(char c:s.toCharArray())
                hashMapS[c-'a']++;
            for(char c:t.toCharArray())
                hashMapT[c-'a']++;
            //get the difference value on each position
            for(int i=0;i<26;i++)
            {
                sum+=Math.abs(hashMapS[i]-hashMapT[i]);
            }
            return sum/2;
        }
    }
}
