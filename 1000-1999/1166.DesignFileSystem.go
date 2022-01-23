/*
You are asked to design a file system which provides two functions:
create(path, value): Creates a new path and associates a value to it if possible and returns True. Returns False if the path already exists or its parent path doesn't exist.
get(path): Returns the value associated with a path or returns -1 if the path doesn't exist.
The format of a path is one or more concatenated strings of the form: / followed by one or more lowercase English letters. For example, /leetcode and /leetcode/problems are valid paths while an empty string and / are not.
Implement the two functions.

Please refer to the examples for clarifications.

Example 1:
Input: 
["FileSystem","create","get"]
[[],["/a",1],["/a"]]
Output: 
[null,true,1]
Explanation: 
FileSystem fileSystem = new FileSystem();
fileSystem.create("/a", 1); // return true
fileSystem.get("/a"); // return 1

Example 2:
Input: 
["FileSystem","create","create","get","create","get"]
[[],["/leet",1],["/leet/code",2],["/leet/code"],["/c/d",1],["/c"]]
Output: 
[null,true,true,2,false,-1]
Explanation: 
FileSystem fileSystem = new FileSystem();
fileSystem.create("/leet", 1); // return true
fileSystem.create("/leet/code", 2); // return true
fileSystem.get("/leet/code"); // return 2
fileSystem.create("/c/d", 1); // return false because the parent path "/c" doesn't exist.
fileSystem.get("/c"); // return -1 because this path doesn't exist.
 
Constraints:
The number of calls to the two functions is less than or equal to 10^4 in total.
2 <= path.length <= 100
1 <= value <= 10^9
*/

type FileSystem struct {
    Map map[string]int
}


func Constructor() FileSystem {
    m := make(map[string]int)
    m[""] = 0
    return FileSystem{Map: m}
}


func (this *FileSystem) Create(path string, value int) bool {
    if path == "" || path =="/" {
        return false
    }
    i := len(path)-1
    for ; i >= 0 && path[i] != '/'; i-- {}
    if i < 0 || this.Get(path[:i]) == -1 {
        return false
    }
    this.Map[path] = value
    return true
}


func (this *FileSystem) Get(path string) int {
    val, ok := this.Map[path]
    if ok {
        return val
    }
    return -1
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Create(path,value);
 * param_2 := obj.Get(path);
 */
