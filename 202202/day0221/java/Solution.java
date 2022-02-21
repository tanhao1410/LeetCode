class Solution {
    //5. 最长回文子串
    public String longestPalindrome(String s) {
        //思路：dp[i][j] s[i][j] 是否是回文。dp[i][i] = 1;
        // dp[i][j] = 0 ,dp[i + 1][j-1] + 2,
        int[][] dp = new int[s.length()][s.length()];
        for(int i = 0;i < dp.length;i ++) dp[i][i] = 1;
        String res = s.substring(0,1);
        for(int i = s.length() - 1;i >=0;i --){
            for(int j = i + 1;j < s.length();j ++){
                //看前后是否相等
                if (s.charAt(i) == s.charAt(j)){
                    if (i + 1 == j){
                        dp[i][j] = 2;
                    }else if (dp[i + 1][j - 1] > 0){
                        dp[i][j] = dp[i + 1][j - 1] + 2;
                    }
                }else{
                    dp[i][j] = 0;
                }

                if(dp[i][j] > res.length()){
                    res = s.substring(i,j + 1);
                }
            }
        }
        return res;
    }
    //838. 推多米诺
    public String pushDominoes(String dominoes) {
        byte[] bytes = dominoes.getBytes();
        int start = 0;
        while (start < bytes.length){
            boolean isR = bytes[start] == 'R';
            int end = start + 1;
            while (end < bytes.length && bytes[end] == '.') end ++;
            if (end == bytes.length){
                if (isR){
                    for(int i = start + 1;i < end;i ++) bytes[i] = 'R';
                }
                break;
            }

            if(isR){
                if(bytes[end] == 'R'){
                    for(int i = start + 1;i < end;i ++) bytes[i] = 'R';
                }else{
                    for(int i = 1;i <= (end - start - 1)/2;i ++){
                        bytes[start + i] = 'R';
                        bytes[end - i] = 'L';
                    }
                }
            }else{
                if (bytes[end] == 'L'){
                    for(int i = start;i < end;i ++) bytes[i] = 'L';
                }
            }
            start = end;
        }
        return new String(bytes);
    }
    //1615. 最大网络秩
    public int maximalNetworkRank(int n, int[][] roads) {
        //将路径图改为每一个点连接的所有点 的图
        List<Integer>[] graph = new ArrayList[n];
        Set<Integer> set = new HashSet();
        for(int i = 0;i < n;i ++) graph[i] = new ArrayList();
        for(int[] road:roads){
            int src = road[0];
            int dst = road[1];
            set.add(101 * src + dst);
            //set.add(101 * dst + src);
            graph[src].add(dst);
            graph[dst].add(src);
        }
        int res = 0;
        //两个不相连的城市也可以
        for(int i = 0;i < n;i ++){
            //有多少与i相连的城市，
            int iCount = graph[i].size();
            for(int j = i + 1;j < n;j ++){
                //有多少与j相连的城市
                int jCount = graph[j].size();
                //i，j是否相连，如果相连，结果-1
                if (set.contains(i * 101 + j) || (set.contains(j * 101 + i))){
                    jCount --;
                }
                res = Math.max(res,iCount + jCount);
           }
        }
        return res;
    }
}