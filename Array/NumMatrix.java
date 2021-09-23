package array;

public class NumMatrix {
    public static void main(String[] args) {
        int[][] arr = new int[][] { { 1, 1, 1, 1, 1 }, { 1, 1, 1, 1, 1 }, };
        NumMatrix obj = new NumMatrix(arr);
        System.out.println(obj.sumRegion(1, 1, 1, 1));
    }

    /*
     * 给定一个二维矩阵 matrix，以下类型的多个请求： 计算其子矩形范围内元素的总和，该子矩阵的左上角为 (row1, col1) ，右下角为
     * (row2, col2) 。 实现 NumMatrix 类： NumMatrix(int[][] matrix) 给定整数矩阵 matrix 进行初始化
     * int sumRegion(int row1, int col1, int row2, int col2) 返回左上角
     * (row1, col1) 、右下角 (row2, col2) 的子矩阵的元素总和
     */
    int[][] prefixSum;
    int m, n;

    public NumMatrix(int[][] matrix) {
        int m = matrix.length;
        int n = matrix[0].length;
        this.m = m;
        this.n = n;
        this.prefixSum = new int[m][n];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (i == 0 && j == 0) {
                    prefixSum[i][j] = matrix[i][j];
                    continue;
                }
                if (i == 0) {
                    this.prefixSum[i][j] = matrix[i][j] + this.prefixSum[i][j - 1];
                    continue;
                }
                if (j == 0) {
                    this.prefixSum[i][j] = matrix[i][j] + this.prefixSum[i - 1][j];
                    continue;
                }
                this.prefixSum[i][j] = this.prefixSum[i][j - 1] + this.prefixSum[i - 1][j]
                        - this.prefixSum[i - 1][j - 1] + matrix[i][j];
            }
        }
    }

    public int sumRegion(int row1, int col1, int row2, int col2) {
        int big, small, left, top;
        big = this.prefixSum[row2][col2];
        if (row1 > 0 && col1 > 0) {
            small = this.prefixSum[row1 - 1][col1 - 1];
        } else {
            small = 0;
        }
        if (row1 == 0) {
            top = 0;
        } else {
            top = this.prefixSum[row1 - 1][col2];
        }
        if (col1 == 0) {
            left = 0;
        } else {
            left = this.prefixSum[row2][col1 - 1];
        }
        return big - top - left + small;
    }
}
