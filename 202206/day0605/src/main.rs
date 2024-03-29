use rand::random;

fn main() {
    println!("Hello, world!");
}


impl Solution {
    //1528. 重新排列字符串
    pub fn restore_string(s: String, indices: Vec<i32>) -> String {
        let mut zip = s.chars().zip(indices.into_iter()).collect::<Vec<_>>();
        zip.sort_by_key(|(_, i)| *i);
        zip.into_iter().map(|e| e.1).collect()
    }
    //1417. 重新格式化字符串
    pub fn reformat(s: String) -> String {
        let letters = s.chars().filter(|c| c.is_ascii_alphabetic()).collect::<Vec<_>>();
        let nums = s.chars().filter(|c| !c.is_ascii_alphabetic()).collect::<Vec<_>>();
        if (letters.len() as i32 - nums.len() as i32).abs() > 1 {
            return "".to_string();
        }
        let mut res = letters
            .iter()
            .zip(nums.iter())
            .flat_map(|(i, j)| vec![i, j].into_iter())
            .collect::<String>();
        if letters.len() > nums.len() {
            res.push(*letters.last().unwrap());
        } else if nums.len() > letters.len() {
            res.insert(0, *nums.last().unwrap());
        }
        res
    }
    //375. 猜数字大小 II
    pub fn get_money_amount(n: i32) -> i32 {
        let n = n + 1;
        let mut dp = vec![vec![0; n as usize]; n as usize];
        for start in (0..n as usize).rev() {
            for end in start + 1..n as usize {
                let mut max_spend = i32::MAX;
                for i in start..=end {
                    let mut cur_spend = 0;
                    //选择了i后，有三种可能性
                    if i > start {
                        cur_spend = cur_spend.max(i as i32 + dp[start][i - 1])
                    }
                    if i < end {
                        cur_spend = cur_spend.max(i as i32 + dp[i + 1][end])
                    }
                    max_spend = max_spend.min(cur_spend);
                }
                dp[start][end] = max_spend;
            }
        }
        dp[0][n as usize - 1]
    }

    //478. 在圆内随机生成点
    fn new(radius: f64, x_center: f64, y_center: f64) -> Self {
        Self { radius, x_center, y_center }
    }

    fn rand_point(&self) -> Vec<f64> {
        let in_circle = |point: (f64, f64)| {
            self.radius.powi(2) >= (point.0 - self.x_center).powi(2) + (point.1 - self.y_center).powi(2)
        };
        let (x, y) = (self.x_center - self.radius, self.y_center - self.radius);
        let rand_point = (random::<f64>() * self.radius * 2.0 + x, random::<f64>() * self.radius * 2.0 + y);
        if in_circle(rand_point) {
            return vec![rand_point.0, rand_point.1];
        }
        self.rand_point()
    }
}

struct Solution {
    radius: f64,
    x_center: f64,
    y_center: f64,
}