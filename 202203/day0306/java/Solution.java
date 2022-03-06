class Solution {
    //49. 字母异位词分组
    public List<List<String>> groupAnagrams(String[] strs) {
        //hashmap ,key 为字母排序后的
        Map<String,List<String>> map = new HashMap();
        for(String str: strs){
            String sortedStr =sortStr(str);
            if(map.containsKey(sortedStr)){
                List<String> item = map.get(sortedStr);
                item.add(str);
            }else{
                List<String> item = new ArrayList();
                item.add(str);
                map.put(sortedStr,item);
            }
        }
        List<List<String>> res = new ArrayList();
        for(Map.Entry<String,List<String>> entry:map.entrySet()){
            res.add(entry.getValue());
        }
        return res;
    }

    private String sortStr(String str){
        byte[] bytes = str.getBytes();
        Arrays.sort(bytes);
        return new String(bytes);
    }
}