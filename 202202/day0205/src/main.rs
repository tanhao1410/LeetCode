fn main() {
    println!("Hello, world!");
}

//1143. 最长公共子序列
pub fn longest_common_subsequence(text1: String, text2: String) -> i32 {
    //思路：dp[i][j] , text1[..i]与text2[..j]的最长公共子序列
    // dp[0][j] = 0 或1 ， if dp[0][j] == 1 ,则 dp[0][j++] = 1
    // if text1[i] == text2[j] ,则 的dp[i][j] = dp[i - 1][j-1] + 1
    let bytes1 = text1.as_bytes();
    let bytes2 = text2.as_bytes();
    let mut dp = vec![vec![0; text2.len()]; text1.len()];
    //求两者第一个字母相遇的位置
    let mut l1 = 0;
    while l1 < text2.len() && bytes2[l1] != bytes1[0] {
        l1 += 1;
    }
    for i in l1..text2.len() {
        dp[0][i] = 1;
    }
    let mut l2 = 0;
    while l2 < text1.len() && bytes1[l2] != bytes2[0] {
        l2 += 1;
    }
    for i in l2..text1.len() {
        dp[i][0] = 1;
    }

    for i in 1..bytes1.len() {
        for j in 1..bytes2.len() {
            if bytes1[i] == bytes2[j] {
                dp[i][j] = dp[i - 1][j - 1] + 1;
            } else {
                dp[i][j] = dp[i - 1][j].max(dp[i][j - 1]);
            }
        }
    }

    dp[text1.len() - 1][text2.len() - 1]
}

//392. 判断子序列
pub fn is_subsequence(s: String, t: String) -> bool {
    //思路：双指针法，对于任意一个s中的字母，要在t中找到，找到后，s指针往前走
    let s_bytes = s.as_bytes();
    let t_byts = t.as_bytes();
    let mut s = 0;
    let mut t = 0;
    while s < s_bytes.len() && t < t_byts.len() {
        while t < t_byts.len() && s_bytes[s] != t_byts[t] {
            t += 1;
        }
        if t < t_byts.len() {
            s += 1;
            t += 1;
        }
    }
    s == s_bytes.len()
}