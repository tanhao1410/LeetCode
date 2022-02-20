fn main() {
    println!("Hello, world!");
    println!("{}", Solution::add_strings("9999".to_string(), "1".to_string()));
}

struct Solution;

impl Solution {
    //415. 字符串相加
    pub fn add_strings(num1: String, num2: String) -> String {
        let mut res = vec![];
        let mut num1 = num1.as_bytes();
        let mut num2 = num2.as_bytes();
        let mut index1 = num1.len() as i32 - 1;
        let mut index2 = num2.len() as i32 - 1;
        let mut flag = 0;
        while index1 >= 0 || index2 >= 0 {
            let mut bit_res = num1.get(index1 as usize).unwrap_or(&b'0') - b'0'
                + num2.get(index2 as usize).unwrap_or(&b'0') - b'0' + flag;
            flag = bit_res / 10;
            bit_res %= 10;
            res.push(bit_res + b'0');
            index1 -= 1;
            index2 -= 1;
        }
        if flag == 1 {
            res.push(b'1');
        }
        String::from_utf8(res.into_iter().rev().collect()).unwrap()
    }
    //1557. 可以到达所有点的最少点数目
    pub fn find_smallest_set_of_vertices(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let mut can_reach = vec![false; n as usize];
        for edge in &edges {
            can_reach[edge[1] as usize] = true;
        }
        can_reach
            .into_iter()
            .enumerate()
            .filter_map(|e| match e.1 {
                false => Some(e.0 as i32),
                _ => None
            })
            .collect()
    }
    //997. 找到小镇的法官
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut map = vec![(0, 0); n as usize];
        for v in trust {
            map[v[0] as usize - 1].0 += 1;
            map[v[1] as usize - 1].1 += 1;
        }
        for i in 0..n as usize {
            if map[i].0 == 0 && map[i].1 == n - 1 {
                return i as i32 + 1;
            }
        }
        return -1;
    }
}