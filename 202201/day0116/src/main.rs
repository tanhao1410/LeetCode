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

    //310. 最小高度树
    pub fn find_min_height_trees(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        use std::collections::HashSet;
        use std::collections::HashMap;

        if n == 1 {
            return vec![0];
        }

        let mut map = HashMap::<i32, HashSet<i32>>::new();
        for edge in edges {
            let p1 = edge[0];
            let p2 = edge[1];
            let entry1 = map.entry(p1).or_insert(HashSet::new());
            entry1.insert(p2);
            let entry2 = map.entry(p2).or_insert(HashSet::new());
            entry2.insert(p1);
        }
        //去除只有一个连接的叶子节点
        while map.len() > 2 {
            //要删除哪些
            let remove_keys = map
                .iter()
                .filter_map(|(k, v)| match v.len() {
                    1 => Some(*k),
                    _ => None
                })
                .collect::<Vec<i32>>();

            for remove_key in remove_keys{
                let set = map.remove(&remove_key).unwrap();
                let other_point = set.iter().next().unwrap();
                //找到对方
                map.get_mut(other_point).unwrap().remove(&remove_key);
            }
        }

        map
            .keys()
            .map(|v| *v)
            .collect()
    }

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

