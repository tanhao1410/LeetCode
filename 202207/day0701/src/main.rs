fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //241. 为运算表达式设计优先级
    pub fn diff_ways_to_compute(expression: String) -> Vec<i32> {
        let bytes = expression.as_bytes();
        let mut res = vec![];
        for i in 0..bytes.len() {
            if bytes[i] > b'9' || bytes[i] < b'0' {
                //找到操作符了
                let pre_res = Self::diff_ways_to_compute(String::from_utf8_lossy(&bytes[..i]).into_owned());
                let pro_res = Self::diff_ways_to_compute(String::from_utf8_lossy(&bytes[i + 1..]).into_owned());
                for a in &pre_res {
                    for b in &pro_res {
                        res.push(Self::get_oper(bytes[i])(*a, *b))
                    }
                }
            }
        }
        if res.len() == 0 {
            res.push(expression.parse::<i32>().unwrap());
        }
        res
    }

    fn get_oper(oper: u8) -> fn(i32, i32) -> i32 {
        match oper {
            b'+' => |a: i32, b: i32| a + b,
            b'-' => |a: i32, b: i32| a - b,
            _ => |a: i32, b: i32| a * b
        }
    }
}