class Solution {
    //213. 打家劫舍 II
    public int rob(int[] nums) {
        if (nums.length < 2) return nums[0];
        // 思路：最后一个房间的偷了，则第一个房间不能偷。分两种情况。
        //int[] dp = new int[nums.length];//可以偷最后一个房间的最大
        int pre2 = 0;//第一个房间不能偷
        int pre = nums[1];
        int cur = pre;
        int res = Math.max(cur,pre2);
        for(int i = 2;i < nums.length;i ++){
            cur = Math.max(pre,nums[i] + pre2);
            pre2 = pre;
            pre = cur;
            res = Math.max(res,cur);
        }
        //不偷最后一个房间的最大值
        pre2 = nums[0];
        pre = Math.max(nums[0],nums[1]);
        cur = pre;
        res = Math.max(res,cur);
        for(int i = 2;i < nums.length - 1;i ++){
            cur = Math.max(pre,nums[i] + pre2);
            pre2 = pre;
            pre = cur;
            res = Math.max(cur,res);
        }
        return res;
    }
}