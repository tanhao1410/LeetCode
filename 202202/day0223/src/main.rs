fn main() {
    println!("Hello, world!");
    println!("{}", Solution::multiply("9".to_string(), "9".to_string()));
}

struct Solution;

impl Solution {
    //剑指 Offer II 021. 删除链表的倒数第 n 个结点
    pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
        ////两个指针，一个先走n步，然后再走一步，第二个指针接着走，若第一个指针为最后一个元素了，则第二个指针的下一个元素即要删除的元素
        //只有一个元素的时候
        if head.as_ref().unwrap().next.is_none() {
            return None;
        }
        //先走一步
        let mut fast = head.as_ref().unwrap();
        //再走N - 1步
        for _ in 1..n {
            fast = fast.next.as_ref().unwrap();
        }
        //再走1步，走不动的情况下，说明移除的是第一个元素
        if fast.next.is_none() {
            return head.unwrap().next;
        }
        fast = fast.next.as_ref().unwrap();
        //慢的走一步
        let mut slow = head.as_ref().unwrap();
        while fast.next.is_some() {
            fast = fast.next.as_ref().unwrap();
            slow = slow.next.as_ref().unwrap();
        }
        //slow指针指向的即是要删除的，slow.next = slow.next.next
        //slow是不可变指针
        let mut slow = slow.as_ref() as *const ListNode as *mut ListNode;
        unsafe {
            let remove_node = (*slow).next.take();
            (*slow).next = remove_node.unwrap().next.take();
        }
        head
    }
    //673. 最长递增子序列的个数
    pub fn find_number_of_lis(nums: Vec<i32>) -> i32 {
        //最长子数组长度，形成这样的子数组的个数
        let mut dp = vec![(1, 1); nums.len()];
        let mut max = 1;
        for i in 1..nums.len() {
            let mut item = (0, 0);
            for j in 0..i {
                if nums[i] > nums[j] {
                    if dp[j].0 > item.0 {
                        item = dp[j];
                    } else if dp[j].0 == item.0 {
                        item.1 += dp[j].1;
                    }
                }
            }
            dp[i] = (item.0 + 1, item.1);
            max = max.max(dp[i].0);
        }
        dp
            .into_iter()
            .filter(|&e| e.0 == max)
            .map(|e| e.1)
            .sum()
    }
    //5. 最长回文子串
    pub fn longest_palindrome(s: String) -> String {
        let mut dp = vec![vec![false; s.len()]; s.len()];
        //dp[i][j]
        let bytes = s.as_bytes();
        for i in 0..s.len() - 1 {
            dp[i][i] = true;
        }
        let mut locaion = (0, 0);
        for i in (0..s.len() - 1).rev() {
            for j in (i + 1..s.len()) {
                if bytes[i] == bytes[j] {
                    if i + 1 == j {
                        dp[i][j] = true;
                    } else if dp[i + 1][j - 1] {
                        dp[i][j] = true;
                    }
                }
                if dp[i][j] && j - i > locaion.1 - locaion.0 {
                    locaion = (i, j);
                }
            }
        }
        String::from_utf8_lossy(&bytes[locaion.0..locaion.1 + 1]).to_string()
    }

    //917. 仅仅反转字母
    pub fn reverse_only_letters(s: String) -> String {
        let mut bytes = s.clone().into_bytes();
        let mut start = 0;
        let mut end = bytes.len() - 1;
        let is_letter = |b: u8| {
            (b <= b'z' && b >= b'a') || (b >= b'A' && b <= b'Z')
        };
        while end > start {
            while start < end && !is_letter(bytes[start]) {
                start += 1;
            }
            while end > start && !is_letter(bytes[end]) {
                end -= 1;
            }
            if end > start {
                let temp = bytes[start];
                bytes[start] = bytes[end];
                bytes[end] = temp;
                end -= 1;
                start += 1;
            }
        }
        String::from_utf8(bytes).unwrap()
    }
    //43. 字符串相乘
    pub fn multiply(num1: String, num2: String) -> String {
        let multiply = |num1: &String, num2: i32, zero: usize| {
            let bytes = num1.as_bytes();
            let mut res = vec![];
            for _ in 0..zero {
                res.push(b'0');
            }
            let mut flag = 0;
            for i in (0..bytes.len()).rev() {
                let item_res = (bytes[i] - b'0') as i32 * num2 + flag;
                flag = item_res / 10;
                res.push((item_res % 10) as u8 + b'0');
            }
            if flag > 0 {
                res.push(b'0' + flag as u8);
            }
            res.reverse();
            String::from_utf8(res).unwrap().trim_start_matches('0').to_string()
        };
        let sum = |num1: &str, num2: &str| {
            let mut index = 0;
            let mut res = vec![];
            let mut flag = 0;
            let bytes1 = num1.as_bytes();
            let bytes2 = num2.as_bytes();
            while index < num1.len() && index < num2.len() {
                let item_res = bytes1[num1.len() - 1 - index] - b'0'
                    + bytes2[num2.len() - 1 - index] - b'0'
                    + flag;
                flag = item_res / 10;
                res.push(item_res % 10 + b'0');
                index += 1;
            }
            while index < num1.len() {
                let item_res = bytes1[num1.len() - 1 - index] - b'0' + flag;
                flag = item_res / 10;
                res.push(item_res % 10 + b'0');
                index += 1;
            }
            while index < num2.len() {
                let item_res = bytes2[num2.len() - 1 - index] - b'0' + flag;
                flag = item_res / 10;
                res.push(item_res % 10 + b'0');
                index += 1;
            }
            if flag > 0 {
                res.push(b'1');
            }
            res.reverse();
            let res = String::from_utf8(res).unwrap();
            println!("{} + {} = {}", num1, num2, res);
            res
        };
        let bytes = num2.as_bytes();
        let mut res = "0".to_string();
        for i in (0..bytes.len()).rev() {
            let cur_num = (bytes[i] - b'0') as i32;
            let mutiply_num = multiply(&num1, cur_num, bytes.len() - 1 - i);
            res = sum(&res, &mutiply_num);
            println!("{} * {} = {}", cur_num, num1, mutiply_num);
        }
        res
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