package main

import "fmt"

type ListNode struct {
	data int
	next *ListNode
}

func printList(head *ListNode) {
	for (head != nil) {
		fmt.Print(head.data)
		fmt.Print(" ")
		head = head.next
	}
	fmt.Println();
}

func merge(head1, head2 *ListNode) ListNode {

	var head ListNode
	var cursor *ListNode = &head;

	for (head1 != nil || head2 != nil) {
		if (head1 == nil) {
			cursor.data = head2.data
			if (head2.next != nil) {
				cursor.next = &ListNode{};
				cursor = cursor.next;
			} else {
				cursor = nil;
			}

			head2 = head2.next
			continue
		}


		if (head2 == nil) {
			cursor.data = head1.data;

			if (head1.next != nil) {
				cursor.next = new(ListNode);
				cursor = cursor.next;
			} else {
				cursor = nil;
			}

			head1 = head1.next;
			continue;
		}

		if (head1.data <= head2.data) {
			cursor.data = head1.data;
			head1 = head1.next;
		} else {
			cursor.data = head2.data;
			head2 = head2.next;
		}

		cursor.next = &ListNode{};
		cursor = cursor.next;
	}

	return head;
}

func testMerge() {
	var list1 ListNode
	list1.data = 2

	var item1 ListNode
	item1.data = 5
	list1.next = &item1

	var item2 ListNode
	item2.data = 8
	item1.next = &item2
	item2.next = nil
	printList(&list1)

	var list2 ListNode
	list2.data = 3

	var item3 ListNode
	item3.data = 6;
	list2.next = &item3

	var item4 ListNode
	item4.data = 7
	item3.next = &item4
	item4.next = nil
	printList(&list2)

	var list3 = merge(&list1, &list2)
	printList(&list3)
}

