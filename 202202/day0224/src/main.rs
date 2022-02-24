fn main() {
    println!("Hello, world!");
}

impl Solution {
    //435. 无重叠区间
    pub fn erase_overlap_intervals(mut intervals: Vec<Vec<i32>>) -> i32 {
        // 动态规划，选择尽量多的不重叠的区间,以区间 i 为最后一个区间，可以选出的区间数量的最大值
        // intervals.sort_unstable_by_key(|e| e[0]);
        // let mut dp = vec![1; intervals.len()];
        // for i in 1..dp.len() {
        //     for j in 0..i {
        //         if intervals[i][0] >= intervals[j][1] {
        //             dp[i] = dp[i].max(dp[j] + 1);
        //         }
        //     }
        // }
        // intervals.len() as i32 - dp.into_iter().max().unwrap()
        //贪心算法，按照右边进行排序，然后，依次判断是否可以加入进来，如果能加入进来，则+1，不能加入进来 ，继续下一个。
        intervals.sort_unstable_by_key(|e| e[1]);
        let mut right = intervals[0][1];
        let mut count = 1;
        for i in 1..intervals.len() {
            //大于最右边，说明可以加入进来
            if intervals[i][0] >= right {
                count += 1;
                right = intervals[i][1];
            }
        }
        intervals.len() as i32 - count
    }
    //496. 下一个更大元素 I
    pub fn next_greater_element(nums1: Vec<i32>, mut nums2: Vec<i32>) -> Vec<i32> {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        let mut stack = vec![];
        for i in (0..nums2.len()).rev() {
            let cur = nums2[i];
            map.insert(cur, i);
            while !stack.is_empty() && *stack.last().unwrap() < cur {
                stack.pop();
            }
            if let Some(top) = stack.last() {
                nums2[i] = *top;
            } else {
                nums2[i] = -1;
            }
            stack.push(cur);
        }
        nums1
            .into_iter()
            .map(|v| nums2[*map.get(&v).unwrap()])
            .collect()
    }
    //1706. 球会落何处
    pub fn find_ball(grid: Vec<Vec<i32>>) -> Vec<i32> {
        //球 怎么表示哪一个格子没了呢
        let mut balls = vec![(false, 0); grid[0].len()];
        for i in 0..balls.len() {
            balls[i] = (true, i);
        }
        Self::find_ball2(&grid, balls)
    }
    fn find_ball2(grid: &[Vec<i32>], balls: Vec<(bool, usize)>) -> Vec<i32> {
        //说明走到最后一层了，返回结果
        if grid.len() == 0 {
            let mut res = vec![-1; balls.len()];
            for i in 0..balls.len() {
                //有球，谁的球
                if balls[i].0 {
                    res[balls[i].1] = i as i32;
                }
            }
            return res;
        }
        //球的情况
        let mut new_balls = vec![(false, 0); balls.len()];
        //往下一层开始滚动
        for i in 0..grid[0].len() {
            //先判断有没有球
            if balls[i].0 {
                //往哪流呢
                if grid[0][i] == 1 {//右边
                    if i != balls.len() - 1 && grid[0][i + 1] != -1 {
                        new_balls[i + 1] = (true, balls[i].1);
                    }
                } else {
                    if i != 0 && grid[0][i - 1] != 1 {
                        new_balls[i - 1] = (true, balls[i].1);
                    }
                }
            }
        }
        Self::find_ball2(&grid[1..], new_balls)
    }
    //2. 两数相加
    pub fn add_two_numbers(mut l1: Option<Box<ListNode>>, mut l2: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        //每一次创建一个
        let mut head = Some(Box::new(ListNode::new(0)));
        let mut p = head.as_ref().unwrap().as_ref();
        let mut flag = 0;
        while l1.is_some() && l2.is_some() {
            let val = l1.as_ref().unwrap().val + l2.as_ref().unwrap().val + flag;
            flag = val / 10;
            let new_node = Some(Box::new(ListNode::new(val % 10)));
            unsafe {
                let p_ptr = p as *const ListNode as *mut ListNode;
                (*p_ptr).next = new_node;
                p = (*p_ptr).next.as_ref().unwrap().as_ref();
            }
            l1 = l1.as_mut().unwrap().next.take();
            l2 = l2.as_mut().unwrap().next.take();
        }
        let mut l3 = None;
        if l2.is_some() {
            l3 = l2;
        }
        if l1.is_some() {
            l3 = l1;
        }
        while let Some(mut node) = l3 {
            let val = node.val + flag;
            flag = val / 10;
            let new_node = Some(Box::new(ListNode::new(val % 10)));
            unsafe {
                let p_ptr = p as *const ListNode as *mut ListNode;
                (*p_ptr).next = new_node;
                p = (*p_ptr).next.as_ref().unwrap().as_ref();
            }
            l3 = node.next.take();
        }
        if flag > 0 {
            let new_node = Some(Box::new(ListNode::new(flag)));
            unsafe {
                let p_ptr = p as *const ListNode as *mut ListNode;
                (*p_ptr).next = new_node;
            }
        }
        head.as_mut().unwrap().next.take()
    }

    //20. 有效的括号
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![];
        for s in s.chars() {
            match s {
                '(' | '[' | '{' => stack.push(s),
                _ => {
                    if stack.is_empty() {
                        let pop = stack.pop().unwrap();
                        if (pop == '(' && s != ')') || (pop == '{' && s != '}') || (pop == '[' && s != ']') {
                            return false;
                        }
                    } else {
                        return false;
                    }
                }
            }
        }
        stack.len() == 0
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