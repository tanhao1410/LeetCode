fn main() {
    println!("Hello, world!");
    println!("{}", integer_break(58));
}

//1405. 最长快乐字符串
pub fn longest_diverse_string(a: i32, b: i32, c: i32) -> String {
    //思路：谁最大，先放谁。如果前面该字母超过了两个，不能放。然后放次大的。
    let mut res = String::new();
    let (mut a, mut b, mut c) = (a, b, c);
    let mut count = 0;
    let mut pre_letter = -1;//0,1,2

    while a != 0 || b != 0 || c != 0 {
        match (a >= b, a >= c, b >= c) {
            (true, true, false) => {
                if pre_letter == 0 && count == 2 {
                    if b == 0 {
                        break;
                    }
                    //选择b
                    b -= 1;
                    res.push('b');
                    pre_letter = 1;
                    count = 1;
                } else if pre_letter == 0 {
                    a -= 1;
                    res.push('a');
                    count += 1;
                } else {
                    a -= 1;
                    res.push('a');
                    count = 1;
                    pre_letter = 0;
                }
            }
            (true, true, true) => {
                if pre_letter == 0 && count == 2 {
                    //选择b
                    c -= 1;
                    res.push('c');
                    pre_letter = 2;
                    count = 1;
                } else if pre_letter == 0 {
                    a -= 1;
                    res.push('a');
                    count += 1;
                } else {
                    a -= 1;
                    res.push('a');
                    count = 1;
                    pre_letter = 0;
                }
            }

            (false, true, true) => {
                if pre_letter == 1 && count == 2 {
                    if a == 0 {
                        break;
                    }
                    //选择b
                    a -= 1;
                    res.push('a');
                    pre_letter = 0;
                    count = 1;
                } else if pre_letter == 1 {
                    b -= 1;
                    res.push('b');
                    count += 1;
                } else {
                    b -= 1;
                    res.push('b');
                    count = 1;
                    pre_letter = 1;
                }
            }
            (false, false, true) => {
                if pre_letter == 1 && count == 2 {
                    //选择b
                    c -= 1;
                    res.push('c');
                    pre_letter = 2;
                    count = 1;
                } else if pre_letter == 1 {
                    b -= 1;
                    res.push('b');
                    count += 1;
                } else {
                    b -= 1;
                    res.push('b');
                    count = 1;
                    pre_letter = 1;
                }
            }

            (true, false, false) => {
                if pre_letter == 2 && count == 2 {
                    if a == 0 {
                        break;
                    }
                    //选择b
                    a -= 1;
                    res.push('a');
                    pre_letter = 0;
                    count = 1;
                } else if pre_letter == 2 {
                    c -= 1;
                    res.push('c');
                    count += 1;
                } else {
                    c -= 1;
                    res.push('c');
                    count = 1;
                    pre_letter = 2;
                }
            }
            (false, false, false) => {
                if pre_letter == 2 && count == 2 {
                    //选择b
                    b -= 1;
                    res.push('b');
                    pre_letter = 1;
                    count = 1;
                } else if pre_letter == 2 {
                    c -= 1;
                    res.push('c');
                    count += 1;
                } else {
                    c -= 1;
                    res.push('c');
                    count = 1;
                    pre_letter = 2;
                }
            }
            //三者相等
            _ => unreachable!()
        }
    }
    res
}

//343. 整数拆分
pub fn integer_break(n: i32) -> i32 {
    //思路：动态规划，dp[n] = dp[2] * dp[n - 2] ,dp[3] * dp[n -3] ... max
    let mut dp = vec![0, 1, 1, 2];
    for i in 4..=n {
        let mut cur = 0;
        for j in 2..=i / 2 {
            cur = cur.max(j * dp[i as usize - j as usize]);
            cur = cur.max(j * (i - j));
        }
        dp.push(cur);
    }
    dp[n as usize]
}

//279. 完全平方数
pub fn num_squares(n: i32) -> i32 {
    //完全平方数有哪些【1,4,9,16,...,10000】共100个。
    // dp[上述] = 1；
    //for [上述] 遍历 完全平方数 + 1 dp[]
    //dp[j] = dp[j - nums[i]] + 1 .min(self);
    //从小往大的开始算
    //dp[j + nums[i]] = dp[j] + 1 .min(self)
    let mut nums = vec![];
    for i in 1..10000 {
        if i * i > n {
            break;
        }
        nums.push(i * i);
    }
    let mut dp = vec![i32::MAX; n as usize + 1];
    for i in 0..nums.len() {
        dp[nums[i] as usize] = 1;
    }
    for j in 1..=n {
        for i in 0..nums.len() {
            if j + nums[i] <= n {
                dp[j as usize + nums[i] as usize] = dp[j as usize + nums[i] as usize].min(dp[j as usize] + 1);
            }
        }
    }
    dp[n as usize]
}

//377. 组合总和 Ⅳ
pub fn combination_sum4(nums: Vec<i32>, target: i32) -> i32 {
    // let mut dp = vec![0; target as usize + 1];
    // dp[0] = 1;
    // for i in 0..nums.len() {
    //     for j in nums[i]..=target {
    //         dp[j] += dp[(j - nums[i]) as usize];
    //     }
    // }
    // dp[target as usize]
    //上面是不在乎顺序的组合，对于与顺序相关的组合？
    let mut dp = vec![0; target as usize + 1];
    dp[0] = 1;
    for i in 0..=target as usize {
        for j in 0..nums.len() {
            if i >= nums[j] as usize {
                dp[i] += dp[i - nums[j] as usize];
            }
        }
    }
    dp[target as usize]
}