fn main() {
    println!("Hello, world!");
}
use std::rc::Rc;
use std::cell::RefCell;

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
      right: None
    }
  }
}
struct Solution {}

impl Solution {

    //每日一题：129.求根到叶子节点数字之和
    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        //思路：深度优先遍历法或递归法

        //
        fn dfs(root: Option<Rc<RefCell<TreeNode>>>,parent_val : i32) -> Vec<i32>{
            let mut nums = vec![];
            if let Some(node) = root{
                let val = node.borrow().val + 10 * parent_val;
                //得到左右根节点的
                if node.borrow().left.is_some() || node.borrow().right.is_some(){
                    let mut left_nums = dfs(node.borrow().left.clone(),  val);
                    let mut right_nums = dfs(node.borrow().right.clone(),val);
                    nums.append(&mut left_nums);
                    nums.append(&mut right_nums);
                }else{
                    nums.push(val);
                }
            }
            nums
        }
        let nums = dfs(root,0);
        let i = nums.iter().sum();
        i
    }

}
