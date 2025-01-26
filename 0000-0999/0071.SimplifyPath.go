func simplifyPath(path string) string {
    dirs := strings.Split(path, "/")
    i := 0
    for _, dir := range dirs {
        if dir == ".." {
            if i > 0 {
                i--
            }
        } else if dir != "" && dir != "." {
            dirs[i] = dir
            i++
        }
    }
    return "/" + strings.Join(dirs[:i], "/")
}