func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
    tmp := tomatoSlices - 2*cheeseSlices
    if tmp >= 0 && tmp&1 == 0 && cheeseSlices-tmp>>1 >= 0 {
        return []int{tmp>>1, cheeseSlices-tmp>>1}
    }
    return nil
}