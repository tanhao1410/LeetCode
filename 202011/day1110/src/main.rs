fn main() {
    println!("Hello, world!");
    println!("{:?}",Solution::next_permutation(&mut vec![3, 2, 1]));
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

struct Solution {}

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {

    //每日一题：31. 下一个排列
    pub fn next_permutation(nums: &mut Vec<i32>) {
        //思路：就是找对应的数值位，然后后面的按顺序摆放即可。
        let nums_size = nums.len();
        if nums_size < 2{
            return;
        }
        let mut i = nums_size as i32- 2;
        while i >= 0{
            //从后面找到一个比它大的最小的一个数
            let mut more_min = i  as usize;
            for j in i  as usize.. nums_size{
                if nums[j] > nums[i as usize]{
                    if more_min == i  as usize{
                        more_min = j;
                    }else if nums[j ] < nums[more_min ]{
                        more_min = j;
                    }
                }
            }
            if more_min != i  as usize{
                //把该数放置在i位置上，后面按顺序摆放即可
                let temp = nums[i as usize];
                nums[i as usize] = nums[more_min ];
                nums[more_min ] = temp;
                //后面的进行排序
                for j in i  as usize + 1 .. nums_size-1{
                    for k in j +1 ..nums_size{
                        if nums[k ] < nums[j ]{
                            let temp = nums[j ];
                            nums[j ] = nums[k];
                            nums[k ] = temp;
                        }
                    }
                }
                return;
            }else{
                //若不存在这样的数，则 i -= 1;如果直到结束都不存在,说明数组是递减的，反转即可。
                i -= 1;
            }
        }
        nums.reverse();
    }
}
