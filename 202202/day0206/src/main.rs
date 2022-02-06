fn main() {
    println!("Hello, world!");
}

//1748. 唯一元素的和
pub fn sum_of_unique(nums: Vec<i32>) -> i32 {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    for num in nums {
        let entry = map.entry(num).or_insert(0);
        *entry += 1;
    }
    map
        .into_iter()
        .filter(|(_, v)| *v < 2)
        .map(|e| e.0)
        .sum()
}
