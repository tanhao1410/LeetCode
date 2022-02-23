class Solution {
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