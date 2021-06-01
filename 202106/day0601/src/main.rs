fn main() {
    println!("Hello, world!");
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