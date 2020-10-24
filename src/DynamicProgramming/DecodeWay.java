package DynamicProgramming;
//#91
public class DecodeWay {
    public static void main(String[] args) {

        System.out.println(Integer.parseInt("12121"));
    }
    static class Solution {
        public int numDecodings(String s) {
            if(s.charAt(0)=='0')return 0;
            int [] nums = new int [s.length()+2];
            nums[0]=0;
            nums[1]=1;
            nums[2]=1;
            for(int i=1;i<s.length();i++)
            {
                boolean flag=s.charAt(i-1)=='0'||Integer.parseInt(s.substring(i-1,i+1))>26;
                if(s.charAt(i)!='0')
                {
                    if(flag)
                        nums[i+2]=nums[i+2-1];
                    else
                        nums[i+2]=nums[i+2-1]+nums[i+2-2];
                }
                else
                {
                    if(flag)
                        return 0;
                    else
                        nums[i+2]=nums[i+2-2];
                }
            }
            return nums[nums.length-1];
        }
    }
}
