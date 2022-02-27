class Solution {
    //720. 词典中最长的单词
    public String longestWord(String[] words) {
        //先找一个字母的，
        ArrayList[] countWords = new ArrayList[31];
        for(int i = 0;i < 31;i ++) countWords[i] = new ArrayList();
        for(String word:words){
            countWords[word.length() - 1].add(word);
        }
        for(int i = 1;i < 31;i ++){
            ArrayList<String> strs = countWords[i];
            //过滤掉一些
            ArrayList<String> pre = countWords[i - 1];
            var iter = strs.iterator();
            while(iter.hasNext()){
                boolean isPre = false;
                String cur = iter.next();
                for(String p :pre){
                    if(cur.startsWith(p)){
                        isPre = true;
                        break;
                    }
                }
                if(!isPre) iter.remove();
            }
            if(strs.size() == 0){
                if(pre.size() == 0) return "";
                Collections.sort(pre);
                return pre.get(0);
            }
        }
        return "";
    }
    //72. 编辑距离
    public int minDistance(String word1, String word2) {
        //如果两个单词中有一个为长度0，则直接返回另一个长度即可。
        if(word1.length() == 0) return word2.length();
        if(word2.length() == 0) return word1.length();
        //采用dp[i][j] 即word1[0..i][0..j]=>word转换需要多少步
        int[][] dp = new int[word1.length() + 1][word2.length() + 1];
        for(int[] arr:dp) Arrays.fill(arr,1000);
        for(int i = 0;i < word1.length() + 1;i ++){
            dp[i][0] = i;
            for(int j = 1;j < word2.length() + 1;j ++){
                if(i == 0){
                    dp[i][j] = j;
                }else{
                    //通过增加字母，通过删除字母，通过替换字母
                    //dp[i][j]与  dp[i][j-1]相比，后面多了一个字母，
                    //dp[i - 1][j]
                    dp[i][j] = Math.min(dp[i][j-1] + 1,dp[i][j]);
                    dp[i][j] = Math.min(dp[i-1][j] + 1,dp[i][j]);
                    //改为替换而来
                    dp[i][j] =Math.min( dp[i - 1][j - 1] + 1,dp[i][j]);
                    //新增的这个字母，和原来的最后一个相等
                    if(word2.charAt(j - 1) == word1.charAt(i - 1)){
                        //改为替换而来
                        dp[i][j] =Math.min( dp[i - 1][j - 1],dp[i][j]);
                    }
                }
            }
        }
        return dp[word1.length()][word2.length()];
    }
    //202. 快乐数
    public boolean isHappy(int n) {
        HashSet<Integer> set = new HashSet();
        set.add(n);
        while(n > 0){
            n = getNum(n);
            if(set.contains(n)){
                break;
            }
            set.add(n);
        }
        return set.contains(1);
    }

    private int getNum(int n){
        int res = 0;
        while (n > 0){
            res += (n % 10) * (n % 10);
            n /= 10;
        }
        return res;
    }
}