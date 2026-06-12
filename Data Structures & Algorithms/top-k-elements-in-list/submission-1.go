func topKFrequent(nums []int, k int) []int {

		n := len(nums)
	hashMap := make(map[int]int)
	for i := 0; i < n; i++ {
		hashMap[nums[i]]++
	}

	ans := make([]int, 0)

	temp := make([]Pair, 0)

	for key, value := range hashMap {
		temp = append(temp, Pair{
			Count: value,
			Num: key,
		})
	}

	sort.Slice(temp, func(i, j int) bool {
		return temp[i].Count > temp[j].Count

	})

	for i := 0; i < k; i++ {
		ans = append(ans, temp[i].Num)
	}

	return ans


}


type Pair struct {
	Count int 
	Num int 
}
