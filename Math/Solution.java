package Math;
public class Solution {
    public static void main(String[] args){
        Solution o = new Solution();
        System.out.println(o.mySqrt(99));
    }
    public int mySqrt(int x) {
        int l=0,r=x/2;
        while(l<r){
            int mid = (l+r)/2;
            int result = mid * mid;
            if(result>x){
                r = mid -1;
            }else if (result<x){
                l = mid+1;
            }else{
                return mid;
            }
        }
        return l;
    }
}
