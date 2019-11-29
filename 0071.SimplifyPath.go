func simplifyPath(path string) string {
    dir := strings.Split(path, "/")
    i := 0
    for j := 0; j < len(dir); j++ {
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