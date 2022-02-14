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
}

struct Solution;

use std::rc::Rc;
use std::cell::RefCell;


impl Solution {
    //572. 另一棵树的子树
    pub fn is_subtree(root: Option<Rc<RefCell<TreeNode>>>, sub_root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        Self::is_subtree_ref(&root, &sub_root)
    }

    pub fn is_subtree_ref(root: &Option<Rc<RefCell<TreeNode>>>, sub_root: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        if let Some(root) = root {
            if root.borrow().val == sub_root.as_ref().unwrap().borrow().val
                && Self::is_equal(&root.borrow().left, &sub_root.as_ref().unwrap().borrow().left)
                && Self::is_equal(&root.borrow().right, &sub_root.as_ref().unwrap().borrow().right) {
                return true;
            }
            return Self::is_subtree_ref(&root.borrow().left, sub_root)
                || Self::is_subtree_ref(&root.borrow().right, sub_root);
        }
        false
    }

    pub fn is_equal(root: &Option<Rc<RefCell<TreeNode>>>, sub_root: &Option<Rc<RefCell<TreeNode>>>) -> bool {
        if sub_root.is_none() {
            return root.is_none();
        }
        if root.is_none() || root.as_ref().unwrap().borrow().val != sub_root.as_ref().unwrap().borrow().val {
            return false;
        }
        Self::is_equal(&root.as_ref().unwrap().borrow().left, &sub_root.as_ref().unwrap().borrow().left)
            && Self::is_equal(&root.as_ref().unwrap().borrow().right, &sub_root.as_ref().unwrap().borrow().right)
    }
}

//540. 有序数组中的单一元素
pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
    //思路，如果就一个元素，直接返回，否则，肯定是奇数个。从中间开始，看该数是否是单个，如果是单个，返回。
    //否则，如果和前面相同，
    // 用递归的方式，更简单
    if nums.len() == 1 {
        return nums[0];
    }
    let mut start = 0;
    let mut end = nums.len() - 1;
    let mut mid = (end + start) / 2;
    while start < end {
        //看中间的和前面相等还是后面相等
        if nums[mid] == nums[mid - 1] {
            if mid % 2 == 0 {
                //在前面
                end = mid - 2;
            } else {
                //在后面
                start = mid + 1;
            }
        } else if nums[mid] == nums[mid + 1] {
            if mid % 2 == 0 {
                //在后面
                start = mid + 2;
            } else {
                end = mid - 1;
            }
        } else {
            return nums[mid];
        }
        mid = (end + start) / 2;
    }
    nums[mid]
}