use std::cmp::Ordering;

fn main() {
    println!("Hello, world!");
}

struct Solution {}

impl Solution {

    pub fn slowest_key(release_times: Vec<i32>, keys_pressed: String) -> char {
        keys_pressed
            .chars()
            .enumerate()
            .map(|(i, c)| (release_times[i] - release_times.get(i - 1).unwrap_or(&0), c))
            .fold((0, 'a'), |r, e| if e.0 > r.0 || (e.0 == r.0 && e.1 > r.1) { e } else { r })
            .1
    }

    //参考python简洁解法思路：max(zip(map(int.__sub__,releaseTimes,[0]+releaseTimes),keysPressed))[1]
    pub fn slowest_key2(release_times: Vec<i32>, keys_pressed: String) -> char {
        vec![0]
            .iter()
            .chain(release_times.iter())
            .zip(release_times.iter())
            .map(|(p,q)| q - p)
            .zip(keys_pressed.chars())
            .max_by(|h,t|match (h.0.cmp(&t.0),h.1.cmp(&t.1)){
                (Ordering::Greater,_)|(Ordering::Equal,Ordering::Greater) => Ordering::Greater,
                _=>Ordering::Less
            })
            .unwrap()
            .1
    }

}