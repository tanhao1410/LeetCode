class Solution {
    //56. 合并区间
    public int[][] merge(int[][] intervals) {
        List<int[]> res = new ArrayList();
        Arrays.sort(intervals,(e1,e2)->{
            return e1[0] - e2[0];
        });
        int[] pre_interval = new int[]{intervals[0][0],intervals[0][1]};
        for (int i = 1;i < intervals.length;i ++){
            int[] cur_interval = new int[]{intervals[i][0],intervals[i][1]};
            if(cur_interval[0] > pre_interval[1]){
                //无重叠
                res.add(pre_interval);
                pre_interval = cur_interval;
            }else if (cur_interval[1] > pre_interval[1]){
                pre_interval[1] = cur_interval[1];
            }
        }
        res.add(pre_interval);
        return res.toArray(new int[res.size()][2]);
    }
    //22. 括号生成
    public List<String> generateParenthesis(int n) {
        //思路：每一次可以选择(或)。有时只可以选一种。递归的思路解决吧
        return generateParenthesis(n,n,new StringBuilder());
    }

    private List<String> generateParenthesis(int left,int right,StringBuilder pre){
        List<String> res = new ArrayList();
        if (left == 0 || right == 0){
            for(int i = 0;i < right;i ++) pre.append(")");
            res.add(pre.toString());
            return res;
        }
        //可以选择left,或right
        if (left < right){
            //可以都进行选择
            StringBuilder pre2 = new StringBuilder(pre);
            res.addAll(generateParenthesis(left ,right - 1,pre2.append(")")));
        }
        res.addAll(generateParenthesis(left - 1,right,pre.append("(")));
        return res;
    }
    //1306. 跳跃游戏 III
    public boolean canReach(int[] arr, int start) {
        //广度优先策论。访问到有0的即可
        boolean[] readed = new boolean[arr.length];
        List<Integer> queue = new ArrayList();
        queue.add(start);
        readed[start] = true;
        while(queue.size() > 0){
            int len = queue.size();
            for(int i = 0;i < len;i ++){
                int cur = queue.remove(0);
                if (arr[cur] == 0) return true;
                //可以往前走，也可以往后走
                int[] dirs = new int[]{cur + arr[cur],cur - arr[cur]};
                for(int dir : dirs){
                    if (dir >= 0 && dir < arr.length && !readed[dir]) {
                        queue.add(dir);
                        readed[dir] = true;
                    }
                }
            }
        }
        return false;
    }

    //1791. 找出星型图的中心节点
    public int findCenter(int[][] edges) {
        //广度优先策略。从最外层开始遍历
        int n = edges.length + 1;
        boolean[] readed = new boolean[n];
        LinkedList<Integer> queue = new LinkedList();
        List<Integer>[] maps = new LinkedList[n];
        for(int i = 0;i < n;i ++)maps[i] = new LinkedList();
        for(int[] edge :edges){
            maps[edge[0] - 1].add(edge[1] - 1);
            maps[edge[1] - 1].add(edge[1] - 1);
        }
        //找到只有一个连接的点，加入到队列中
        for(int i = 0;i < n;i ++){
            if (maps[i].size() == 1){
                queue.add(i);
                readed[i] =true;
            }
        }

        int res = 0;
        for(int l = queue.size(); l > 0;l = queue.size()){
            for(int i = 0;i < l;i ++){
                Integer node = queue.remove(0);
                res = node;
                for(Integer j : maps[node]){
                    if (!readed[j]){
                        queue.add(j);
                        readed[j] = true;
                    }
                }
            }
        }
        return res + 1;
    }

    //240. 搜索二维矩阵 II
    public boolean searchMatrix(int[][] matrix, int target) {
        //从右上角开始，往左下角开始走，走不动了，返回false
        int x = 0;
        int y = matrix[0].length - 1;
        while(x < matrix.length && y >= 0){
            if (matrix[x][y] == target) return true;
            if (matrix[x][y] > target){
                y --;
            }else{
                x ++;
            }
        }
        return false;
    }

    //59. 螺旋矩阵 II
    public int[][] generateMatrix(int n) {
        int[][] res = new int[n][n];
        int num = 1;
        int rank = 0;
        int x = 0;
        int y = 0;
        while(num <= n * n){
            //往右边走
            //res[x][y] = num++;
            while(y < n - rank) res[x][y++] = num++;
            y --;
            //往下走
            while( x + 1 < n - rank) res[++x][y] = num++;
            //往左走
            while( y - 1 >= rank) res[x][--y] = num++;
            //往上走
            while (x - 1 >= 1 + rank) res[--x][y] = num ++;
            //开始新一轮循环
            rank ++;
            y ++;
        }
        return res;
    }
}