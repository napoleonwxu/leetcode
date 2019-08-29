type Data struct {
    Val string
    Time int
}

type TimeMap struct {
    Map map[string][]Data
}


/** Initialize your data structure here. */
func Constructor() TimeMap {
    return TimeMap{Map: make(map[string][]Data)}
}

// O(1)
func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.Map[key] = append(this.Map[key], Data{Val: value, Time: timestamp})
}

// O(logN)
func (this *TimeMap) Get(key string, timestamp int) string {
    arr, ok := this.Map[key]
    if !ok {
        return ""
    }
    idx := binSearch(arr, timestamp)
    if arr[idx].Time <= timestamp {
        return arr[idx].Val
    }
    return ""
}

func binSearch(arr []Data, timestamp int) int {
    left, right := 0, len(arr)-1
    for left < right {
        mid := (left + right) >> 1
        if arr[mid].Time == timestamp {
            return mid
        } else if arr[mid].Time < timestamp {
            if arr[mid+1].Time > timestamp {
                return mid
            }
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return left
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */