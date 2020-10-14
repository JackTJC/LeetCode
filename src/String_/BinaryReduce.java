package String_;
//1404

public class BinaryReduce {
    public static void main(String[] args) {
        new Solution().numSteps(
                "10");
    }
    static class Solution {
        public int numSteps(String s) {
            int count=0;
            char [] cArr=s.toCharArray();
            for (int i=s.length()-1;i>0;)
            {
                while(cArr[i]=='0')
                {
                    count++;
                    i--;
                    if(i==0)
                        break;
                }
                if(i==0)
                    break;
                while (cArr[i]=='1')
                {
                    count++;
                    if(i==0)
                    {
                        count++;
                        break;
                    }
                    i--;
                    if(cArr[i]=='0')
                    {
                        cArr[i]='1';
                        count++;
                        break;
                    }
                }
            }
            return count;
        }
    }
}
