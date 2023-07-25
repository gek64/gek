package gString

import (
    "path/filepath"
    "strings"
)

// StringTrimSuffix remove suffix
func StringTrimSuffix(file string) string {
    fileExt := filepath.Ext(file)
    if fileExt != "" {
        file = strings.TrimSuffix(file, fileExt)
    }
    return file
}

// StringFindInSlice find string's position in a slice
func StringFindInSlice(slice []string, str string) int {
    for position, i := range slice {
        if str == i {
            return position
        }
    }
    return -1
}
