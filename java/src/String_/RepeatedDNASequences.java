package String_;
//#187
import java.util.ArrayList;
import java.util.List;
//time limit exceeded
public class RepeatedDNASequences {
    public static void main(String[] args) {
        new Solution().findRepeatedDnaSequences("AAAAAAAAAAA");
    }
    static class Solution {
        public List<String> findRepeatedDnaSequences(String s) {
            List<String> result=new ArrayList<>();
            List<String> waitMatch=new ArrayList<>();
            String sub="";
            for(int i=0;i<=s.length()-10;i++)
            {
                sub=s.substring(i,i+10);
                if(!waitMatch.contains(sub))
                    waitMatch.add(sub);
                else
                {
                    if(!result.contains(sub))
                        result.add(sub);
                    waitMatch.remove(sub);
                }
            }
            return result;
        }
    }
}
