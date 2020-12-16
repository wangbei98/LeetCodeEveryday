// 贪心算法
func partitionLabels(S string) []int {
	// 存储26个字母最后出现的位置
	// postion['x'-'a'] 代表字母x最后出现的位置
	position := make([]int, 26)
	for i := 0; i < 26; i++ {
		position[i] = -1 //字母没出现过，为-1
	}
	for i := 0; i < len(S); i++ {
		position[S[i]-'a'] = i
	}
	res := make([]int, 0) //存储最终结果
	start, end := 0, 0    //区间开始位置和结束位置
	p := 0                //指针
	for start < len(S) {
		// 当start==len(S)时跳出
		for p <= end {
			// 当p>end 的时候，当前区间结束
			// 把当前区间结束位置更新为现有元素最右边的出现位置
			end = max(end, position[S[p]-'a'])
			p++
		}
		res = append(res, end-start+1)
		start = end + 1
		end = start
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}