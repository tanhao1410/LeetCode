fn main() {
    println!("Hello, world!");
}

//905. 按奇偶排序数组
pub fn sort_array_by_parity(nums: Vec<i32>) -> Vec<i32> {
    //思路：双指针法，第一个指针往后走，找到一个为奇数的点，第二个指针，从后往前走，找到一个为偶数的地方，
    let mut i = 0;
    let mut j = nums.len()-1;
    let mut nums = nums;
    while i < j{
        while i < nums.len() && nums[i] % 2 == 0{
            i += 1;
        }
        while j > 0 && nums[j] % 2== 1{
            j -=1;
        }
        if i < j{
            let temp = nums[i];
            nums[i]  = nums[j];
            nums[j] = temp;
        }
    }

    nums
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