fn main() {
    println!("Hello, world!");
}

impl Solution {
    //997. 找到小镇的法官
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        let mut map = vec![(0,0);n as usize];
        for v in trust{
            map[v[0] as usize - 1].0 += 1;
            map[v[1] as usize - 1].1 += 1;
        }
        for i in 0..n as usize{
            if map[i].0 == 0 && map[i].1 == n - 1{
                return i as i32 + 1;
            }
        }
        return -1 ;
    }
}