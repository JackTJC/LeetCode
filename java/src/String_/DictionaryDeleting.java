package String_;
//unfinished
import java.util.ArrayList;
import java.util.List;

public class DictionaryDeleting {
    public static void main(String[] args) {
        List<String> test=new ArrayList<>();
        String s="wordgoodgoodgoodbestword";
        test.add("word");
        test.add("good");
        test.add("best");
        test.add("good");
        new Solution().findLongestWord(s,test);
    }
    static class Solution {
        public String findLongestWord(String s, List<String> d) {
            int maxLen=0;
            String result="";
            for(String i:d)
            {
                int pointerInS=0;
                int pointerInI=0;
                while(pointerInI<=i.length()-1&&pointerInS<=s.length()-1)
                {
                    if(s.charAt(pointerInS)==i.charAt(pointerInI))
                        pointerInI++;
                    pointerInS++;
                }
                if(pointerInI==maxLen)
                {
                    for(int j=0;j<maxLen;j++)
                    {
                        if(i.charAt(j)<result.charAt(j))
                        {
                            result=i;
                            break;
                        }
                    }
                }
                if(pointerInI>maxLen)
                {
                    maxLen=pointerInI;
                    result=i;
                }
            }
            return result;
        }
    }
}
