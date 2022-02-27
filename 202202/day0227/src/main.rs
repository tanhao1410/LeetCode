fn main() {
    println!("Hello, world!");
}

impl Solution {
    //553. 最优除法
    pub fn optimal_division(nums: Vec<i32>) -> String {
        //递归方式，求最大值，或最小值
        Self::max_or_min(&nums[..], true).1
    }

    fn max_or_min(nums: &[i32], is_max: bool) -> (f64, String) {
        //如果更新了最大值，还需要返回应有的计算方式。
        let mut s = String::new();
        if nums.len() == 1 {
            return (nums[0] as f64, nums[0].to_string());
        }
        let mut res = 0.0f64;
        for i in 1..nums.len() {
            //前面的求最大值，后面求最小值
            let cur = Self::max_or_min(&nums[..i], is_max).0 / Self::max_or_min(&nums[i..], !is_max).0;
            if (is_max && cur > res) || (!is_max && cur < res) {
                res = cur;
                s = Self::create_express(&nums[..i], &nums[i..]);
            }
        }
        (res, s)
    }

    fn create_express(pre: &[i32], pro: &[i32]) -> String {
        let mut s = String::new();
        if pre.len() > 1 {
            s.push('(');
            for &num in pre {
                s.push_str(&num.to_string());
                s.push('/');
            }
            s.remove(s.len() - 1);
            s.push(')');
        } else {
            s.push_str(&pre[0].to_string());
        }
        s.push('/');
        if pro.len() > 1 {
            s.push('(');
            for &num in pro {
                s.push_str(&num.to_string());
                s.push('/');
            }
            s.remove(s.len() - 1);
            s.push(')');
        } else {
            s.push_str(&pro[0].to_string());
        }
        s
    }
    //149. 直线上最多的点数
    pub fn max_points(points: Vec<Vec<i32>>) -> i32 {
        //选择一个点，再选择一个点，然后找到所有在这上面的点。
        let is_one_line = |p1: usize, p2: usize, p3: usize| {
            let p1 = (points[p1][0], points[p1][1]);
            let p2 = (points[p2][0], points[p2][1]);
            let p3 = (points[p3][0], points[p3][1]);
            //p2.x - p1.x / p2.y - p1.y = p3.x - p1.x /p3.y - p1.y
            (p2.0 - p1.0) * (p3.1 - p1.1) == (p2.1 - p1.1) * (p3.0 - p1.0)
        };
        let mut res = 2;
        if points.len() < 3 {
            return points.len() as i32;
        }
        for i in 0..points.len() - 2 {
            for j in i + 1..points.len() - 1 {
                let mut item = 2;
                for k in j + 1..points.len() {
                    if is_one_line(i, j, k) {
                        item += 1;
                    }
                }
                res = res.max(item);
            }
        }
        res
    }
    //3. 无重复字符的最长子串
    pub fn length_of_longest_substring(s: String) -> i32 {
        // 用一个map 记录每一个字母再前面的位置，-1代表无。
        //子串最大为26，用一个滑动窗口，开始时是1，然后依次加入新的字母，看该字母在前面是否存在，如果存在，看是否在窗口中
        //如果不在窗口里，则继续。否则，窗口开始位置定为它的下一个位置
        let mut pre_index = vec![-1; 128];
        let bytes = s.as_bytes();
        let mut start = 0;
        let mut res = 0;
        for i in 0..s.len() {
            //在窗口中
            if pre_index[bytes[i] as usize] != -1 && pre_index[bytes[i] as usize] >= start as i32 {
                start = pre_index[bytes[i] as usize] as usize + 1;
            }
            pre_index[bytes[i] as usize] = i as i32;
            res = res.max(i - start + 1);
        }
        res as i32
    }
    //1367. 二叉树中的列表
    pub fn is_sub_path(mut head: Option<Box<ListNode>>, root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(node) = head {
            if let Some(tree_node) = root {
                let mut ref_mut = tree_node.borrow_mut();
                if node.val == tree_node.borrow().val {
                    if Self::is_equal(&node.next, &ref_mut.left)
                        || Self::is_equal(&node.next, &ref_mut.right) {
                        return true;
                    }
                }
                Self::is_sub_path((Some(node.clone())), ref_mut.left.take())
                    || Self::is_sub_path(Some(node), ref_mut.right.take())
            } else {
                false
            }
        } else {
            true
        }
    }
    //必须是接连相等
    fn is_equal(mut head: &Option<Box<ListNode>>, root: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(node) = head {
            if let Some(tree_node) = root {
                if node.val != tree_node.borrow().val {
                    return false;
                }
                Self::is_equal(&node.next, &tree_node.borrow().left)
                    || Self::is_equal(&node.next, &tree_node.borrow().right)
            } else {
                false
            }
        } else {
            true
        }
    }
}

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

struct Solution;

use std::rc::Rc;
use std::cell::RefCell;
