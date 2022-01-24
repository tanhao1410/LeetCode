fn main() {
    println!("Hello, world!");
}

//519. 随机翻转矩阵
pub fn reverse_string(s: &mut Vec<char>) {
    let mut i = 0;
    let mut j = s.len() - 1;
    while j > i{
        let temp = s[i];
        s[i] = s[j];
        s[j] = temp;
        i += 1;
        j -= 1;
    }
}