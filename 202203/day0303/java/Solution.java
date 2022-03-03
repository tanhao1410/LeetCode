class Solution {
    //258. 各位相加
    public int addDigits(int num) {
        if(num <= 9) return num;
        int next = 0;
        while(num > 0){
            next += num % 10;
            num /= 10;
        }
        return addDigits(next);
    }
}