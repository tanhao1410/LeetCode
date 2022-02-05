fn main() {
    println!("Hello, world!");
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