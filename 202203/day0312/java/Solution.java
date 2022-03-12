class Solution {
    //剑指 Offer II 076. 数组中的第 k 大的数字
    public int findKthLargest(int[] nums, int k) {
        //快速排序的方案。以第一个字母为中心，将数字分为两部分，大于的，小于等于的。然后判断前面的个数，可能存在
        //三种情况，在前面，在中间，在后面。
        //在前面的话，递归调用前面的，后面的话，递归调用后面的， 在中间直接返回。
        return findKthLargest(nums,0,nums.length - 1,k);
    }

    private int findKthLargest(int[] nums,int start,int end,int k){
        if (end == start) return nums[start];
        int mid = nums[start];
        int pre = start;
        int pro = end;
        while (pre < pro){
            //对于所有大于等于mid的，跳过
            while(pre < end && nums[pre] >= mid) pre ++;
            while(pro > start && nums[pro] < mid) pro --;
            if (pro > pre) {
                int temp = nums[pre];
                nums[pre] = nums[pro];
                nums[pro] = temp;
            }
        }
        //把最后一个进行交换上。
        nums[start] = nums[pro];
        nums[pro] = mid;
        //中间的位置是pro了。pro前面有多少个：pro - start个
        if(pro - start == k - 1){
            return nums[pro];
        }else if (pro - start < k - 1){
            return findKthLargest(nums,pro + 1,end,k - pro + start - 1);
        }else{
            return findKthLargest(nums,start,pro - 1,k);
        }
    }
}