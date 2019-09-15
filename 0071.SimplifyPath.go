func simplifyPath(path string) string {
    dir := strings.Split(path, "/")
    i, j := 0, 0
    for ; j < len(dir); j++ {
        if dir[j] == "" || dir[j] == "." {
            continue
        } else if dir[j] == ".." {
            if i > 0 {
                i--
            }
        } else {
            dir[i] = dir[j]
            i++
        }
    }
    return "/" + strings.Join(dir[:i], "/")
}