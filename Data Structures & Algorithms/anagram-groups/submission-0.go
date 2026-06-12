func groupAnagrams(strs []string) [][]string {
    n := len(strs)
    m := make(map[string][]string)
    for i :=0 ; i<n; i ++ {
        str := sortString(strs[i])
        m[str] = append(m[str], strs[i])
    }
    ans := make([][]string, 0, len(m))
    for _, g := range m {   
        ans = append(ans, g)
    }
    return ans 
}

func sortString (str string) string{
    temp := []byte(str)
      sort.Slice(temp, func(i, j int) bool {
        return temp[i] < temp[j]
    })
    return string(temp)
}

// act - 2
// hat - 1
// stop - 3 