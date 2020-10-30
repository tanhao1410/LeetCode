fn main() {
    println!("Hello, world!");
}

struct Solution{

}

impl Solution {

    //每日一题；463 岛屿的周长
    pub fn island_perimeter(grid: Vec<Vec<i32>>) -> i32 {
        if grid.len() == 0 || grid[0].len() == 0{
            return 0
        }
        //思路：从四个方向分别求出对应的边长，然后相加
        let  (mut left_num,mut right_num,mut up_num,mut down_num) = (0,0,0,0);
        for i in 0..grid.len(){
            let (mut left_flag,mut right_flag) = (true,true);
            for j in 0..grid[0].len(){
                if grid[i][j] == 1 && left_flag{
                    left_num +=1;
                    left_flag = false;
                }
                if grid[i][j] == 0{
                    left_flag = true;
                }
            }
            for j in (0..grid[0].len()).rev(){
                if grid[i][j] == 1 && right_flag{
                    right_num +=1;
                    right_flag = false;
                }
                if grid[i][j] == 0{
                    right_flag = true;
                }
            }
        }
        for i in 0..grid[0].len(){
            let (mut up_flag,mut down_flag) = (true,true);
            for j in 0..grid.len(){
                if grid[j][i] == 1 && up_flag{
                    up_num +=1;
                    up_flag = false;
                }
                if grid[j][i] == 0{
                    up_flag = true;
                }
            }
            for j in 0..grid.len(){
                if grid[j][i] == 1 && down_flag{
                    down_num +=1;
                    down_flag = false;
                }
                if grid[j][i] == 0{
                    down_flag = true;
                }
            }
        }
        left_num+right_num+up_num+down_num
    }
}