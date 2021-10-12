package list_op;

public class Solution {

    public static void main(String[] args) {
        ListNode n1, n2, n3;
        n1 = new ListNode(1);
        n2 = new ListNode(2);
        n3 = new ListNode(3);
        n1.next = n2;
        n2.next = n3;
        n3.next = null;
        Solution o = new Solution();
        ListNode ret = o.reverseList(n1);
        ListNode p = ret;
        while (p != null) {
            System.out.println(p.val);
            System.out.println("->");
            p = p.next;
        }
    }

    public ListNode reverseList(ListNode head) {
        if (head == null || head.next == null) {
            return head;
        }
        ListNode ret = this.reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return ret;
    }

    /*
     * url:https://leetcode-cn.com/problems/lMSNwu/ . 剑指 Offer II 025. 链表中的两数相加 给定两个
     * 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。 可以假设除了数字
     * 0 之外，这两个数字都不会以零开头。
     */
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode rl1, rl2;
        rl1 = this.reverseList(l1);
        rl2 = this.reverseList(l2);
        int cout = 0;
        ListNode p, dupHead = new ListNode();
        p = dupHead;
        while (rl1 != null || rl2 != null) {
            int a, b;
            a = (rl1 == null) ? 0 : rl1.val;
            b = (rl2 == null) ? 0 : rl2.val;
            p.next = new ListNode((a + b + cout) % 10);
            cout = (a + b + cout) / 10;
            p = p.next;
            if (rl1 != null) {
                rl1 = rl1.next;
            }
            if (rl2 != null) {
                rl2 = rl2.next;
            }
        }
        if (cout != 0) {
            p.next = new ListNode(cout);
        }
        return this.reverseList(dupHead.next);
    }

    /*
     * https://leetcode-cn.com/problems/aMhZSa/
     *  给定一个链表的 头节点 head ，请判断其是否为回文链表。
     * 如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。
     */
    public boolean isPalindrome(ListNode head) {
        ListNode fast, slow;
        fast = slow = head;
        while (fast.next != null && fast.next.next != null) {
            slow = slow.next;
            fast = fast.next.next;
        }
        ListNode after = this.reverseList(slow.next);
        while (after != null) {
            if (after.val != head.val) {
                return false;
            }
            after = after.next;
            head = head.next;
        }
        return true;
    }
}