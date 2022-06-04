fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //779. 第K个语法符号
    pub fn kth_grammar(n: i32, k: i32) -> i32 {
        if n == 1 {
            return 0;
        }
        let two_n_2 = 2_i32.pow(n as u32 - 2);
        if k <= two_n_2 {
            return Self::kth_grammar(n - 1, k);
        }
        return Self::kth_grammar(n - 1, two_n_2 * 2 - k + 1) ^ (1 - (n % 2));
    }
    //1572. 矩阵对角线元素的和
    pub fn diagonal_sum(mat: Vec<Vec<i32>>) -> i32 {
        let n = mat.len();
        let res = (0..n).map(|i| mat[i][i] + mat[i][n - 1 - i]).sum::<i32>();
        match n % 2 == 0 {
            true => res,
            _ => res - mat[n / 2][n / 2]
        }
    }
    //929. 独特的电子邮件地址
    pub fn num_unique_emails(emails: Vec<String>) -> i32 {
        use std::collections::HashSet;
        emails.into_iter()
            .map(|email| {
                let mut split = email.split('@');
                split.next().unwrap().split('+').next().unwrap().replace('.', "") + "@" + split.next().unwrap().as_ref()
            })
            .collect::<HashSet<_>>()
            .len() as i32
    }
}