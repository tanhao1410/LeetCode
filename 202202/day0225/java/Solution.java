class Solution {
    //537. 复数乘法
    public String complexNumberMultiply(String num1, String num2) {
        StringBuilder res = new StringBuilder();
        //先算整数
        String[] num1s = num1.substring(0,num1.length() - 1).split("\\+");
        String[] num2s = num2.substring(0,num2.length() - 1).split("\\+");
        int numPart = Integer.parseInt(num1s[0]) * Integer.parseInt(num2s[0]);
        int iMulti = Integer.parseInt(num1s[1]) * Integer.parseInt(num2s[0]) + Integer.parseInt(num2s[1])* Integer.parseInt(num1s[0]);
        int numPart2 = Integer.parseInt(num1s[1]) * Integer.parseInt(num2s[1]) * -1;
        res.append(numPart + numPart2);
        res.append("+");
        res.append(iMulti);
        res.append("i");
        return res.toString();
    }
}