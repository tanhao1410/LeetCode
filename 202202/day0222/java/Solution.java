class Solution {
    //70. 爬楼梯
    public int climbStairs(int n) {
        int pre = 1;
        int prePre = 0;
        for(int i = 1;i <= n;i ++){
            pre += prePre;
            prePre = pre - prePre;
        }
        return pre;
    }
    //509. 斐波那契数
    public int fib(int n) {
        if (n < 2) return n;
        int[] dp = new int[n+1];
        dp[1] = 1;
        for(int i = 2;i <= n;i ++) dp[i] = dp[i - 1] + dp[i - 2];
        return dp[n];
    }
    //49. 字母异位词分组
    public List<List<String>> groupAnagrams(String[] strs) {
        //思路：String =》 求它的按字母序版本，insert一下即可。
        HashMap<String,List<String>> map = new HashMap();
        for(String s:strs){
            String newS = getLetterSeqStr(s);
            if (map.containsKey(newS)){
                map.get(newS).add(s);
            }else{
                List<String> item = new ArrayList();
                item.add(s);
                map.put(newS,item);
            }
        }
        List<List<String>> res = new ArrayList();
        for(Map.Entry<String,List<String>> entry : map.entrySet()) res.add(entry.getValue());
        return res;
    }

    private String getLetterSeqStr(String s){
        int[] chars = new int[26];
        for(int i = 0;i < s.length();i ++) chars[s.charAt(i) - 'a'] ++;
        StringBuilder sb = new StringBuilder();
        for(int i = 0;i < 26;i ++){
            for(int j = 0;j < chars[i];j ++) sb.append((char)(i + 'a'));
        }
        return sb.toString();
    }
    //139. 单词拆分
    public boolean wordBreak(String s, List<String> wordDict) {
        //dp[i] 是否可以拆分出来。s[..=i] = true;
        // if s[..i] { 则dp[0..j]  = s[i+1..j] 是否在set中，}
        Set<String> set = new HashSet();
        for(String word:wordDict) set.add(word);
        boolean[] dp = new boolean[s.length()];
        for(int i = 0; i < s.length();i ++){
            //直接包含的话。
            if(set.contains(s.substring(0,i + 1))){
                dp[i] = true;
            }else{
                for(int j = 0;j < i;j ++){
                    //遍历dp，从dp为true的地方截断看是否包含
                    if (dp[j] && set.contains(s.substring(j + 1,i + 1))){
                        dp[i] = true;
                        break;
                    }
                }
            }
        }
        return dp[dp.length - 1];
    }
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