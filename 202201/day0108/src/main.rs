#![feature(int_abs_diff)]

fn main() {
    for i in 1..17{
        let vec = Solution::gray_code_rebuild2(i);
        assert!(Solution::verify(vec));
    }
}

struct Solution{}
#[allow(dead_code)]
impl Solution {

    pub fn gray_code_old(n: i32) -> Vec<i32> {
        let mut res = vec![0];
        (1..=n).for_each(|i|res.append(&mut res.iter().rev().map(|&num| num + (1<<i - 1)).collect()));
        res
    }
    pub fn gray_code(n: i32) -> Vec<i32> {
        let mut res = vec![0;1<<n];
        for i in 1..=n{
            let start = 1 << i - 1;
            for index in 0..start {
                res[index + start] = start as i32+ res[start - 1 - index];
            }
        }
        res
    }

    pub fn gray_code_rebuild(n: i32) -> Vec<i32> {
        let mut res = vec![0;1<<n];
        (1..=n)
            .map(|e|1<<e-1)
            .for_each(|e|
                (0..e)
                    .for_each(|i|res[i + e] = e as i32 + res[e - 1 - i])
            );
        res
    }

    pub fn gray_code_rebuild2(n: i32) -> Vec<i32> {
        (1..=n)
            .map(|e|1<<e-1)
            .fold(vec![0; 1 << n], |mut res, e| {
                (0..e)
                    .for_each(|i|res[i + e] = e as i32 + res[e - 1 - i]);
                res
            })
    }

    pub fn verify(nums:Vec<i32>)->bool{
        //需要知道上一个元素
        nums.iter().fold((1,true),|(pre,res),e|{
            (*e,res && pre.count_ones().abs_diff(e.count_ones())==1)
        }).1
    }
}
