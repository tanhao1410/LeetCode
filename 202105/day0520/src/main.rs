fn main() {
    println!("Hello, world!");
}

//448. 找到所有数组中消失的数字
pub fn find_disappeared_numbers(mut nums: Vec<i32>) -> Vec<i32> {
    let BASE = nums.len() + 1;

    for i in 0..nums.len(){
        let index = nums[i] % BASE as i32 - 1 ;
        nums[index as usize] += BASE as i32
    }

    let mut res =vec![];
    for i in 0..nums.len(){
        if nums[i] < BASE as i32 {
            res.push(i as i32 + 1);
        }
    }

    res
}

//442. 数组中重复的数据
pub fn find_duplicates(mut nums: Vec<i32>) -> Vec<i32> {
    const BASE:i32 = 1_000_000_000;
    for i in 0..nums.len(){
        let i = nums[i];
        nums[(i % BASE) as usize - 1] += BASE
    }
    //nums.into_iter().filter(|&n| n / BASE == 2).map(|i| i % BASE).collect()
    nums.iter().enumerate().filter_map(|(i,&n)|match n/BASE {
        2=>Some(i as i32 + 1),
        _=>None
    }).collect()
}

//692. 前K个高频单词
pub fn top_k_frequent(words: Vec<String>, k: i32) -> Vec<String> {
    // let mut dic = std::collections::HashMap::new();
    // words.iter().for_each(|&w| *dic.entry(w).or_insert(0)+= 1);
    let mut dic = words.iter()
        .fold(&mut std::collections::HashMap::new(), |mut m, w| {
            *m.entry(w).or_insert(0) += 1;
            m
        })
        .iter()
        .map(|(&k, &v)| (k, v))
        .collect::<Vec<_>>();
    dic.sort_by(|pre, pro|
        match pro.1.partial_cmp(&pre.1).unwrap() {
            std::cmp::Ordering::Equal => pre.0.partial_cmp(&pro.0).unwrap(),
            _ => pro.1.partial_cmp(&pre.1).unwrap()
        }
    );
    dic.iter()
        .take(k as usize)
        .map(|wc| wc.0.clone())
        .collect::<Vec<_>>()
}