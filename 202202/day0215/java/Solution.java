class Solution {
    //797. 所有可能的路径
    public List<List<Integer>> allPathsSourceTarget(int[][] graph) {
        //思路：深度优先递归遍历
        List<Integer> pre = new ArrayList();
        pre.add(0);
        return allPath(graph,pre,0);
    }

    private List<List<Integer>> allPath(int[][] graph,List<Integer> pre,int start){
        List<List<Integer>> res = new ArrayList();
        //从start处开始
        if (start == graph.length - 1){
            //即走到了终点
            res.add(pre);
            return res;
        }
        //还没有走到终点，则看从starr能走到哪
        for(int dst : graph[start]){
            List<Integer> new_pre = new ArrayList();
            new_pre.addAll(pre);
            new_pre.add(dst);
            res.addAll(allPath(graph,new_pre,dst));
        }
        return res;
    }
    //547. 省份数量
    public int findCircleNum(int[][] isConnected) {
        int res = 0;
        Stack<Integer> stack = new Stack();
        //思路：用一个数组记录，哪些城市已经加入省份了。加入过的就不再统计了。用多个set表示每一个省份
        //每加入一个城市时，将其中相连的，未统计过的加入进来
        boolean[] citys = new boolean[isConnected.length];
        //从第一个城市开始
        for (int i = 0;i < isConnected.length;i ++){
            //未加入任何省份才进行统计
            if (!citys[i]){
                res ++;
                stack.push(i);
                citys[i] = true;
                while (stack.size() > 0){
                    Integer j = stack.pop();
                    //将j城相连的城市都加入进来
                    for (int k = 0;k < isConnected.length;k ++){
                        if (!citys[k] && isConnected[j][k] == 1){
                            stack.push(k);
                            citys[k] = true;
                        }
                    }
                }
            }
        }
        return res;
    }
}