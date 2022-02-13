fn main() {
    println!("Hello, world!");
}

//438. 找到字符串中所有字母异位词
pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
    //思路：先统计p中字母的数量。用一个数组统计
    // 在s中用一个窗口，每一次，进一个字母，出一个字母。
    let mut p_set = vec![0; 26];
    for &i in p.as_bytes() {
        p_set[(i - b'a') as usize] += 1;
    }
    let s_bytes = s.as_bytes();
    let mut s_set = vec![0; 26];
    for &i in s.as_bytes().iter().take(p.len()) {
        s_set[(i - b'a') as usize] += 1;
    }
    let is_equal = |p: &[i32], s: &[i32]| -> bool{
        p.iter().zip(s.iter()).all(|(i, j)| i == j)
    };
    let mut res = vec![];
    if s.len() < p.len() {
        return res;
    }
    for i in 0..=s.len() - p.len() {
        if i > 0 {
            s_set[(s_bytes[i - 1] - b'a') as usize] -= 1;
            s_set[(s_bytes[i - 1 + p.len()] - b'a') as usize] += 1;
        }
        //加入一个词，减去一个词
        if is_equal(&p_set, &s_set) {
            res.push(i as i32);
        }
    }
    res
}

//1189. “气球” 的最大数量
pub fn max_number_of_balloons(text: String) -> i32 {
    //balloon
    let bytes = text.as_bytes();
    let (mut a, mut b, mut n, mut l, mut o) = (0, 0, 0, 0, 0);
    for &byte in bytes {
        match byte {
            b'a' => a += 1,
            b'b' => b += 1,
            b'n' => n += 1,
            b'l' => l += 1,
            b'o' => o += 1,
            _ => {}
        }
    }
    a.min(b).min(n).min(l / 2).min(o / 2)
}
