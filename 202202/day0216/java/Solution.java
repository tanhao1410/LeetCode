class Solution {

    //剑指 Offer II 004. 只出现一次的数字
    public int singleNumber(int[] nums) {
        //记录每一位的个数，每一位的个数 % 3 即为结果 该位的值
        int res = 0;
        for(int i = 0;i < 32;i ++){
            int offset = 1<<i;
            int count = 0;
            for(int num : nums){
                if ((num & offset) != 0){
                    count ++;
                }
            }
            res += (count % 3) << i;
        }
        return res;
    }

    //169. 多数元素
    public int majorityElement(int[] nums) {
        int res = 0;
        int count = 0;
        for(int i = 0;i < nums.length;i ++){
            if(count == 0){
                res = nums[i];
                count ++;
            }else{
                if(res == nums[i]){
                    count ++;
                }else{
                    count --;
                }
            }
        }
        return res;
    }

    //1319. 连通网络的操作次数
    public int makeConnected(int n, int[][] connections) {
        if (n > connections.length + 1) return -1;
        // 先处理connections，改成一个数组，数组里放的是 set，里面是它能进入的地方
        HashSet<Integer>[] numSet = new HashSet[n];
        for(int[] connection : connections){
            if (numSet[connection[0]] == null) numSet[connection[0]] = new HashSet();
            if (numSet[connection[1]] == null) numSet[connection[1]] = new HashSet();
            numSet[connection[0]].add(connection[1]);
            numSet[connection[1]].add(connection[0]);
        }

        boolean[] isRead = new boolean[n];
        int netCount = 0;
        //深度优先遍历
        Stack<Integer> stack = new Stack();

        for(int i = 0;i < numSet.length;i ++){
            if (!isRead[i]){
                netCount ++;
                stack.push(i);
                while (stack.size() > 0){
                    int cur = stack.pop();
                    isRead[cur] = true;
                    if (numSet[cur] != null)for(Integer j :numSet[cur]){
                        if (!isRead[j]){
                            stack.push(j);
                        }
                    }
                }
            }
        }
        for(boolean b : isRead){
            if (!b) netCount ++;
        }
        return netCount - 1;
    }

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