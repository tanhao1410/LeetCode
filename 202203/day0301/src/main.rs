fn main() {
    println!("Hello, world!");
}

impl Solution {
    //787. K 站中转内最便宜的航班
    pub fn find_cheapest_price(n: i32, flights: Vec<Vec<i32>>, src: i32, dst: i32, k: i32) -> i32 {
        use std::collections::VecDeque;
        let mut queue = VecDeque::new();
        //最多k站，广度优先遍历
        let mut used = vec![i32::MAX; n as usize];
        //下标：出发城市，内容：能够到达的地方，价格
        let mut flight_sets = vec![vec![]; n as usize];
        for flight in &flights {
            flight_sets[flight[0] as usize].push((flight[1] as usize, flight[2]));
        }
        let mut step = 0;
        let mut res = -1;
        queue.push_back((src as usize, 0));
        used[src as usize] = 0;
        while step <= k && !queue.is_empty() {
            let len = queue.len();
            for _ in 0..len {
                let cur = queue.pop_front().unwrap();
                //能都到哪呢？
                for &reach in &flight_sets[cur.0] {
                    //不对，如果走的一条路比原来的要近，是可以更新的
                    //从当前路径到该点需要的费用
                    let cur_money = reach.1 + cur.1;
                    if cur_money < used[reach.0] {
                        queue.push_back((reach.0, reach.1 + cur.1));
                        used[reach.0] = cur_money;
                    }
                    if reach.0 == dst as usize {
                        res = res.max(cur.1 + reach.1);
                    }
                }
            }
            step += 1;
        }
        res
    }
    //108. 将有序数组转换为二叉搜索树
    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        //思路：从中间切开。中间节点就是根节点。
        Self::sorted_slice_to_bst(&nums)
    }

    fn sorted_slice_to_bst(nums: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        if nums.len() == 0 {
            return None;
        }
        //找中间节点
        let mid = nums.len() / 2;
        let mut root = TreeNode::new(nums[mid]);
        root.left = Self::sorted_slice_to_bst(&nums[..mid]);
        root.right = Self::sorted_slice_to_bst(&nums[mid + 1..]);
        Some(Rc::new(RefCell::new(root)))
    }
    //103. 二叉树的锯齿形层序遍历
    pub fn zigzag_level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        //层序遍历
        use std::collections::VecDeque;
        let mut flag = true;
        let mut res = vec![];
        let mut queue = VecDeque::new();
        if let Some(node) = &root {
            queue.push_back(node.clone());
        };
        while !queue.is_empty() {
            let mut len = queue.len();
            let mut v = vec![];
            for _ in 0..len {
                let cur = queue.pop_front().unwrap();
                let cur = cur.borrow();
                if cur.left.is_some() {
                    queue.push_back(cur.left.as_ref().unwrap().clone());
                }
                if cur.right.is_some() {
                    queue.push_back(cur.right.as_ref().unwrap().clone());
                }
                v.push(cur.val);
            }
            if flag {
                res.push(v);
            } else {
                res.push(v.into_iter().rev().collect());
            }
            flag = !flag;
        }
        res
    }
    //105. 从前序与中序遍历序列构造二叉树
    pub fn build_tree(preorder: Vec<i32>, inorder: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        Self::build_tree_by_slice(&preorder, &inorder)
    }
    fn build_tree_by_slice(preorder: &[i32], inorder: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        //思路：前序 ：中左右。中序 左中右。所有很容易得到哪个是中间
        //然后递归的方式，分别构造左子树与右子树
        if preorder.len() == 0 {
            return None;
        }
        //找到根节点
        let root_val = preorder[0];
        let mut index = 0;
        while inorder[index] != root_val {
            index += 1;
        }
        let mut root = TreeNode::new(root_val);
        if index > 0 {
            root.left = Self::build_tree_by_slice(&preorder[1..index + 1], &inorder[..index]);
        }
        if index < preorder.len() - 1 {
            root.right = Self::build_tree_by_slice(&preorder[index + 1..], &inorder[index + 1..]);
        }
        Some(Rc::new(RefCell::new(root)))
    }
    //1042. 不邻接植花
    pub fn garden_no_adj(n: i32, paths: Vec<Vec<i32>>) -> Vec<i32> {
        use std::collections::VecDeque;
        let mut path = vec![vec![]; n as usize];
        for p in &paths {
            path[p[0] as usize - 1].push(p[1] as usize - 1);
            path[p[1] as usize - 1].push(p[0] as usize - 1);
        }
        let mut queue = VecDeque::new();
        let mut res = vec![0; n as usize];
        for i in 0..n as usize {
            if res[i] == 0 {
                queue.push_back(i);
                while !queue.is_empty() {
                    let len = queue.len();
                    for _ in 0..len {
                        let cur = queue.pop_front().unwrap();
                        let mut temp = vec![0; 4];
                        //应该用哪种颜色
                        for &arround in &path[cur] {
                            if res[arround] == 0 {
                                queue.push_back(arround);
                            } else {
                                temp[res[arround] as usize - 1] = 1;
                            }
                        }
                        for i in 0..4 {
                            if temp[i] == 0 {
                                res[cur] = i as i32 + 1;
                                break;
                            }
                        }
                    }
                }
            }
        }
        res
    }
    //58. 最后一个单词的长度
    pub fn length_of_last_word(s: String) -> i32 {
        let bytes = s.as_bytes();
        let mut end = s.len() - 1;
        while end >= 0 && bytes[end] == b' ' {
            end -= 1;
        }
        let mut start = end;
        while start >= 0 && bytes[start] != b' ' {
            start -= 1;
        }
        (end - start) as i32
    }
    //739. 每日温度
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut res = vec![0; temperatures.len()];
        let mut stack: Vec<(i32, usize)> = vec![];
        //单调栈,栈中存放后面比它大的数及坐标
        for i in (0..temperatures.len()).rev() {
            //如果栈顶元素比当前元素小或等于，则弹出元素。
            while stack.len() > 0 && stack[stack.len() - 1].0 <= temperatures[i] {
                stack.pop();
            }
            //如果栈不为空，则说明存在了一个比当前大的数，且是最近的
            if stack.len() > 0 {
                res[i] = stack[stack.len() - 1].1 as i32 - i as i32;
            }
            stack.push((temperatures[i], i));
        }
        res
    }
    //6. Z 字形变换
    pub fn convert(s: String, num_rows: i32) -> String {
        if num_rows == 1 {
            return s;
        }
        let mut m = vec![vec![]; num_rows as usize];
        let mut x = 0;
        let mut y = 0;
        let mut is_down = true;//方向有往下和往右上
        let bytes = s.as_bytes();
        for &b in bytes {
            m[x].push(b);
            if is_down {
                if x < num_rows as usize - 1 {
                    x += 1;
                } else {
                    is_down = false;
                    x -= 1;
                    y += 1;
                }
            } else {
                if x == 0 {
                    is_down = true;
                    x += 1;
                } else {
                    x -= 1;
                    y += 1;
                }
            }
        }
        String::from_utf8(m
            .into_iter()
            .flat_map(|v| v.into_iter())
            .collect::<Vec<_>>()).unwrap()
    }
}

struct Solution;

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