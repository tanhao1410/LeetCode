fn main() {
    println!("Hello, world!");
}

//557. 反转字符串中的单词 III
pub fn reverse_words(s: String) -> String {
    s
        .split(" ")
        .map(|s| s.chars().rev().collect::<String>())
        .collect::<Vec<_>>()
        .join(" ")
}

//519. 随机翻转矩阵
pub fn reverse_string(s: &mut Vec<char>) {
    let mut i = 0;
    let mut j = s.len() - 1;
    while j > i {
        let temp = s[i];
        s[i] = s[j];
        s[j] = temp;
        i += 1;
        j -= 1;
    }
}