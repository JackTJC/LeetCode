package String_;
//#165
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class CompareVersionNumbers {
    static class Solution {
        public int compareVersion(String version1, String version2) {
            version1="."+version1;
            version2="."+version2;
            int sum1=0,sum2=0;
            int base=1;
            List<Integer> v1List=new ArrayList<>();
            List<Integer> v2List=new ArrayList<>();
            getVersionVal(version1,v1List);
            getVersionVal(version2,v2List);
            for(;v1List.size()<v2List.size();v1List.add(0));
            for(;v2List.size()<v1List.size();v2List.add(0));
            for(int i=0;i<v1List.size();i++)
            {
                if(v1List.get(i)>v2List.get(i))return 1;
                else if (v1List.get(i)<v2List.get(i))return -1;
            }
            return 0;
        }
        public void getVersionVal(String version,List<Integer> list)
        {
            int sum=0;
            int base=1;
            for(int i=version.length()-1;i>=0;i--)
            {
                if(version.charAt(i)=='.')
                {
                    list.add(sum);
                    sum=0;
                    base=1;
                }
                else
                {
                    sum+=Integer.parseInt(version.substring(i,i+1))*base;
                    base*=10;
                }
            }
            Collections.reverse(list);
        }
    }
}
