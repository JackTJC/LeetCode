package Tree;

import java.util.ArrayList;
import java.util.List;
import java.util.logging.Level;

//#1361
//time exceed
public class ValidateTree {
    public static void main(String[] args) {
        new Solution().validateBinaryTreeNodes(4,new int[]{3,-1,1,-1},new int[]{-1,-1,0,-1});
    }
    static class Solution {
        public boolean validateBinaryTreeNodes(int n, int[] leftChild, int[] rightChild) {
            Integer i= 0;
            List<Integer> count = new ArrayList<Integer>(1);
            count.add(0);
            transverse(0,leftChild,rightChild,count);
            return count.get(0)==n;
        }
        void transverse(int root,int[] leftChild,int [] rightChild,List<Integer> count)
        {
            count.set(0,count.get(0)+1);
            if(count.get(0)>leftChild.length)
            {
                count.set(0,leftChild.length+1);
                return;
            }
            if(leftChild[root]!=-1)
                transverse(leftChild[root],leftChild,rightChild,count);
            if(rightChild[root]!=-1)
                transverse(rightChild[root],leftChild,rightChild,count);
        }
    }
}
