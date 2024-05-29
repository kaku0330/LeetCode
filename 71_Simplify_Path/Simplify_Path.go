package simplifypath

import (
	"strings"
)

func SimplifyPath(path string) string {
	//第一版  BEAT 100%
	splitpath := strings.Split(path, "/")
	i := 0
	for i < len(splitpath) {
		if splitpath[i] == "" || splitpath[i] == "." {
			splitpath = append(splitpath[:i], splitpath[i+1:]...)
		} else if splitpath[i] == ".." {
			if i >= 1 {
				splitpath = append(splitpath[:i-1], splitpath[i+1:]...)
				i--
			} else {
				splitpath = append(splitpath[:i], splitpath[i+1:]...)
			}

		} else {
			i++
		}
	}
	result := strings.Join(splitpath, "/")
	if result == "" {
		return "/"
	}
	return "/" + result
}
