use std::cmp::Ordering;

fn main() {
    println!("Hello, world!");
}

struct Solution;

#[derive(Eq)]
struct CouplEle<'a> {
    first: usize,
    second: usize,
    nums1: &'a Vec<i32>,
    nums2: &'a Vec<i32>,
}

impl<'a> CouplEle<'a> {
    fn new(first: usize, second: usize, nums1: &'a Vec<i32>, nums2: &'a Vec<i32>) -> Self {
        Self {
            first,
            second,
            nums1,
            nums2,
        }
    }
}

impl<'a> Ord for CouplEle<'a> {
    fn cmp(&self, other: &Self) -> Ordering {
        (other.nums1[other.first] + other.nums2[other.second])
            .cmp(&(self.nums1[self.first] + self.nums2[self.second]))
    }
}

impl<'a> PartialOrd for CouplEle<'a> {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some((other.nums1[other.first] + other.nums2[other.second])
            .cmp(&(self.nums1[self.first] + self.nums2[self.second])))
    }
}

impl<'a> PartialEq for CouplEle<'a> {
    fn eq(&self, other: &Self) -> bool {
        other.nums1[other.first] + other.nums2[other.second] ==
            self.nums1[self.first] + self.nums2[self.second]
    }
}


impl Solution {
    //373. 查找和最小的K对数字
    pub fn k_smallest_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        use std::collections::BinaryHeap;

        let mut heap = BinaryHeap::new();
        let (m, n) = (nums1.len(), nums2.len());
        for num1_index in 0..m.min(k as usize) {
            heap.push(CouplEle::new(num1_index, 0, &nums1, &nums2));
        }
        let mut res = vec![];
        let mut k = k;
        loop {
            if k < 0 || heap.is_empty() {
                break;
            }
            if let CouplEle { first, second, .. } = heap.pop().unwrap() {
                res.push(vec![nums1[first], nums2[second]]);
                if second + 1 < n {
                    heap.push(CouplEle::new(first, second + 1, &nums1, &nums2));
                }
            }
        }
        res
    }

    pub fn k_smallest_pairs2(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> Vec<Vec<i32>> {
        use std::collections::BinaryHeap;
        let mut heap = BinaryHeap::new();
        let (m, n) = (nums1.len(), nums2.len());
        for num1_index in 0..m.min(k as usize) {
            heap.push(CouplEle::new(num1_index, 0, &nums1, &nums2));
        }
        let mut res = vec![];
        while let Some(CouplEle { first, second, .. }) = heap.pop() {
            res.push(vec![nums1[first], nums2[second]]);
            if res.len() == k as usize{
                break;
            }
            if second + 1 < n {
                heap.push(CouplEle::new(first, second + 1, &nums1, &nums2));
            }
        }
        res
    }
}