fn main() {
    println!("Hello, world!");
}

impl Solution {
    //521. 最长特殊序列 Ⅰ
    pub fn find_lu_slength(a: String, b: String) -> i32 {
        if a == b {
            -1
        } else {
            a.len().max(b.len()) as i32
        }
    }
}

struct Solution;