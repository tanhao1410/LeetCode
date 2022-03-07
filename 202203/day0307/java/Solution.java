class Solution {
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