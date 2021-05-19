fn main() {
    println!("Hello, world!");
}

//409. 最长回文串
pub fn longest_palindrome(s: String) -> i32 {
    //偶数个的直接加进来，奇数个的+上，-1，最后+1
    let mut m = std::collections::HashMap::new();
    s.as_bytes().iter().for_each(|&c| {
        let mut count = m.entry(c).or_insert(0);
        *count += 1;
    });
    m.iter().fold((0, true), |(count, is_first_odd), (_, &v)|
        match (v & 1, is_first_odd) {
            (1, true) => (count + v, false),
            (1, _) => (count + v - 1, false),
            _ => (count + v, is_first_odd)
        }).0
}

//65. 有效数字
pub fn is_number(s: String) -> bool {
    //一个 小数 或者 整数
    // （可选）一个 'e' 或 'E' ，后面跟着一个 整数
    //先确定e/E,
    if s.contains(&['e', 'E'][..]) {
        let ss: Vec<&str> = s.split(&['e', 'E'][..]).collect();
        match ss.len() {
            2 => Self::is_small_number(ss[0].to_string())
                && Self::is_zhengshu(ss[1].to_string()),
            _ => false
        }
    } else {
        Self::is_small_number(s)
    }
}

//判断字符串是不是小数 +.8
pub fn is_small_number(mut s: String) -> bool {
    let ss: Vec<&str> = s.split('.').collect();
    match ss.len() {
        //一个的情况下，必须是数
        1 => Self::is_zhengshu(ss[0].to_string()),
        //前面是数，后面可以为空，或数
        2 => (Self::is_zhengshu(ss[0].to_string())
            && (Self::is_shuzi(ss[1].to_string()) || ss[1].len() == 0))
            || ((ss[0].len() == 0 || ss[0] == "+" || ss[0] == "-") && Self::is_shuzi(ss[1].to_string())),
        _ => false
    }
}

//判断字符串是不是整数
pub fn is_zhengshu(mut s: String) -> bool {
    if s.len() == 0 {
        return false;
    }
    match s.remove(0) {
        '+' | '-' => Self::is_shuzi(s),
        '0'...'9' => s.len() == 0 || Self::is_shuzi(s),
        _ => false,
    }
}
//判断是不是数字
pub fn is_shuzi(s: String) -> bool {
    for i in s.as_bytes() {
        if *i > b'9' || *i < b'0' {
            return false;
        }
    }
    s.len() != 0
}

//1738. 找出第 K 大的异或坐标值
pub fn kth_largest_value(matrix: Vec<Vec<i32>>, k: i32) -> i32 {
    //思路：v[m][n] = v[m-1][n] ^ v[m][n-1]^ v[m-1][n-1] ^ matrix[m][n]
    let mut v: Vec<Vec<i32>> = matrix;
    let mut nums = vec![];
    //一行一行求
    for i in 0..v.len() {
        for j in 0..v[0].len() {
            if i >= 1 {
                v[i][j] ^= v[i - 1][j]
            }
            if j >= 1 {
                v[i][j] ^= v[i][j - 1]
            }
            if i >= 1 && j >= 1 {
                v[i][j] ^= v[i - 1][j - 1]
            }
            nums.push(v[i][j]);
        }
    }


    nums.sort();
    nums[nums.len() - k as usize]
}