package gSlice

// SliceClear 清空切片
func SliceClear(s *[]any) {
	*s = (*s)[0:0]
}

// SliceDelete 删除元素
func SliceDelete(slice *[]string, index int) error {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
	return nil
}
