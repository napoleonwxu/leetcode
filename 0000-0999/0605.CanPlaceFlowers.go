func canPlaceFlowers(flowerbed []int, n int) bool {
    cnt, Len := 0, len(flowerbed)
    for i := 0; i < Len && cnt < n; i++ {
        if flowerbed[i]==0 && (i==0 || flowerbed[i-1]==0) && (i==Len-1 || flowerbed[i+1] == 0) {
            flowerbed[i] = 1
            cnt++
        }
    }
    return cnt == n
}