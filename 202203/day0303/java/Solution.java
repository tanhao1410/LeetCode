class Solution {
    //54. 螺旋矩阵
    public List<Integer> spiralOrder(int[][] matrix) {
        List<Integer> res = new ArrayList();
        int count = 0;
        int m = matrix.length;
        int n = matrix[0].length;
        int x = 0;
        int y = 0;
        int cycle = 0;
        while(true){
            //往右走
            while(y < n - cycle){
                res.add(matrix[x][y++]);
                count ++;
                if(count == m * n) return res;
            }
            //往下走，此时y已经越界，注意
            y --;
            x ++;
            while(x < m - cycle){
                res.add(matrix[x++][y]);
                count ++;
                if(count == m * n) return res;
            }
            //往左走
            x --;
            y --;
            while (y >= cycle){
                res.add(matrix[x][y --]);
                count ++;
                if(count == m * n) return res;
            }
            //往上
            y++;
            x --;
            while(x >= cycle + 1){
                res.add(matrix[x--][y]);
                count ++;
                if(count == m * n) return res;
            }
            //第二圈的开始位置
            x++;
            y++;
            cycle ++;
        }
    }
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