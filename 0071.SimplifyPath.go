import "strings"

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	i := 0
	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		} else if dir == ".." {
			if i > 0 {
				i--
			}
		} else {
			dirs[i] = dir
			i++
		}
	}
	return "/" + strings.Join(dirs[:i], "/")
}