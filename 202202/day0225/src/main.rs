fn main() {
    println!("Hello, world!");
}

impl Solution {
    //82. 删除排序链表中的重复元素 II
    pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        //如果有重复的，则都要删除，采用弹出，加入方法
        // 每一次必须知道本次的情况下，才能加入上一个。
        let mut res = Some(Box::new(ListNode::new(0)));
        let mut p = res.as_mut().unwrap();
        let mut pre = 101;
        while let Some(mut node) = head {
            head = node.next.take();
            //如果当前访问的值与下一个值相等或与上一个值相等，则当前值不加进去。
            if (head.is_some() && head.as_ref().unwrap().val == node.val)
                || node.val == pre {
                pre = node.val;
            } else {
                pre = node.val;
                p.next = Some(node);
                p = p.next.as_mut().unwrap();
            }
        }
        res.as_mut().unwrap().next.take()
    }
    //322. 零钱兑换
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut dp = vec![vec![-1; amount as usize + 1]; coins.len()];
        for i in 0..dp.len() {
            dp[i][0] = 0;
            for j in 1..dp[0].len() {
                if j >= coins[i] as usize && dp[i][j - coins[i] as usize] != -1 {
                    dp[i][j] = dp[i][j - coins[i] as usize] + 1;
                    if i > 0 && dp[i - 1][j] != -1 && dp[i - 1][j] < dp[i][j] {
                        dp[i][j] = dp[i - 1][j]
                    }
                } else if i > 0 {
                    dp[i][j] = dp[i - 1][j];
                }
            }
        }
        *dp.last().unwrap().last().unwrap()
    }
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

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode {
            next: None,
            val,
        }
    }
}