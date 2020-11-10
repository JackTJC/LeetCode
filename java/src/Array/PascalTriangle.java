package Array;

import java.util.ArrayList;
import java.util.List;

public class PascalTriangle {
    public static void main(String[] args) {
        new Solution().getRow(3);
    }
    static class Solution {
        public List<Integer> getRow(int rowIndex) {
            List<Integer> last=new ArrayList<>();
            List<Integer> current=new ArrayList<>();
            last.add(1);
            if(rowIndex==0)return last;
            for(int i=0;i<rowIndex;i++)
            {
                last.add(0,0);
                last.add(last.size(),0);
                for(int j=0;j<last.size()-1;j++)
                    current.add(last.get(j)+last.get(j+1));
                last.clear();
                last.addAll(current);
                current.clear();
            }
            return current;
        }
    }
}
