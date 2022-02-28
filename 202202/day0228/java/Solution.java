class Solution {

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