func productExceptSelf(nums []int) []int {

	n := len(nums)
	ans:= make([]int,n)
	for i :=0; i<n; i++ {
		mul := 1 
		for j :=0; j<n; j ++ {
			if j != i {
				mul *= nums[j]
			}
		}
		ans[i] = mul
	}
	return ans 
}


