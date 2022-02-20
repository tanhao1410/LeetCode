class Solution {
    //997. 找到小镇的法官
    public int findJudge(int n, int[][] trust) {
        //用一个数组来表示自己信任的人的数量，用一个数组来表示信任自己的人的数量
        int[] trustOhter = new int[n];
        int[] trustSelf = new int[n];
        for(int[] nums:trust){
            trustOhter[nums[0] - 1] ++;
            trustSelf[nums[1] - 1] ++;
        }
        //查看符合条件的
        for(int i = 0;i < n;i ++){
            if (trustSelf[i] == n - 1 && trustOhter[i] == 0) return i + 1;
        }
        return -1;
    }
    //717. 1比特与2比特字符
    public boolean isOneBitCharacter(int[] bits) {
        for(int i = 0;;i ++){
            if (i >= bits.length - 1) return i == bits.length-1;
            if (bits[i] == 1) i ++;
        }
    }
}