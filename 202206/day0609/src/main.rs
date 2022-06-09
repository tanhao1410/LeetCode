use rand::Rng;

fn main() {
    println!("Hello, world!");
}

//497. 非重叠矩形中的随机点
struct Solution {
    areas: i64,
    sum_areas: Vec<i64>,
    rects: Vec<Vec<i32>>,
}

impl Solution {
    fn new(rects: Vec<Vec<i32>>) -> Self {
        let mut sum_areas = vec![0i64; rects.len()];
        for i in 0..sum_areas.len() {
            let cur_area = (rects[i][3] - rects[i][1] + 1) * (rects[i][2] - rects[i][0] + 1);
            sum_areas[i] = cur_area as i64;
            if i > 0 {
                sum_areas[i] += sum_areas[i - 1];
            }
        }

        let areas = *sum_areas.last().unwrap();
        Self { sum_areas, rects, areas }
    }


    fn pick(&self) -> Vec<i32> {
        let mut rng = rand::thread_rng();
        let rand: i64 = rng.gen_range(0..self.areas);
        let mut i = 0;
        loop {
            if rand < self.sum_areas[i] {
                let rect = &self.rects[i];
                let y: i32 = rng.gen_range(rect[1]..=rect[3]);
                let x: i32 = rng.gen_range(rect[0]..=rect[2]);
                return vec![x, y];
            }
            i += 1;
        }
    }
}