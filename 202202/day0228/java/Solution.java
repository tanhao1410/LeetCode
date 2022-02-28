class Solution {
    //989. 数组形式的整数加法
    public List<Integer> addToArrayForm(int[] num, int k) {
        //都转成数组吧
        List<Integer> num2List = new ArrayList();
        while (k > 0){
            num2List.add(k % 10);
            k /= 10;
        }
        // int[] num2 = new int[num2List.size()];
        // for(int i = 0;i < num2.length;i ++) num2[i] = num2List.get(num2.length - 1 - i);
        int index = 0;
        int flag = 0;
        LinkedList<Integer> res = new LinkedList();
        while(index < num2List.size() || index < num.length){
            int cur = flag;
            if (index < num2List.size() && index < num.length){
                cur += num[num.length - 1 - index]  + num2List.get(index);
            }else if (index < num2List.size()){
                cur += num2List.get(index);
            }else{
                cur += num[num.length - 1 - index];
            }
            res.add(0,cur % 10);
            flag = cur / 10;
            index ++;
        }
        if (flag == 1){
            res.add(0,1);
        }
        return res;
    }

    //1601. 最多可达成的换楼请求数目
    public int maximumRequests(int n, int[][] requests) {
        //思路：递归思路，总共就16个请求，对于每一个请求都有两个选择，要或不要。共2^16种情况。可以暴力解决
        return maximumRequests(requests,new boolean[requests.length],0,n);
    }
    private int canResolve(int[][] requests,boolean[] used,int n){
        int[] buildings = new int[n];
        int res = 0;
        for(int i = 0;i < used.length;i ++){
            if(used[i]){
                res ++;
                int[] request = requests[i];
                buildings[request[0]] --;
                buildings[request[1]] ++;
            }
        }
        for(int b : buildings){
            if(b != 0) return 0;
        }
        return res;
    }
    private int maximumRequests(int[][] requests,boolean[] used,int curIndex,int n){
        if (curIndex == requests.length) return canResolve(requests,used,n);
        //对于当前的请求，两种选择
        used[curIndex] = false;
        int res = maximumRequests(requests,used,curIndex + 1,n);
        used[curIndex] = true;
        res = Math.max(res,maximumRequests(requests,used,curIndex + 1,n));
        return res;
    }
}