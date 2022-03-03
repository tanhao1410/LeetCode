class Solution {
    //48. 旋转图像
    public void rotate(int[][] matrix) {
        for(int i = 0;i < matrix.length/2;i ++){
            for(int j = 0;j < (matrix.length+1) / 2;j ++){
                int temp = matrix[i][j];
                //0,0 <- 2,0
                matrix[i][j] = matrix[matrix.length - 1 - j][i];
                matrix[matrix.length - 1 - j][i] = matrix[matrix.length - 1 - i][matrix.length - 1- j];
                matrix[matrix.length - 1- i][matrix.length - 1 - j] = matrix[j][matrix.length - 1- i];
                matrix[j][matrix.length - 1 - i] = temp;
            }
        }
    }
    //258. 各位相加
    public int addDigits(int num) {
        if(num <= 9) return num;
        int next = 0;
        while(num > 0){
            next += num % 10;
            num /= 10;
        }
        return addDigits(next);
    }
}