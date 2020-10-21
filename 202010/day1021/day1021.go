package main

func main() {
	
}

//每日一题：925.长按键入
func isLongPressedName(name string, typed string) bool {
	//双指针法解决
	i,j := 0,0
	for ;i < len(name) && j < len(typed);{

		if name[i] != typed[j]{
			return false
		}
		//与后一个字母不相等
		if i+1 < len(name) {
			if name[i] != name[i+1]{
				for ;j < len(typed) && typed[j] == name[i];j++{
				}
			}else{
				j++
			}
			i++
		}else{
			//是最后一个子目了
			for ;j < len(typed) && typed[j] == name[i];j++{
			}
			return j == len(typed)
		}
	}

	return i==len(typed) && j == len(typed)
}
