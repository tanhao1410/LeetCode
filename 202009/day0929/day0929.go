package main

func main() {

}

//数字转罗马数字 1-3999
func intToRoman(num int) string {

	processNum := func(count int, nine, five, four, one string) string {
		res := ""
		for count > 0 {
			if count == 9 {
				res += nine
				count -= 9
			}

			if count >= 5 {
				res += five
				count -= 5
			}

			if count == 4 {
				res += four
				count -= 4
			}

			for count < 4 && count > 0 {
				res += one
				count--
			}
		}
		return res
	}

	//1.先确实是几千
	res := processNum(num/1000, "", "", "", "M")
	num %= 1000
	//2.确定是几百
	res += processNum(num/100, "CM", "D", "CD", "C")
	num %= 100
	//3.确定是几十
	res += processNum(num/10, "XC", "L", "XL", "X")
	num %= 10
	//4.确定是几
	res += processNum(num, "IX", "V", "IV", "I")

	return res
}
