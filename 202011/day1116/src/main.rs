use std::cmp::Ordering;

fn main() {
    println!("Hello, world!");
    // println!("{:?}", Solution::reconstruct_queue(
    //     vec![vec![7, 0], vec![4, 4], vec![7, 1], vec![5, 0], vec![6, 1], vec![5, 2]]))
    println!("{}",Solution::largest_rectangle_area(vec![2,1,5,6,2,3]))
}

impl Solution {
    //84. 柱状图中最大的矩形
    pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
        //思路：从第一个数开始，依次求每一个数能围成的最大的面积。
        // 如何求，如果比自己大或相等，求面积往前伸，直到碰到小于自己的，可以求面积了，
        //如果比自己小的话，求面积往前伸，直到碰到比这个还要小的，求面积。
        //止于某个比自己小或比前一个还要小的后，以该数为基准，继续往前伸，求面积，直到走到最后为止。
        //至于前面的数，可以不用管，因为能和前面数组成矩形的，都已经算过了。只需要考虑后面的即可。
        //如果前面的比自己大或相等的话，不用算了，肯定比自己大，可以跳过。
        let mut res = 0;
        let mut i = 0;
        while i < heights.len() {
            let mut cur = heights[i];
            //比前一个数小的话，肯定没有前一个组成的面积大，跳过去。
            if i > 0 && cur < heights[i - 1] {
                i += 1;
                continue;
            }

            let mut s = 0;
            let mut next = i;
            while next < heights.len() {

                let h = heights[next];
                cur = heights[next];
                while next < heights.len() && heights[next] >= cur {
                    next += 1;
                }
                if (next - i) * h as usize > s {
                    s = (next - i) * h as usize;
                }
            }

            if s > res{
                res = s;
            }

            i += 1;

        }

        res as i32
    }

    //每日一题：406. 根据身高重建队列
    pub fn reconstruct_queue(mut people: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        //思路：从前往后一个一个实验是否可以，不行的话，要么是往前面放，要么是往后面放 ，往前面放了的话，会导致之前本来合适的现在变
        //的不合适，需要从它插入的位置开始，从新开始看，
        //排序people count比较大的尽量放前面
        // fn compare(person1 :&Vec<i32>,person2:&Vec<i32>)->Ordering{
        //     if person1[1] > person2[1]{
        //         person1[1].cmp(&person2[1])
        //     }else if person1[1] == person2[1]{
        //         Ordering::Equal
        //     }
        //     Ordering::Greater
        // }
        // people.sort_by(compare);
        if people.len() == 0 {
            return people;
        }
        //优化2：先排好序，120ms ->90ms
        for i in 0..people.len() - 1 {
            for j in i..people.len() {
                if people[j][1] < people[i][1] || (people[j][1] == people[i][1] && people[j][0] < people[i][0]) {
                    let temp = people[j].clone();
                    people[j] = people[i].clone();
                    people[i] = temp;
                }
            }
        }

        let mut cur = 0;
        while cur < people.len() {
            let person = people[cur].clone();
            //看它前面的身高比他高的有多少
            let mut count = 0;
            for i in 0..cur {
                if people[i][0] >= person[0] {
                    //也可以这样优化，遇到和自己相等的就能知道了
                    count += 1;
                }
            }
            //该数的位置正好
            if count == person[1] {
                cur += 1;
            } else if count < person[1] {
                count = person[1] - count;
                //需要将该人往后面放
                let mut index = cur + 1;
                while count > 0 {
                    if people[index][0] >= person[0] {
                        count -= 1;
                    }
                    index += 1;
                }
                //交换位置
                people.remove(cur);
                if index - 1 > people.len() {
                    people.push(person);
                } else {
                    people.insert(index - 1, person);
                }
            } else {
                //需要将该人往前面放，然后从插入的位置从新开始。
                //优化，尽量不要有回溯？怎样能呢？count比较大的尽量放前面
                count = 0;
                let mut index = 0;
                while count < person[1] {
                    if people[index][0] >= person[0] {
                        count += 1;
                    }
                    index += 1;
                }
                //优化1：尽量往后面走 300ms->100ms
                while people[index][0] < person[0] {
                    index += 1;
                }
                people.remove(cur);
                people.insert(index, person);
                cur = index + 1;
            }
        }
        people
    }
}

struct Solution {}