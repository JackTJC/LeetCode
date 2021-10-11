package list_op;

public class Solution {
    public static void main(String[] args){
        ListNode n1,n2,n3;
        n1=new ListNode(1);
        n2=new ListNode(2);
        n3=new ListNode(3);
        n1.next=n2;
        n2.next=n3;
        n3.next=null;
        Solution o = new Solution();
        o.reverseList(n1);
        ListNode p  = n1;
        while(p!=null){
            System.out.println(p.val);
            System.out.println("->");
            p=p.next;
        }
    }
    public ListNode reverseList(ListNode head) {
        if (head==null||head.next==null){
            return head;
        }
        ListNode ret = this.reverseList(head.next);
        head.next.next=head;
        head.next=null;
        return ret;
    }
}

