fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1071. 字符串的最大公因子
    pub fn gcd_of_strings(str1: String, str2: String) -> String {
        let is_child_str = |parent: &str, child: &str| {
            if parent.len() % child.len() != 0 {
                return false;
            }
            for i in 0..parent.len() / child.len() {
                if &parent[i * child.len()..(i + 1) * child.len()] != child {
                    return false;
                }
            }
            true
        };
        for i in 1..str1.len() + 1 {
            if str1.len() % i == 0
                && is_child_str(str1.as_str(), &str1[..str1.len() / i])
                && is_child_str(str2.as_str(), &str1[..str1.len() / i]) {
                return str1[..str1.len() / i].to_string();
            }
        }
        "".to_string()
    }
    //473. 火柴拼正方形
    pub fn makesquare(matchsticks: Vec<i32>) -> bool {
        let all_len = matchsticks.iter().sum::<i32>();
        if all_len % 4 == 0 && matchsticks.iter().all(|&l| l <= all_len / 4) {
            return Self::select_next(matchsticks.as_slice(), all_len / 4, &mut vec![0; 4], 0);
        }
        false
    }
    fn select_next(matchsticks: &[i32], res_len: i32, edges: &mut Vec<i32>, cur: usize) -> bool {
        if cur == matchsticks.len() {
            return edges.iter().all(|&l| l == res_len);
        }
        for i in 0..4 {
            if edges[i] + matchsticks[cur] <= res_len {
                edges[i] += matchsticks[cur];
                let inner_res = Self::select_next(matchsticks, res_len, edges, cur + 1);
                if inner_res {
                    return true;
                }
                edges[i] -= matchsticks[cur]
            }
        }
        false
    }
}