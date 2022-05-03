use std::cmp::Ordering;

fn main() {
    println!("Hello, world!");
}

impl Solution {
    //1295. 统计位数为偶数的数字
    pub fn find_numbers(nums: Vec<i32>) -> i32 {
        nums
            .into_iter()
            .map(|mut num| {
                let mut res = 0;
                while num > 0 {
                    res += 1;
                    num /= 10;
                }
                res
            })
            .filter(|num| num % 2 == 0)
            .count() as i32
    }
    //1281. 整数的各位积和之差
    pub fn subtract_product_and_sum(mut n: i32) -> i32 {
        let mut nums = vec![];
        while n != 0 {
            nums.push(n % 10);
            n /= 10;
        }
        nums.iter().fold(1, |p, num| p * *num) - nums.iter().sum::<i32>()
    }
    //937. 重新排列日志文件
    pub fn reorder_log_files(mut logs: Vec<String>) -> Vec<String> {
        let is_num_log = |s: &String| {
            let last = s.as_bytes()[s.len() - 1];
            last >= b'0' && last <= b'9'
        };
        logs.sort_by(|s1, s2| {
            //如果两者都是数字日志，则相等，如果有一个是的，则是的小。如果两个都不是呢，判断
            match (is_num_log(s1), is_num_log(s2)) {
                (true, true) => Ordering::Equal,
                (true, false) => Ordering::Greater,
                (false, true) => Ordering::Less,
                (_, _) => {
                    //内容相同，比标志，内容不同，比内容
                    if s1.split(' ').skip(1).cmp(s2.split(' ').skip(1)) == Ordering::Equal {
                        s1.split(' ').take(1).cmp(s2.split(' ').take(1))
                    } else {
                        s1.split(' ').skip(1).cmp(s2.split(' ').skip(1))
                    }
                }
            }
        });
        logs
    }
}