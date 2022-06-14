fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1093. 大样本统计
    pub fn sample_stats(count: Vec<i32>) -> Vec<f64> {
        let (mut num_count, mut sum) = (0, 0i64);
        let mut mode_count = 0;
        let mut mode = 0;
        let (mut min, mut max) = (256, 0);
        for i in 0..count.len() {
            num_count += count[i];
            sum += i as i64 * count[i] as i64;
            if count[i] > mode_count {
                mode_count = count[i];
                mode = i;
            }
            if count[i] > 0 {
                if i < min {
                    min = i;
                }
                if i > max {
                    max = i;
                }
            }
        }

        //求中间数,取 num_count/2,num_count/2 -1这两个数，跳过多少个呢
        let mut median = 0;
        let mut median2 = 256;
        let mut skip_num = 0;
        for i in 0..count.len() {
            skip_num += count[i];
            if skip_num > num_count / 2 - 1 && median2 == 256 {
                median2 = i;
            }
            if skip_num > num_count / 2 {
                median = i;
                break;
            }
        }
        if num_count % 2 == 1 {
            median2 = median;
        }
        vec![min as f64, max as f64, sum as f64 / num_count as f64, (median + median2) as f64 / 2.0, mode as f64]
    }

    //1161. 最大层内元素和
    pub fn max_level_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut queue = std::collections::VecDeque::new();
        queue.push_back(root.unwrap());
        let mut res = 1;
        let mut max = i32::MIN;
        let mut layer = 0;
        while queue.len() > 0 {
            layer += 1;
            let mut cur_sum = 0;
            let len = queue.len();
            for _ in 0..len {
                let head = queue.pop_front().unwrap();
                let head_ref = head.borrow();
                if head_ref.left.is_some() {
                    queue.push_back(head_ref.left.as_ref().unwrap().clone());
                }
                if head_ref.right.is_some() {
                    queue.push_back(head_ref.right.as_ref().unwrap().clone());
                }
                cur_sum += head_ref.val;
            }
            if cur_sum > max {
                max = cur_sum;
                res = layer;
            }
        }
        res
    }

    //498. 对角线遍历
    pub fn find_diagonal_order(mat: Vec<Vec<i32>>) -> Vec<i32> {
        let mut res = vec![];
        let (m, n) = (mat.len(), mat[0].len());
        let mut cur = (0, 0);
        let mut is_up = true;
        while res.len() < m * n {
            res.push(mat[cur.0][cur.1]);
            if is_up {
                if cur.0 as i32 - 1 >= 0 && cur.1 + 1 < n {
                    cur.0 -= 1;
                    cur.1 += 1;
                } else if cur.1 + 1 < n {
                    is_up = !is_up;
                    cur.1 += 1;
                } else {
                    is_up = !is_up;
                    cur.0 += 1;
                }
            } else {
                if cur.0 + 1 < m && cur.1 as i32 - 1 >= 0 {
                    cur.0 += 1;
                    cur.1 -= 1;
                } else if cur.0 + 1 < m {
                    cur.0 += 1;
                    is_up = !is_up;
                } else {
                    is_up = !is_up;
                    cur.1 += 1;
                }
            }
        }
        res
    }
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

use std::rc::Rc;
use std::cell::RefCell;