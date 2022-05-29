fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //468. 验证IP地址
    pub fn valid_ip_address(query_ip: String) -> String {
        let is_ip4_num = |num_str: &str| {
            if num_str.len() == 0 || (num_str.len() > 1 && num_str.starts_with('0')) {
                return false;
            }
            match num_str.parse::<i32>() {
                Ok(num) => num >= 0 && num <= 255,
                _ => false
            }
        };

        let is_ip6_num = |num_str: &str| {
            if num_str.len() >= 1 && num_str.len() <= 4 {
                if num_str.chars().all(|c| "0123456789abcdefABCDEF".contains(c)) {
                    return true;
                }
            }
            false
        };

        if query_ip.contains('.') {
            let mut split = query_ip.split('.').collect::<Vec<&str>>();
            if split.len() == 4 && split.into_iter().all(is_ip4_num) {
                return "IPv4".to_string();
            }
        } else if query_ip.contains(':') {
            let mut split = query_ip.split(':').collect::<Vec<&str>>();
            if split.len() == 8 && split.into_iter().all(is_ip6_num) {
                return "IPv6".to_string();
            }
        }
        "Neither".to_string()
    }
}