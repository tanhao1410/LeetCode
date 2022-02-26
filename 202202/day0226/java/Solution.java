class Solution {
    //150. 逆波兰表达式求值
    public int evalRPN(String[] tokens) {
        //采用栈的方式求值
        Stack<Integer> stack = new Stack();
        for(String s : tokens){
            if(s.length() == 1 && (s.charAt(0) < '0' || s.charAt(0) > '9')){
                int b = stack.pop();
                int a = stack.pop();
                if(s.equals("+")){
                    stack.push(a + b);
                }else if (s.equals("-")){
                    stack.push(a - b);
                }else if (s.equals("*")){
                    stack.push(a *b);
                }else{
                    stack.push(a / b);
                }
            }else{
                stack.push(Integer.parseInt(s));
            }
        }
        return stack.pop();
    }
    //459. 重复的子字符串
    public boolean repeatedSubstringPattern(String s) {
        return (s + s).substring(1,s.length() * 2 - 1).contains(s);
    }
}