fn main() {
    println!("Hello, world!");

    assert_eq!(max_product(vec![0,2]),2);
    assert_eq!(max_product(vec![2,-5,-2,-4,3]),24);
    assert_eq!(max_product(vec![-2,3,-4]),24);
    assert_eq!(max_product(vec! [-2,0,-1]),0);
    assert_eq!(max_product(vec! [2,3,-2,4]),6);
}

//283. 移动零
pub fn move_zeroes(nums: &mut Vec<i32>) {
    //双指针，i指向第一个为0的地方，j指向i + 1，j不停的往前走，碰到不为0的就将它放到i处，同时，i 放到j位置
    let mut i = 0;
    while i < nums.len() && nums[i] != 0{
        i += 1;
    }
    let mut j = i + 1;
    while j < nums.len(){
        if nums[j] != 0 {
            nums[i] = nums[j];
            nums[j] = 0;
            i += 1;
        }
        j += 1;
    }
}

//152. 乘积最大子数组
pub fn max_product(nums: Vec<i32>) ->  i32 {
    let mut dp_max = vec![0; nums.len()]; //以nums[i]结尾的最大子数组之积
    let mut dp_min = vec![0; nums.len()]; //以nums[i]结尾的负的最小值。

    dp_max[0] = nums[0];
    if nums[0] < 0{
        dp_min[0] = nums[0];
    }

    for i in 1..nums.len() {
        if nums[i] >= 0 {
            //看它的最大值是否大于0，如果大于0，则dp_max[i] = nums[i] * dp[i - 1]
            //如果它的前一个最大值小于0，自己又大于0，所以，dp_max[i] = nums[i]
            //如果它的前一个等于0，自己又大于0，dp_max[i] = nums[i]
            //更新dp_min
            //dp_min[i] = nums[i] * dp_min[i - 1]，因为自己大于0，而前面的最大负，乘上正数，肯定更小
            if dp_max[i - 1] > 0 {
                dp_max[i] = nums[i] * dp_max[i - 1];
            } else {
                dp_max[i] = nums[i];
            }
            dp_min[i] = nums[i] * dp_min[i - 1];
        } else {
            //如果当前值小于0
            //如果dp_max[i - 1] > 0,dp_min[i - 1] < 0,=>dp_max[i] = nums[i] * dp_min[i - 1]
            //如果dp_max[i - 1] > 0,_ =>dp_max[i] = nums[i],
            //
            if dp_max[i-1] > 0{
                if dp_min[i - 1] < 0{
                    //更新最大值
                    dp_max[i] = nums[i] * dp_min[i - 1];
                }else{
                    dp_max[i] = nums[i];
                }
                //更新负的最小值
                dp_min[i] = nums[i] * dp_max[i - 1];
            }else{
                dp_max[i] = dp_max[i - 1].min(dp_min[i - 1]) * nums[i];
                dp_min[i] = nums[i];
            }
        }
    }
    dp_max.into_iter().max().unwrap()
}