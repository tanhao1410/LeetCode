fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1324. 竖直打印单词
    pub fn print_vertically(s: String) -> Vec<String> {
        //先求单词的最大长度
        let words = s.split(' ').map(|w| w.as_bytes()).collect::<Vec<_>>();
        let max_len = words.iter().map(|s| s.len()).max().unwrap_or(0);
        (0..max_len)
            .map(|i| {
                let mut res = String::new();
                for word in &words {
                    if word.len() > i {
                        res.push(word[i] as char);
                    } else {
                        res.push(' ');
                    }
                }
                res.trim_end()
                res.trim_end().to_string()
            })
            .collect()
    }
}

struct Solution;