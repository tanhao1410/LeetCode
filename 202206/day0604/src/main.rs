fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
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