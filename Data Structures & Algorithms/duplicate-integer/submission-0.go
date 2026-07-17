func hasDuplicate(nums []int) bool {
    a := map[int]bool{}
    for _, item := range nums {
        _, ok := a[item] 
        if ok {
            return true
        }
        a[item] = true
    }
    return false
}
