use std::collections::BTreeMap;

fn main() {
    println!("Hello, world!");
    println!("{:?}", Solution::subsets_with_dup(vec![1, 2, 2]));
}

struct Solution;

impl Solution {

    //90. 子集 II
    pub fn subsets_with_dup(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut map = BTreeMap::new();
        for &num in &nums{
            let entry = map.entry(num).or_insert(0);
            *entry += 1;
        }
        let nums = map.into_iter().map(|(k, v)| (k, v)).collect::<Vec<(i32, i32)>>();
        Self::subsets_with_dup2(vec![],&nums)
    }

    fn subsets_with_dup2(mut pre:Vec<Vec<i32>>, nums:&[(i32,i32)]) ->Vec<Vec<i32>>{
        if nums.len() == 0{
            return pre;
        }
        let mut new_pre = vec![];
        let (k,v ) = nums[0];
        if pre.len() == 0{
            for i in 0..=v{
                new_pre.push(vec![k;i as usize]);
            }
        }else{
            for item in &pre{
                for i in 1..=v{
                    let mut new_item = item.clone();
                    new_item.append(&mut vec![k;i as usize]);
                    new_pre.push(new_item);
                }
            }
            new_pre.append(&mut pre);
        }
        Self::subsets_with_dup2(new_pre,&nums[1..])
    }

    //78. 子集
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        Self::subsets2(vec![], &nums)
    }

    fn subsets2(mut pre: Vec<Vec<i32>>, nums: &[i32]) -> Vec<Vec<i32>> {
        if nums.len() == 0 {
            return pre;
        }
        let mut new_pre = vec![];
        if pre.len() == 0 {
            new_pre.push(vec![]);
            new_pre.push(vec![nums[0]]);
        } else {
            for item in &pre{
                let mut new_item = item.clone();
                new_item.push(nums[0]);
                new_pre.push(new_item);
            }
            new_pre.append(&mut pre);
        }
        Self::subsets2(new_pre, &nums[1..])
    }
}