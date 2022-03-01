class Solution {
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