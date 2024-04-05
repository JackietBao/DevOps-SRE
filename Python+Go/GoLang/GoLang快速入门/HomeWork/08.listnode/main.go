package main

import "fmt"

// ListNode represents a node in a linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two sorted linked lists
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // Dummy node to simplify code
	current := dummy

	// Compare elements from both lists and add the smaller one
	// to the new list until one of the lists becomes empty
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// Add remaining elements from the non-empty list
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}

	return dummy.Next
}

// Utility function to print a linked list
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example lists
	list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
	list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}

	// Merge the two lists
	mergedList := MergeTwoLists(list1, list2)

	// Print the merged list
	fmt.Println("Merged list:")
	PrintList(mergedList)
}
