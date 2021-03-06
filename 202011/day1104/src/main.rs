fn main() {
    println!("Hello, world!");
    println!("{:?}", Solution::insert(vec![vec![1,5]], vec![0, 1]));
    println!("{}",Solution::is_match("".to_string(),"b*a*".to_string()))
}

impl Solution {
    //57. 插入区间
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        //思路：先从前找，找到一个结尾比插入区间大或等于的,如果都找不到，说明不存在，直接插入原来的即可
        let mut res = Vec::new();
        if intervals.len() == 0{
            res.push(new_interval);
            return res;
        }
        //插入区间比所有区间小，之间放在前面即可
        let (start, end) = (new_interval[0], new_interval[1]);
        if end < intervals[0][0]{
            res.push(new_interval);
            res.append(&mut intervals.clone());
            return res;
        }

        //找到区间的第一个交集
        let mut i = 0;
        while i < intervals.len() && intervals[i][1] < start {
            res.push(intervals[i].clone());
            i += 1;
        }

        //插入的区间大于任何区间
        if i == intervals.len() {
            res.push(new_interval);
            return res;
        }

        let mut new_start = intervals[i][0];
        if start < new_start {
            new_start = start;
        }

        //下一步找下一个区间开头比end大于或等于的
        let mut is_same = true;//开始区间和结束区间是同一个
        while i < intervals.len() && intervals[i][0] < end {
            i += 1;
            is_same = false;
        }

        //插入区间的末尾大于任何区间的开始数
        if i == intervals.len() {
            if end <= intervals[i - 1][1] {
                res.push(vec![new_start, intervals[i - 1][1]]);
            } else {
                res.push(vec![new_start, end]);
            }
            return res;
        }else if intervals[i][0] < start && end < intervals[i][1] && is_same{
            //在某区间的内部
            res.push(new_interval);
            new_start = intervals[i][0];
        }else if i > 0 && intervals[i -1][1] < start && end < intervals[i][0]{
            //与任何区间无关联
            res.push(new_interval);
            new_start = intervals[i][0];
        }

        let new_end;
        if is_same {
            //说明新插入的区间在原某区间内部
            new_end = intervals[i][1];
            i+=1;
        } else if intervals[i][0] == end {
            new_end = intervals[i][1];
            i += 1;//因为这次加入相当于把该区间加入进来了，下面不用重复加入
        } else {
            if end < intervals[i - 1][1] {
                new_end = intervals[i - 1][1];
            } else {
                new_end = end;
            }
        }
        res.push(vec![new_start, new_end]);
        //剩余的区间加入返回集合
        while i < intervals.len() {
            res.push(intervals[i].clone());
            i += 1;
        }
        res
    }

    //剑指 Offer 19. 正则表达式匹配
    pub fn is_match(s: String, mut p: String) -> bool {
        if s.is_empty() {
            //c*c*这种
            while !p.is_empty(){
                if let Some('*') = p.pop(){
                    p.pop();
                }else{
                    return false;
                }
            }
            return true;
        }
        //先判断第一个字符
        let first_char = s.chars().nth(0).unwrap();
        match p.chars().nth(0) {
            Some(p_first) if (p_first == '.' || p_first == first_char) && Some('*') == p.chars().nth(1) => {
                //*取任意数时候
                let mut i = 1;
                while Some(first_char) == s.chars().nth(i) || (p_first == '.' && s.chars().nth(i).is_some()) {
                    if Solution::is_match(s[i..].to_string(), p[2..].to_string()) {
                        return true;
                    }
                    i += 1;
                }
                //下一个判别
                if i > s.len() {
                    i = s.len();
                }
                Solution::is_match(s[i..].to_string(), p[2..].to_string())
                    || Solution::is_match(s.clone(), p[2..].to_string()) //当*取0个时进行判别
            }
            Some(p_first) if p_first == '.' || p_first == first_char =>
                Solution::is_match(s[1..].to_string(), p[1..].to_string()),
            Some(_) if Some('*') == p.chars().nth(1) =>
                Solution::is_match(s.clone(), p[2..].to_string()),
            _ => false
        }
    }
    //剑指 Offer 34. 二叉树中和为某一值的路径
    pub fn path_sum(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        //直到叶节点
        if let Some(node) = root {
            //该节点为叶子节点
            if node.borrow().left.is_none() && node.borrow().right.is_none() {
                if sum == node.borrow().val {
                    res.push(vec![sum]);
                    return res;
                }
                return res;
            }
            match node.borrow().right.clone() {
                Some(right) => {
                    let mut right_vec = Solution::path_sum(node.borrow().right.clone(), sum - node.borrow().val);
                    for i in 0..right_vec.len() {
                        let mut item = vec![node.borrow().val];
                        if right_vec[i].is_empty() {
                            continue;
                        }
                        item.append(&mut right_vec[i]);
                        res.push(item);
                    }
                }
                _ => {}
            }
            if node.borrow().left.is_some() {
                let mut left_vec = Solution::path_sum(node.borrow().left.clone(), sum - node.borrow().val);
                for i in 0..left_vec.len() {
                    if left_vec[i].is_empty() {
                        continue;
                    }
                    let mut item = vec![node.borrow().val];
                    item.append(&mut left_vec[i]);
                    res.push(item);
                }
            }
        }
        res
    }
}

struct Solution {}
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
            right: None,
        }
    }
}