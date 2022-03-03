class Solution {
    //973. 最接近原点的 K 个点
    public int[][] kClosest(int[][] points, int k) {
        //排序
        Arrays.sort(points,(e1,e2)->{
            int e1Dis = e1[0] * e1[0] + e1[1] * e1[1];
            int e2Dis = e2[0] * e2[0] + e2[1] * e2[1];
            return e1Dis - e2Dis;
        });
        int[][] res = new int[k][2];
        for(int i = 0;i < k;i ++){
            res[i] = points[i];
        }
        return res;
    }
    //122. 买卖股票的最佳时机 II
    public int maxProfit(int[] prices) {
        //思路：每天可以得到选择，买，卖，不动。
        // 如果还没买，明天价格上升， 今天买。价格下降或不变，不买。
        //如果已经买了，明天价格上升， 不卖，下跌，卖掉。
        //最后一天，只能卖了。
        //如果采用动态规划呢？dp1[i] 股票在手中时的最大价值，dp2[i] 股票不在手中时的
        int dp = -prices[0];
        int dp2 = 0;
        for(int i = 1;i < prices.length;i ++){
            //是否购买股票
            int tempDp =Math.max( dp,dp2 - prices[i]);
            //是否卖出
            dp2 = Math.max(dp2,dp + prices[i]);
            dp = tempDp;
        }
        return dp2;
    }
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