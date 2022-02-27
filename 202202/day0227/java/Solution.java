class Solution {
    //202. 快乐数
    public boolean isHappy(int n) {
        HashSet<Integer> set = new HashSet();
        set.add(n);
        while(n > 0){
            n = getNum(n);
            if(set.contains(n)){
                break;
            }
            set.add(n);
        }
        return set.contains(1);
    }

    private int getNum(int n){
        int res = 0;
        while (n > 0){
            res += (n % 10) * (n % 10);
            n /= 10;
        }
        return res;
    }
}