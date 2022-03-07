fn main() {
    println!("Hello, world!");
}

impl Solution {
    //剑指 Offer II 005. 单词长度的最大乘积
    pub fn max_product(words: Vec<String>) -> i32 {
        //一个单词对应一个数，26位，分别表示是否含有a,b,c等等。如果两个单词生成的数 & == 0,说明不含同一个字母
        let as_int = |word: &String| {
            let mut res = 0;
            let mut letters = vec![false; 26];
            let bytes = word.as_bytes();
            for &b in bytes {
                letters[(b - b'a') as usize] = true;
            }
            for i in 0..26 {
                if letters[i] {
                    res |= (1 << i);
                }
            }
            res
        };
        let word_ints = words.iter().map(as_int).collect::<Vec<_>>();
        let mut res = 0;
        for i in 0..words.len() - 1 {
            for j in i..words.len() {
                if word_ints[i] & word_ints[j] == 0 {
                    res = res.max(words[i].len() * words[j].len())
                }
            }
        }
        res as i32
    }
    //剑指 Offer II 063. 替换单词
    pub fn replace_words(dictionary: Vec<String>, sentence: String) -> String {
        let mut trie = Trie::new();
        for word in dictionary {
            trie.insert(word);
        }
        sentence
            .split(" ")
            .map(|word| trie.preword(word))
            .collect::<Vec<_>>()
            .join(" ")
    }

    //剑指 Offer II 056. 二叉搜索树中两个节点之和
    pub fn find_target(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> bool {
        let mut v = vec![];
        Self::read_tree(&root, &mut v);
        let mut i = 0;
        let mut j = v.len() - 1;
        while j > i {
            if v[i] + v[j] == k {
                return true;
            } else if v[i] + v[j] > k {
                j -= 1;
            } else {
                i += 1;
            }
        }
        return false;
    }

    fn read_tree(root: &Option<Rc<RefCell<TreeNode>>>, v: &mut Vec<i32>) {
        if root.is_some() {
            Self::read_tree(&root.as_ref().unwrap().borrow().left, v);
            v.push(root.as_ref().unwrap().borrow().val);
            Self::read_tree(&root.as_ref().unwrap().borrow().right, v);
        }
    }
    //剑指 Offer II 012. 左右两边子数组的和相等
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        //先求总和，然后从零开始依次求和sum，若sum = (sum2 - cur )/2。返回结果
        let sum = nums.iter().sum::<i32>();
        let mut s = 0;
        for i in 0..nums.len() {
            let cur = nums[i];
            if s * 2 == sum - cur {
                return i as i32;
            }
            s += cur;
        }
        -1
    }
    //剑指 Offer II 038. 每日温度
    pub fn daily_temperatures(temperatures: Vec<i32>) -> Vec<i32> {
        let mut stack: Vec<(i32, usize)> = vec![];
        let mut res = vec![0; temperatures.len()];
        for i in (0..temperatures.len()).rev() {
            let cur = temperatures[i];
            while !stack.is_empty() && stack[stack.len() - 1].0 <= cur {
                stack.pop();
            }
            if !stack.is_empty() {
                res[i] = stack[stack.len() - 1].0 as i32 - i as i32;
            }
            stack.push((cur, i));
        }
        res
    }
    //剑指 Offer II 035. 最小时间差
    pub fn find_min_difference(time_points: Vec<String>) -> i32 {
        //每一分钟做个对应的话，也不过24*60个。需要注意的是存在循环。
        use std::collections::HashSet;
        let mut set = HashSet::new();
        let time_to_int = |time: &str| {
            let mut res = 0;
            let mut split = time.split(":");
            res += split.next().unwrap().parse::<i32>().unwrap() * 60;
            res += split.next().unwrap().parse::<i32>().unwrap();
            res
        };
        for time in time_points {
            let time = time_to_int(&time);
            if set.contains(&time) {
                return 0;
            }
            set.insert(time);
        }
        let mut times = set.into_iter().collect::<Vec<_>>();
        times.sort_unstable();
        let mut min = times[0] + 1440 - times[times.len() - 1];
        for i in 1..times.len() {
            min = min.min(times[i] - times[i - 1]);
        }
        min
    }
    //剑指 Offer II 036. 后缀表达式
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        use std::ops::Add;
        use std::ops::Mul;
        use std::ops::Div;
        use std::ops::Sub;
        let mut stack = vec![];
        let opera = |stack: &mut Vec<i32>, op: fn(i32, i32) -> i32| {
            let mut a = stack.pop().unwrap();
            let mut b = stack.pop().unwrap();
            stack.push(op(b, a));
        };
        for token in tokens {
            //如果是数就进栈
            match token.as_str() {
                "+" => opera(&mut stack, i32::add),
                "-" => opera(&mut stack, i32::sub),
                "*" => opera(&mut stack, i32::mul),
                "/" => opera(&mut stack, i32::div),
                _ => stack.push(token.parse().unwrap())
            }
        }
        stack.pop().unwrap()
    }
    //451. 根据字符出现频率排序
    pub fn frequency_sort(s: String) -> String {
        use std::collections::HashMap;
        let mut map = HashMap::new();
        s.chars().for_each(|c| *map.entry(c).or_insert(0) += 1);
        let mut v = map.into_iter().collect::<Vec<_>>();
        v.sort_unstable_by_key(|&e| -e.1);
        let mut res = String::new();
        for (c, count) in v {
            for _ in 0..count {
                res.push(c);
            }
        }
        res
    }
}

struct Solution;

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

use std::rc::Rc;
use std::cell::RefCell;

//剑指 Offer II 062. 实现前缀树
#[derive(Clone)]
struct Trie {
    child: Vec<Option<Trie>>,
    end_flag: bool,
}

impl Trie {
    fn new() -> Self {
        Self { child: vec![None; 26], end_flag: false }
    }

    fn insert(&mut self, word: String) {
        let bytes = word.as_bytes();
        let mut next = self;
        for &b in bytes {
            if next.child[(b - b'a') as usize].is_none() {
                next.child[(b - b'a') as usize] = Some(Trie::new());
            }
            next = next.child[(b - b'a') as usize].as_mut().unwrap();
        }
        next.end_flag = true;
    }

    fn search(&self, word: String) -> bool {
        let bytes = word.as_bytes();
        let mut next = self;
        for &b in bytes {
            if next.child[(b - b'a') as usize].is_some() {
                next = next.child[(b - b'a') as usize].as_ref().unwrap();
            } else {
                return false;
            }
        }
        next.end_flag
    }

    fn starts_with(&self, prefix: String) -> bool {
        let bytes = prefix.as_bytes();
        let mut next = self;
        for &b in bytes {
            if next.child[(b - b'a') as usize].is_some() {
                next = next.child[(b - b'a') as usize].as_ref().unwrap();
            } else {
                return false;
            }
        }
        true
    }

    pub fn preword(&self, word: &str) -> String {
        let bytes = word.as_bytes();
        let mut res = vec![];
        let mut next = self;
        for &b in bytes {
            //找到了前缀词
            if next.end_flag {
                return String::from_utf8(res).unwrap();
            }
            if next.child[(b - b'a') as usize].is_some() {
                res.push(b);
                next = next.child[(b - b'a') as usize].as_ref().unwrap();
            } else {
                break;
            }
        }
        word.to_string()
    }
}
