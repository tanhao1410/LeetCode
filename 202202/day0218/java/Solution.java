class Solution {
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