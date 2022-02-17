class Solution {

    //剑指 Offer II 006. 排序数组中两个数字之和
    public int[] twoSum(int[] numbers, int target) {
        int[] res = new int[2];
        int i = 0;
        int j = numbers.length - 1;
        while(j > i){
            if (numbers[i] + numbers[j] == target){
                res[0] = i;
                res[1] = j;
                return res;
            }else if (numbers[i] + numbers[j] > target){
                j --;
            }else{
                i ++;
            }
        }
        return res;
    }

    //688. 骑士在棋盘上的概率
    public double knightProbability(int n, int k, int row, int column) {
        //动态规划算法呢？
        //dp[k][x][y],k代表步数。dp[0].. = 1;
        // dp[1][x][y] = 从x,y 能到达哪?
        //最终返回dp[k][row][column]
        double[][][] dp = new double[k + 1][n][n];
        //在棋盘内的一步都不走，个数是1
        for(int i = 0;i <= n-1 ;i ++){
            for (int j =  0;j <= n- 1;j ++){
                dp[0][i][j] = 1.0;
            }
        }
        int[] vector = new int[]{-2,-1,1,2};
        for(int i = 1;i <= k;i ++){
            for(int x = 0;x < n ; x ++){
                for (int y = 0; y < n ;y ++){
                    //看(x,y)能走到哪,8个方向
                    for(int u:vector){
                        for(int v:vector){
                            if( (Math.abs(u) + Math.abs(v) == 3)
                                && (x + u >= 0 && x + u < n && y + v >= 0 && y + v < n)
                            ){
                                dp[i][x][y] += dp[i - 1][x + u][y + v];
                            }
                        }
                    }
                }
            }
        }

        double gross = 1.0;
        for(int i = 0;i < k;i ++) gross *= 8.0;
        return dp[k][row][column]/gross;
    }
    //39. 组合总和
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        Arrays.sort(candidates);
        return combinationSum(candidates,target,0);
    }

    public List<List<Integer>> combinationSum(int[] candidates, int target,int pre) {
        //思路:递归思想，每次用一个，然后target-cad
        //去重，可以按递增的顺序去增加数字
        List<List<Integer>> res = new ArrayList();
        //选择第一个数
        for(int num : candidates){
            if (num < pre){
                continue;
            }
            if (target > num){
                List<List<Integer>> innerRes = combinationSum(candidates,target - num,num);
                for(List<Integer> inner : innerRes){
                    inner.add(num);
                }
                res.addAll(innerRes);
            }else if (target == num){
                List<Integer> item = new ArrayList();
                item.add(num);
                res.add(item);
            }else{
                break;
            }
        }
        return res;
    }
}