package mypkg

import "container/list"

// ListSearchFirstElement Search for the first matching element in the list
func ListSearchFirstElement(source *list.List, target string) *list.Element {
	node := source.Front()
	for i := 0; i < source.Len(); i++ {
		if node.Value == target {
			return node
		}
		node = node.Next()
	}
	return nil
}
