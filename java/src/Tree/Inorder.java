package Tree;
//#94
import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

public class Inorder {
    static class Solution {
        public List<Integer> inorderTraversal(TreeNode root) {
            Stack<TreeNode> nodeStack=new Stack<>();
            List<Integer> result=new ArrayList<>();
            if(root==null)return result;
            nodeStack.push(root);
            boolean flag=true;
            while(!nodeStack.empty())
            {
                if(nodeStack.peek().left!=null&&flag)
                    nodeStack.push(nodeStack.peek().left);
                else
                {
                    flag=false;
                    TreeNode t=nodeStack.pop();
                    result.add(t.val);
                    if(t.right!=null)
                    {
                        nodeStack.push(t.right);
                        flag = true;
                    }
                }
            }
            return result;
        }
    }
}
