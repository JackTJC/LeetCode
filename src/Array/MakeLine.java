package Array;
//interview problem 0508
public class MakeLine {
    public static void main(String[] args) {
        new Solution().drawLine(24285,259040,105946,219220,1);
    }
    static class Solution {
        public int[] drawLine(int length, int w, int x1, int x2, int y) {
            int begin=0x80000000;
            int [] mask=new int[32];
            int [] ans=new int[length];
            int startPos=y*(w/32);
            for(int i=0;i<32;i++)
                mask[i]=begin>>i;
            for(int i=0;i<w/32;i++)
            {
                if(x1>32*i+31||x2<32*i)
                    ans[startPos+i]=0;
                else if(x2>=32*i+31&&x1<=32*i)
                    ans[startPos+i]=-1;
                else if(x1>=32*i&&x2<=32*i+31)
                    ans[startPos+i]= mask[x2 - 32 * i] & ~mask[x1 - 32 * i-1];
                else if(x1<32*i&&x2>32*i+31)
                    ans[startPos+i]=mask[x2-32*i];
                else
                    ans[startPos+i]=~mask[x1-32*i-1];
            }
            return ans;
        }
    }
}
