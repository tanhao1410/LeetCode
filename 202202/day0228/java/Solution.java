class Solution {
    //97. 交错字符串
    public boolean isInterleave(String s1, String s2, String s3) {
        //递归思路：很容易超时。
        // if(s3.length() == 0) return true;
        // boolean res = false;
        // if(s1.length() > 0 && s3.charAt(0) == s1.charAt(0)){
        //     res |= isInterleave(s1.substring(1,s1.length()),s2,s3.substring(1,s3.length()));
        // }
        // if( s2.length() > 0 && s3.charAt(0) == s2.charAt(0)){
        //     res |= isInterleave(s1,s2.substring(1,s2.length()),s3.substring(1,s3.length()));
        // }
        // return res;

        //问题是当一个开头，即可以用s1，又可以用s2，产生了分歧。
        // dp[i][j][k] dp[0].. = true;
        // dp[i][0][k]  看是否相等
        // dp[i][k][0]  一样
        // dp[i][j][k]  dp[i - 1] [j - 1][k] , dp[i-1][j][k-1],
        boolean[][][] dp = new boolean[s3.length()+1][s1.length()+1][s2.length()+1];
        for(int i = 0;i <= s3.length();i ++){
            for(int j = 0;j <= s1.length();j ++){
                for(int k = 0;k <= s2.length();k ++){
                    if(i != j + k){
                        dp[i][j][k] = false;
                    }else if(i == 0){
                        dp[i][j][k] = true;
                    }else{
                        if (j == 0){
                            dp[i][0][k] = dp[i -1][0][k - 1] && s3.charAt(i - 1) == s2.charAt(k - 1);
                        }else if (k == 0){
                            dp[i][j][k] = dp[i - 1][j - 1][k] && s3.charAt(i - 1) == s1.charAt(j - 1);
                        }else{
                            dp[i][j][k] |= dp[i - 1][j - 1][k] && s3.charAt(i - 1) == s1.charAt(j - 1);
                            dp[i][j][k] |= dp[i - 1][j][k-1] && s3.charAt(i - 1) == s2.charAt(k - 1);
                        }
                    }
                }
            }
        }
        return dp[s3.length()][s1.length()][s2.length()];
    }
    //1823. 找出游戏的获胜者
    public int findTheWinner(int n, int k) {
        boolean[] failed = new boolean[n];
        int failNum = 0;
        int start = 0;
        while(failNum < n - 1){
            //需要淘汰一位
            for(int i = 0;i < k - 1;i ++){
                start = (start + 1) % n;
                while(failed[start]) start = (start + 1) % n;
            }
            failed[start] = true;
            //下一个开始的地方
            start = (start + 1) % n;
            while(failed[start]) start = (start + 1) % n;
            failNum ++;
        }
        for(int i = 0;i < n;i ++){
            if(!failed[i]) return i + 1;
        }
        return -1;
    }
    //1249. 移除无效的括号
    public String minRemoveToMakeValid(String s) {
        //思路：用一个变量记录(的出现次数，
        int count = 0;
        StringBuilder sb = new StringBuilder();
        Stack<Integer> stack = new Stack();
        Set<Integer> set = new HashSet();
        for(int i = 0;i < s.length();i ++){
            //前面可能需要移除(
            char cur = s.charAt(i);
            if(cur == '('){
                count ++;
                //需要记住这个位置
                stack.push(i);
            }else if (cur == ')'){
                if (count > 0){
                    count --;
                    set.add(i);
                }
            }
        }
        //说明前面还需要继续删除一些(
        while ( count > 0) {
            stack.pop();
            count --;
        }
        while(stack.size() > 0) set.add(stack.pop());
        for(int i = 0;i < s.length();i ++){
            char cur = s.charAt(i);
            if(cur == '(' || cur == ')'){
                if (set.contains(i)){
                    sb.append(cur);
                }
            }else{
                sb.append(cur);
            }
        }

        return sb.toString();
    }
    //524. 通过删除字母匹配到字典里最长单词
    public String findLongestWord(String s, List<String> dictionary) {
        //先根据长度倒排序，在根据字母序排序，找到合适的，就返回
        Collections.sort(dictionary,(e1,e2)->{
            if (e1.length() != e2.length()) return e2.length() - e1.length();
            return e1.compareTo(e2);
        });
        for(String word : dictionary){
            if (isSubStr(s,word)){
                return word;
            }
        }
        return "";
    }

    private boolean isSubStr(String sup,String sub){
        if(sup.length() < sub.length()) return false;
        int supIndex = 0;
        int subIndex = 0;
        while(supIndex < sup.length() && subIndex < sub.length()){
            if(sup.charAt(supIndex) == sub.charAt(subIndex)){
                supIndex ++;
                subIndex ++;
            }else{
                supIndex++;
            }
        }
        return subIndex == sub.length();
    }
    //989. 数组形式的整数加法
    public List<Integer> addToArrayForm(int[] num, int k) {
        //都转成数组吧
        List<Integer> num2List = new ArrayList();
        while (k > 0){
            num2List.add(k % 10);
            k /= 10;
        }
        // int[] num2 = new int[num2List.size()];
        // for(int i = 0;i < num2.length;i ++) num2[i] = num2List.get(num2.length - 1 - i);
        int index = 0;
        int flag = 0;
        LinkedList<Integer> res = new LinkedList();
        while(index < num2List.size() || index < num.length){
            int cur = flag;
            if (index < num2List.size() && index < num.length){
                cur += num[num.length - 1 - index]  + num2List.get(index);
            }else if (index < num2List.size()){
                cur += num2List.get(index);
            }else{
                cur += num[num.length - 1 - index];
            }
            res.add(0,cur % 10);
            flag = cur / 10;
            index ++;
        }
        if (flag == 1){
            res.add(0,1);
        }
        return res;
    }

    //1601. 最多可达成的换楼请求数目
    public int maximumRequests(int n, int[][] requests) {
        //思路：递归思路，总共就16个请求，对于每一个请求都有两个选择，要或不要。共2^16种情况。可以暴力解决
        return maximumRequests(requests,new boolean[requests.length],0,n);
    }
    private int canResolve(int[][] requests,boolean[] used,int n){
        int[] buildings = new int[n];
        int res = 0;
        for(int i = 0;i < used.length;i ++){
            if(used[i]){
                res ++;
                int[] request = requests[i];
                buildings[request[0]] --;
                buildings[request[1]] ++;
            }
        }
        for(int b : buildings){
            if(b != 0) return 0;
        }
        return res;
    }
    private int maximumRequests(int[][] requests,boolean[] used,int curIndex,int n){
        if (curIndex == requests.length) return canResolve(requests,used,n);
        //对于当前的请求，两种选择
        used[curIndex] = false;
        int res = maximumRequests(requests,used,curIndex + 1,n);
        used[curIndex] = true;
        res = Math.max(res,maximumRequests(requests,used,curIndex + 1,n));
        return res;
    }
}
//155. 最小栈
class MinStack {
    //采用单调栈。push 进来一个数的时候，与最小栈的栈顶比较，若小于等于则，push进来，若大于，什么都不做，
    // 弹出数时，判断是否与最小栈顶部一样
    private Stack<Integer> data;
    private Stack<Integer> small;
    public MinStack() {
        this.data = new Stack();
        this.small = new Stack();
    }
    public void push(int val) {
        this.data.push(val);
        if (small.size() == 0){
            small.push(val);
        }else if (small.peek() >= val){
            small.push(val);
        }
    }
    public void pop() {
        int val = this.data.pop();
        if (val == small.peek()){
            small.pop();
        }
    }
    public int top() {
        return this.data.peek();
    }
    public int getMin() {
        return this.small.peek();
    }
}