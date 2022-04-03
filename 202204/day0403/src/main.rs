fn main() {
    println!("Hello, world!");
}

impl Solution {
    //744. 寻找比目标字母大的最小字母
    pub fn next_greatest_letter(letters: Vec<char>, target: char) -> char {
        if *letters.last().unwrap() <= target {
            return letters[0];
        }
        let mut l = 0;
        let mut r = letters.len() - 1;
        let mut m = (l + r) / 2;
        while (l < r) {
            if letters[m] <= target {
                l = m + 1;
            } else {
                r = m;
            }
        }
        letters[m]
    }
}

struct Solution;