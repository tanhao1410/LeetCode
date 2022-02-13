fn main() {
    println!("Hello, world!");
}

//547. 省份数量
pub fn find_circle_num(is_connected: Vec<Vec<i32>>) -> i32 {
    let mut res = 0;
    //已经加入的城市省份
    let mut provinces = vec![0; is_connected.len()];
    for i in 0..is_connected.len() {
        //说明该省份尚未加入任何省份
        if provinces[i] == 0 {
            res += 1;
            provinces[i] = 1;
            //能到达的城市有：
            let mut stack = vec![];
            stack.push(i);
            //深度遍历
            while let Some(n) = stack.pop() {
                //深度遍历，并把遍历到的城市置位1
                let connecteds = &is_connected[n];
                for j in 0..connecteds.len() {
                    if connecteds[j] == 1 && provinces[j] == 0 {
                        stack.push(j);
                        provinces[j] = 1;
                    }
                }
            }
        }
    }
    res
}

//713. 乘积小于K的子数组
pub fn num_subarray_product_less_than_k(nums: Vec<i32>, k: i32) -> i32 {
    //思路：先计算出第一个窗口中，小于k的字数组长度，
    let mut res = 0;
    let mut cap = 0;
    let mut sum = 0;
    for i in 0..nums.len() {
        //遇到比k大的数
        if nums[i] >= k {
            cap = 0;
            sum = 0;
            continue;
        }
        if cap == 0 {
            cap = 1;
            sum = nums[i];
        } else {
            //去掉一个，加上一个。
            sum /= nums[i - 1];
            //补上后面的
            cap -= 1;
        }
        while i + cap < nums.len() && sum * nums[cap + i] < k {
            sum *= nums[cap + i];
            cap += 1;
        }
        res += cap;
    }
    res as i32
}

//438. 找到字符串中所有字母异位词
pub fn find_anagrams(s: String, p: String) -> Vec<i32> {
    //思路：先统计p中字母的数量。用一个数组统计
    // 在s中用一个窗口，每一次，进一个字母，出一个字母。
    let mut p_set = vec![0; 26];
    for &i in p.as_bytes() {
        p_set[(i - b'a') as usize] += 1;
    }
    let s_bytes = s.as_bytes();
    let mut s_set = vec![0; 26];
    for &i in s.as_bytes().iter().take(p.len()) {
        s_set[(i - b'a') as usize] += 1;
    }
    let is_equal = |p: &[i32], s: &[i32]| -> bool{
        p.iter().zip(s.iter()).all(|(i, j)| i == j)
    };
    let mut res = vec![];
    if s.len() < p.len() {
        return res;
    }
    for i in 0..=s.len() - p.len() {
        if i > 0 {
            s_set[(s_bytes[i - 1] - b'a') as usize] -= 1;
            s_set[(s_bytes[i - 1 + p.len()] - b'a') as usize] += 1;
        }
        //加入一个词，减去一个词
        if is_equal(&p_set, &s_set) {
            res.push(i as i32);
        }
    }
    res
}

//1189. “气球” 的最大数量
pub fn max_number_of_balloons(text: String) -> i32 {
    //balloon
    let bytes = text.as_bytes();
    let (mut a, mut b, mut n, mut l, mut o) = (0, 0, 0, 0, 0);
    for &byte in bytes {
        match byte {
            b'a' => a += 1,
            b'b' => b += 1,
            b'n' => n += 1,
            b'l' => l += 1,
            b'o' => o += 1,
            _ => {}
        }
    }
    a.min(b).min(n).min(l / 2).min(o / 2)
}
