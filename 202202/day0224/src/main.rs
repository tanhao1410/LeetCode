fn main() {
    println!("Hello, world!");
}

impl Solution {
    //20. 有效的括号
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![];
        for s in s.chars() {
            match s {
                '(' | '[' | '{' => stack.push(s),
                _ => {
                    if stack.is_empty() {
                        let pop = stack.pop().unwrap();
                        if (pop == '(' && s != ')') || (pop == '{' && s != '}') || (pop == '[' && s != ']') {
                            return false;
                        }
                    } else {
                        return false;
                    }
                }
            }
        }
        stack.len() == 0
    }
}

struct Solution;