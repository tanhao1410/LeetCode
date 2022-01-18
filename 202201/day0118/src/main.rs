fn main() {
    println!("Hello, world!");
}

//539. 最小时间差
// 采用数组排序的方式计算
pub fn find_min_difference(time_points: Vec<String>) -> i32 {
    const MAX: i32 = 24 * 60;
    if time_points.len() > MAX as usize {
        return 0;
    }
    let mut times = time_points
        .iter()
        .map(|time| time[0..2].parse::<i32>().unwrap() * 60 + time[3..5].parse::<i32>().unwrap())
        .collect::<Vec<i32>>();

    times.sort_unstable();
    let res = times
        .iter()
        .fold((-MAX, MAX), |(pre, res), &time| (time, res.min(time - pre)))
        .1;
    res.min(times[0] + MAX - times.last().unwrap())
}