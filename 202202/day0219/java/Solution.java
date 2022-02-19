class Solution {
    //1654. 到家的最少跳跃次数
    public int minimumJumps(int[] forbidden, int a, int b, int x) {
        //广度优先遍历法
        //不能跳的地方，与跳过的地方不再进行跳跃，用一个set来存储
        //有些位置是禁止走的。
        int maxDistance = x + b;
        Set<Integer> set = new HashSet();
        Set<Integer> backSet = new HashSet();
        Set<Integer> forbidderSet = new HashSet();
        for(int i : forbidden) {
            forbidderSet.add(i);
            maxDistance = Math.max(maxDistance,i + a +b);
        }
        LinkedList<Integer> queue = new LinkedList();
        set.add(0);
        queue.add(0);
        int step = 0;
        while (queue.size() > 0){
            if (set.contains(x) || backSet.contains(x)) return step;
            int size = queue.size();
            for (int i = 0;i < size;i ++){
                int cur = queue.remove(0);
                //System.out.print(cur);
                //System.out.print("   ");
                //不能往后连跳两次。如果上次是往后，那么这次不能接着往后了，所以需要记录上次是怎么跳到这的。
                //问题：如果是倒退到了某个位置，在这个位置不能继续往后退。但如果是正常走到这个位置，是可以后退的。还是有区分的。

                //往前走,不能无休止的往前走，走的最大位置不能超过 x + b
                //走到很远，一下子退回到x位置呢？
                if (!set.contains(cur + a) && cur <= maxDistance && !forbidderSet.contains(cur + a)){
                    queue.add(cur + a);
                    set.add(cur + a);
                }
                //往后走到某位置，当前位置没被反向走过。或者当前位置可以正向走到
                //这里不对！后退走到该位置，不代表不能在该位置继续后退。
                if (cur - b > 0 && !forbidderSet.contains(cur - b)  && (!backSet.contains(cur) || set.contains(cur))){
                    queue.add(cur - b);
                    backSet.add(cur - b);
                }

            }
            //System.out.println();
            step ++;
        }
        return -1;
    }
    //433. 最小基因变化
    public int minMutation(String start, String end, String[] bank) {
        LinkedList<String> queue = new LinkedList();
        boolean[] readed = new boolean[bank.length];
        for(int i = 0;i < bank.length;i ++){
            if (bank[i].equals(start)) readed[i] = true;
        }
        queue.add(start);
        int layer = 1;
        while (queue.size() > 0){
            int len = queue.size();
            for (int i = 0;i < len;i ++){
                String s = queue.remove(0);
                //测试s突变成什么
                for(int j = 0;j < bank.length;j ++){
                    if(!readed[j] && one_change(s,bank[j])){
                        if(bank[j].equals(end)){
                            return layer;
                        }
                        queue.add(bank[j]);
                        readed[j] = true;
                    }
                }
            }
            layer ++;
        }
        return -1;
    }

    private boolean one_change(String s1,String s2){
        int res = 0;
        for (int i = 0;i < 8;i ++){
            if(s1.charAt(i) != s2.charAt(i)) res ++;
        }
        return res == 1;
    }
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