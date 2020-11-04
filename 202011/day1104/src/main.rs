fn main() {
    println!("Hello, world!");
    println!("{:?}", Solution::insert(vec![vec![1,5]], vec![0, 1]));
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
}

struct Solution {}