fn main() {
    println!("Hello, world!");
}

impl Solution {
    //剑指 Offer II 036. 后缀表达式
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        use std::ops::Add;
        use std::ops::Mul;
        use std::ops::Div;
        use std::ops::Sub;
        let mut stack = vec![];
        let opera = |stack: &mut Vec<i32>, op: fn(i32, i32) -> i32| {
            let mut a = stack.pop().unwrap();
            let mut b = stack.pop().unwrap();
            stack.push(op(b, a));
        };
        for token in tokens {
            //如果是数就进栈
            match token.as_str() {
                "+" => opera(&mut stack, i32::add),
                "-" => opera(&mut stack, i32::sub),
                "*" => opera(&mut stack, i32::mul),
                "/" => opera(&mut stack, i32::div),
                _ => stack.push(token.parse().unwrap())
            }
        }
        stack.pop().unwrap()
    }
    //451. 根据字符出现频率排序
    pub fn frequency_sort(s: String) -> String {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        s.chars().for_each(|c| *map.entry(c).or_insert(0) += 1);
        let mut v = map.into_iter().collect::<Vec<_>>();
        v.sort_unstable_by_key(|&e| -e.1);
        let mut res = String::new();
        for (c, count) in v {
            for _ in 0..count {
                res.push(c);
            }
        }
        res
    }
}

struct Solution;