class Solution {
    public char nextGreatestLetter(char[] letters, char target) {
        if(letters[letters.length - 1] <= target) return letters[0];
        //二分查找的方式
        int l = 0;
        int r = letters.length - 1;
        int m = (l + r) / 2;
        while(l < r){
            if(letters[m] <= target){
                l = m + 1;
            }else{
                r = m;
            }
            m = (l + r)/2;
        }
        return letters[m];
    }
}