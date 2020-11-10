package BitOperation;
//#1318
public class MinOr {
    static class Solution {
        public int minFlips(int a, int b, int c) {
            int flipCount=0;
            int lastBitOfA;
            int lastBitOfB;
            int lastBitOfC;
            while(a!=0||b!=0||c!=0)
            {
                lastBitOfA=a%2;
                a/=2;
                lastBitOfB=b%2;
                b/=2;
                lastBitOfC=c%2;
                c/=2;
                if(lastBitOfC==0)
                {
                    if(lastBitOfA==0&&lastBitOfB==0)
                        continue;
                    else if(lastBitOfA==1&&lastBitOfB==1)
                        flipCount+=2;
                    else
                        flipCount++;
                }
                else
                {
                    if(lastBitOfA==1||lastBitOfB==1)
                        continue;
                    else
                        flipCount++;
                }
            }
            return flipCount;
        }
    }
}
