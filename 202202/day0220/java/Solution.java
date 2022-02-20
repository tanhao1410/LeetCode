class Solution {
    //847. 访问所有节点的最短路径
    public int shortestPathLength(int[][] graph) {
        //用java怎么表示被访问过呢，用一个二维数组来表示,第一维表示编号，第二维表示整个图的访问情况（即一个int(元素最多12个，因此最大为)）
        int n = graph.length;
        boolean[][] used = new boolean[n][1<<(n + 1)];
        LinkedList<int[]> queue = new LinkedList();
        for(int i = 0;i < n;i ++){
            queue.add(new int[]{i,1<<i,0});
            used[i][1<<i] = true;
        }
        while(true){
            int size = queue.size();
            for(int i = 0;i < size;i ++){
                int[] cur = queue.remove(0);
                //访问了所有了
                if (countOne(cur[1]) == n){
                    return cur[2];
                }
                // 能访问到的数
                for(int next : graph[cur[0]]){
                    if(!used[next][cur[1] | (1 << next) ]){
                        used[next][cur[1] | (1 <<next)] = true;
                        queue.add(new int[]{next,cur[1] | (1 << next),cur[2] + 1});
                    }
                }
            }
        }
    }

    private int countOne(int num){
        int res = 0;
        while(num != 0){
            res += num & 1;
            num >>= 1;
        }
        return res;
    }
    //409. 最长回文串
    public int longestPalindrome(String s) {
        int res = 0;
        int[] table = new int[128];
        for(int i = 0;i < s.length();i ++){
            table[s.charAt(i)] ++;
        }
        int flag = 0;
        for(int t :table){
            if (t % 2 == 0) {
                res += t;
            }else{
                res += t - 1;
                flag = 1;
            }
        }
        return res + flag;
    }
    //127. 单词接龙
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Map<String,Boolean> map = new HashMap();
        for(String word:wordList) map.put(word,false);
        LinkedList<String> queue = new LinkedList();
        queue.add(beginWord);
        map.put(beginWord,true);
        int step = 1;
        while (queue.size() > 0){
            if(map.getOrDefault(endWord,false)) return step;
            int size = queue.size();
            for(int i = 0;i < size;i ++){
                String cur = queue.remove(0);
                for(String word:wordList){
                    if (!map.get(word) && isOneDiff(word,cur) ){
                        queue.add(word);
                        map.put(word,true);
                    }
                }
            }
            step ++;
        }
        return 0;
    }

    private boolean isOneDiff(String s1,String s2){
        int res = 0;
        for (int i = 0;i < s1.length();i ++){
            if (s1.charAt(i) != s2.charAt(i)) res ++;
        }
        return res == 1;
    }
    //415. 字符串相加
    public String addStrings(String num1, String num2) {
        //进位标志
        int flag = 0;
        StringBuilder res = new StringBuilder();
        //注意需要从低位开始加
        int index1 = num1.length() - 1;
        int index2 = num2.length() - 1;
        while(index1 >= 0 && index2 >= 0){
            int bitRes = num1.charAt(index1) - '0' + num2.charAt(index2) - '0' + flag;
            flag = bitRes/10;
            bitRes %= 10;
            res.append(bitRes);
            index2 --;
            index1 --;
        }
        while (index1 >= 0){
            int bitRes = num1.charAt(index1) - '0' + flag;
            flag = bitRes/10;
            bitRes %= 10;
            res.append(bitRes);
            index1 --;
        }
        while (index2 >= 0){
            int bitRes = num2.charAt(index2) - '0' + flag;
            flag = bitRes/10;
            bitRes %= 10;
            res.append(bitRes);
            index2--;
        }
        if (flag == 1) res.append(1);
        return res.reverse().toString();
    }
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