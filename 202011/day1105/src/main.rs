fn main() {
    println!("Hello, world!");
    let mut word_list = vec!["hot".to_string(),"dot".to_string(),"dog".to_string(),"lot".to_string(),"log".to_string(),"cog".to_string()];
    println!("{}",Solution::ladder_length("dot".to_string(),"cog".to_string(),word_list));
}

struct Solution {}

impl Solution {
    //每日一题：127. 单词接龙
    pub fn ladder_length(begin_word: String, end_word: String, mut word_list: Vec<String>) -> i32 {
        //先确定 结束的单词是否在其中，不在，直接返回0
        //准备一个map，key为word，value为距离end_word的距离，先 找最近的放入其中。value=1
        //将剩余的和begin_word 再依次变化至原来的map中value = 对应的+1
        //结束条件为word_list为空，或者，begin_word放在了map中
        if !word_list.contains(&begin_word) {
            word_list.push(begin_word.clone());
        }
        if !word_list.contains(&end_word) {
            return 0;
        }
        fn get_index(word: &String, list: &Vec<String>) -> usize {
            for i in 0..list.len() {
                if list[i] == *word {
                    return i;
                }
            }
            list.len()
        }

        fn is_one_step(word1: &String, word2: &String) -> bool {
            let mut count = 0;
            let u1 = word1.as_bytes();
            let u2 = word2.as_bytes();
            for i in 0..u1.len() {
                if u1[i] != u2[i] {
                    count += 1;
                }
            }
            count == 1
        }

        use std::collections::HashMap;
        use std::i32::MAX;
        let mut dp = HashMap::new();
        let end_word = word_list.remove(get_index(&end_word, &word_list));
        dp.insert(end_word, 1);
        let mut no_one  = false;
        while word_list.len() > 0 && !no_one{
            no_one = true;
            for i in 0..word_list.len() {
                let mut value = MAX;
                for j in &dp {
                    if is_one_step(&word_list[i], j.0) {
                        if value > j.1 + 1 {
                            value = j.1 + 1;
                        }
                    }
                }
                if value != MAX{
                    no_one=false;
                    dp.insert(word_list.remove(get_index(&word_list[i], &word_list)), value);
                    if let Some(v) = dp.get(&begin_word) {
                        return *v;
                    }
                    break;
                }
            }
        }
        0
    }
}