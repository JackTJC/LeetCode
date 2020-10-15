package List;

import java.util.Arrays;

//817
public class Component {
    public static void main(String[] args) {
        ListNode test=new ListNode(new int[] {0,3,2,4,1});
        new Solution().numComponents(test,new int[]{0,3,2});
    }
    static class Solution {
        public int numComponents(ListNode head, int[] G) {
            ListNode cur=head;
            Arrays.sort(G);
            int count=0;
            boolean flag;
            while (cur!=null)
            {
                flag=false;
                while(cur!=null&&Arrays.binarySearch(G,cur.val)>=0)
                {
                    flag=true;//when attach this code block, set flag true
                    cur=cur.next;
                }
                count+=(flag)?1:0;//there is a component when flag is true
                if(cur== null)
                    break;
                cur=cur.next;
            }
            return count;
        }
    }
}
