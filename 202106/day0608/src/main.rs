fn main() {
    println!("Hello, world!");
}

//324. 摆动排序 II
pub fn wiggle_sort(nums: &mut Vec<i32>) {
    //摆动排序思路：先排序，然后将最最小的插入到每个中间。
    let mut temp = nums.clone();
    temp.sort();
    let mut k = temp.len() - 1;
    //奇数位插入大的
    for i in (1..temp.len()).step_by(2){
        nums[i] = temp[k];
        k -= 1;
    }
    for i in (0..temp.len()).step_by(2){
        nums[i] = temp[k];
        k -= 1;
    }
}

//239. 滑动窗口最大值
pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
    //思路：第一次的时候，找到最大值及所在的位置 k
    //每往前移动一个，带入一个新的数字，若该数字小于最大值， 最大值没有被踢掉，则依旧为原最大值
    //若数字小于最大值，最大值被踢掉 ？重新寻找最大值。
    //采用TreeMap
    #![feature(map_first_last)]
    use std::collections::BTreeMap;
    let mut tree_map = BTreeMap::new();
    //初始化TreeMap
    for i in 0..k as usize {
        *tree_map.entry(nums[i]).or_insert(0) += 1;
    }
    let mut res = vec![*tree_map.last_entry().unwrap().key()];

    for i in k as usize..nums.len() {
        //进入一个值，出来一个值
        let remove = nums[i - k as usize];
        let come = nums[i];
        *tree_map.entry(remove).or_insert(0) -= 1;
        if *tree_map.get(&remove).unwrap() < 1{
            tree_map.remove(&remove);
        }
        *tree_map.entry(come).or_insert(0) += 1;
        res.push(*tree_map.last_entry().unwrap().key())
    }
    res
}