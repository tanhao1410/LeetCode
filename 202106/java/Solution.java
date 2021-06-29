import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.TreeMap;

/**
 * @author tanhao
 * @date 2021/06/29 14:06
 */
public class Solution {

    //791. 自定义字符串排序
    public String customSortString(String order, String str) {
        //思路：每个字母对应一个数字，0-25，重新排序即可。

        Map<Integer,Character> int2Char = new TreeMap();
        Map<Character,Integer> char2Int = new HashMap();
        for(int i = 0;i < order.length();i ++){
            int2Char.put(i,order.charAt(i));
            char2Int.put(order.charAt(i),i);
        }

        //各个字母的数量
        StringBuilder tail = new StringBuilder();
        int[] strCharCount= new int[26];
        for (int i = 0;i< str.length();i ++){
            char c = str.charAt(i);
            if(char2Int.containsKey(c)){
                strCharCount[c - 'a'] ++;
            }else{
                tail.append(c);
            }
        }

        StringBuilder res = new StringBuilder();
        // for(int i = 0;i < 26;i ++){
        //     if(int2Char.containsKey(i)){
        //         for (int j = 0;j < strCharCount[int2Char.get(i) - 'a'];j ++){
        //             res.append(int2Char.get(i));
        //         }
        //     }
        // }

        int2Char.entrySet()
                .stream()
                .forEach(e->{
                    //value是字母
                    for (int j = 0;j < strCharCount[e.getValue() - 'a'];j ++){
                        res.append(e.getValue());
                    }
                });


        return res.append(tail).toString();
    }

    //792. 匹配子序列的单词数
    public int numMatchingSubseq(String s, String[] words) {
        HashSet<String> setY = new HashSet();
        HashSet<String> setN = new HashSet();
        int res = 0;
        for(String word : words){
            if(setN.contains(word)){
                break;
            }else if(setY.contains(word)){
                res ++;
            }else if(isChild(s,word)){
                res ++;
                setY.add(word);
            }else{
                setN.add(word);
            }
        }
        return res;
    }

    //效率低了。
    //怎么样更高呢
    private boolean isChild(String s,String child){
        if (child.length() > s.length()){
            return false;
        }
        int i = 0;
        int j = 0;
        while(j < child.length() && i < s.length()){
            if(s.charAt(i) == child.charAt(j)){
                i ++;
                j ++;
            }else{
                i ++;
            }
        }
        return j == child.length();
    }

    //820. 单词的压缩编码
    public int minimumLengthEncoding(String[] words) {
        //1.word 去重
        //2.一个单词是某个单词的结尾，去除
        //3.每个单词补充一个#
        //问题转化为：高效率查找那些是结尾的//2000个单词，最多6个尾部，2000 * 6 = 12000个
        HashSet<String> dict = new HashSet();
        for(String word : words){
            dict.add(word);
        }
        HashSet<String> remove = new HashSet();
        dict.forEach(e->{
            for(int i = 1;i < e.length();i ++){
                if(dict.contains(e.substring(i))){
                    remove.add(e.substring(i));
                }
            }
        });
        dict.removeAll(remove);
        return dict.size() + dict.stream()
                .map(String::length)
                .reduce(0,(i,j)->{
                    return i + j;
                });
    }
}
