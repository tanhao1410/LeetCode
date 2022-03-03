use std::cell::RefCell;
use std::rc::Rc;

fn main() {
    println!("Hello, world!");
}

impl Solution {
    //剑指 Offer II 050. 向下的路径节点之和
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> i32 {
        Self::path_sum_parent(&root, target_sum, vec![])
    }

    fn path_sum_parent(root: &Option<Rc<RefCell<TreeNode>>>, target: i32, mut parent: Vec<i32>) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.as_ref().unwrap();
        let val = root.borrow().val;
        let mut res = 0;
        for p in &mut parent {
            if val + *p == target {
                res += 1;
            }
            *p += val;
        }
        parent.push(val);
        if val == target {
            res += 1;
        }
        res += Self::path_sum_parent(&root.borrow().left, target, parent.clone());
        res + Self::path_sum_parent(&root.borrow().right, target, parent)
    }
    //剑指 Offer II 049. 从根节点到叶节点的路径数字之和
    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //思路：遍历时增加一个层级。parent代表前面的数
        let mut res = 0;
        Self::sum_numbers_give_parent(root.as_ref().unwrap().clone(), 0, &mut res);
        res
    }
    fn sum_numbers_give_parent(root: Rc<RefCell<TreeNode>>, parent: i32, res: &mut i32) {
        let cur_val = parent * 10 + root.borrow().val;
        if root.borrow().left.is_none() && root.borrow().right.is_none() {
            *res += cur_val;
        } else {
            if root.borrow().left.is_some() {
                Self::sum_numbers_give_parent(root.borrow().left.as_ref().unwrap().clone(), cur_val, res);
            }
            if root.borrow().right.is_some() {
                Self::sum_numbers_give_parent(root.borrow().right.as_ref().unwrap().clone(), cur_val, res);
            }
        }
    }
    //剑指 Offer II 047. 二叉树剪枝
    pub fn prune_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        //什么样的节点删除呢？值为0 ，所有子节点都是0.
        if let Some(node) = root {
            //看它的左子树会不会被删
            let left = Self::prune_tree(node.borrow_mut().left.take());
            let right = Self::prune_tree(node.borrow_mut().right.take());
            if left.is_none() && right.is_none() && node.borrow().val == 0 {
                return None;
            }
            node.borrow_mut().left = left;
            node.borrow_mut().right = right;
            return Some(node);
        }
        None
    }
    //剑指 Offer II 045. 二叉树最底层最左边的值
    pub fn find_bottom_left_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        use std::collections::VecDeque;
        let mut queue = VecDeque::new();
        queue.push_back(root.as_ref().unwrap().clone());
        let mut res = 0;
        while !queue.is_empty() {
            let l = queue.len();
            for i in 0..l {
                let node = queue.pop_front().unwrap();
                if i == 0 {
                    res = node.borrow().val;
                }
                if node.borrow().left.is_some() {
                    queue.push_back(node.borrow().left.as_ref().unwrap().clone());
                }
                if node.borrow().right.is_some() {
                    queue.push_back(node.borrow().right.as_ref().unwrap().clone());
                }
            }
        }
        res
    }
    //剑指 Offer II 051. 节点之和最大的路径
    pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //思路：递归。
        //以根节点为起始的最大路径。
        //以根节点为中心的最大路径。
        if root.is_none() {
            return 0;
        }
        //左 + 中 + 右，与起始的相比
        let (start_max, mid_max) = Self::max_path_start_root(&root);
        start_max.max(mid_max)
    }
    fn max_path_start_root(root: &Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
        if root.is_none() {
            return (-1001, -1001);
        }
        let cur_val = root.as_ref().unwrap().borrow().val;
        //看它左边的大，还是右边的大，
        let left = Self::max_path_start_root(&root.as_ref().unwrap().borrow().left);
        let right = Self::max_path_start_root(&root.as_ref().unwrap().borrow().right);
        //以root为起点的最大值为，若左右两边为负，可以都不加
        let start_max = 0.max(left.0.max(right.0)) + cur_val;
        //结果的最大值为：
        let mid_max = start_max.max(left.0 + right.0 + cur_val).max(left.1).max(right.1);
        (start_max, mid_max)
    }
    //剑指 Offer II 044. 二叉树每层的最大值
    pub fn largest_values(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = vec![];
        let mut queue = std::collections::VecDeque::new();
        if root.is_none() {
            return res;
        }
        queue.push_back(root.as_ref().unwrap().clone());
        while !queue.is_empty() {
            let len = queue.len();
            let mut max = i32::MIN;
            for _ in 0..len {
                let node = queue.pop_front().unwrap();
                max = max.max(node.borrow().val);
                if node.borrow().left.is_some() {
                    queue.push_back(node.borrow().left.as_ref().unwrap().clone());
                }
                if node.borrow().right.is_some() {
                    queue.push_back(node.borrow().right.as_ref().unwrap().clone());
                }
            }
            res.push(max);
        }
        res
    }
    //973. 最接近原点的 K 个点
    pub fn k_closest(mut points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        points.sort_unstable_by_key(|v| v[0] * v[0] + v[1] * v[1]);
        points
            .into_iter()
            .take(k as usize)
            .collect()
    }
    //122. 买卖股票的最佳时机 II
    pub fn max_profit2(prices: Vec<i32>) -> i32 {
        let mut res = 0;
        let mut buy_res = -prices[0];
        for i in 0..prices.len() {
            buy_res = buy_res.max(res - prices[i]);
            res = res.max(buy_res + prices[i]);
        }
        res
    }
    //121. 买卖股票的最佳时机
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        //写一个dp,表示后面的最大值
        let mut dp = vec![0; prices.len()];
        let mut res = 0;
        for i in (0..prices.len() - 1).rev() {
            dp[i] = dp[i + 1].max(prices[i + 1]);
            res = res.max(dp[i] - prices[i]);
        }
        res
    }
    //48. 旋转图像
    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();
        for i in 0..n / 2 {
            for j in 0..(n + 1) / 2 {
                let temp = matrix[i][j];
                matrix[i][j] = matrix[n - 1 - j][i];
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j];
                matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i];
                matrix[j][n - 1 - i] = temp;
            }
        }
    }
    //258. 各位相加
    pub fn add_digits(num: i32) -> i32 {
        if num < 9 {
            return num;
        }
        if num % 9 == 0 {
            return 9;
        }
        return num % 9;
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

//173. 二叉搜索树迭代器
struct BSTIterator {
    index: usize,
    datas: Vec<i32>,
}

impl BSTIterator {
    fn new(root: Option<Rc<RefCell<TreeNode>>>) -> Self {
        let mut res = Self {
            index: 0,
            datas: vec![],
        };
        res.dfs(&root);
        res
    }

    fn dfs(&mut self, root: &Option<Rc<RefCell<TreeNode>>>) {
        if root.is_some() {
            self.dfs(&root.as_ref().unwrap().borrow().left);
            self.datas.push(root.as_ref().unwrap().borrow().val);
            self.dfs(&root.as_ref().unwrap().borrow().right);
        }
    }


    fn next(&mut self) -> i32 {
        self.index += 1;
        self.datas[self.index - 1]
    }

    fn has_next(&self) -> bool {
        self.index < self.datas.len()
    }
}

//剑指 Offer II 048. 序列化与反序列化二叉树
struct Codec {}

impl Codec {
    fn new() -> Self {
        Self {}
    }

    fn serialize(&self, root: Option<Rc<RefCell<TreeNode>>>) -> String {
        //[1,null,1]
        let mut res = String::new();
        res.push('[');
        //采用层级遍历
        use std::collections::VecDeque;
        let mut queue = VecDeque::new();
        queue.push_back(root);
        while !queue.is_empty() {
            let len = queue.len();
            for _ in 0..len {
                let cur = queue.pop_front().unwrap();
                if let Some(mut node) = cur {
                    res.push_str(node.borrow().val.to_string().as_str());
                    res.push(',');
                    queue.push_back(node.borrow_mut().left.take());
                    queue.push_back(node.borrow_mut().right.take());
                } else {
                    res.push_str("nil,");
                }
            }
        }
        res.remove(res.len() - 1);
        res.push(']');
        res
    }

    fn deserialize(&self, data: String) -> Option<Rc<RefCell<TreeNode>>> {
        let data = data.replace('[', "");
        let data = data.replace(']', "");
        let data = data.split(',').collect::<Vec<_>>();
        if data[0].eq("nil") {
            return None;
        }
        let mut root = Some(Rc::new(RefCell::new(TreeNode::new(data[0].parse().unwrap()))));
        use std::collections::VecDeque;
        let mut queue = VecDeque::new();
        queue.push_back(root.as_ref().unwrap().clone());
        for i in (1..data.len()).step_by(2) {
            let parent = queue.pop_front().unwrap();
            for j in 0..2 {
                if !data[i + j].eq("nil") {
                    let val: i32 = data[i + j].parse().unwrap();
                    let mut node = Some(Rc::new(RefCell::new(TreeNode::new(val))));
                    queue.push_back(node.as_ref().unwrap().clone());
                    //该节点挂在哪个下面
                    if j == 0 {
                        parent.borrow_mut().left = node;
                    } else {
                        parent.borrow_mut().right = node;
                    }
                }
            }
        }
        root
    }
}

//剑指 Offer II 043. 往完全二叉树添加节点
struct CBTInserter {
    root: Rc<RefCell<TreeNode>>,
    //下一个要插入的地方
    pre_nodes: std::collections::VecDeque<Rc<RefCell<TreeNode>>>,
}

impl CBTInserter {
    fn new(root: Option<Rc<RefCell<TreeNode>>>) -> Self {
        let root = root.as_ref().unwrap().clone();
        //倒数第二层的，只要最后一层。
        let mut queue = std::collections::VecDeque::new();
        queue.push_back(root.clone());
        let mut pre_nodes = std::collections::VecDeque::new();
        while !queue.is_empty() {
            let len = queue.len();
            for _ in 0..len {
                let node = queue.pop_front().unwrap();
                //如果node的左右节点有一个为空，则将node加入进pre
                if node.borrow().left.is_none() || node.borrow().right.is_none() {
                    pre_nodes.push_back(node.clone());
                }
                if node.borrow().left.is_some() {
                    queue.push_back(node.borrow().left.as_ref().unwrap().clone());
                }
                if node.borrow().right.is_some() {
                    queue.push_back(node.borrow().right.as_ref().unwrap().clone());
                }
            }
        }
        Self {
            root,
            pre_nodes,
        }
    }

    fn insert(&mut self, v: i32) -> i32 {
        let node = Rc::new(RefCell::new(TreeNode::new(v)));
        self.pre_nodes.push_back(node.clone());
        let top = self.pre_nodes[0].clone();
        let mut res = top.borrow().val;
        if top.borrow().left.is_none() {
            top.borrow_mut().left = Some(node);
        } else {
            top.borrow_mut().right = Some(node);
            self.pre_nodes.pop_front();
        }
        res
    }

    fn get_root(&self) -> Option<Rc<RefCell<TreeNode>>> {
        Some(self.root.clone())
    }
}

struct Solution;