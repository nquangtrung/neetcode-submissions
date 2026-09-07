func hasDuplicate(nums []int) bool {
    var found map[int]bool = map[int]bool{}

    for _, i := range nums {
        if f := found[i]; f {
            return true
        }

        found[i] = true
    }

    return false
}
