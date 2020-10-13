package List.SingleLinkedList;

public class AddTwoNumbers {
}


// Definition for singly-linked list.
class ListNode {
     int val;
     ListNode next;
     ListNode() {}
     ListNode(int val) { this.val = val; }
     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 }

class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode cur1=l1;
        ListNode cur2=l2;
        int cout=0;
        ListNode cur3=new ListNode();
        ListNode result=cur3;
        while(cur1!=null||cur2!=null)
        {
            int val1,val2;
            val1=(cur1==null)?0:cur1.val;
            val2=(cur2==null)?0:cur2.val;
            cur3.val=(val1+val2+cout)%10;
            if(val1+val2+cout>=10)
                cout=1;
            else
                cout=0;
            cur1=(cur1==null)?null:cur1.next;
            cur2=(cur2==null)?null:cur2.next;
            if(cur1!=null||cur2!=null){
                cur3.next=new ListNode();
                cur3=cur3.next;
            }
        }
        if(cout==1)
            cur3.next=new ListNode(1);
        return result;
    }
}

