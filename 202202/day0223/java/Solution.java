class Solution {
    //2149. 按符号重排数组
    public int[] rearrangeArray(int[] nums) {
        int[] res = new int[nums.length];
        int i = 0;
        int j = 0;
        int k = 0;
        while (k < nums.length){
            while(i < nums.length && nums[i] < 0) i ++;
            while(j < nums.length && nums[j] > 0) j ++;
            if (k % 2 == 0){
                res[k++] = nums[i++];
            }else{
                res[k++] = nums[j++];
            }
        }
        return res;
    }
    //79. 单词搜索
    public boolean exist(char[][] board, String word) {
        boolean[][] used = new boolean[board.length][board[0].length];
        for(int i = 0;i < board.length;i ++){
            for (int j = 0;j <board[0].length;j ++){
                used[i][j] = true;
                if (check(board,used,i,j,word)){
                    return true;
                }
                used[i][j] = false;
            }
        }
        return false;
    }

    private boolean check(char[][] board,boolean[][] used,int x,int y,String word){
        if (word.length() == 1){
            return word.charAt(0) == board[x][y];
        }
        if (word.charAt(0) != board[x][y]) return false;
        int[][] dirctions = new int[][]{{1,0},{-1,0},{0,1},{0,-1}};
        for(int[] dirct : dirctions){
            if (x + dirct[0] >= 0 && x + dirct[0] < board.length && y + dirct[1] >= 0 && y + dirct[1] < board[0].length){
                if (!used[x + dirct[0]][y + dirct[1]]){
                    used[x + dirct[0]][y+dirct[1]] =  true;
                    if (check(board,used,x + dirct[0],y + dirct[1],word.substring(1,word.length()))){
                        return true;
                    }
                    used[x + dirct[0]][y + dirct[1]] = false;
                }
            }
        }
        return false;
    }
    //416. 分割等和子集
    public boolean canPartition(int[] nums) {
        //用一个hashset
        int sum = 0;
        for(int i :nums) sum += i;
        if (sum % 2== 1) return false;
        int target = sum / 2;
        HashSet<Integer> set = new HashSet();
        set.add(nums[0]);
        for(int i = 1;i <nums.length;i ++){
            if (nums[i] < target){
                HashSet<Integer> newSet = new HashSet();
                for(Integer num : set){
                    if (num + nums[i] == target){
                        return true;
                    }else if (num + nums[i] < target){
                        newSet.add(num + nums[i]);
                    }
                }
                set.addAll(newSet);
            }
        }
        return set.contains(target);
    }

    //322. 零钱兑换
    public int coinChange(int[] coins, int amount) {
        // dp[i][j] 只用前i种硬币，凑j元钱的最少次数
        // 凑不出来是-1
        int[][] dp = new int[coins.length][amount + 1];
        for(int i = 0;i < dp.length;i ++){
            for(int j = 1;j < amount + 1;j ++){
                if (coins[i] == j){
                    dp[i][j] = 1;
                }else{
                    dp[i][j] = -1;
                }
            }
        }
        for(int i = 0;i < dp.length;i ++){
            for(int j = 0;j < amount + 1;j ++){
                //不能用本次的硬币
                if (j < coins[i]){
                    if (i > 0){
                        dp[i][j] = dp[i-1][j];
                    }
                    continue;
                }
                //用一个当前货币
                if (dp[i][j - coins[i]] >= 0){
                    dp[i][j] = dp[i][j - coins[i]] + 1;
                }

                if (i > 0 && dp[i - 1][j] >= 0){
                    if (dp[i][j] == -1){
                        dp[i][j] = dp[i - 1][j];
                    }else{
                        // 不用当前硬币的情况
                        dp[i][j] = Math.min(dp[i - 1][j],dp[i][j]);
                    }
                }
            }
        }
        return dp[coins.length - 1][amount];
    }
    //剑指 Offer II 022. 链表中环的入口节点
    public ListNode detectCycle(ListNode head) {
        //快慢指针，先找到环，一个一次走两步，一个一次走一步
        //在两个指针走到重合后，一个从头开始走，一个就在原地开始走，下一次重合的位置即环开始的地方
        if (head == null || head.next == null) return null;
        ListNode fast = head.next;//fast走两步
        ListNode slow = head;//slow走一步
        while(fast != null){
            //fast 走两步
            fast = fast.next;
            if (fast == null) return null;
            fast = fast.next;
            //slow走一步
            slow = slow.next;
            //两者碰面了，说明环存在了
            if (fast == slow){
                //此时，fast从开头开始走，slow在原位值开始走
                fast = head;
                slow = slow.next;
                while (fast != slow){
                    slow = slow.next;
                    fast = fast.next;
                }
                return fast;
            }
        }
        return null;
    }

    //剑指 Offer II 021. 删除链表的倒数第 n 个结点
    class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
    public ListNode removeNthFromEnd(ListNode head, int n) {
        //两个指针，一个先走n步，然后再走一步，第二个指针接着走，若第一个指针为最后一个元素了，则第二个指针的下一个元素即要删除的元素
        if (head.next == null) return null;

        //fast走一步
        ListNode fast = head;
        //fast再走n-1步
        while (n > 1){
            fast = fast.next;
            n --;
        }
        //删除第一个元素
        if (fast.next == null) return head.next;
        //快的再走一步
        fast = fast.next;
        //慢的走一步
        ListNode slow = head;
        while (fast.next != null){
            fast = fast.next;
            slow = slow.next;
        }
        slow.next = slow.next.next;
        return head;
    }
    //53. 最大子数组和
    public int maxSubArray(int[] nums) {
        //dp[i] ,以nums[i] 结尾的最大子数组，若前面是负的，不加，前面是正的，加上
        int[] dp = new int[nums.length];
        dp[0] = nums[0];
        int res = dp[0];
        for(int i = 1;i < nums.length;i ++){
            if (dp[i - 1] > 0){
                dp[i] = nums[i] + dp[i - 1];
            }else{
                dp[i] = nums[i];
            }
            res = Math.max(res,dp[i]);
        }
        return res;
    }
    //673. 最长递增子序列的个数
    public int findNumberOfLIS(int[] nums) {
        int[] dp = new int[nums.length];
        int[] dpCount = new int[nums.length];
        int dpMax = 1;
        dp[0] = 1;
        dpCount[0] = 1;
        for(int i = 1;i < nums.length;i ++){
            int max = 0;
            int maxCount = 0;
            //需要记录有几个可以产生最大子序列情况。
            for(int j = 0;j < i;j ++){
                if (nums[i] > nums[j] && dp[j] > max){
                    max = dp[j];
                    maxCount = 0;
                }
                if (max == dp[j]) maxCount += dpCount[j];
            }
            maxCount = Math.max(1,maxCount);
            dp[i] = max + 1;
            dpCount[i] = maxCount;
            dpMax = Math.max(dpMax,dp[i]);
        }
        //找到最大的值
        int res = 0;
        for(int i = 0;i < nums.length;i ++) {
            if (dp[i] == dpMax) res += dpCount[i];
        }
        return res;
    }
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