// Similar with: 981. Time Based Key-Value Store
type SnapshotArray struct {
    nums [][][2]int
    snap int
}

func Constructor(length int) SnapshotArray {
    arr := make([][][2]int, length)
    for i := 0; i < length; i++ {
        arr[i] = [][2]int{[2]int{0, 0}}
    }
    return SnapshotArray{nums: arr, snap: 0}
}

// O(1)
func (this *SnapshotArray) Set(index int, val int)  {
    this.nums[index] = append(this.nums[index], [2]int{this.snap, val})
}


func (this *SnapshotArray) Snap() int {
    this.snap++
    return this.snap-1
}

// O(logN)
func (this *SnapshotArray) Get(index int, snap_id int) int {
    idx := sort.Search(len(this.nums[index]), func(i int) bool {
        return this.nums[index][i][0] > snap_id
    })
    return this.nums[index][idx-1][1]
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */