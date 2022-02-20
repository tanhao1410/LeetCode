class Solution {
    //717. 1比特与2比特字符
    public boolean isOneBitCharacter(int[] bits) {
        for(int i = 0;;i ++){
            if (i >= bits.length - 1) return i == bits.length-1;
            if (bits[i] == 1) i ++;
        }
    }
}