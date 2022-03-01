class Solution {

    //968. 监控二叉树
    public int minCameraCover(TreeNode root) {
        //TreeNode =>
        HashMap<TreeNode,Pair[]>  map = new HashMap();
        return minCameraCover(root,false,false,map);

    }
    private int minCameraCover(TreeNode root,boolean need,boolean parent,HashMap<TreeNode,Pair[]> map){

        if (root == null){
            //必须要怎么办？
            if(need) return 1000;
            return 0;
        };

        if(map.containsKey(root)){
            Pair[] paris = map.get(root);
            //System.out.println("vv");
            for(Pair p : paris){
                if (p.need == need && p.parent == parent){
                    //
                    //System.out.println("找到了");
                    return p.res;
                }
            }
        }

        int res = 1 + minCameraCover(root.left,false,true,map) + minCameraCover(root.right,false,true,map);
        //不在当前节点布置
        if (!need){
            int temp = 0;
            if(parent){
                temp = minCameraCover(root.left,false,false,map) + minCameraCover(root.right,false,false,map);
            }else{
                temp = Math.min((minCameraCover(root.left,true,false,map) + minCameraCover(root.right,false,false,map)),
                (minCameraCover(root.left,false,false,map) + minCameraCover(root.right,true,false,map)));
            }
            res = Math.min(res,temp);
        }
       Pair[] pairs = map.getOrDefault(root,new Pair[4]);
       for(int i = 0;i < 4;i ++){
           if (pairs[i] == null) pairs[i] = new Pair(need,parent,res);
       }
       map.put(root,pairs);
        return res;
    }
    class Pair{
        public boolean need;
        public boolean parent;
        public int res = 0;
        public Pair(boolean need,boolean parent,int res){
            this.need = need;
            this.parent = parent;
            this.res = res;
        }
    }
    //1042. 不邻接植花
    public int[] gardenNoAdj(int n, int[][] paths) {
        //最多有三个邻居。先将数组paths转换下，转换成一个每个花坛对应的
        List<Integer>[] path = new ArrayList[n];
        for(int i = 0;i < n;i ++) path[i] = new ArrayList();
        for(int[] p : paths){
            path[p[0] - 1].add(p[1]-1);
            path[p[1] - 1].add(p[0] - 1);
        }
        int[] res = new int[n];
        //总共四种花，优先种1，与它相邻的种2
        Stack<Integer> stack = new Stack();
        for(int i = 0;i < n;i ++){
            if(res[i] == 0){
                stack.push(i);
                while(stack.size() > 0){
                    int cur = stack.pop();
                    int[] temp = new int[4];
                    //应该种什么花呢？看它周围都是种了什么
                    for(int arround : path[cur]){
                        if (res[arround] != 0){
                            //说明有花了,标记下哪种花种了
                            temp[res[arround] - 1] = 1;
                        }else{
                            stack.push(arround);
                        }
                    }
                    for(int j = 0;j < 4;j ++){
                        if (temp[j] == 0){
                            res[cur] = j + 1;
                            break;
                        }
                    }
                }
            }
        }

        return res;
    }
    //58. 最后一个单词的长度
    public int lengthOfLastWord(String s) {
        //不使用trim，split等方式
        int end = s.length() - 1;
        //最后一个单词字母所在的位置
        while(end >= 0 && s.charAt(end)==' ') end --;
        int start = end;
        while (start >= 0 && s.charAt(start) != ' ') start --;
        return end - start;
    }
}