class Solution {
    //300. 最长递增子序列
    public int lengthOfLIS(int[] nums) {
        //dp[i] 代表以nums[i]结尾的最长子序列长度
        int[] dp = new int[nums.length];
        int res = 1;
        dp[0] = 1;
        for(int i = 1;i < nums.length;i ++){
            int item = 0;
            for(int j = 0;j < i;j ++){
                if(nums[i] > nums[j] && dp[j] > item) item = dp[j];
            }
            dp[i] = item + 1;
            res = Math.max(dp[i],res);
        }
        return res;
    }
    //187. 重复的DNA序列
    public List<String> findRepeatedDnaSequences(String s) {
        //思路：hashmap，存储每一次10个长度，共2^11 2048种。
        //优化：将字符串变为数字，每一次后移，其实就是移动两位，加上新的数而已。看数字是否出现过即可。
        // acgt => 0 1 2 3
        Map<Character,Integer> map = new HashMap();
        Map<Integer,Integer> map2 = new HashMap();
        map.put('A',0);
        map.put('C',1);
        map.put('G',2);
        map.put('T',3);
        List<String> res = new ArrayList();
        if (s.length() < 10) return res;
        //先求前十位
        int window = 0;
        for(int i = 0;i < 10;i ++){
            window <<= 2;
            window += map.get(s.charAt(i));
        }
        map2.put(window,1);
        for(int i = 10;i < s.length();i ++){
            //求移动后新的数
            window <<= 2;
            //只保留前二十位
            window &= 0x000fffff;
            window += map.get(s.charAt(i));
            int count = map2.getOrDefault(window,0);
            if (count == 1){
                res.add(s.substring(i - 9,i + 1));
            }
            map2.put(window,count + 1);
        }
        return res;
    }
    //917. 仅仅反转字母
    public String reverseOnlyLetters(String s) {
        byte[] bytes = s.getBytes();
        //
        int start = 0;
        int end = bytes.length - 1;
        while (end > start){
            //start往前走，
            while(start < end && !isLetter(bytes[start])) start++;
            while(end > start && !isLetter(bytes[end])) end --;
            if(end > start) {
                byte temp = bytes[start];
                bytes[start] = bytes[end];
                bytes[end] = temp;
                start ++;
                end --;
            }
        }
        return new String(bytes);
    }

    private boolean isLetter(byte b){
        return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z');
    }

    //43. 字符串相乘
    public String multiply(String num1, String num2) {
        //按照数学的乘法来计算，每一位计算相加，后一位的
        if(num1.equals("0") || num2.equals("0")) return "0";
        String res = "0";
        for(int i = num1.length() - 1;i >= 0;i --){
            int cur = num1.charAt(i) - '0';
            String itemRes = mutiply(num2,cur,num1.length() - 1 - i);
            //System.out.println(itemRes);
            String temp = sum(itemRes,res);
            //System.out.println(itemRes + " + " + res + " = "+ temp);
            res = temp;
        }
        return res;
    }
    //一个大数 * 一个数
    private String mutiply(String num1,int num2,int ten){
        StringBuilder sb = new StringBuilder();
        //补充0
        for(;ten > 0;ten --) sb.append("0");
        //相乘，从低位开始乘
        int flag = 0;
        for(int i = num1.length() - 1;i >= 0;i --){
            int cur = num1.charAt(i) - '0';
            int item = cur * num2 + flag;
            flag = item/10;
            sb.append(item % 10);
        }
        if (flag != 0) sb.append(flag);
        return sb.reverse().toString();
    }
    //多个大数相加
    private String sum(String num1,String num2){
        StringBuilder res = new StringBuilder();
        int index = 0;
        int flag = 0;
        while(index < num1.length() && index < num2.length()){
            int cur1 = num1.charAt(num1.length() - 1 - index) - '0';
            int cur2 = num2.charAt(num2.length() - 1 - index) - '0';
            int item  = cur1 + cur2 + flag;
            flag = item / 10;
            res.append(item % 10);
            index ++;
        }
        while(index < num1.length()){
            int cur = num1.charAt(num1.length() - 1 - index) - '0';
            int item = cur + flag;
            flag = item / 10;
            res.append(item % 10);
            index ++;
        }
        while(index < num2.length()){
            int cur = num2.charAt(num2.length() - 1 - index) - '0';
            int item = cur + flag;
            flag = item / 10;
            res.append(item % 10);
            index ++;
        }
        if (flag > 0) res.append(flag);
        return res.reverse().toString();
    }
}