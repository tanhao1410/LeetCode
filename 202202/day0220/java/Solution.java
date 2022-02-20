class Solution {
    //62. 不同路径
    public int uniquePaths(int m, int n) {
        int[][] dp = new int[m][n];
        for(int i = 0;i < m;i ++){
            for(int j = 0;j <n ;j ++){
                if (i > 0 && j > 0){
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                }else if (i > 0){
                    dp[i][j] = dp[i-1][j];
                }else if (j > 0){
                    dp[i][j] = dp[i][j-1];
                }else{
                    dp[i][j] = 1;
                }
            }
        }
        return dp[m - 1][n - 1];
    }
    //45. 跳跃游戏 II
    public int jump(int[] nums) {
        //动态规划，广度优先策略
        int[] dp = new int[nums.length];
        for(int i = 0;i < dp.length - 1;i ++){
            for(int j = 1;j <= nums[i];j ++){
                if (i + j == nums.length - 1) return dp[i] + 1;
                if (dp[i + j] == 0){
                    dp[i +j] = dp[i] + 1;
                }else{
                    dp[i + j] = Math.min(dp[i + j],dp[i] + 1);
                }
            }
        }
        return dp[dp.length - 1];
    }
    //1557. 可以到达所有点的最少点数目
    public List<Integer> findSmallestSetOfVertices(int n, List<List<Integer>> edges) {
        //思路：查看哪些节点可以从其他地方访问到，去掉这些就是结果
        boolean[] canReach = new boolean[n];
        for(List<Integer> edge:edges){
            canReach[edge.get(1)] = true;
        }
        List<Integer> res = new ArrayList();
        for(int i = 0;i < n;i ++){
            if (!canReach[i]) res.add(i);
        }
        return res;
    }
    //997. 找到小镇的法官
    public int findJudge(int n, int[][] trust) {
        //用一个数组来表示自己信任的人的数量，用一个数组来表示信任自己的人的数量
        int[] trustOhter = new int[n];
        int[] trustSelf = new int[n];
        for(int[] nums:trust){
            trustOhter[nums[0] - 1] ++;
            trustSelf[nums[1] - 1] ++;
        }
        //查看符合条件的
        for(int i = 0;i < n;i ++){
            if (trustSelf[i] == n - 1 && trustOhter[i] == 0) return i + 1;
        }
        return -1;
    }
    //717. 1比特与2比特字符
    public boolean isOneBitCharacter(int[] bits) {
        for(int i = 0;;i ++){
            if (i >= bits.length - 1) return i == bits.length-1;
            if (bits[i] == 1) i ++;
        }
    }
}