class Solution {
    //896. 单调数列
    public boolean isMonotonic(int[] nums) {
        boolean down = false;
        boolean up  = false;
        for(int i = 1;i < nums.length;i ++){
            if (nums[i]>nums[i - 1]) up = true;
            if (nums[i] < nums[i - 1]) down = true;
            if (up && down) return false;
        }
        return true;
    }
    //583. 两个字符串的删除操作
    public int minDistance(String word1, String word2) {
        int[][] dp = new int[word1.length()+1][word2.length() + 1];
        for(int i = 0;i < word1.length();i ++){
            for (int j = 0;j < word2.length();j ++){
                if(word1.charAt(i) == word2.charAt(j)){
                    dp[i + 1][j + 1] = dp[i][j] + 1;
                }else{
                    dp[i + 1][j + 1] = Math.max(dp[i][j+1],dp[i + 1][j]);
                }
            }
        }
        return word1.length() + word2.length() - 2 * dp[word1.length()][word2.length()];
    }
    //1143. 最长公共子序列
    public int longestCommonSubsequence(String text1, String text2) {
        //动态规划算法；
        int[][] dp = new int[text1.length()][text2.length()];
        for(int i = 0;i < text1.length();i ++){
            for(int j =0;j < text2.length();j ++){
                if (text1.charAt(i) == text2.charAt(j)){
                    if (i > 0 && j > 0) {
                        dp[i][j] = dp[i - 1][j - 1] + 1;
                    }else{
                        dp[i][j] = 1;
                    }
                }else{
                    if (i > 0 && j > 0){
                        dp[i][j] = Math.max(dp[i - 1][j],dp[i][j-1]);
                    }else if (i > 0){
                        dp[i][j] = dp[i -1][j];
                    }else if (j > 0){
                        dp[i][j] = dp[i][j - 1];
                    }

                }
            }
        }
        return dp[dp.length - 1][dp[0].length - 1];
    }
    //496. 下一个更大元素 I
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        //元素的位置可以用一个map来快读获得。下一个比自己大的元素怎么获取呢？
        int[] next = new int[nums2.length];
        Map<Integer,Integer> map = new HashMap();
        Stack<Integer> stack = new Stack();
        for(int i = nums2.length - 1;i >= 0;i --){
            int cur = nums2[i];
            map.put(cur,i);
            if (stack.size() == 0){
                stack.push(cur);
                next[i] = -1;
            }else{
                //看stack顶部的数是否大于自己，如果大于自己
                int top = stack.peek();
                if (top > cur){
                    stack.push(cur);
                    next[i] = top;
                }else{
                    while(stack.size() > 0 && stack.peek() < cur) stack.pop();
                    if (stack.size() == 0){
                        stack.push(cur);
                        next[i] = -1;
                    }else{
                        next[i] = stack.peek();
                        stack.push(cur);
                    }
                }
            }
        }
        int[] res = new int[nums1.length];
        for(int i = 0;i < res.length;i ++){
            res[i] = next[map.get(nums1[i])];
        }
        return res;
    }
    //1706. 球会落何处
    public int[] findBall(int[][] grid) {
        //思路：对于每一个格子来说，它上面的球会往哪个方向走呢？
        // 如果是 1 ，则球往下一个，否则往左边走
        // 什么情况下，球下不去了呢？往右边走时，它右边的格子是-1，往左边走时，它左边的格子是+1
        //一层一层往下掉，
        boolean[] balls = new boolean[grid[0].length];
        int[] ballsRes  = new int[balls.length];
        for(int i = 0;i < balls.length;i ++){
            balls[i] = true;
            ballsRes[i] = i;
        }
        return findBall(grid,0,balls,ballsRes);

    }
    //balls :每一个格是否有球，ballsRes：每一格球的最开始来源
    private int[] findBall(int[][] grid,int layer,boolean[] balls,int[] ballsRes){
        if (layer == grid.length){
            int[] res = new int[grid[0].length];
            for(int i = 0;i < res.length;i ++) res[i] = -1;
            for(int i = 0;i < res.length;i ++){
                if (balls[i]){
                    res[ballsRes[i]] = i;
                }
            }
            return res;
        }
        boolean[] nextLayerBalls = new boolean[balls.length];
        int[] nextLayerBallsRes = new int[balls.length];
        for(int i = 0;i < ballsRes.length;i ++){
            //判断是否有球
            if(balls[i]){
                if (grid[layer][i] == 1){//往右边走
                    if (i != balls.length - 1 && grid[layer][i + 1] != -1){
                        //可以往下落
                        nextLayerBalls[i + 1] = true;
                        nextLayerBallsRes[i + 1] = ballsRes[i];
                    }
                }else{
                    if (i != 0 && grid[layer][i - 1] != 1){
                        nextLayerBalls[i - 1] = true;
                        nextLayerBallsRes[i - 1] = ballsRes[i];
                    }
                }
            }
        }
        return findBall(grid,layer + 1,nextLayerBalls,nextLayerBallsRes);
    }
    //142. 环形链表 II
    public ListNode detectCycle(ListNode head) {
        if (head == null || head.next == null) return null;
        //快慢指针先找=判断是否有换，
        ListNode fast = head.next;
        ListNode slow = head;
        while(fast != null && slow != null){
            fast = fast.next;
            if (fast == null) return null;
            fast = fast.next;
            slow = slow.next;
            if (fast == slow){
                //成环了
                //此时，fast指针跳到head，一步一步走即可
                fast = head;
                slow = slow.next;
                while(fast != slow){
                    slow = slow.next;
                    fast = fast.next;
                }
                return fast;
            }
        }
        return null;
    }
    //2. 两数相加
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        //用l1返回结果
        ListNode res = l1;
        int flag = 0;
        ListNode pre = null;
        while(l1 != null && l2 != null){
            l1.val += l2.val + flag;
            flag = l1.val/10;
            l1.val %= 10;
            pre = l1;
            l1 = l1.next;
            l2 = l2.next;
        }
        if (l1 == null) {
            pre.next = l2;
            l1 = pre.next;
        }

        while(l1 != null){
            l1.val += flag;
            flag = l1.val/10;
            l1.val %= 10;
            pre = l1;
            l1 = l1.next;
        }

        if (flag > 0){
            pre.next = new ListNode(flag);
        }

        return res;
    }
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
    class ListNode {
          int val;
          ListNode next;
          ListNode() {}
          ListNode(int val) { this.val = val; }
          ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
}