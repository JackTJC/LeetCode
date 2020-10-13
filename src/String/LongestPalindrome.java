package String;
//time limit exceeded
public class LongestPalindrome {

    public static void main(String[] args) {
        System.out.println(new Solution().longestPalindrome("dtgrtoxuybwyfskikukcqlvprfipgaygawcqnfhpxoifwgpnzjfdnhpgmsoqzlpsaxmeswlhzdxoxobxysgmpkhcylvqlzenzhzhnakctrliyyngrquiuvhpcrnccapuuwrrdufwwungaevzkvwbkcietiqsxpvaaowrteqgkvovcoqumgrlsxvouaqzwaylehybqchsgpzbkfugujrostopwhtgrnrggocprnxwsecmvofawkkpjvcchtxixjtrnngrzqpiwywmnbdnjwqpmnifdiwzpmabosrmzhgdwgcgidmubywsnshcgcrawjvfiuxzyzxsbpfhzpfkjqcpfyynlpshzqectcnltfimkukopjzzmlfkwlgbzftsddnxrjootpdhjehaafudkkffmcnimnfzzjjlggzvqozcumjyazchjkspdlmifvsciqzgcbehqvrwjkusapzzxyiwxlcwpzvdsyqcfnguoidiiekwcjdvbxjvgwgcjcmjwbizhhcgirebhsplioytrgjqwrpwdciaeizxssedsylptffwhnedriqozvfcnsmxmdxdtflwjvrvmyausnzlrgcchmyrgvazjqmvttabnhffoe"));
    }
    static class Solution {
        public String longestPalindrome(String s) {
            int windowSize=2;
            int startPos;
            String result = s.substring(0,1);
            for(;windowSize<=s.length();windowSize++)
            {
                for(startPos=0;startPos<s.length()-windowSize+1;startPos++)
                {
                    String subString=s.substring(startPos,startPos+windowSize);
                    if(new StringBuffer(subString).reverse().toString().equals(subString))
                    {
                        result=subString;
                        break;
                    }
                }
            }
            return result;
        }
    }
}

