fn main() {
    println!("Hello, world!");
}

//128. 最长连续序列
pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
    use std::iter::FromIterator;
    let mut m = std::collections::HashSet::<i32>::from_iter(nums.clone().into_iter());
    let mut res = 0;
    for i in &nums {
        if !m.contains(&(i - 1)) {
            let mut cur_len = 1;
            while m.contains(&(i + cur_len)) {
                cur_len += 1;
            }
            res = res.max(cur_len );
        }
    }
    res
}

//1013. 将数组分成和相等的三个部分
pub fn can_three_parts_equal_sum(arr: Vec<i32>) -> bool {
    let sum: i32 = arr.iter().sum();
    sum % 3 == 0 &&
        arr.iter().take(arr.len() - 1).fold((false, false, 0), |r, &num| {
            match (r.0, r.1) {
                (true, false) => (true, num + r.2 == sum * 2 / 3, num + r.2),
                (true, true) => (true, true, num + r.2),
                _ => (r.2 + num == sum / 3, false, num + r.2),
            }
        }) == (true, true, sum - arr.last().unwrap())
}

//1744. 你能在你最喜欢的那天吃到你最喜欢的糖果吗？
pub fn can_eat(candies_count: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<bool> {
    ///思路：怎么样才能算迟到type类型的 1，type类型前面最多(不包括type)有多少糖果，每天最多吃多少， < day * cap
    ///  2.前面最少有多少糖果：每天至少吃一个糖果，那么前面的糖果数不能小于i
    ///
    let mut pre = 0u64;
    let candies_sum = candies_count.iter().map(|&i| {
        pre += i as u64;
        pre
    }).collect::<Vec<u64>>();
    queries.iter().map(|v| {
        let (fav_type, day, cap) = (v[0] as u64, v[1] as u64 + 1, v[2] as u64);
        (fav_type > 0 && candies_sum[fav_type as usize - 1] < day * cap && candies_sum[fav_type as usize] >= day)
            || (fav_type == 0 && candies_sum[fav_type as usize] >= day)
    }).collect()
}