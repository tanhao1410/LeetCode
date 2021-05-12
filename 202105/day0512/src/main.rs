fn main() {
    println!("Hello, world!");
}

pub struct Solution {}

impl Solution {

    //1310. 子数组异或查询
    pub fn xor_queries2(arr: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        //前缀异或数组解决 前缀异或数组不包含自身。所以，[2,6]^ = pre[2]^pre[6+1]
        let mut pre = vec![0;arr.len() + 1];
        // for i in 1..arr.len()+1{
        //     pre[i] = pre[i - 1] ^ arr[i - 1]
        // }
        (1..arr.len() + 1).for_each(|i| pre[i] = pre[i - 1] ^ arr[i - 1]);

        queries.iter().map(|v|{
            pre[v[0] as usize] ^ pre[v[1] as usize + 1]
        }).collect()

    }

    //1310. 子数组异或查询
    pub fn xor_queries(arr: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        //暴力法
        queries.iter().map(|v| (v[0]..v[1] + 1).fold(0, |i, j| i ^ arr[j as usize])).collect()
    }

    //229. 求众数 II-间复杂度为 O(n)、空间复杂度为 O(1)
    pub fn majority_element2(nums: Vec<i32>) -> Vec<i32> {
        //最简单的思路，采用map
        // let mut map = std::collections::HashMap::new();
        // nums.iter().for_each(|&num| { * map.entry(num).or_insert(0) += 1; });
        // map.iter().filter(|&(_, &v)| { v > nums.len() / 3 }).map(|(&k, _)| { k }).collect()

        let mut vec: Vec<(i32, i32)> = vec![];
        nums.iter().for_each(|&num| {
            //投票数组为空，或只有一个时
            if vec.is_empty() || (vec.len() == 1 && vec[0].0 != num) {
                vec.push((num, 1))
            } else if vec[0].0 == num {
                //与第一个相同
                vec[0].1 += 1;
            } else if vec[1].0 == num {
                //与第二个相同
                vec[1].1 += 1;
            } else {
                //与两个都不相等
                vec[0].1 -= 1;
                vec[1].1 -= 1;
                if vec[0].1 < 1 {
                    vec.remove(0);
                }
                if vec[1].1 < 1 {
                    vec.remove(0);
                }
                if vec.len() < 2 {
                    vec.push((num, 1));
                }
            }
        });
        //未保证一定存在两个
        vec.iter().map(|&(num, _)| num).collect()
    }

    //229. 求众数 II
    pub fn majority_element(nums: Vec<i32>) -> Vec<i32> {
        //最简单的思路，采用map
        let mut map = std::collections::HashMap::new();
        let count = nums.len() / 3;
        nums.iter().for_each(|&num| {
            let entry = map.entry(num).or_insert(0);
            *entry += 1;
        });
        map.iter().filter(|&(_, &v)| { v > count }).map(|(&k, _)| { k }).collect()
    }

    //260. 只出现一次的数字 III
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        //有两个元素只出现一次，其余的出现两次。
        //思路：把数组分为两部分，一部分包含一个，其中相同的数包含在同一部分中
        //将所有的数进行异或，得到的结果是两个不同的数异或的值
        let two = nums.iter().fold(0, |i, j| i ^ j);
        //从异或的值中找到某一位为1，以此进行分割数组
        let mut v = two;
        let mut master = -1;
        while v.count_ones() > 1 {
            v &= master;
            master <<= 1;
        }
        let res1 = nums.iter().filter(|&&i| i & v == 0).fold(0, |i, &j| i ^ j);
        vec![res1, res1 ^ two]
    }
}
