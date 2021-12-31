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

    //689. 三个无重叠子数组的最大和
    pub fn max_sum_of_three_subarrays(nums: Vec<i32>, k: i32) -> Vec<i32> {
        //思路：可以先简化问题，当只要一个无重叠时，从后往前，以每一个位置为起始的数组的最大和易求
        //当上升到两个数组时，从后往前，开始计算第一个数组的
        let (mut v1,mut v2,mut v3) = (vec![0;nums.len()],vec![0;nums.len()],vec![0;nums.len()]);
        //记录切割字符串时的下标
        let (mut l1,mut l2,mut l3) = (vec![0;nums.len()],vec![0;nums.len()],vec![0;nums.len()]);
        //先计算只划分一个数组情况
        v1[nums.len() - k as usize] = nums.iter().rev().take(k as usize).sum::<i32>();
        l1[nums.len() - k as usize] = nums.len() - k as usize;
        for i in (2 * k as usize ..nums.len() - k as usize).rev(){
            let sum = nums.iter().skip(i).take(k as usize).sum::<i32>();
            if sum >= v1[i + 1]{
                v1[i] = sum;
                l1[i] = i;
            }else{
                v1[i] = v1[i + 1];
                l1[i] = l1[i + 1];
            }
        }
        //计算当划分为两个数组时
        v2[nums.len() - 2 * k as usize] = nums.iter().skip(nums.len() - 2 * k as usize).take(k as usize).sum::<i32>() + v1[nums.len() - k as usize];
        l2[nums.len() - 2 * k as usize] = nums.len() - 2 * k as usize;
        for i in (k as usize..nums.len() - 2 * k as usize).rev(){
            let sum = nums.iter().skip(i).take(k as usize).sum::<i32>() + v1[i + k as usize];
            if sum >= v2[i + 1]{
                v2[i] = sum;
                l2[i] = i;
            }else{
                v2[i] = v2[i + 1];
                l2[i] = l2[i + 1];
            }
        }
        //计算当分为三个数组时
        v3[nums.len() - 3 * k as usize] = nums.iter().skip(nums.len() - 3 * k as usize).take(k as usize).sum::<i32>() + v2[nums.len() - 2 * k as usize];
        l3[nums.len() - 3 * k as usize] = nums.len() - 3 * k as usize;
        for i in (0..nums.len() - 3 * k as usize).rev(){
            let sum = nums.iter().skip(i).take(k as usize).sum::<i32>() + v2[i + k as usize];
            if sum >= v3[i + 1]{
                v3[i] = sum;
                l3[i] = i;
            }else{
                v3[i] = v3[i + 1];
                l3[i] = l3[i + 1];
            }
        }
        vec![l3[0] as i32,l2[l3[0] + k as usize] as i32,l1[l2[l3[0] + k as usize] + k as usize] as i32]
    }
}
