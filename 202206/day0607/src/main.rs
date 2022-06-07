fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //2119. 反转两次的数字
    pub fn is_same_after_reversals(num: i32) -> bool {
        num == 0 || num % 10 != 0
    }

    //875. 爱吃香蕉的珂珂
    pub fn min_eating_speed(piles: Vec<i32>, h: i32) -> i32 {
        let mut min = 1;
        let mut max = *piles.iter().max().unwrap();
        let mut mid = (max + min) / 2;
        while min < max {
            //用max肯定可以，看mid够不够
            let need_times = piles.iter().map(|&i| (i + mid - 1) / mid).sum::<i32>();
            if need_times <= h {
                //说明mid也算大的了，可以继续减小
                max = mid;
            } else {
                //说明，mid不够，需要加大
                min = mid + 1;
            }
            mid = (max + min) / 2;
        }
        min
    }
}