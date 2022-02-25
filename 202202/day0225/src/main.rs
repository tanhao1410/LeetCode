fn main() {
    println!("Hello, world!");
}

impl Solution {
    //537. 复数乘法
    pub fn complex_number_multiply(num1: String, num2: String) -> String {
        // a b c d
        let transfer = |num: &str| {
            let i = num.find('+').unwrap();
            (num[..i].parse::<i32>().unwrap(), num[i + 1..num1.len() - 1].parse::<i32>().unwrap())
        };
        let (a, b) = transfer(&num1);
        let (c, d) = transfer(&num2);
        // let mut res = String::new();
        // res.push_str(&(a * c - b * d).to_string());
        // res.push('+');
        // res.push_str(&(a * d + b * c).to_string());
        // res.push('i');
        // res
        format!("{}+{}i", (a * c - b * d), (a * d + b * c))
    }
}

struct Solution;
