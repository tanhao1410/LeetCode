class Solution {
    //556. 下一个更大元素 III
    public int nextGreaterElement(int n) {
        //思路：如果是非递增顺序的，则返回-1；
        //00187264，从后往前找到第一个递增的两个数字。然后将后面的最小的，但比本身大的数放到该位置，剩下的按递增顺序摆放
        int[] bits = new int[10];
        for(int i = 0;i < 10;i ++) bits[i] = -1;
        int i = 9;
        while( n > 0){
            bits[i--] = n % 10;
            n /= 10;
        }
        //从后往前
        for(i = 8;i >= 0;i --){
            if(bits[i] == -1) return -1;
            if (bits[i] < bits[i + 1]){
                //找到了要替换的位置了。
                int min = i + 1;
                for(int j = i + 1;j < 10;j ++){
                    if(bits[j] > bits[i] && bits[j] < bits[min]){
                        min = j;
                    }
                }
                //替换
                swap(bits,i,min);
                //后面的排序。
                for(int j = i + 1;j < 10;j ++){
                    for(int k = j + 1;k < 10;k ++){
                        //找出最小的
                        if(bits[k] < bits[j]){
                            swap(bits,j,k);
                        }
                    }
                }
                long res = 0l;
                for(int k = 0;k < 10;k ++){
                    if (bits[k] != -1){
                        res *= 10;
                        res += bits[k];
                    }
                }
                if (res > Integer.MAX_VALUE){
                    return -1;
                }else{
                    return (int)res;
                }
            }
        }
        return -1;
    }
    private void swap(int[] bits,int src,int dst){
        int temp = bits[src];
        bits[src] = bits[dst];
        bits[dst] = temp;
    }
}