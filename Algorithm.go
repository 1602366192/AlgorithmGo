package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func printList(head *Node) {
	for (head != nil) {
		fmt.Print(head.data)
		fmt.Print(" ")
		head = head.next
	}
	fmt.Println();
}

func main() {
	var list1 Node
	list1.data = 2

	var item1 Node
	item1.data = 5
	list1.next = &item1

	var item2 Node
	item2.data = 8
	item1.next = &item2
	item2.next = nil
	printList(&list1)

	var list2 Node
	list2.data = 3

	var item3 Node
	item3.data = 6;
	list2.next = &item3

	var item4 Node
	item4.data = 7
	item3.next = &item4
	item4.next = nil
	printList(&list2)

}
