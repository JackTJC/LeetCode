package Graph;
//#886

import java.util.Arrays;

public class Bipartition {
    public static void main(String[] args) {

    }
    static  class Solution {
        public boolean possibleBipartition(int N, int[][] dislikes) {
            //use color method
            int [] colorOf = new int[N+1];//0 for none color
            for(int i=1;i<N+1;i++)
            {
                if(colorOf[i]==0)
                {
                    if(!dfs(i,1,colorOf,dislikes))
                        return false;
                }
            }
            return true;
        }
        public boolean dfs(int node,int color,int [] colorOf,int [][] dislikes)
        {
            colorOf[node]=color;
            for(int [] dislike :dislikes) {
                if (dislike[0] == node)
                {
                    if (colorOf[dislike[1]] == color)
                        return false;
                    if (colorOf[dislike[1]] == 0 && !dfs(dislike[1], -color, colorOf, dislikes))
                        return false;
                }
            }
            return true;
        }
    }
}
