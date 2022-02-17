fn main() {
    println!("Hello, world!");
}

impl Solution {
    //47. 全排列 II
    pub fn permute_unique(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort_unstable();
        //选择一个数，然后选择第二个数，如果接下来的数和自己相同，跳过该数。
        Self::select_next_num(nums, vec![])
    }

    fn select_next_num(nums: Vec<i32>, pre: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        if nums.len() == 0 {
            res.push(pre);
            return res;
        }
        let mut pre_num = 11;
        for i in 0..nums.len() {
            if nums[i] != pre_num {
                //选择当前数字
                let mut new_pre = pre.clone();
                new_pre.push(nums[i]);
                let mut new_nums = vec![0; nums.len() - 1];
                for j in 0..i {
                    new_nums[j] = nums[j];
                }
                for j in i + 1..nums.len() {
                    new_nums[j - 1] = nums[j];
                }
                res.append(&mut Self::select_next_num(new_nums, new_pre));
                pre_num = nums[i];
            }
        }
        res
    }
}