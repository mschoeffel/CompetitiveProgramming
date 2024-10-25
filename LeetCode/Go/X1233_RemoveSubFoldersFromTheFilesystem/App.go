package X1233_RemoveSubFoldersFromTheFilesystem

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func Main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}

	result := removeSubfolders(folder)
	resultStr, _ := json.Marshal(result)
	fmt.Println("Result Remove Sub-Folders from the Filesystem: " + string(resultStr))
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)

	result := []string{folder[0]}

	for i := 1; i < len(folder); i++ {
		lastFolder := result[len(result)-1] + "/"
		if !strings.HasPrefix(folder[i], lastFolder) {
			result = append(result, folder[i])
		}
	}

	return result
}
