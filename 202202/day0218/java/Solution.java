class Solution {
    //59. 螺旋矩阵 II
    public int[][] generateMatrix(int n) {
        int[][] res = new int[n][n];
        int num = 1;
        int rank = 0;
        int x = 0;
        int y = 0;
        while(num <= n * n){
            //往右边走
            //res[x][y] = num++;
            while(y < n - rank) res[x][y++] = num++;
            y --;
            //往下走
            while( x + 1 < n - rank) res[++x][y] = num++;
            //往左走
            while( y - 1 >= rank) res[x][--y] = num++;
            //往上走
            while (x - 1 >= 1 + rank) res[--x][y] = num ++;
            //开始新一轮循环
            rank ++;
            y ++;
        }
        return res;
    }
}