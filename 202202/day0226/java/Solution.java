class Solution {
    //459. 重复的子字符串
    public boolean repeatedSubstringPattern(String s) {
        return (s + s).substring(1,s.length() * 2 - 1).contains(s);
    }
}