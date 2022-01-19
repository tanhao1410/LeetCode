fn main() {
    println!("Hello, world!");
}

pub fn contains_nearby_duplicate2(nums: Vec<i32>, k: i32) -> bool {
    use std::collections::HashSet;
    //滑动窗口的大小为k,每一次都往里添加元素，如果添加元素后，窗口小于k，说明又重复的了。
    let mut set = HashSet::new();
    for i in 0..nums.len(){
        if set.len() > k as usize{
            set.remove(&nums[i -1 - k as usize]);
        }
        if set.contains(&nums[i]){
            return true;
        }
        set.insert(nums[i]);
    }
    false
}

//219. 存在重复元素 II
pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
    use std::collections::HashMap;
    let mut map = HashMap::new();
    nums
        .iter()
        .enumerate()
        .for_each(|(i, n)| {
            let vec = map.entry(n).or_insert(Vec::new());
            vec.push(i)
        });
    map
        .iter()
        .any(|(_, v)| {
            v.len() > 1
                && v.iter()
                .fold((-k - 1, false), |(pre, res), &n| (n as i32, res || n as i32 - pre <= k))
                .1
        })
}