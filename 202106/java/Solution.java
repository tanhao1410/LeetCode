import java.util.HashSet;

/**
 * @author tanhao
 * @date 2021/06/29 14:06
 */
public class Solution {

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
