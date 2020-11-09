use std::i32::MAX;

fn main() {
    println!("Hello, world!");
    Solution::k_closest(vec![vec![1, 3], vec![-2, 2]], 1);
}

struct Solution {}

impl Solution {
    //973. 最接近原点的 K 个点
    pub fn k_closest(points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        //思路：用一个vec记录最接近的点，发现更小的后，清空原来的，加入新的
        //找出k个，而不仅仅是最小的几个。插入排序法
        let (mut res, mut cur_dis) = (Vec::new(), std::i32::MAX);

        fn get_dis(point: &Vec<i32>) -> i32 {
            point[0] * point[0] + point[1] * point[1]
        }
        for i in 0..k {
            let mut index = 0;
            while index < res.len() && get_dis(&res[index]) < get_dis(&points[i as usize]) {
                index += 1;
            }
            res.insert(index, points[i as usize].clone());
        }
        for i in k..points.len() as i32 {
            let mut index = k - 1;
            while index >= 0 && get_dis(&res[index as usize]) > get_dis(&points[i as usize]) {
                index -= 1;
            }
            if index < 0 {
                res.insert(0  ,points[i as usize].clone());
                res.pop();
            } else if index < k - 1 {
                res.insert(index as usize + 1, points[i as usize].clone());
                res.pop();
            }
        }
        res
    }
}
