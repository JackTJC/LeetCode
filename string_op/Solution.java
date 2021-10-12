
public class Solution {
    public static void main(String[] args) {
        Solution obj = new Solution();
        System.out.println(obj.strToInt("    -100000"));
    }

    public boolean checkInclusion(String s1, String s2) {
        int[] primes = new int[26];
        primes[0] = 2;
        for (int i = 1; i < primes.length; i++) {
            int start = primes[i - 1] + 1;
            while (true) {
                if (Solution.isPrime(start)) {
                    primes[i] = start;
                    break;
                }
                start++;
            }
        }
        int s1Prd = 1;
        for (int i = 0; i < s1.length(); i++) {
            s1Prd *= primes[s1.charAt(i) - 'a'];
        }
        for (int i = 0; i < s2.length(); i++) {
            int prd = 1;
            for (int j = 0; j < s1.length() && j < s2.length(); j++) {
                prd *= primes[s2.charAt(i + j) - 'a'];
            }
            if (prd == s1Prd) {
                return true;
            }
        }
        return false;
    }

    public static boolean isPrime(int n) {
        for (int i = 2; i < n; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }

    public int strToInt(String str) {
        int start = 0;
        // find start
        char c = '0';
        while (start < str.length()) {
            c = (char) str.indexOf(start);
            boolean isPlusOrMinus = c == '+' || c == '-';
            if (!Character.isDigit(c) && !isPlusOrMinus && c != ' ') {
                return 0;
            } else if (isPlusOrMinus) {
                start++;
                break;
            }
            start++;
        }
        int ret = 0;
        boolean negative = false;
        if (c == '-') {
            negative = true;
        }
        while (start < str.length() && Character.isDigit(str.charAt(start))) {
            int oldRet = ret;
            int newRet = 10 * ret + str.charAt(start) - '0';
            if (newRet < oldRet) {
                return Integer.MAX_VALUE;
            }
            ret = newRet;
        }
        if (negative) {
            return -ret;
        }
        return ret;
    }
}
