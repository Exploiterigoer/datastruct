package main

import (
	"fmt"

	"github.com/Exploiterigoer/datastruct/linearlist/chainlist"
)

func main() {
	// Initializes list
	list := chainlist.InitList()
	fmt.Println(list, "list init")

	// Inserts elements into list
	chainlist.ListInsert(list, 1, 12)
	fmt.Println("ListInsert")

	// Counts length of list
	length := chainlist.ListLength(list)
	fmt.Println(length, "list length")

	chainlist.ListInsert(list, 1, 3)
	length = chainlist.ListLength(list)
	fmt.Println(length, "list length")

	chainlist.ListDelete(list, 2)
	length = chainlist.ListLength(list)
	fmt.Println(length, "list length")

	list = chainlist.CreateListHead(3)
	length = chainlist.ListLength(list)
	fmt.Println(length, "list length")

	list = chainlist.CreateListTail(4)
	length = chainlist.ListLength(list)
	fmt.Println(length, "list length")

	data := chainlist.GetElement(list)
	fmt.Println(data, "list data")

	length = chainlist.ListLength(list)
	fmt.Println(length, "list length before")
	chainlist.ClearList(list)
	length = chainlist.ListLength(list)
	fmt.Println(length, "list length after", list)

	data = chainlist.GetElement(list)
	fmt.Println(data, "list data")
}
