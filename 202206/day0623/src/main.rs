fn main() {
    println!("Hello, world!");
}

struct Solution;

impl Solution {
    //1387. 将整数按权重排序
    pub fn get_kth(lo: i32, hi: i32, k: i32) -> i32 {
        use std::collections::HashMap;
        let map = (0..31).map(|i| (1 << i, i)).collect::<HashMap<i32, i32>>();
        let become_one_count = |num: &i32| -> i32{
            let mut num = *num;
            let mut res = 0;
            while !map.contains_key(&num) {
                if num % 2 == 0 {
                    num /= 2;
                } else {
                    num = num * 3 + 1;
                }
                res += 1;
            }
            res + map.get(&num).unwrap()
        };
        let mut nums = (lo..=hi).collect::<Vec<i32>>();
        nums.sort_by_key(become_one_count);
        nums[k as usize - 1]
    }
}

use std::collections::HashMap;

//1146. 快照数组
struct SnapshotArray {
    snap_num: i32,
    datas: HashMap<i32, Vec<(i32, i32)>>,
}

impl SnapshotArray {
    fn new(length: i32) -> Self {
        Self { datas: HashMap::new(), snap_num: 0 }
    }

    fn set(&mut self, index: i32, val: i32) {
        let entry = self.datas.entry(index).or_insert(vec![]);
        //当前的snap中是否已经存了值了，如果存了，更新，没存，加入一个值。
        if let Some((snap_id, v)) = entry.last_mut() {
            if *snap_id == self.snap_num {
                *v = val;
                return;
            }
        }
        entry.push((self.snap_num, val));
    }
    fn snap(&mut self) -> i32 {
        self.snap_num += 1;
        self.snap_num - 1
    }

    fn get(&self, index: i32, snap_id: i32) -> i32 {
        self.datas
            .get(&index)
            .map_or(0, |v| v
                .into_iter()
                .take_while(|e| e.0 <= snap_id)
                .last()
                .unwrap_or(&(0, 0))
                .1,
            )
    }
}