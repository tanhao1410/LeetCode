fn main() {
    println!("Hello, world!");
}

use std::collections::BTreeSet;

//729. 我的日程安排表 I
struct MyCalendar {
    times: BTreeSet<(i32, i32)>
}


impl MyCalendar {
    fn new() -> Self {
        Self {
            times: BTreeSet::new()
        }
    }

    fn book(&mut self, start: i32, end: i32) -> bool {
        let l = self.times.range((start, 0)..).next().map_or(true, |&(_, e)| end <= e);
        let r = self.times.range(..(start, 0)).rev().next().map_or(true, |&(_, e)| start >= e);
        l && r && self.times.insert((start, end))
    }
}
