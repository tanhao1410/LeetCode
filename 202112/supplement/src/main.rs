fn main() {
    println!("Hello, world!");
}

struct Solution{}

impl Solution {
    //12-31-每日一题：507. 完美数
    pub fn check_perfect_number(num: i32) -> bool {
        (1..)
            .take_while(|&n| n * n <= num)
            .filter(|&n|num % n == 0 && num != n)
            .map(|n| n + num / n)
            .sum::<i32>()
            == 2 * num
    }

    //825. 适龄的朋友
    pub fn num_friend_requests(ages: Vec<i32>) -> i32 {
        let mut ages_num = vec![0; 121];
        ages.iter().for_each(|&age| {
            ages_num[age as usize] += 1;
        });
        //判断x是否向y发送消息
        let is_send_to = |x, y| -> bool{
            x >= y && y > x / 2 + 7 && (y <= 100 || x >= 100)
        };
        //年龄为x的所有人，发送的消息量
        let age_send_count = |x| -> i32{
            (1..x + 1).filter(|&y| is_send_to(x, y))
                .map(|y| ages_num[y])
                .sum::<i32>()
                * ages_num[x]
                - ages_num[x]
        };
        (1..121)
            .map(|i|ages_num[i])
            .filter(&i32::is_positive)
            .map(age_send_count)
            .filter(i32::is_positive)
            .sum()
    }
}
