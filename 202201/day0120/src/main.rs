use std::arch::x86_64::_rdrand64_step;

fn main() {
    println!("Hello, world!");
}

//478. 在圆内随机生成点
struct Solution {
    diameter: f64,
    left_down_point: (f64, f64),
    center: (f64, f64),
}

impl Solution {
    fn new(radius: f64, x_center: f64, y_center: f64) -> Self {
        //平台自带的随机方式为随机一个u64,将随机到的u64转变为一个0-1之间的f64
        Self {
            diameter: 2.0 * radius,
            left_down_point: (x_center - radius, y_center - radius),
            center: (x_center, y_center),
        }
    }

    fn rand_f64(&self) -> f64 {
        let mut num = 0;
        unsafe {
            _rdrand64_step(&mut num)
        };
        (num as f64 / u64::MAX as f64) * self.diameter
    }

    fn is_out_circle(&self, (x, y): (f64, f64)) -> bool {
        (x - self.center.0).powi(2) + (y - self.center.1).powi(2)
            > self.diameter * self.diameter / 4.0
    }


    fn rand_point(&self) -> Vec<f64> {
        let x = self.rand_f64() + self.left_down_point.0;
        let y = self.rand_f64() + self.left_down_point.1;
        if self.is_out_circle((x, y)) {
            return self.rand_point();
        }
        vec![x, y]
    }
}

//740. 删除并获得点数
pub fn delete_and_earn(nums: Vec<i32>) -> i32 {
    //每一个数都对应有一个值，先求nums中的最大值以及最小值，然后依次遍历，转化成打家劫舍问题
    let min = *nums.iter().min().unwrap();
    let max = *nums.iter().max().unwrap();
    //创建一个数组，
    let mut nums2 = vec![0; (max - min) as usize + 1];
    for num in nums {
        nums2[(num - min) as usize] += num;
    }

    //即就一种数字
    if nums2.len() == 1 {
        return nums2[0];
    }

    //转化成了rob问题。
    let mut pre = nums2[0];
    let mut cur = nums2[1];

    for i in 2..nums2.len() {
        let temp = cur.max(pre + nums2[i]);
        pre = cur.max(pre);
        cur = temp;
    }
    pre.max(cur)
}

//213. 打家劫舍 II
pub fn rob2(nums: Vec<i32>) -> i32 {
    //区别在于是环形的 dp[i] = max{dp[i-2] + nums[i],dp[i - 1]} 限制是，dp[0] 与dp[len - 1]不可共存
    //思路1：转化为198，dp[n-1] 或  去除掉nums[0] 后，求dp[n-1]。大的为返回结果。
    let rob = |nums: &[i32]| {
        if nums.len() < 2 {
            return nums[0];
        }
        let mut pre1 = nums[0];
        let mut pre2 = nums[1];
        for i in 2..nums.len() {
            let temp = pre2.max(pre1 + nums[i]);
            pre1 = pre2.max(pre1);
            pre2 = temp;
        }
        pre1.max(pre2)
    };
    if nums.len() == 1 {
        return nums[0];
    }
    rob(&nums[1..]).max(rob(&nums[..nums.len() - 1]))
}

//198. 打家劫舍
pub fn rob(mut nums: Vec<i32>) -> i32 {
    //dp[i] 偷了第i家的情况下的最大值,则第i-1家不能偷。dp[i] = max{nums[i] + dp[i - 2] or dp[i - 3]}
    for i in 2..nums.len() {
        nums[i] += nums[i - 2].max(*nums.get(i - 3).unwrap_or(&0));
    }
    nums.into_iter().rev().take(2).max().unwrap()
}

//2029. 石子游戏 IX
pub fn stone_game_ix(stones: Vec<i32>) -> bool {
    //如果类型 0 的石子的个数为偶数，那么 Alice 获胜当且仅当类型 1和类型 2 的石子至少都有 1 个；
    //
    // 如果类型 0 的石子的个数为奇数，那么 Alice 获胜当且仅当
    // 「在没有类型 0 石子的情况下，Bob 获胜且原因不是因为所有石子都被移除」。
    // 对应到上面的分析即为「类型 1 的石子比类型 2 多超过 2个」或者「类型 2 的石子比类型 1 多超过 2 个」。
    //
    let stones = stones
        .into_iter()
        .fold((0, 0, 0i32), |(i, j, k), n| match n % 3 {
            0 => (i + 1, j, k),
            1 => (i, j + 1, k),
            _ => (i, j, k + 1)
        });
    match stones {
        (i, j, k) if i % 2 == 0 => j >= 1 && k >= 1,
        (_, j, k) => (j - k).abs() > 2
    }
}