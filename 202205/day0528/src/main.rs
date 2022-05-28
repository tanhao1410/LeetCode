fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1021. 删除最外层的括号
    pub fn remove_outer_parentheses(s: String) -> String {
        let mut left_count = 0;
        let bytes = s.as_bytes();
        let mut res = String::new();
        for i in 0..s.len() {
            if bytes[i] == b'(' {
                left_count += 1;
                if left_count > 1 {
                    res.push(bytes[i] as char);
                }
            } else {
                left_count -= 1;
                if left_count != 0 {
                    res.push(bytes[i] as char);
                }
            }
        }
        res
    }
}