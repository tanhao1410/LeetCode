fn main() {
    println!("Hello, world!");
}

//1190. 反转每对括号间的子串
pub fn reverse_parentheses(s: String) -> String {
    //思路：如果，栈不为空，按顺序，进栈，遇到）出栈，直到一个（出来
    //如果栈为空，遇到字母，直接输出，遇到(开始进栈
    let mut stack = Vec::new();
    s.as_bytes().iter().fold(String::new(), |mut res, &c| {
        match (stack.len(), c) {
            (0, c) if c != b'(' => res.push(c as char),
            //栈为空
            (_, b')') => {
                //需要出栈，直到(,
                let mut temp = vec![];
                while let c = stack.pop().unwrap() {
                    if c == b'(' {
                        break;
                    }
                    temp.push(c);
                }
                match stack.len() {
                    0 => res.push_str(&String::from_utf8(temp).unwrap()),
                    _ => stack.append(&mut temp)
                }
            }
            //栈不为空，
            _ => stack.push(c)
        }
        res
    })
}