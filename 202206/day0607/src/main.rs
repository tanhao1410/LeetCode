fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1025. 除数博弈
    pub fn divisor_game(n: i32) -> bool {
        //可以选择哪些数，或者这样说，选择哪些数，爱丽丝会赢。 1的时候，爱丽丝输，2的时候赢，3 的时候输，
        //4的时候，可以选择的数有 1 2 ，选择1 的话，
        //爱丽丝可以赢的数有：
        let mut win_num = vec![2];
        'outer: for i in 3..1000 {
            for j in 1..i {
                if i % j == 0 && !win_num.contains(&(i - j)) {
                    //看i - j是否在里面，如果不在里面，则必赢
                    win_num.push(i);
                    continue 'outer;
                }
            }
        }
        win_num.contains(&n)
    }

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