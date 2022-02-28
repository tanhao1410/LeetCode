fn main() {
    println!("Hello, world!");
    let mut list = MyLinkedList::new();
    list.add_at_head(0);
    list.add_at_index(1, 4);
    list.add_at_tail(8);
    list.add_at_head(5);
    list.add_at_index(4, 3);
    list.add_at_tail(0);
    list.add_at_tail(5);
    list.add_at_index(6, 3);
    list.delete_at_index(7);
    list.delete_at_index(5);
    list.add_at_tail(4);
}

//707. 设计链表
struct MyLinkedList {
    head: Option<Box<ListNode>>,
    tail: *mut ListNode,
    size: i32,
}

impl MyLinkedList {
    fn new() -> Self {
        Self {
            head: None,
            tail: std::ptr::null_mut(),
            size: 0,
        }
    }

    fn get(&self, index: i32) -> i32 {
        if index >= self.size || self.size == 0 {
            return -1;
        }
        let mut p = &self.head;
        for _ in 0..index {
            p = &p.as_ref().unwrap().next;
        }
        p.as_ref().unwrap().val
    }

    fn add_at_head(&mut self, val: i32) {
        let mut node = Some(Box::new(ListNode::new(val)));
        node.as_mut().unwrap().next = self.head.take();
        self.head = node;
        if self.tail.is_null() {
            self.tail = self.head.as_mut().unwrap().as_mut() as *mut _;
        }
        self.size += 1;
    }

    fn add_at_tail(&mut self, val: i32) {
        if self.size == 0 {
            self.add_at_head(val);
        } else {
            let mut node = Some(Box::new(ListNode::new(val)));
            unsafe {
                (*self.tail).next = node;
                self.tail = (*self.tail).next.as_mut().unwrap().as_mut() as *mut _;
            }
            self.size += 1;
        }
    }

    fn add_at_index(&mut self, index: i32, val: i32) {
        if index == self.size {
            self.add_at_tail(val);
        } else if index == 0 {
            self.add_at_head(val);
        } else if index < self.size {
            let mut p = self.head.as_mut().unwrap();
            for _ in 0..index - 1 {
                p = p.next.as_mut().unwrap();
            }
            let temp = p.next.take();
            let mut node = Some(Box::new(ListNode::new(val)));
            node.as_mut().unwrap().next = temp;
            p.next = node;
            self.size += 1;
        }
    }

    fn delete_at_index(&mut self, index: i32) {
        if index < self.size && self.size > 0 {
            if index == 0 {
                //删除头结点
                self.head = self.head.as_mut().unwrap().next.take();
            } else {
                let mut p = self.head.as_mut().unwrap();
                for _ in 0..index - 1 {
                    p = p.next.as_mut().unwrap();
                }
                p.next = p.next.as_mut().unwrap().next.take();
                //删除最后一个的时候需要处理下
                if index == self.size - 1 {
                    self.tail = p.as_mut() as *mut _;
                }
            }
            self.size -= 1;
            if self.size == 0 {
                self.tail = std::ptr::null_mut();
            }
        }
    }
}

impl Solution {
    //25. K 个一组翻转链表
    pub fn reverse_k_group(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        //求长度
        let list_len = |mut head: &Option<Box<ListNode>>| -> i32{
            let mut res = 0;
            while let Some(node) = head {
                head = &node.next;
                res += 1;
            }
            res
        };
        //翻转
        let reverse_list = |mut head: Option<Box<ListNode>>| -> Option<Box<ListNode>>{
            let mut res = None;
            while let Some(mut node) = head {
                head = node.next.take();
                node.next = res;
                res = Some(node);
            }
            res
        };
        let list_len = list_len(&head);
        let mut res_head = Some(Box::new(ListNode::new(0)));
        let mut res_tail = res_head.as_mut().unwrap();

        //走k步，循环多少次呢?
        for _ in 0..list_len / k {
            let mut p = head.as_mut().unwrap();
            //走k步
            for _ in 0..k - 1 {
                let next = p.next.as_mut().unwrap();
                p = next;
            }
            //需要翻转的是
            //先把后面的需要取出来
            let tail = p.next.take();
            let reversed = reverse_list(head);
            res_tail = Self::append_list(res_tail, reversed);
            head = tail;
        }
        //最后剩的没翻转的需要补充上去
        res_tail.next = head;
        res_head.unwrap().next.take()
    }

    fn append_list(mut tail: &mut Box<ListNode>, other: Option<Box<ListNode>>) -> &mut Box<ListNode> {
        tail.next = other;
        while tail.next.is_some() {
            tail = tail.next.as_mut().unwrap();
        }
        tail
    }
    //97. 交错字符串
    pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
        if s1.len() + s2.len() != s3.len() {
            return false;
        }
        let (s1, s2, s3) = (s1.as_bytes(), s2.as_bytes(), s3.as_bytes());
        //动态规划，两重循环是否可行呢？s1的前i个元素，与s2的前j 个元素
        let mut dp = vec![vec![false; s2.len() + 1]; s1.len() + 1];
        for i in 0..dp.len() {
            for j in 0..dp[0].len() {
                if i == 0 && j == 0 {
                    dp[i][0] = true;
                } else if i == 0 {
                    dp[i][j] = dp[i][j - 1] && s2[j - 1] == s3[j - 1];
                } else if j == 0 {
                    dp[i][j] = dp[i - 1][j] && s1[i - 1] == s3[i - 1];
                } else {
                    dp[i][j] |= dp[i][j - 1] && s2[j - 1] == s3[j + i - 1];
                    dp[i][j] |= dp[i - 1][j] && s1[i - 1] == s3[i + j - 1];
                }
            }
        }
        dp[s3.len()][s1.len()]
    }
    //1823. 找出游戏的获胜者
    pub fn find_the_winner(n: i32, k: i32) -> i32 {
        let mut fails = vec![false; n as usize];
        let mut fail_count = 0;
        let mut start = 0;
        while fail_count < n - 1 {
            //第一步是走在了start位置，所以，此处用k - 1
            for _ in 0..k - 1 {
                //往前走一步
                start = (start + 1) % n;
                //但此时有可能走在了已经淘汰过的位置上了,如果是，继续往前走，直到走到没有淘汰的地方
                while fails[start as usize] {
                    start = (start + 1) % n;
                }
            }
            //此时，start即要被淘汰的地方
            fails[start as usize] = true;
            fail_count += 1;
            //走到下一个开始位置
            start = (start + 1) % n;
            while fails[start as usize] {
                start = (start + 1) % n;
            }
        }
        start
    }
    //67. 二进制求和
    pub fn add_binary(a: String, b: String) -> String {
        let bytes1 = a.as_bytes();
        let bytes2 = b.as_bytes();
        let mut res = vec![b'1'; a.len().max(b.len()) + 1];
        let n = res.len();
        let mut index = 0;
        let mut flag = 0;
        while index < a.len() || index < b.len() {
            let mut item = 0;
            if (index < a.len() && index < b.len()) {
                item = bytes1[a.len() - 1 - index] - b'0' + bytes2[b.len() - 1 - index] - b'0' + flag;
            } else if (index < a.len()) {
                item = bytes1[a.len() - 1 - index] - b'0' + flag;
            } else {
                item = bytes2[b.len() - 1 - index] - b'0' + flag;
            }
            res[n - 1 - index] = item % 2 + b'0';
            flag = item / 2;
            index += 1;
        }
        if flag == 1 {
            String::from_utf8(res).unwrap()
        } else {
            String::from_utf8_lossy(&res[1..]).to_string()
        }
    }
    //1601. 最多可达成的换楼请求数目
    pub fn maximum_requests(n: i32, requests: Vec<Vec<i32>>) -> i32 {
        //使用java用递归解决了，用栈来解决呢？深度优先遍历
        //用一个数来表示哪些请求使用了或没使用，因为最多16个，所以，一个i32的数足矣
        //(请求的序号，是否要，已经要了的请求）
        let mut stack = vec![(0, 1, 1), (0, 0, 0)];
        //总共16个深度，每步有两种走法，选择或不选择
        let mut res = 0;
        while let Some((index, used, count)) = stack.pop() {
            //走到头了
            if index == requests.len() - 1 {
                let mut item = 0;
                let mut buildings = vec![0; n as usize];
                //还需要判断是否满足要求
                for i in 0..requests.len() {
                    //该位被使用了
                    if used & (1 << i) > 0 {
                        item += 1;
                        buildings[requests[i][0] as usize] -= 1;
                        buildings[requests[i][1] as usize] += 1;
                    }
                }
                if buildings.into_iter().all(|v| v == 0) {
                    res = res.max(item);
                }
            } else {
                stack.push((index + 1, used | (1 << index + 1), count + 1));
                stack.push((index + 1, used, count));
            }
        }
        res
    }
}

struct Solution;

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