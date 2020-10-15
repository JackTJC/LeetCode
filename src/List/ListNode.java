package List;

public class ListNode {
    int val;
    ListNode next;
    ListNode() {}
    ListNode(int val) { this.val = val; }
    ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    ListNode(int [] v){
        ListNode cur=this;
        for(int i:v)
        {
            cur.val=i;
            cur.next=new ListNode();
            cur=cur.next;
        }
        cur.next=null;
    }
}
