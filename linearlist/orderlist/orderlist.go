// Package orderlist a package of linear table, which storaging element in order.
package orderlist

const maxSize = 10

// OrderList the struct of OrderList
type OrderList struct {
	Data []int
}

// InitList Initializes OrderList
func InitList() *OrderList {
	list := new(OrderList)
	list.Data = make([]int, maxSize)
	return list
}

// ListEmpty Checks OrderList empty or not
func ListEmpty(list *OrderList) bool {
	return len(list.Data) == 0
}

// ClearList Clears the OrderList
func ClearList(list *OrderList) {
	list.Data = make([]int, 0)
}

// GetElem Gets the element form OrderList by the given index
func GetElem(list *OrderList, index int) int {
	if index < 0 || index > len(list.Data) {
		return -1
	}
	return list.Data[index]
}

// LocateElem Gets the order number of the given element,which is contained in OrderList,
// order number = index + 1, if not found,it will return -1.
func LocateElem(list *OrderList, ele int) int {
	for k, v := range list.Data {
		if v == ele {
			return k + 1
		}
	}
	return -1
}

// ListInsert Inserts the given element into the OrderList at the given index,
// it will return true when inserted successfully,or return false.
func ListInsert(list *OrderList, index int, ele int) bool {
	dataLen := len(list.Data)
	if index < 0 || index > dataLen {
		return false
	}

	if index < dataLen {
		for i := dataLen - 1; i > index-1; i-- {
			list.Data[i] = list.Data[i-1]
		}
	}

	list.Data[index-1] = ele

	return true
}

// ListDelete Deletes the given element from OrderList, it returns true
// when deleted successfully,or is false.
func ListDelete(list *OrderList, index int) int {
	dataLen := len(list.Data)
	if index < 0 || index > dataLen {
		return -1
	}

	ele := list.Data[index-1]
	if index < dataLen {
		for i := index; i < dataLen; i++ {
			list.Data[i-1] = list.Data[i]
		}
	}
	list.Data[dataLen-1] = 0

	return ele
}
