class Solution {

    //39. 组合总和
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        Arrays.sort(candidates);
        return combinationSum(candidates,target,0);
    }

    public List<List<Integer>> combinationSum(int[] candidates, int target,int pre) {
        //思路:递归思想，每次用一个，然后target-cad
        //去重，可以按递增的顺序去增加数字
        List<List<Integer>> res = new ArrayList();
        //选择第一个数
        for(int num : candidates){
            if (num < pre){
                continue;
            }
            if (target > num){
                List<List<Integer>> innerRes = combinationSum(candidates,target - num,num);
                for(List<Integer> inner : innerRes){
                    inner.add(num);
                }
                res.addAll(innerRes);
            }else if (target == num){
                List<Integer> item = new ArrayList();
                item.add(num);
                res.add(item);
            }else{
                break;
            }
        }
        return res;
    }
}