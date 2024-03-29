class Solution {
    //785. 判断二分图
    public boolean isBipartite(int[][] graph) {
        //用两个集合来存储
        Set<Integer> set1 = new HashSet();
        Set<Integer> set2 = new HashSet();
        boolean[] used = new boolean[graph.length];
        Stack<Integer> stack1 = new Stack();
        Stack<Integer> stack2 = new Stack();
        for(int i = 0;i < used.length;i ++){
            if (!used[i]){
                //放入到set1中
                set1.add(i);
                stack1.push(i);
                while(stack1.size() > 0 || stack2.size() > 0){
                    while(stack1.size() > 0){
                        int cur = stack1.pop();
                        for(int next : graph[cur]){
                            if (!set2.contains(next)){
                                stack2.push(next);
                                set2.add(next);
                                used[next] = true;
                            }
                        }
                    }

                    while (stack2.size() > 0){
                        int cur = stack2.pop();
                        for(int next : graph[cur]){
                            if (!set1.contains(next)){
                                stack1.push(next);
                                set1.add(next);
                                used[next] = true;
                            }
                        }
                    }
                }
            }
        }
        return set1.size() + set2.size() == graph.length;
    }
    //763. 划分字母区间
    public List<Integer> partitionLabels(String s) {
        List<Integer> res = new ArrayList();
        //思路：后面没有出现前面的字母时，就从此处断开。
        //记录每一个字母最后出现的位置。
        int[] lastIndexs = new int[26];
        for(int i = 0;i < s.length();i ++){
            lastIndexs[s.charAt(i) - 'a'] = i;
        }
        //记录前面的字母最后出现的位置
        int preIndex = -1;
        int lastIndex = 0;
        for(int i = 0;i < s.length();i ++){
            lastIndex = Math.max(lastIndex,lastIndexs[s.charAt(i) - 'a']);
            if (lastIndex == i){
                res.add(lastIndex - preIndex);
                preIndex = lastIndex;
            }
        }
        return res;
    }
    //290. 单词规律
    public boolean wordPattern(String pattern, String s) {
        HashMap<Character,String> map = new HashMap();
        //字母数与单词数也要对应
        HashSet<String> set = new HashSet();
        String[] ss = s.split(" ");
        for(String word : ss){
            set.add(word);
        }
        if (pattern.length() != ss.length) return false;
        for(int i = 0;i < ss.length;i ++){
            char cur = pattern.charAt(i);
            if(map.containsKey(cur)){
                if (!map.get(cur).equals(ss[i])){
                    return false;
                }
            }else{
                map.put(cur,ss[i]);
            }
        }
        return set.size() == map.size();
    }
    //413. 等差数列划分
    public int numberOfArithmeticSlices(int[] nums) {
        //思路：先找最长等差数列。然后，从最长等差数列中 最后一个，重新开始找。
        // 数量=1 + 2 + .. （长度 - 2）
        int res = 0;
        int start = 0;
        //从start开始后面要至少有两个数存在
        while (start + 2 < nums.length){
            int diff = nums[start + 1] - nums[start];
            int end = start + 2;
            while (end < nums.length && nums[end] - nums[end - 1] == diff) end ++;
            res += childNum(end - start);
            start = end - 1;
        }
        return res;
    }

    private int childNum(int length){
        int res = 0;
        for(int i = 1;i <= length - 2;i ++) res += i;
        return res;
    }
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