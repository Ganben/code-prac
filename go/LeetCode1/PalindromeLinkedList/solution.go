package PalindromeLinkedList

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}

	curr1 := head
	curr2 := head
	// ifOdd := true
	arr := []int{}
	for curr2.Next != nil {
		arr = append(arr, curr1.Val)
		if curr2.Next.Next != nil {
			curr2 = curr2.Next.Next

			curr1 = curr1.Next
			continue
		}

		// ifOdd = false
		break
	}

	// if ifOdd {
	// 	arr = arr[:len(arr)-1]

	// }
	curr1 = curr1.Next
	for i := 0; i < len(arr); i++ {
		if arr[len(arr)-1-i] != curr1.Val {
			return false
		}
		curr1 = curr1.Next
	}
	return true
}
