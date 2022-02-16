fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
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