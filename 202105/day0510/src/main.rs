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

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
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