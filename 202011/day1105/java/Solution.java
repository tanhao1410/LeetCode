import java.util.*;

/**
 * @author tanhao
 * @date 2020/11/05 13:51
 */
public class Solution {
    public static void main(String[] args) {
        List wordList = new ArrayList();
        wordList.add("hot");
        wordList.add("dot");
        wordList.add("dog");
        wordList.add("lot");
        wordList.add("log");
        wordList.add("cog");
        System.out.println(new Solution().ladderLength("hit","cog",wordList));
    }

    //优化版本
    public int ladderLength2(String beginWord, String endWord, List<String> wordList) {
            HashSet<String> wordSet = new HashSet<String>(wordList);
            if (!wordSet.contains(endWord)) {
                return 0;
            }
            HashMap<Integer, HashSet<String>> dp = new HashMap<>();
            HashSet<String> one = new HashSet();
            one.add(endWord);
            wordSet.remove(endWord);
            wordSet.add(beginWord);
            dp.put(1, one);
            int count = 1;
            while (dp.get(count) != null) {
                HashSet<String> pre = dp.get(count);
                for (String s : pre) {
                    for (int j = 0; j < s.length(); j++) {
                        char[] chars = s.toCharArray();
                        for (char c = 'a'; c <= 'z'; c++) {
                            chars[j] = c;
                            String canReach = new String(chars);
                            if (wordSet.contains(canReach)) {
                                //可以到达
                                wordSet.remove(canReach);
                                if (canReach.equals(beginWord)) {
                                    return count + 1;
                                }
                                if (dp.get(count + 1) == null) {
                                    HashSet<String> nextStrings = new HashSet<>();
                                    nextStrings.add(canReach);
                                    dp.put(count + 1, nextStrings);
                                } else {
                                    dp.get(count + 1).add(canReach);
                                }
                            }
                        }
                    }
                }
                count += 1;
            }
            return 0;
        }

    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        HashSet<String> wordSet = new HashSet<String>(wordList);
        if (!wordSet.contains(endWord)){
            return 0;
        }
        HashMap<Integer,HashSet<String>> dp = new HashMap<>();
        HashSet<String> one = new HashSet();
        one.add(endWord);
        wordSet.remove(endWord);
        wordSet.add(beginWord);
        dp.put(1,one);
        int count = 1;

        while(dp.get(count) != null){
            Iterator<String> pre = dp.get(count).iterator();
            while (pre.hasNext()){
                String next1 = pre.next();
                Iterator<String> remain = wordSet.iterator();
                while (remain.hasNext()){
                    String next = remain.next();
                    if (this.oneStep(next1, next)){
                        if (next.equals(beginWord)){
                            return count + 1;
                        }
                        remain.remove();
                        if(dp.get(count + 1)==null){
                            HashSet<String> nextStrings = new HashSet<>();
                            nextStrings.add(next);
                            dp.put(count+1,nextStrings);
                        }else{
                            dp.get(count+1).add(next);
                        }
                    }
                }
            }
            count +=1;
        }
        return 0;
    }

    public boolean oneStep(String s1,String s2){
        int count = 0;
        for (int i =0;i < s1.length();i ++){
            if(s1.charAt(i) != s2.charAt(i)){
                count ++;
            }
        }
        return count == 1;
    }
}
