// 852. Peak Index in a Mountain Array
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */
 func findInMountainArray(target int, mountainArr *MountainArray) int {
    peak, right := 0, mountainArr.length()-1
    for peak < right {
        mid := (peak+right) >> 1
        if mountainArr.get(mid) < mountainArr.get(mid+1) {
            peak = mid + 1
        } else {
            right = mid
        }
    }
    l, r := 0, peak
    for l <= r {
        mid := (l+r) >> 1
        if mountainArr.get(mid) == target {
            return mid
        } else if mountainArr.get(mid) < target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    l, r = peak, mountainArr.length()-1
    for l <= r {
        mid := (l+r) >> 1
        if mountainArr.get(mid) == target {
            return mid
        } else if mountainArr.get(mid) > target {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return -1
}