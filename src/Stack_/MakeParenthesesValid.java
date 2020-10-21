package Stack_;
//#921
import java.util.Stack;

public class MakeParenthesesValid {
    static class Solution {
        public int minAddToMakeValid(String S) {
            Stack<Character> s=new Stack<>();
            for(char c:S.toCharArray())
            {
                if(c=='(')
                    s.push(c);
                else
                {
                    if(s.empty()||s.peek()==')')
                        s.push(c);
                    if(s.peek()=='(')
                        s.pop();
                }
            }
            return s.size();
        }
    }
}
