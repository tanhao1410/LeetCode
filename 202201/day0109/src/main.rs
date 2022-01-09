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
}