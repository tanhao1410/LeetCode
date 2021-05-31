fn main() {
    println!("Hello, world!");
}

//16. 最接近的三数之和
pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
    use std::cmp::Ordering;
    //求两数之间的距离
    let distance = |num1: i32| {
        match target > num1{
            true=>target - num1,
            false=>num1 - target
        }
    };
    //放置越界
    let mut res = i32::MAX - 10000;
    //先排序
    let mut nums = nums;
    nums.sort();
    //先从第一个数开始枚举，采用双指针，一个指向它后面的第一个，一个指向最后一个，如果和大于目标，那么p往后走，否则，q往前走。
    for i in 0..nums.len() {
        let (mut m, mut n) = (i + 1, nums.len() - 1);
        //目标，寻找最接近target - nums[i]的值
        while m < n {
            //三者之和
            let sum = nums[i] + nums[m] + nums[n];
            //比原来的距离小
            if distance(res) > distance(sum) {
                res = sum;
            }
            match sum.cmp(&target) {
                Ordering::Equal => return sum,
                Ordering::Greater => n -= 1,
                Ordering::Less => m += 1
            }
        }
    }
    res
}

//717. 1比特与2比特字符
pub fn is_one_bit_character(bits: Vec<i32>) -> bool {
    bits.iter().fold(0, |p, q| {
        match p {
            1 => *q + 2,
            _ => *q
        }
    }) == 0
}

