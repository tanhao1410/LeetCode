fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1047. 删除字符串中的所有相邻重复项
    pub fn remove_duplicates(s: String) -> String {
        if s.len() == 0 {
            return s;
        }
        let bytes = s.as_bytes();
        let mut res = vec![bytes[0]];
        let mut flag = false;
        for i in bytes.into_iter().skip(1) {
            if i == res.last().unwrap_or(&b'*') {
                res.remove(res.len() - 1);
                flag = true;
            } else {
                res.push(*i);
            }
        }
        if flag {
            return Self::remove_duplicates(String::from_utf8(res).unwrap());
        }
        String::from_utf8(res).unwrap()
    }
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