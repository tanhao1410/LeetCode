class Solution {

    //78. 子集
    public List<List<Integer>> subsets(int[] nums) {
        // 思路：递归思路，对于每一个元素都有两种可能性，加、不加
        return subsets(null,nums,0);
    }

    private List<List<Integer>> subsets(List<List<Integer>> pre,int[] nums,int index){
        if (index == nums.length) return pre;
        List<List<Integer>> newPre = new ArrayList();
        if (pre == null) {
            newPre.add(new ArrayList());
            List<Integer> item = new ArrayList();
            item.add(nums[index]);
            newPre.add(item);
        }else{
            //对于pre中的每一个集合，都存在两种可能性，加入该元素，不加入该元素
            for(List<Integer> item: pre){
                List<Integer> newItem = new ArrayList(item);
                item.add(nums[index]);
                newPre.add(item);
                newPre.add(newItem);
            }
        }
        return subsets(newPre,nums,index + 1);
    }
}