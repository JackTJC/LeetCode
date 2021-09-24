package string_op;

public class Solution {
    public static void main(String[] args) {
        Solution obj = new Solution();
        System.out.println(obj.checkInclusion("ab", "eidbaooo"));
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
}