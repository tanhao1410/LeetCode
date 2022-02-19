class Solution {
    //969. 煎饼排序
    public List<Integer> pancakeSort(int[] arr) {
        //思路：每一次通过两次翻转，把最大的数放入到合适的位置
        List<Integer> res = new ArrayList();
        for (int i = arr.length;i >= 1;i --){
            if (arr[i - 1] != i){
                //需要翻转两次，找到i位置，
                for(int j = 0;j < i;j ++){
                    if (arr[j] == i){
                        if(j != 0){
                            //找到了
                            res.add(j + 1);
                            //翻转数组
                            reverseArr(arr,0,j);
                        }
                        res.add(i);
                        reverseArr(arr,0,i - 1);
                        break;
                    }
                }
            }
        }
        return res;
    }

    private void reverseArr(int[] arr,int start,int end){
        while (start < end){
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] =temp;
            start ++;
            end --;
        }
    }
    //55. 跳跃游戏
    public boolean canJump(int[] nums) {
        //if (nums.length == 1) return true;
        //i能走到的位置有
        int maxDistance = nums[0];
        for(int j = 1;j <= maxDistance && j < nums.length ;j ++){
            maxDistance = Math.max(maxDistance,nums[j] + j);
            if (maxDistance >= nums.length - 1) return true;
        }
        return maxDistance >= nums.length - 1;
    }

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