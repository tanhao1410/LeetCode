class Solution {
    //1615. 最大网络秩
    public int maximalNetworkRank(int n, int[][] roads) {
        //将路径图改为每一个点连接的所有点 的图
        List<Integer>[] graph = new ArrayList[n];
        Set<Integer> set = new HashSet();
        for(int i = 0;i < n;i ++) graph[i] = new ArrayList();
        for(int[] road:roads){
            int src = road[0];
            int dst = road[1];
            set.add(101 * src + dst);
            //set.add(101 * dst + src);
            graph[src].add(dst);
            graph[dst].add(src);
        }
        int res = 0;
        //两个不相连的城市也可以
        for(int i = 0;i < n;i ++){
            //有多少与i相连的城市，
            int iCount = graph[i].size();
            for(int j = i + 1;j < n;j ++){
                //有多少与j相连的城市
                int jCount = graph[j].size();
                //i，j是否相连，如果相连，结果-1
                if (set.contains(i * 101 + j) || (set.contains(j * 101 + i))){
                    jCount --;
                }
                res = Math.max(res,iCount + jCount);
           }
        }
        return res;
    }
}