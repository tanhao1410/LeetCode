class Solution {
    //343. 整数拆分
    public int integerBreak(int n) {
        //每一个数都可以拆分成2 3 4 。。。
        // dp[i] 代表i 能拆分出来后，所形成的的最大数
        int[] dp = new int[n + 1];
        if(n < 4){
            return n - 1;
        }
        dp[2] = 1;
        dp[3] = 2;
        for(int i = 4;i <= n;i ++){
            int max = i;
            for(int j = 2;j < i ;j ++){
                max = Math.max(max,j * Math.max(dp[i - j],i - j));
            }
            dp[i] = max;
        }
        return dp[n];
    }
    //322. 零钱兑换
    public int coinChange(int[] coins, int amount) {
        // dp[i][j] 前i种硬币，能凑出j元的最少硬币数
        int[][] dp = new int[coins.length][amount + 1];
        for(int i = 0;i < dp.length;i ++) Arrays.fill(dp[i],-1);
        for(int i = 0;i < dp.length;i ++){
            int money = coins[i];
            for(int j = 0;j <= amount;j ++){
                //钱是0元时
                if (j == 0){
                    dp[i][j] = 0;
                    continue;
                }
                //使用一个当前硬币
                if (j >= money && dp[i][j - money] != -1){
                    dp[i][j] = dp[i][j - money] + 1;
                }

                //不用当前硬币的情况
                if (i > 0 && dp[i - 1][j] != -1){
                    if (dp[i][j] != -1){
                        dp[i][j] = Math.min(dp[i][j],dp[i-1][j]);
                    }else{
                        dp[i][j] = dp[i-1][j];
                    }
                }
            }
        }
        return dp[coins.length - 1][amount];
    }
    //537. 复数乘法
    public String complexNumberMultiply(String num1, String num2) {
        StringBuilder res = new StringBuilder();
        //先算整数
        String[] num1s = num1.substring(0,num1.length() - 1).split("\\+");
        String[] num2s = num2.substring(0,num2.length() - 1).split("\\+");
        int numPart = Integer.parseInt(num1s[0]) * Integer.parseInt(num2s[0]);
        int iMulti = Integer.parseInt(num1s[1]) * Integer.parseInt(num2s[0]) + Integer.parseInt(num2s[1])* Integer.parseInt(num1s[0]);
        int numPart2 = Integer.parseInt(num1s[1]) * Integer.parseInt(num2s[1]) * -1;
        res.append(numPart + numPart2);
        res.append("+");
        res.append(iMulti);
        res.append("i");
        return res.toString();
    }
}