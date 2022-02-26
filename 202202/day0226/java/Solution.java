class Solution {
    //209. 长度最小的子数组
    public int minSubArrayLen(int target, int[] nums) {
        int res = nums.length+ 1;
        int end = 0;
        int start = 0;
        int sum = nums[0];
        while (end < nums.length){
            //比sum大的情况
            if(sum >= target){
                //减小窗口
                while(sum >= target){
                    sum -= nums[start];
                    start ++;
                }
                //新窗口大小
                res = Math.min(end - start + 2,res);
            }else{
                //比sum小
                end ++;
                if (end == nums.length) break;
                sum += nums[end];
            }
        }
        return res == nums.length + 1?0:res;
    }
    //201. 数字范围按位与
    public int rangeBitwiseAnd(int left, int right) {
        //只看最高位的1到0位置
        // 从后往前看，如果遇到
        int res = 0;
        for(int i = 0;i < 32;i ++){
            int mask = 1 << 31 - i;
            if ((left & mask) != 0 && (right & mask) != 0) res += mask;
            if  ((left & mask) != (right & mask) ) break;
        }
        return res;
    }
    //456. 132 模式
    public boolean find132pattern(int[] nums) {
        //
        if(nums.length < 3) return false;
        int[] preMin = new int[nums.length];
        preMin[0] = 100000002;
        for(int i = 1;i < nums.length;i ++){
            preMin[i] = Math.min(preMin[i - 1],nums[i - 1]);
        }
        Stack<Integer> stack = new Stack();
        stack.push(nums[nums.length - 1]);
        for(int i = nums.length - 2;i >= 1; i --){
            //stack不可能为空
            int top = stack.peek();
            while( stack.size() > 0 && stack.peek() < nums[i]){
                top = stack.pop();
            }
            if (nums[i] > top && top > preMin[i]) return true;
            stack.push(nums[i]);
        }
        return false;
    }
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