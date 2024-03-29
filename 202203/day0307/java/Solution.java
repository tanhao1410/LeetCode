class Solution {
    //713. 乘积小于K的子数组
    public int numSubarrayProductLessThanK(int[] nums, int k) {
        //滑动窗口？先计算一个窗口，小于k,然后移动。每加入一个数，看是否可以，如果不可以了，减去一个数，直到小于或敞口为0
        int start = 0;
        int end = 0;
        int res = 0;
        int mul = nums[0];
        while(end < nums.length){
            //看窗口内的积是否大于k
            if(mul < k){
                //意思是这个数可以放进来，但要加几呢？
                //
                res += (end - start + 1);
                end ++;
                if(end < nums.length){
                    mul *= nums[end];
                }
            }else{
                //说明窗口内的积大于k,start向前走
                mul /= nums[start];
                start ++;
                //如果此时start == end了，即
                if (start > end){
                    end = start;
                    if(start < nums.length) mul = nums[start];
                }
            }
        }
        return res;
    }
    //438. 找到字符串中所有字母异位词
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> res = new ArrayList();
        //用一个数组表示各位的情况。
        int[] bits = new int[26];
        byte[] pBytes = p.getBytes();
        for(int i = 0;i < pBytes.length;i ++) bits[pBytes[i] - 'a']++;
        for(int i = 0;i < p.length() && i < s.length();i ++) bits[s.charAt(i) - 'a'] --;
        if(allZero(bits)) res.add(0);
        //确定是否相同看的就是每位是否都是0
        // 从0开始算起，进来一个词，减去一个词
        for(int i = p.length();i < s.length();i ++){
            //加上一个词，减去一个词
            bits[s.charAt(i) - 'a'] --;
            bits[s.charAt(i - p.length()) - 'a'] ++;
            if(allZero(bits)){
                res.add(i - p.length() + 1);
            }
        }
        return res;
    }

    private boolean allZero(int[] bits){
        for(int i = 0;i < bits.length;i ++){
            if(bits[i] != 0) return false;
        }
        return true;
    }
    //504. 七进制数
    public String convertToBase7(int num) {
        if (num == 0)return "0";
        if (num < 0){
            return "-"+convertToBase7(-num);
        }
        //贪心算法，最大8次方
        int[] ss = new int[9];
        ss[0] = 1;
        for(int i = 1;i < 9;i ++){
            ss[i] = ss[i-1]*7;
        }
        StringBuilder res = new StringBuilder();
        int len = 8;
        for(;len >= 0;len --){
            if (num >= ss[len]){
                break;
            }
        }
        //计算每一位应该是多少；
        byte[] bits = new byte[len + 1];
        for(int i = 0;i < len + 1;i ++) bits[i] = '0';
        for(int i = len;i >= 0;i --){
            if (num >= ss[i]){
                bits[len-i] ++;
                num -= ss[i];
                i ++;
            }
        }
        return new String(bits);
    }
}