class Solution {

    //1994. 好子集的数目
    private int mod = 1000000007;
    public int numberOfGoodSubsets(int[] nums) {
        int[] oddNums = new int[]{2,3,5,7,11,13,17,19,23,29};
        //记录数字出现的次数
        int[] count = new int[31];
        for(int num : nums){
            count[num]++;
        }

        //先算仅质数存在的情况下数量
        long res = getOddCount(oddNums,count,1);

        //再算两个质数相乘得到的数出现的情况。如 6 10 14 。。。
        long count6 = getOddCount(oddNums,count,6);
        res += count6 * count[6] + count[6];
        res %= mod;

        long count10 = getOddCount(oddNums,count,10);//2 * 5 ,可以和 3 * 7 共存
        res += count10 * count[10] + count[10];
        //还可以和21 共同存在。
        long count10And21 = getOddCount(oddNums,count,10 * 21);
        res += (count10And21+1) * count[10] * count[21] ;
        res %= mod;

        long count14 = getOddCount(oddNums,count,14);// 2 * 7 可以和 3 * 5 共存
        res += (count14 + 1) * count[14];
        long count14And15 = getOddCount(oddNums,count,14 * 15);
        res += (count14And15 + 1) * count[14]*count[15];
        res %= mod;

        long count22 = getOddCount(oddNums,count,22);// 2 * 11 可以和 3 * 7 或3 * 5 共存
        res += (count22 + 1) * count[22];
        long count22And21 = getOddCount(oddNums,count,22 * 21);
        res += (count22And21 + 1) * count[22] * count[21];
        long count22And15 = getOddCount(oddNums,count,22 * 15);
        res += (count22And15 + 1) * count[22] * count[15];
        res %= mod;

        long count26 = getOddCount(oddNums,count,26);// 2 * 13
        res += (count26 + 1) * count[26];
        long count26And21 = getOddCount(oddNums,count,26 * 21);
        res += (count26And21 + 1) * count[26] * count[21];
        long count26And15 = getOddCount(oddNums,count,26 * 15);
        res += (count26And15 + 1) * count[26] * count[15];
        res %= mod;

        long count15 = getOddCount(oddNums,count,15);
        res += (count15 + 1) * count[15];
        res %= mod;

        long count21 = getOddCount(oddNums,count,21);
        res += (count21 + 1) * count[21];
        res %= mod;

        //还有个特殊存在 30
        long count30 = getOddCount(oddNums,count,30);
        res += (count30 + 1) * count[30];
        res %= mod;

        //最后算含1的情况
        for(int i = 0;i < count[1];i ++) {
            res *= 2;
            res %= mod;
        }
        return (int)res;
    }

    private long getOddCount(int[] oddNums,int[] counts ,int excludeNum){
        long res = 0l;
        for(int oddNum : oddNums){
            int n = counts[oddNum];
            if (n > 0 && excludeNum % oddNum != 0){
                res = res * (n + 1) + n;
                res %= mod;
            }
        }
        return res % mod;
    }
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