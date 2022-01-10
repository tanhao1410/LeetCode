fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    pub fn is_additive_number(num: String) -> bool {
        (1..20)
            .any(|i|
                (1..20)
                    .filter(|&j| i + j < num.len())
                    .filter(|&j| !num.starts_with("0") || i >= 1)
                    .filter(|&j| j <= 1 || num.as_bytes()[i] != b'0')
                    .map(|j| i + j)
                    .map(|j| ((&num.as_str()[..i]).parse::<i64>().unwrap(), (&num.as_str()[i..j]).parse::<i64>().unwrap(), j))
                    .map(|(x, y, j)| (x, y, num.as_str()[j..].to_string(), num.as_str()[..j].to_string()))
                    .filter(|(x, y, sum, _)| sum.starts_with(&(x + y).to_string()))
                    .map(|(mut x, mut y, _, mut res)|{
                        while res.len() < num.len() {
                            res.push_str(&(x + y).to_string());
                            y += x;
                            x = y - x;
                        }
                        res
                    })
                    .any(|res| res.eq(&num))
            )
    }
}
