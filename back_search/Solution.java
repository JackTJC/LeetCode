package back_search;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.graalvm.compiler.lir.amd64.AMD64Move.LeaDataOp;

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

    /*
     * 153. 寻找旋转排序数组中的最小值 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。
     * 
     * 例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到： 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2] 若旋转 7
     * 次，则可以得到 [0,1,2,4,5,6,7]
     * 
     * 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2],
     * ..., a[n-2]] 。
     * 
     * 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
     */
    public int findMin(int[] nums) {
        if (nums[nums.length - 1] >= nums[0]) {
            return nums[0];
        }
        int left, right;
        left = nums[0];
        right = nums[nums.length - 1];
        int l = left;
        int r = right;
        while (l < r) {
            int mid = (l + r) / 2;
            if (nums[mid] <= left && nums[mid] <= right) {
                r = mid;
            } else if (nums[mid] >= left && nums[mid] >= right) {
                l = mid + 1;
            }
        }
        return nums[l] + 1;
    }
}