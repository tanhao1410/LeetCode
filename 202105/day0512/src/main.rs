fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {
    //260. 只出现一次的数字 III
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        //有两个元素只出现一次，其余的出现两次。
        //思路：把数组分为两部分，一部分包含一个，其中相同的数包含在同一部分中
        //将所有的数进行异或，得到的结果是两个不同的数异或的值
        let two = nums.iter().fold(0, |i, j| i ^ j);
        //从异或的值中找到某一位为1，以此进行分割数组
        let mut v = two;
        let mut master = -1;
        while v.count_ones() > 1 {
            v &= master;
            master <<= 1;
        }
        let res1 = nums.iter().filter(|&&i| i & v == 0).fold(0, |i, &j| i ^ j);
        vec![res1, res1 ^ two]
    }
}
