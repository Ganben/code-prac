package VerifyPreorderSerializationOfBTree

import "strings"

func isValidSerialization(preorder string) bool {
	node_arr := strings.Split(preorder, ",")
	slot := 1
	for i := 0; i < len(node_arr); i++ {
		slot -= 1
		if slot > 0 {
			return false
		}
		if node_arr[i] != "#" {
			slot += 2
		}
	}
	return slot == 0
}
