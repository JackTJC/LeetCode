from typing import List
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        # 旋转一个格子
        def helper(mat,startpos,num):
            temp=mat[startpos][startpos]
            for i in range(num):
                mat[startpos+i][startpos]=mat[startpos+i+1][startpos]
            for i in range(num):
                mat[startpos+num][i]=mat[startpos+num][i+1]
            for i in range(num):
                mat[startpos+num-i][num]=mat[startpos+num-i-1][num]
            for i in range(num):
                if i==num-1:
                    mat[startpos][num-i]=temp
                    break
                mat[startpos][num-i]=mat[startpos][num-i-1]
        for i in range(int(len(matrix)/2)):
            n=len(matrix)
            for j in range(n-2*i-1):
                helper(matrix,i,n-2*i-1)

if __name__ == '__main__':
    Solution.rotate(None,[[ 5, 1, 9,11],[ 2, 4, 8,10],[13, 3, 6, 7],[15,14,12,16]])