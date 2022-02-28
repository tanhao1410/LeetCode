fn main() {
    println!("Hello, world!");
}

impl Solution {
    //97. 交错字符串
    pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
        if s1.len() + s2.len() != s3.len() {
            return false;
        }
        let (s1, s2, s3) = (s1.as_bytes(), s2.as_bytes(), s3.as_bytes());
        //动态规划，两重循环是否可行呢？s1的前i个元素，与s2的前j 个元素
        let mut dp = vec![vec![false; s2.len() + 1]; s1.len() + 1];
        for i in 0..dp.len() {
            for j in 0..dp[0].len() {
                if i == 0 && j == 0 {
                    dp[i][0] = true;
                } else if i == 0 {
                    dp[i][j] = dp[i][j - 1] && s2[j - 1] == s3[j - 1];
                } else if j == 0 {
                    dp[i][j] = dp[i - 1][j] && s1[i - 1] == s3[i - 1];
                } else {
                    dp[i][j] |= dp[i][j - 1] && s2[j - 1] == s3[j + i - 1];
                    dp[i][j] |= dp[i - 1][j] && s1[i - 1] == s3[i + j - 1];
                }
            }
        }
        dp[s3.len()][s1.len()]
    }
    //1823. 找出游戏的获胜者
    pub fn find_the_winner(n: i32, k: i32) -> i32 {
        let mut fails = vec![false; n as usize];
        let mut fail_count = 0;
        let mut start = 0;
        while fail_count < n - 1 {
            //第一步是走在了start位置，所以，此处用k - 1
            for _ in 0..k - 1 {
                //往前走一步
                start = (start + 1) % n;
                //但此时有可能走在了已经淘汰过的位置上了,如果是，继续往前走，直到走到没有淘汰的地方
                while fails[start as usize] {
                    start = (start + 1) % n;
                }
            }
            //此时，start即要被淘汰的地方
            fails[start as usize] = true;
            fail_count += 1;
            //走到下一个开始位置
            start = (start + 1) % n;
            while fails[start as usize] {
                start = (start + 1) % n;
            }
        }
        start
    }
    //67. 二进制求和
    pub fn add_binary(a: String, b: String) -> String {
        let bytes1 = a.as_bytes();
        let bytes2 = b.as_bytes();
        let mut res = vec![b'1'; a.len().max(b.len()) + 1];
        let n = res.len();
        let mut index = 0;
        let mut flag = 0;
        while index < a.len() || index < b.len() {
            let mut item = 0;
            if (index < a.len() && index < b.len()) {
                item = bytes1[a.len() - 1 - index] - b'0' + bytes2[b.len() - 1 - index] - b'0' + flag;
            } else if (index < a.len()) {
                item = bytes1[a.len() - 1 - index] - b'0' + flag;
            } else {
                item = bytes2[b.len() - 1 - index] - b'0' + flag;
            }
            res[n - 1 - index] = item % 2 + b'0';
            flag = item / 2;
            index += 1;
        }
        if flag == 1 {
            String::from_utf8(res).unwrap()
        } else {
            String::from_utf8_lossy(&res[1..]).to_string()
        }
    }
    //1601. 最多可达成的换楼请求数目
    pub fn maximum_requests(n: i32, requests: Vec<Vec<i32>>) -> i32 {
        //使用java用递归解决了，用栈来解决呢？深度优先遍历
        //用一个数来表示哪些请求使用了或没使用，因为最多16个，所以，一个i32的数足矣
        //(请求的序号，是否要，已经要了的请求）
        let mut stack = vec![(0, 1, 1), (0, 0, 0)];
        //总共16个深度，每步有两种走法，选择或不选择
        let mut res = 0;
        while let Some((index, used, count)) = stack.pop() {
            //走到头了
            if index == requests.len() - 1 {
                let mut item = 0;
                let mut buildings = vec![0; n as usize];
                //还需要判断是否满足要求
                for i in 0..requests.len() {
                    //该位被使用了
                    if used & (1 << i) > 0 {
                        item += 1;
                        buildings[requests[i][0] as usize] -= 1;
                        buildings[requests[i][1] as usize] += 1;
                    }
                }
                if buildings.into_iter().all(|v| v == 0) {
                    res = res.max(item);
                }
            } else {
                stack.push((index + 1, used | (1 << index + 1), count + 1));
                stack.push((index + 1, used, count));
            }
        }
        res
    }
}

struct Solution;