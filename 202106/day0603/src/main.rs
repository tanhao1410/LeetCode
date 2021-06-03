fn main() {
    println!("Hello, world!");
}

//1160. 拼写单词
pub fn count_characters(words: Vec<String>, chars: String) -> i32 {
    let letter_count = |s: &String| -> std::collections::HashMap<char, i32>{
        s.chars().fold(std::collections::HashMap::new(), |mut m, c| {
            *m.entry(c).or_insert(0) += 1;
            m
        })
    };
    //思路：反向思维，从单词反向出发，看该单词是否能有chars中的字符拼成
    let m = letter_count(&chars);
    words.iter().filter(|&w| {
        letter_count(w).iter().fold(true, |p, (&w, &c)| {
            p && m.contains_key(&w) && m.get(&w).unwrap() >= &c
        })
    }).map(|s| s.len()).sum()
}

//525. 连续数组
pub fn find_max_length(nums: Vec<i32>) -> i32 {
    //记录1与0数量差 最靠前的位置
    let mut m = std::collections::HashMap::new();
    //1的总和，包括自身
    let mut nums = nums;
    let mut res = 0;
    for i in 0..nums.len() {
        let pre = match i {
            0 => 0,
            _ => nums[i - 1]
        };

        if nums[i] == 1 {
            nums[i] = pre + 1
        } else {
            nums[i] = pre - 1
        }
        //从第一个开始算起
        if nums[i] == 0 {
            res = res.max(i + 1);
        } else {
            match m.get(&nums[i]) {
                Some(&preIndex) => {
                    res = res.max(i - preIndex);
                }
                _ => {
                    m.insert(nums[i], i);
                }
            }
        }
    }
    res as i32
}