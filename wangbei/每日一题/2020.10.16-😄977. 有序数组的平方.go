func sortedSquares(A []int) []int {
    i,j := 0,len(A)-1
    cur := len(A)-1
    res := make([]int,len(A))
    for cur >= 0{
        if abs(A[i]) >= abs(A[j]){
            // 左边绝对值更大或者相等
            res[cur] = A[i]*A[i]
            i++
        }else{
            // 右边更大
            res[cur] = A[j]*A[j]
            j--
        }
        cur--
    }
    return res
}
func abs(x int)int{
    if x > 0{
        return x
    }else{
        return -x
    }
}