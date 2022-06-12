fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    // 890.查找和替换模式
    pub fn find_and_replace_pattern(words: Vec<String>, pattern: String) -> Vec<String> {
        let pattern_bytes = pattern.as_bytes();
        //长度一致，对应关系一致
        let match_pattern = |word: &String| {
            let mut match_vec = vec![0u8; 128];
            let mut match_rev = vec![0u8; 128];
            let word_bytes = word.as_bytes();
            word_bytes.iter().zip(pattern_bytes.iter()).all(|(&i, &j)| {
                let res = match_vec[i as usize] == 0 || match_vec[i as usize] == j;
                let res2 = match_rev[j as usize] == 0 || match_rev[j as usize] == i;
                match_vec[i as usize] = j;
                match_rev[j as usize] = i;
                res && res2
            })
        };

        words.into_iter()
            .filter(match_pattern)
            .collect()
    }
}