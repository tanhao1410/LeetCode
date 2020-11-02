fn main() {
    println!("Hello, world!");
    //println!("{}", Solution::find_nth_digit(1000000000));
    //Solution::moving_count(2, 2, 2);
    println!("{}",Solution::length_of_longest_substring("abdib".to_string()))
}

struct Solution {}

impl Solution {
    //每日一题：349. 两个数组的交集
    pub fn intersection(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        //思路：用一个hashset来存储1中的数据，然后，遍历nums2，每次拿数据的时候，按顺序放入结果 集即可避免重复的出现。或者放入到set中，然后，再输出为vec
        use std::collections::HashSet;
        let (mut res, mut set1, mut res_set) = (vec![], HashSet::new(), HashSet::new());
        for i in nums1 {
            set1.insert(i);
        }
        for i in nums2 {
            if set1.contains(&i) {
                res_set.insert(i);
            }
        }
        for i in res_set {
            res.push(i);
        }
        res
    }

    //面试题 01.02. 判定是否互为字符重排
    pub fn check_permutation(s1: String, s2: String) -> bool {
        if s1.len() != s2.len() {
            return false;
        }
        //采用map的形式
        let mut m1 = std::collections::HashMap::new();
        let mut m2 = std::collections::HashMap::new();
        for i in 0..s1.len() {
            if m1.contains_key(&s1.chars().nth(i).unwrap()) {
                m1.insert(s1.chars().nth(i).unwrap(), m1.get(&s1.chars().nth(i).unwrap()).unwrap() + 1);
            } else {
                m1.insert(s1.chars().nth(i).unwrap(), 1);
            }

            if m2.contains_key(&s2.chars().nth(i).unwrap()) {
                m2.insert(s2.chars().nth(i).unwrap(), m2.get(&s1.chars().nth(i).unwrap()).unwrap() + 1);
            } else {
                m2.insert(s2.chars().nth(i).unwrap(), 1);
            }
        }
        if m1.len() != m2.len() {
            return false;
        }
        for i in m2.into_iter() {
            match m1.get(&i.0) {
                None => return false,
                Some(&j) => {
                    if i.1 != j {
                        return false;
                    }
                }
            }
        }
        true
    }

    //剑指 Offer 44. 数字序列中某一位的数字
    pub fn find_nth_digit(n: i32) -> i32 {

        //10->1 //1000000000 ->1
        if n < 10 {
            return n;
        }

        fn get_bit(mut n: i32, mut base: i32, mut index: i32) -> i32 {
            let mut res = n;
            while index >= 0 {
                res = n / base;
                n = n % base;
                base /= 10;
                index -= 1;
            }
            res
        }

        let mut dp: Vec<i64> = vec![0, 10];
        let mut max: i64 = 10;
        let mut i = 1;
        let mut ten_i = 1;
        while max <= n as i64 {
            i += 1;
            ten_i *= 10;
            max = *dp.last().unwrap() + 9 * i * ten_i;
            dp.push(max);
        }
        //前面切割掉
        let new_index = n - dp[dp.len() - 2] as i32;
        //此时的数字的位数都是 i
        let skip_count = new_index / i as i32;
        let skip_index = skip_count + ten_i as i32;
        get_bit(skip_index, ten_i as i32, new_index % i as i32)
    }

    //剑指 Offer 13. 机器人的运动范围
    pub fn moving_count(m: i32, n: i32, k: i32) -> i32 {
        //广度优先遍历，采用队列来解决。
        //怎么记录已经遍历过的呢，采用二维数组
        //问题，同一个节点多次入队。
        fn can_in(x: usize, y: usize, k: i32) -> bool {
            let mut count = 0;
            count += x % 10;
            count += x / 10;
            count += y % 10;
            count += y / 10;
            count as i32 <= k
        }
        let mut res = 0;
        let mut queue = vec![(0usize, 0usize)];
        let mut matrix = vec![vec![0; n as usize]; m as usize];
        while queue.len() > 0{
            //从队列中出队，
            let point = queue.pop().unwrap();
            if matrix[point.0][point.1] == 0{
                res += 1;
                matrix[point.0][point.1] = 1;
            }
            //上下左右是否可以移动到，如果能移动到，就入队
            if point.0 as i32 - 1 >= 0 && matrix[point.0 - 1][point.1] == 0 && can_in(point.0 - 1, point.1, k) {
                queue.push((point.0 - 1, point.1))
            }
            if point.0 as i32 + 1 < m && matrix[point.0 + 1][point.1] == 0 && can_in(point.0 + 1, point.1, k) {
                queue.push((point.0 + 1, point.1))
            }
            if point.1 as i32 - 1 >= 0 && matrix[point.0][point.1 - 1] == 0 && can_in(point.0, point.1 - 1, k) {
                queue.push((point.0, point.1 - 1))
            }
            if point.1 as i32 + 1 < n && matrix[point.0][point.1 + 1] == 0 && can_in(point.0, point.1 + 1, k) {
                queue.push((point.0, point.1 + 1))
            }
        }
        res
    }

    //剑指 Offer 48. 最长不含重复字符的子字符串
    pub fn length_of_longest_substring(s: String) -> i32 {
        if s.len() < 2{
            return s.len() as i32;
        }
        fn has_repit_char(s:&str)->i32{
            let mut end = s.len() as i32;
            let last_char = s.chars().nth(end as usize - 1).unwrap();
            end -= 2;
            while end >= 0{
                if s.chars().nth(end as usize).unwrap() == last_char{
                    return end;
                }
                end -=1;
            }
            -1
        }
        let mut res = 0;
        let (mut i,mut j) = (1,0);
        while i < s.len(){
            //判断s[i..j+1]在前面是否有，如果有，确定它位置
            if has_repit_char(&s[j..i+1]) == -1{
                i += 1;
            }else{
                j = j + has_repit_char(&s[j..i + 1]) as usize + 1;
                i += 1;
            }
            if i -j > res{
                res = i - j
            }
        }
        res as i32
    }
}