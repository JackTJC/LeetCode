package back_search;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

class Solution {
    List<String> Parenthesis;

    public static void main(String[] args) {
        Solution obj = new Solution();
        obj.Parenthesis = new ArrayList<>();
        System.out.println(obj.generateParenthesis2(4));
    }

    /*
     * 剑指 Offer II 085. 生成匹配的括号 正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
     */
    public List<String> generateParenthesis(int n) {
        if (n == 1) {
            return Arrays.asList("()");
        }
        List<String> l = new ArrayList<>();
        l.add("()");
        for (int i = 1; i < n; i++) {
            Set<String> s = new HashSet<>();
            for (String curString : l) {
                for (int j = 0; j < curString.length(); j++) {
                    String sCopy = new String(curString);
                    String ns = sCopy.substring(0, j) + "()" + sCopy.substring(j, curString.length());
                    s.add(ns);
                }
            }
            l = new ArrayList<>(s);
        }
        return l;
    }

    public List<String> generateParenthesis2(int n) {
        this.generateParenthesisHelper(n, n, 0, 0, new String());
        return this.Parenthesis;
    }

    public void generateParenthesisHelper(int l, int r, int realL, int realR, String curS) {
        if (realR > realL) {
            return;
        }
        if (l == 0 && r == 0) {
            this.Parenthesis.add(curS);
        }
        String dupS = new String(curS);
        if (l > 0) {
            this.generateParenthesisHelper(l - 1, r, realL + 1, realR, dupS + "(");
        }
        if (r > 0) {
            this.generateParenthesisHelper(l, r - 1, realL, realR + 1, curS + ")");
        }
    }
}