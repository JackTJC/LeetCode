package Math;

public class Solution extends SolBase {
    public static void main(String[] args) {
        Solution o = new Solution();
        System.out.println(o.mySqrt(99));
    }

    public int mySqrt(int x) {
        int l = 0, r = x / 2;
        while (l < r) {
            int mid = (l + r) / 2;
            int result = mid * mid;
            if (result > x) {
                r = mid - 1;
            } else if (result < x) {
                l = mid + 1;
            } else {
                return mid;
            }
        }
        return l;
    }

    public int rand10() {

    }

    public double angleClock(int hour, int minutes) {
        if (hour == 12) {
            hour = 0;
        }
        if (minutes == 60) {
            minutes = 0;
        }
        double pm = 6.0 * minutes;
        double ph = 30.0 * hour + (minutes / 60.0) * 30;
        double abs = Math.abs(pm - ph);
        return Math.min(abs, 360 - abs);
    }
}
