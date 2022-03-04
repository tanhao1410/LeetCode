class Solution {

    //1630. 等差子数组
    public List<Boolean> checkArithmeticSubarrays(int[] nums, int[] l, int[] r) {
        //判断是否是等差数列，排序，看差值是否相等。
        List<Boolean> res = new ArrayList();
        for(int i = 0;i < l.length;i ++){
            int[] itemArray = new int[r[i] - l[i] + 1];
            for(int j = l[i];j <= r[i];j ++) itemArray[j-l[i]] = nums[j];
            res.add(isSubarrays(itemArray));
        }
        return res;
    }

    //判断一个数组是否是等差数列数组，排序，判断，nlogn；
    private boolean isSubarrays(int[] nums){
        if(nums.length < 2) return false;
        Arrays.sort(nums);
        int dis = nums[1] - nums[0];
        for(int i = 2;i < nums.length;i ++){
            if (nums[i] - nums[i - 1] != dis) return false;
        }
        return true;
    }

    //1987. 不同的好子序列数目
    public int numberOfUniqueGoodSubsequences(String binary) {
        int[] dp0 =  new int[binary.length()];
        int[] dp1 = new int[binary.length()];
        int containsZero = 0;
        if(binary.charAt(binary.length() - 1) == '0'){
            containsZero = 1;
            dp0[binary.length() - 1] = 1;
        }else{
            dp1[binary.length() - 1] = 1;
        }
        System.out.println(binary.length());
        //以0,1开头的子序列的个数
        for(int i = binary.length() - 2;i >= 0;i --){
            if (binary.charAt(i) == '0'){
                containsZero = 1;
                dp0[i] = (dp0[i + 1] + dp1[i + 1] + 1) % 1000000007;
                dp1[i] = dp1[i + 1];
            }else{
                dp1[i] = (dp1[i + 1] + dp0[i + 1] + 1)%1000000007;
                dp0[i] = dp0[i + 1];
            }
        }
        return dp1[0] + containsZero;
    }
    //297. 二叉树的序列化与反序列化
    public class Codec {
        // Encodes a tree to a single string.
        public String serialize(TreeNode root) {
            StringBuilder res = new StringBuilder();
            res.append("[");
            LinkedList<TreeNode> queue = new LinkedList();
            queue.add(root);
            while(queue.size() > 0){
                int size = queue.size();
                for(int i = 0;i < size;i ++){
                    TreeNode head = queue.remove(0);
                    if (head == null){
                        res.append("nil");
                    }else{
                        queue.add(head.left);
                        queue.add(head.right);
                        res.append(head.val);
                    }
                    res.append(",");
                }
            }
            res.deleteCharAt(res.length() - 1);
            res.append("]");
            return res.toString();
        }

        // Decodes your encoded data to tree.
        public TreeNode deserialize(String data) {
            String[] datas = data.substring(1,data.length() - 1).split(",");
            if(datas[0].equals("nil")){
                return null;
            }
            TreeNode root = new TreeNode(Integer.parseInt(datas[0]));
            LinkedList<TreeNode> queue = new LinkedList();
            boolean insertLeft = true;
            queue.add(root);
            for(int i = 1;i < datas.length;i ++){
                if(!datas[i].equals("nil")){
                    //需要把新的节点插入到对应的节点上
                    TreeNode newNode = new TreeNode(Integer.parseInt(datas[i]));
                    if(insertLeft){
                        TreeNode parent = queue.get(0);
                        parent.left = newNode;
                    }else{
                        TreeNode parent = queue.remove(0);
                        parent.right = newNode;
                    }
                    queue.add(newNode);
                }else if (!insertLeft){
                    queue.remove(0);
                }
                insertLeft = !insertLeft;
            }
            return root;
        }
    }
    //429. N 叉树的层序遍历
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> res = new ArrayList();
        LinkedList<Node> queue= new LinkedList();
        if(root != null) queue.add(root);
        while (queue.size() > 0){
            int size = queue.size();
            List<Integer> item = new ArrayList();
            for (int i = 0;i < size;i ++){
                Node node = queue.remove(0);
                item.add(node.val);
                for(Node n : node.children){
                    if (n != null) queue.add(n);
                }
            }
            res.add(item);
        }
        return res;
    }
    class Node {
        public int val;
        public List<Node> children;
    };

    //124. 二叉树中的最大路径和
    public int maxPathSum(TreeNode root) {
        return rootStartMaxAndMaxValue(root)[1];
    }

    // int[0] 是以root开始的最大的值，int[1] 是整个的最大值。
    private int[] rootStartMaxAndMaxValue(TreeNode root){
        //以根节点开始
        if(root == null){
            return new int[]{0,-1001};
        }
        int[] left = rootStartMaxAndMaxValue(root.left);
        int[] right = rootStartMaxAndMaxValue(root.right);
        int[] res = new int[2];
        //以root 为开始的最大值，要么连接左边，要么连接右边，要么自己就行
        res[0] = Math.max(Math.max(0,left[0]),right[0]) + root.val;
        //对于结果的最大值，要么是包括root本身，要么是root作为中间，要么是来自与子树
        res[1] = Math.max(Math.max(res[0],root.val + left[0] + right[0]),Math.max(left[1],right[1]));
        return res;
    }
    public class TreeNode {
         int val;
         TreeNode left;
         TreeNode right;
         TreeNode() {}
         TreeNode(int val) { this.val = val; }
         TreeNode(int val, TreeNode left, TreeNode right) {
             this.val = val;
             this.left = left;
             this.right = right;
         }
     }
}