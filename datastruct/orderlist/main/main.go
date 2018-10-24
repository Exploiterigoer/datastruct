package main

import (
	"fmt"

	"github.com/Exploiterigoer/datastruct/linearlist/orderlist"
)

func main() {
	// Initializes list
	list := orderlist.InitList()
	for i := 0; i < 7; i++ {
		list.Data[i] = i
	}
	fmt.Println(list, "initializing")

	// Gets element from list
	ele := orderlist.GetElem(list, 7)
	fmt.Println(ele, "Getting")

	// Inserts element into list
	flag := orderlist.ListInsert(list, 5, 12)
	fmt.Println(list, flag, "inserting")

	// Deletes element from list
	ele = orderlist.ListDelete(list, 2)
	fmt.Println(ele, list, "Deleting")

	// Gets the given element's index
	i := orderlist.LocateElem(list, 12)
	fmt.Println(i, list, "Locating")

	// Judges list is empty
	isempty := orderlist.ListEmpty(list)
	fmt.Println(isempty, "ListEmpty")

	// CLears the list
	orderlist.ClearList(list)
	isempty = orderlist.ListEmpty(list)
	fmt.Println(isempty, "ListEmpty")
}
