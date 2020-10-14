package Tree;
//1038
import java.util.ArrayList;
import java.util.List;
//#1038
public class BST2GST {
    public static void main(String[] args) {

    }

    static class Solution {
        static int sum;
        public TreeNode bstToGst(TreeNode root) {
            //ldr
            sum=0;
            List<Integer> valList=new ArrayList<Integer>();
            Solution.rdl(root,valList);
            Solution.ldrChange(root,valList);
            return root;
        }
        //ldr transverse and get the value list
        static void rdl(TreeNode root,List<Integer> valList)
        {
            if(root==null)return;
            if(root.right!=null)
                Solution.rdl(root.right,valList);
            valList.add(root.val);
            if(root.left!=null)
                Solution.rdl(root.left,valList);
        }
        static void ldrChange(TreeNode root,List<Integer> valList)
        {
            if(root==null)return;
            if(root.right!=null)
                Solution.ldrChange(root.right,valList);
            Solution.sum+=valList.get(0);
            valList.remove(0);
            root.val=Solution.sum;
            if(root.left!=null)
                Solution.ldrChange(root.left,valList);
        }
    }
}
