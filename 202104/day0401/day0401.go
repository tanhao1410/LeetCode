package main

func main() {

}

//1472. 设计浏览器历史记录
type BrowserHistory struct {
	Pages []string
	Cur   int
	Last  int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		Pages: []string{homepage},
	}
}

func (this *BrowserHistory) Visit(url string) {
	if this.Last == this.Cur {

		//原有的满了
		if this.Last == len(this.Pages)-1 {
			this.Pages = append(this.Pages, url)
			this.Cur++
			this.Last++
		} else {
			this.Cur++
			this.Last++
			this.Pages[this.Cur] = url
		}
	} else {
		this.Cur++
		this.Pages[this.Cur] = url
		this.Last = this.Cur
	}
}

func (this *BrowserHistory) Back(steps int) string {
	if this.Cur-steps <= 0 {
		//回退到
		this.Cur = 0
		return this.Pages[0]
	}
	this.Cur -= steps
	return this.Pages[this.Cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.Cur += steps
	if this.Cur > this.Last {
		this.Cur = this.Last
	}
	return this.Pages[this.Cur]
}

//每日一题：1006. 笨阶乘
func clumsy(N int) int {
	//思路：按照*/+-顺序来
	//共有 几组 N - 1 / 4 组，余 N - 1 % 4项
	//先算*/,以
	if N == 2 {
		return 2
	} else if N == 1 {
		return 1
	} else if N == 3 {
		return 6
	}
	res := N*(N-1)/(N-2) + N - 3
	//用for循环来做
	for N = N - 4; N >= 3; N -= 4 {
		res -= N * (N - 1) / (N - 2)
		//这个是加上
		res += N - 3
	}
	if N > 0 {
		res -= N
	}
	return res
}
