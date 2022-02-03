fn main() {
    println!("Hello, world!");
    println!("{}", longest_palindrome("babbad".to_string()))
}

//516. 最长回文子序列
pub fn longest_palindrome_subseq(s: String) -> i32 {
    //思路：动态规划，dp[i][j]，s[i..j]中最长回文子串。
    //
    let mut dp = vec![vec![0; s.len()]; s.len()];
    let bytes = s.as_bytes();
    for i in (0..s.len()).rev() {
        dp[i][i] = 1;
        for j in i + 1..s.len() {
            //如果两头相等。则dp[i][j] = 2 + dp[i+1][j - 1]
            //如果两头不相等，则dp[i][j] = dp[i - 1][j] 或 dp[i][j - 1]
            if bytes[j] == bytes[i] {
                dp[i][j] = dp[i][j].max(2 + dp[i + 1][j - 1]);
            } else {
                dp[i][j] = dp[i][j - 1].max(dp[i + 1][j])
            }
        }
    }
    dp[0][s.len() - 1]
}

//5. 最长回文子串
pub fn longest_palindrome(s: String) -> String {
    //思路：不用动态规划，强行做，以某个字母为中心，找最大的回文子串
    //也可以用两个字母为中心。
    let mut max = 0;
    let mut index = 0;
    let bytes = s.as_bytes();
    let mut flag = false;
    for i in 0..s.len() {
        //以i为中心
        let mut j = 1;
        while i >= j && i + j < s.len() && bytes[i - j] == bytes[i + j] {
            j += 1;
        }
        if max < (j - 1) * 2 + 1 {
            index = i;
            max = (j - 1) * 2 + 1;
        }
    }
    //以两个字母为中心
    for i in 0..s.len() - 1 {
        //只有与下一个字母相同的情况才能以两个字母为中心
        if bytes[i] == bytes[i + 1] {
            let mut j = 1;
            while i >= j && i + j + 1 < s.len() && bytes[i - j] == bytes[i + j + 1] {
                j += 1;
            }
            if max < j * 2 {
                flag = true;
                max = j * 2;
                index = i;
            }
        }
    }
    let res_bytes;
    if flag {
        //以两个字母为中心
        res_bytes = &bytes[index - max / 2 + 1..index + 1 + max / 2];
    } else {
        res_bytes = &bytes[index - max / 2..index + 1 + max / 2]
    }
    String::from_utf8_lossy(res_bytes).to_string()
}

//136. 只出现一次的数字
pub fn single_number(nums: Vec<i32>) -> i32 {
    nums
        .iter()
        .fold(0, |p, &r| p ^ r)
}

//190. 颠倒二进制位
pub fn reverse_bits(x: u32) -> u32 {
    let mut bytes = vec![0; 32];
    for i in 0..32 {
        bytes[i] = (x & (1 << i)) >> i;
    }
    let mut res = 0;
    for i in 0..32 {
        res += (bytes[i] << (31 - i));
    }
    res
}
