fn main() {
    println!("Hello, world!");
}

use std::collections::BTreeMap;

//732. 我的日程安排表 III
struct MyCalendarThree {
    items: BTreeMap<i32, i32>
}

impl MyCalendarThree {
    fn new() -> Self {
        Self { items: BTreeMap::new() }
    }
    fn book(&mut self, start: i32, end: i32) -> i32 {
        *self.items.entry(start).or_insert(0) += 1;
        *self.items.entry(end).or_insert(0) += 1;
        //每一次都加上该值，求最大的情况即答案
        self.items
            .values()
            .fold((0, 0), |(res, mut sum), i| {
                sum += *i;
                (res.max(sum), sum)
            }).0
    }
}