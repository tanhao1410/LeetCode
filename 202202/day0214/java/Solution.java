class Solution {

    //1162. 地图分析
    public int maxDistance(int[][] grid) {
        int[][] res = new int[grid.length][grid[0].length];
        //陆地的数量与位置
        LinkedList<Location> queue = new LinkedList();
        for(int i = 0;i < res.length;i ++){
            for (int j = 0;j < res[0].length;j ++){
                if (grid[i][j] == 1){
                    queue.add(new Location(i,j));
                    res[i][j] = -1;
                }
            }
        }
        if (queue.size() == 0 || queue.size() == res.length * res[0].length) return -1;
        int distance = 1;
        while (queue.size() > 0){
            //从队列中取出一个
            int size = queue.size();
            for(int i = 0;i <size;i ++ ){
                Location lo = queue.remove(0);
                //将它周围的未设置位置的设置位置
                if (lo.x > 0 && res[lo.x - 1][lo.y] == 0){
                    res[lo.x - 1][lo.y] = distance;
                    queue.add(new Location(lo.x - 1,lo.y));
                }
                if (lo.x < grid.length - 1 && res[lo.x + 1][lo.y] == 0){
                    res[lo.x + 1][lo.y] = distance;
                    queue.add(new Location(lo.x + 1,lo.y));
                }
                if (lo.y > 0 && res[lo.x][lo.y - 1] == 0){
                    res[lo.x][lo.y - 1]  = distance;
                    queue.add(new Location(lo.x,lo.y - 1));
                }
                if (lo.y < grid[0].length - 1 && res[lo.x][lo.y + 1] == 0){
                    res[lo.x][lo.y + 1] = distance;
                    queue.add(new Location(lo.x,lo.y + 1));
                }
            }
            distance ++ ;
        }
        int max = -1;
        for (int i = 0;i < res.length;i ++){
            for (int j = 0;j < res[0].length;j ++){
                max = Math.max(max,res[i][j]);
            }
        }
        return max;
    }

    //1254. 统计封闭岛屿的数目
    public int closedIsland(int[][] grid) {
        //思路：将外围的0以及与0相连的0都变成1,然后剩下的0就是封闭岛屿了
        //计算封闭岛屿个数即可。遍历0，然后将与其相连的0都变为1
        Stack<Location> stack = new Stack();
        for(int i = 0;i < grid[0].length;i ++){
            if (grid[0][i] == 0) stack.push(new Location(0,i));
            if (grid[grid.length - 1][i] == 0) stack.push(new Location(grid.length - 1,i));
        }
        for(int i = 1;i < grid.length - 1;i ++){
            if (grid[i][0] == 0) stack.push(new Location(i,0));
            if (grid[i][grid[0].length - 1] == 0) stack.push(new Location(i,grid[0].length - 1));
        }
        setOne(grid,stack);
        //开始计算岛屿数量
        int res = 0;
        for(int i = 1;i < grid.length - 1;i ++){
            for (int j = 1;j < grid[0].length - 1;j ++){
                if(grid[i][j] == 0){
                    res ++;
                    stack.push(new Location(i,j));
                    setOne(grid,stack);
                }
            }
        }
        return res;
    }

    private void setOne(int[][] grid,Stack<Location> stack){
        //将它周围的0都变成1
        while(!stack.isEmpty()){
            Location lo = stack.pop();
            grid[lo.x][lo.y] = 1;
            if (lo.x > 0 && grid[lo.x - 1][lo.y] == 0) stack.push(new Location(lo.x - 1,lo.y));
            if (lo.x < grid.length - 1 && grid[lo.x + 1][lo.y] == 0) stack.push(new Location(lo.x + 1,lo.y));
            if (lo.y > 0 && grid[lo.x][lo.y - 1] == 0) stack.push(new Location(lo.x,lo.y - 1));
            if (lo.y < grid[0].length - 1 && grid[lo.x][lo.y + 1] == 0) stack.push(new Location(lo.x,lo.y + 1));
        }
    }

    //695. 岛屿的最大面积
    public int maxAreaOfIsland(int[][] grid) {
        int width = grid.length;
        int height = grid[0].length;
        int res = 0;
        for(int i = 0;i < width;i ++){
            for (int j = 0;j < height;j ++){
                if (grid[i][j] == 1){
                    //开始遍历
                    res = Math.max(res,dfs(grid,i,j));
                }
            }
        }
        return res;
    }

    class Location{
        public int x;
        public int y;
        public Location(int x,int y){
            this.x = x;
            this.y = y;
        }
    }

    private int dfs(int[][] grid,int x,int y){
        Stack<Location> stack = new Stack();
        stack.push(new Location(x,y));
        grid[x][y] = 0;
        int res = 0;
        while (!stack.isEmpty()){
            //从stack中弹出一个
            Location l = stack.pop();
            res ++;
            //把它的上下左右加入进来
            if (l.x > 0 && grid[l.x - 1][l.y] == 1) {
                stack.push(new Location(l.x - 1,l.y));
                grid[l.x - 1][l.y] = 0;
            }
            if (l.x < grid.length - 1 && grid[l.x + 1][l.y] == 1) {
                grid[l.x + 1][l.y] = 0;
                stack.push(new Location(l.x + 1,l.y));
            }
            if (l.y > 0 && grid[l.x][l.y - 1] == 1) {
                stack.push(new Location(l.x,l.y - 1));
                grid[l.x][l.y - 1] = 0;
            }
            if (l.y < grid[0].length -1 && grid[l.x][l.y + 1] == 1 ) {
                stack.push(new Location(l.x,l.y + 1));
                grid[l.x][l.y + 1] = 0;
            }

        }

        return res;
    }
    //572. 另一棵树的子树
    public boolean isSubtree(TreeNode root, TreeNode subRoot) {
        //递归思路：写一个方法，判断是否是包含另一个树。则是否存在子树的条件变成了:某一个节点与subRoot的根相同，且其子树包含其左右子树
        if (root == null) return false;
        if (subRoot == null) return true;
        if (root.val == subRoot.val){
            if (isContains(root.left,subRoot.left) && isContains(root.right,subRoot.right)){
                return true;
            }
        }
        return isSubtree(root.left,subRoot) || isSubtree(root.right,subRoot);

    }

    public boolean isContains(TreeNode root,TreeNode subRoot){
        if (subRoot == null) return root == null;
        if (root == null) return false;
        if (root.val != subRoot.val) return false;
        return isContains(root.left,subRoot.left)&& isContains(root.right,subRoot.right);
    }

    //117. 填充每个节点的下一个右侧节点指针 II
    public Node connect(Node root) {
        //思路：采用层序遍历
        LinkedList<Node> queue = new LinkedList();
        if (root == null){
            return root;
        }
        queue.add(root);
        int size = queue.size();
        while (size > 0){
            Node pre = queue.remove(0);
            if (pre.left != null) queue.add(pre.left);
            if (pre.right != null) queue.add(pre.right);
            for (int i = 1;i < size;i ++){
                Node node = queue.remove(0);
                if (node.left != null) queue.add(node.left);
                if (node.right != null) queue.add(node.right);
                pre.next = node;
                pre = node;
            }
            size = queue.size();
        }
        return root;
    }
}