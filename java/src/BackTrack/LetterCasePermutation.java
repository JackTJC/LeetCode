package BackTrack;

import java.util.ArrayList;
import java.util.List;

public class LetterCasePermutation {
    public static void main(String[] args) {
        new Solution().letterCasePermutation("a1b2");
    }
    static class Solution {
        List<String> output=new ArrayList<>();
        public List<String> letterCasePermutation(String S) {
            backtrack("",S);
            return output;
        }
        public void backtrack(String combinations,String next)
        {
            if(next.length()==0)
                output.add(combinations);
            else
            {
                char c=next.charAt(0);
                next=next.substring(1);
                if(Character.isDigit(c))
                    backtrack(combinations+c, next);
                else
                {
                    backtrack(combinations+Character.toLowerCase(c),next);
                    backtrack(combinations+Character.toUpperCase(c), next);
                }
            }
        }
    }
}
