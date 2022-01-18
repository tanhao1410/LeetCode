use std::collections::HashSet;

fn main() {
    println!("Hello, world!");
    println!("{}", can_measure_water(1, 2, 3));
}

//365. 水壶问题
pub fn can_measure_water(jug1_capacity: i32, jug2_capacity: i32, target_capacity: i32) -> bool {
    let mut res_set = std::collections::HashSet::new();
    let mut state_set = std::collections::HashSet::new();
    operate(jug1_capacity,jug2_capacity,0,0,&mut res_set,&mut state_set);
    println!("{:?}", state_set);
    res_set.contains(&target_capacity)
}

fn operate(jug1_capacity: i32, jug2_capacity: i32, jug1: i32, jug2: i32, res_set: &mut HashSet<i32>, state_set: &mut HashSet<(i32, i32)>) {
    //开始进行操作，有三种操作，加满一个，清空任意一个，倒水操作
    res_set.insert(jug1);
    res_set.insert(jug2);
    res_set.insert(jug2 + jug1);
    //给1倒满水
    if jug1 < jug1_capacity && !state_set.contains(&(jug1_capacity, jug2)) {
        state_set.insert((jug1_capacity, jug2));
        operate(jug1_capacity, jug2_capacity, jug1_capacity, jug2, res_set, state_set);
    }
    if jug2 < jug2_capacity && !state_set.contains(&(jug1, jug2_capacity)) {
        //给2倒满水
        state_set.insert((jug1, jug2_capacity));
        operate(jug1_capacity, jug2_capacity, jug1, jug2_capacity, res_set, state_set);
    }
    if !state_set.contains(&(0, jug2)) {
        //清空1
        state_set.insert((0, jug2));
        operate(jug1_capacity, jug2_capacity, 0, jug2, res_set, state_set);
    }
    if !state_set.contains(&(jug1, 0)) {
        //清空2
        state_set.insert((jug1, 0));
        operate(jug1_capacity, jug2_capacity, jug1, 0, res_set, state_set);
    }
    if jug1 > 0 && jug2 < jug2_capacity &&
        jug1 + jug2 <= jug2_capacity && !state_set.contains(&(0, jug2 + jug1)) {
        //从1 => 2 倒满或倒空。条件，1 不能空，2 不能满，倒完情况
        state_set.insert((0, jug2 + jug1));
        operate(jug1_capacity, jug2_capacity, 0, jug1 + jug2, res_set, state_set);
    }
    if jug1 > 0 && jug2 < jug2_capacity &&
        jug1 + jug2 > jug2_capacity && !state_set.contains(&(jug1 - jug2_capacity + jug2, jug2_capacity)) {
        //从1 => 2 倒满或倒空。条件，1 不能空，2 不能满，倒满情况
        state_set.insert((jug1 - jug2_capacity + jug2, jug2_capacity));
        operate(jug1_capacity, jug2_capacity, jug1 - jug2_capacity + jug2, jug2_capacity, res_set, state_set);
    }
    if jug2 > 0 && jug1 < jug1_capacity &&
        jug1 + jug2 <= jug1_capacity && !state_set.contains(&(jug2 + jug1, 0)) {
        //从2 => 1 倒满或倒空。倒完情况
        state_set.insert((jug2 + jug1, 0));
        operate(jug1_capacity, jug2_capacity, jug2 + jug1, 0, res_set, state_set);
    }
    if jug2 > 0 && jug1 < jug1_capacity &&
        jug1 + jug2 > jug1_capacity && !state_set.contains(&(jug1_capacity, jug1 + jug2 - jug1_capacity)) {
        //从2 => 1 倒满或倒空。倒满情况
        state_set.insert((jug1_capacity, jug2 + jug1 - jug1_capacity));
        operate(jug1_capacity, jug2_capacity, jug1_capacity, jug1 - jug2_capacity + jug2, res_set, state_set);
    }
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