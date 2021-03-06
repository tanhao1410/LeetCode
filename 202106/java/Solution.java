import java.util.*;

/**
 * @author tanhao
 * @date 2021/06/29 14:06
 */
public class Solution {

    //395. 至少有 K 个重复字符的最长子串
    public int longestSubstring(String s, int k) {

        ArrayList<Integer> split = new ArrayList();
        int[] chars = new int[26];
        for(int i = 0;i < s.length();i ++){
            char c = s.charAt(i);
            chars[c - 'a'] ++;
        }

        for(int i = 0;i < s.length();i ++){
            char c = s.charAt(i);
            if(chars[ c - 'a'] > 0 && chars[c - 'a'] < k){
                split.add(i);
            }
        }
        if(split.size() == 0){
            return s.length();
        }
        //否则递归去解决
        //思路：先看所有的，看哪些字母不适合加入进来，分隔点便是这些字母的位置。
        int res = 0;
        if(split.get(0) != 0){
            res =  Math.max(res,longestSubstring(s.substring(0,split.get(0)),k));
        }
        for(int i = 0;i < split.size();i ++){
            String subStr = null;
            if(i + 1 == split.size()){
                subStr = s.substring(split.get(i) + 1);
            }else{
                subStr = s.substring(split.get(i) + 1,split.get(i + 1));
            }
            int subRes = longestSubstring(subStr,k);
            res = Math.max(res,subRes);
        }
        return res;
    }

    //149. 直线上最多的点数
    public int maxPoints(int[][] points) {
        //思路：总共有三百个点。包含某点的直线 共有300种，剩余的点在不在其中，最多298中情况，总的时间复杂度 o(n^3)
        //先用 一个点（x1,y1），然后以此点为开始，另找一个点(x2,y2)划线，看还有多少个点在这条线上。
        //怎么确定其他点是否在这条线上，(x3,y3) (y2-y1)*(x3-x1) == (y3-y1)*(x2-x1)
        int res = 1;
        for(int i = 0;i < points.length-1;i ++){
            int x1 = points[i][0];
            int y1 = points[i][1];
            for (int j = i + 1;j < points.length;j ++){
                int x2 = points[j][0];
                int y2 = points[j][1];
                int curRes = 2;
                for (int k = j + 1; k < points.length;k ++){
                    int x3 =  points[k][0];
                    int y3 =  points[k][1];
                    if((y2-y1)*(x3-x1) == (y3-y1)*(x2-x1) ){
                        curRes ++;
                    }
                }
                if (curRes > res){
                    res = curRes;
                }
            }
        }
        return res;
    }

    //815. 公交路线
    public int numBusesToDestination(int[][] routes, int source, int target) {
        if(source == target){
            return 0;
        }

        HashSet<Integer> sourceLine = new HashSet();
        HashSet<Integer> targetLine = new HashSet();
        //新思路，不以正常的点为计算，而是以公交路线来计算，可以走的路线，map中的key为公交路线编号
        Map<Integer,HashSet<Integer>> m = new HashMap();
        //点能到达的路线
        Map<Integer,HashSet<Integer>> point2Line = new HashMap();
        for(int i = 0;i < routes.length;i ++){
            int[] route = routes[i];
            HashSet<Integer> set = new HashSet();
            for(int j = 0;j < route.length;j ++){
                set.add(route[j]);

                HashSet<Integer> point2LineSet = point2Line.get(route[j]);
                if(point2LineSet == null){
                    point2LineSet = new HashSet<>();
                    point2Line.put(route[j],point2LineSet);
                }
                point2LineSet.add(i);
            }
            m.put(i,set);
            if(set.contains(source)){
                sourceLine.add(i);
            }
            if(set.contains(target)){
                targetLine.add(i);
            }
        }


        //记录已经坐过的路线
        HashSet<Integer> oldRoad = new HashSet();
        oldRoad.addAll(sourceLine);

        //在当前步数下可以到达的路线
        HashSet<Integer> curPoint = new HashSet();
        curPoint.addAll(sourceLine);

        int res = 1;
        //不包含目标点，及当前可以到达的点不为空，就一直循环
        while(!containsOne(oldRoad,targetLine) && curPoint.size() > 0){
            res ++;

            //本次可以走向那些路线
            HashSet<Integer> newCurPoint = new HashSet();
            for(Integer p : curPoint){
                //该条路线，能到达的点
                HashSet<Integer> canReach = m.get(p);
                if(canReach!=null){
                    //只有未走过的点加入到newCurPoint
                    for(Integer pp :canReach){
                        //从该点能够达到的路线有
                        HashSet<Integer> canReachLine = point2Line.get(pp);
                        for(Integer line : canReachLine){
                            if(!oldRoad.contains(line)){
                                newCurPoint.add(line);
                            }
                        }
                    }
                }
            }

            //更新已经走过的路线
            oldRoad.addAll(newCurPoint);

            //不能往下走了，并且目标点未到达，返回-1
            if(newCurPoint.size() == 0 && !containsOne(oldRoad,targetLine)){
                return -1;
            }
            //更新当前可在位置
            curPoint = newCurPoint;
        }

        return res;
    }

    public boolean containsOne(HashSet<Integer> set1,HashSet<Integer> set2){
        for(Integer i : set2){
            if(set1.contains(i)){
                return true;
            }
        }
        return false;
    }

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
