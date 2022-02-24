class Solution {
    //20. 有效的括号
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack();
        for(int i = 0;i < s.length();i ++){
            char cur = s.charAt(i);
            if(cur == '(' || cur =='{' || cur == '['){
                stack.push(cur);
            }else if (stack.size() > 0){
                Character pop = stack.pop();
                if (cur == ')' && pop != '(') return false;
                if (cur == ']' && pop != '[') return false;
                if (cur == '}' && pop != '{') return false;
            }else{
                return false;
            }
        }
        return stack.size() == 0;
    }
}