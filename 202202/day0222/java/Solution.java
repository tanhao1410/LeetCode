class Solution {
    //91. 解码方法
    public int numDecodings(String s) {
        //一个数字，有三种情况，单独作为一个编码，作为前面字母的后面编码，与后面一个合在一起做编码
        //用递归方式的话，可能超时。用动态规划的话：dp[i] 代表s[..=i] 所拥有的编码数量
        //if (s.length() == 0) return 1;
        if(s.charAt(0) == '0') return 0;
        int[] dp = new int[s.length()];
        dp[0] = 1;
        for(int i = 1;i < dp.length;i ++ ){
            //得到前一个字母、
            char pre = s.charAt(i - 1);
            char cur = s.charAt(i);
            //dp[i + 1] ,数字与前面结合，数字单独作为一个编码两种情况
            // 0 不能单独作为一个编码
            if(cur == '0' && (pre >= '3' || pre == '0')){
                return 0;
            }else if (cur == '0'){
                if (i > 1) dp[i] = dp[i - 2];
                if (i == 1) dp[i] = dp[i - 1];
            }else if (pre == '1' || (pre == '2' && cur <= '6')){
                if (i > 1){
                    dp[i] = dp[i - 1] + dp[i - 2];
                }else{
                    dp[i] = dp[i - 1] + 1;
                }
            }else{
                dp[i] = dp[i-1];
            }
        }
        return dp[dp.length - 1];
    }
}