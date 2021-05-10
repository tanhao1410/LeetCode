use std::rc::Rc;
use std::cell::RefCell;
fn main() {
    println!("Hello, world!");
}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

//449. 序列化和反序列化二叉搜索树
struct Codec {}

impl Codec {
    fn new() -> Self {
        Codec {}
    }

    fn serialize(&self, root: Option<Rc<RefCell<TreeNode>>>) -> String {
        let mut middle_nums = Vec::new();
        Self::middle_read_tree(&root, &mut middle_nums);
        middle_nums.iter().fold("".to_string(), |mut pre, &pro| {
            pre.push(',');
            pre.push_str(pro.to_string().as_str());
            pre
        })
    }

    fn middle_read_tree(root: &Option<Rc<RefCell<TreeNode>>>, nums: &mut Vec<i32>) {
        let node = root.as_ref();
        if let Some(node) = node {
            nums.push(node.borrow().val);
            Self::middle_read_tree(&node.borrow().left, nums);
            Self::middle_read_tree(&node.borrow().right, nums);
        }
    }

    fn deserialize(&self, data: String) -> Option<Rc<RefCell<TreeNode>>> {
        let pre_order = data.split(",").filter(|s| { s.len() > 0 }).map(|s| {
            s.parse::<i32>().unwrap()
        }).collect::<Vec<i32>>();

        //中序遍历 是有序的
        let mut middle_order = pre_order.clone();
        middle_order.sort();

        //根据先序遍历和中序遍历构造树即可
        Self::create_tree_from_pre_and_middle(&middle_order, &pre_order)
    }

    fn create_tree_from_pre_and_middle(middle_order: &[i32], pre_order: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        if !middle_order.is_empty() {
            let mut middle_index = 0;
            for i in 0..middle_order.len() {
                if middle_order[i] == pre_order[0] {
                    middle_index = i;
                    break;
                }
            }

            let (middle_left, middle_right) = middle_order.split_at(middle_index);
            let middle_right = &middle_right[1..];

            let (pre_left, pre_right) = pre_order.split_at(middle_index + 1);
            let pre_left = &pre_left[1..];

            return Some(Rc::new(RefCell::new(TreeNode {
                val: pre_order[0],
                left: Self::create_tree_from_pre_and_middle(middle_left, pre_left),
                right: Self::create_tree_from_pre_and_middle(middle_right, pre_right),
            })));
        }
        None
    }
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

    //1302. 层数最深叶子节点的和
    pub fn deepest_leaves_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //层次遍历，
        let mut queue = vec![root];
        let mut res = 0;
        while queue.len() > 0 {
            let queue_len = queue.len();
            res = 0;
            for i in 0..queue_len {
                if let Some(node) = queue.remove(0) {
                    res += node.borrow().val;
                    if node.borrow().left.is_some() {
                        queue.push(node.borrow_mut().left.take());
                    }
                    if node.borrow().right.is_some() {
                        queue.push(node.borrow_mut().right.take());
                    }
                }
            }
        }
        res
    }

    //872. 叶子相似的树
    pub fn leaf_similar(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let mut leaf1: Vec<i32> = vec![];
        let mut leaf2: Vec<i32> = vec![];

        Self::read_tree(root1, &mut leaf1);
        Self::read_tree(root2, &mut leaf2);

        leaf1.eq(&leaf2)
    }
}