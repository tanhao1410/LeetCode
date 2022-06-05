use rand::random;

fn main() {
    println!("Hello, world!");
}


impl Solution {
    //478. 在圆内随机生成点
    fn new(radius: f64, x_center: f64, y_center: f64) -> Self {
        Self { radius, x_center, y_center }
    }

    fn rand_point(&self) -> Vec<f64> {
        let in_circle = |point: (f64, f64)| {
            self.radius.powi(2) >= (point.0 - self.x_center).powi(2) + (point.1 - self.y_center).powi(2)
        };
        let (x, y) = (self.x_center - self.radius, self.y_center - self.radius);
        let rand_point = (random::<f64>() * self.radius * 2.0 + x, random::<f64>() * self.radius * 2.0 + y);
        if in_circle(rand_point) {
            return vec![rand_point.0, rand_point.1];
        }
        self.rand_point()
    }
}

struct Solution {
    radius: f64,
    x_center: f64,
    y_center: f64,
}