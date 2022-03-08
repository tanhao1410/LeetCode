public class Solution{
}
//304. 二维区域和检索 - 矩阵不可变
class NumMatrix {
    private int[][] matrixSum;
    public NumMatrix(int[][] matrix) {
        this.matrixSum = new int[matrix.length][matrix[0].length];
        for(int i = 0;i < matrixSum.length;i ++){
            for(int j = 0;j < matrixSum[0].length;j ++){
                if (i == 0){
                    if (j == 0){
                        matrixSum[i][j] = matrix[i][j];
                    }else{
                        matrixSum[i][j] = matrix[i][j] + matrixSum[i][j-1];
                    }
                }else{
                    if (j == 0){
                        matrixSum[i][j] = matrix[i][j] + matrixSum[i - 1][j];
                    }else{
                        matrixSum[i][j] = matrixSum[i - 1][j] + matrixSum[i][j - 1] + matrix[i][j] - matrixSum[i - 1][j - 1];
                    }
                }
            }
        }
    }

    public int sumRegion(int row1, int col1, int row2, int col2) {
        if (row1 > 0 && col1 > 0){
            return matrixSum[row2][col2] - matrixSum[row1 - 1][col2] - matrixSum[row2][col1 - 1] + matrixSum[row1 - 1][col1 - 1];
        }else if (row1 > 0){
            return matrixSum[row2][col2] - matrixSum[row1 - 1][col2];
        }else if (col1 > 0){
            return matrixSum[row2][col2] - matrixSum[row2][col1 - 1];
        }else{
            return matrixSum[row2][col2];
        }
    }
}