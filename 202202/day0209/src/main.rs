fn main() {
    println!("Hello, world!");
}

//2006. 差的绝对值为 K 的数对数目
pub fn count_k_difference(nums: Vec<i32>, k: i32) -> i32 {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    for num in nums {
        let entry = map.entry(num).or_insert(0);
        *entry += 1;
    }
    let mut res = 0;
    for (kay, v) in &map {
        res += *map.get(&(*kay + k)).unwrap_or(&0) * *v;
        res += *map.get(&(*kay - k)).unwrap_or(&0) * *v;
    }
    res / 2
}