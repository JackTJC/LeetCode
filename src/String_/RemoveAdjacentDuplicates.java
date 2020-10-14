package String_;
//#1209
public class RemoveAdjacentDuplicates {
    public static void main(String[] args) {
        System.out.println(new Solution().removeDuplicates("yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy",4));
    }
    static class Solution {
        public String removeDuplicates(String s, int k) {
            for(int i=0;i<=s.length()-k;i++)
            {
                String subString = s.substring(i,i+k);
                if(Solution.isDuplicate(subString))
                {
                    return new Solution().removeDuplicates(s.substring(0,i)+s.substring(i+k),k);
                }
            }
            return s;
        }
        static boolean isDuplicate(String s)
        {
            char c=s.charAt(0);
            for(char c_:s.toCharArray())
            {
                if(c!=c_)
                    return false;
            }
            return true;
        }
    }
}
