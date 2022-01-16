///动态分发与静态分发时间性能测试
fn main() {}


///382. 链表随机节点
/// 如果链表非常大且长度未知，该怎么处理？
// 你能否在不使用额外空间的情况下解决此问题？

// 思路1：计算链表长度，保留链表首，每次随你一个数，从里面拿取，缺点，每一次getRandom都是O(N)
// 思路2: 使用额外空间，一个vec保存，随你后直接获取。缺点，空间使用率
// 思路3：看题解，蓄水池解法
struct Solution {
    size: usize,
    nums: Vec<i32>,
}

impl Solution {
    fn new(mut head: Option<Box<ListNode>>) -> Self {
        let mut nums = vec![];
        let mut size = 0;
        while let Some(mut node) = head {
            nums.push(node.val);
            size += 1;
            head = node.next.take();
        }
        Self {
            size,
            nums,
        }
    }

    fn get_random(&self) -> i32 {
        // let rng = rand::thread_rng();
        // rng.gen_ratio()
        let mut rand_num = 0;
        unsafe {
            use std::arch::x86_64::_rdrand32_step;
            _rdrand32_step(&mut rand_num);
        }


        self.nums[rand_num as usize % self.size]
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

