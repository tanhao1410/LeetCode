class Solution {
    //599. 两个列表的最小索引总和
    public String[] findRestaurant(String[] list1, String[] list2) {
        Map<String,Integer> map = new HashMap();
        for(int i = 0;i < list1.length;i ++){
            map.put(list1[i],i);
        }
        int min = list1.length + list2.length;
        int count = 0;
        for(int i = 0;i < list2.length;i ++){
            if (map.containsKey(list2[i])){
                if (min > i + map.get(list2[i])){
                    min = i + map.get(list2[i]);
                    count = 1;
                }else if (min == i + map.get(list2[i])){
                    count ++;
                }
            }
        }
        String[] res = new String[count];
        int index = 0;
        for(int i = 0;i < list2.length;i ++){
            if (map.containsKey(list2[i]) && i + map.get(list2[i]) == min){
                res[index++] = list2[i];
            }
        }
        return res;
    }
}