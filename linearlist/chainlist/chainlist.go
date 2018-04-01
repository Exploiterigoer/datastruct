package chainlist

// ChainList The ChainList struct
type ChainList struct {
	data int
	next *ChainList
}

// InitList Initializes ChainList
func InitList() *ChainList {
	list := new(ChainList)
	list.next = nil
	return list
}

// ListInsert Inserts an element to chain list
func ListInsert(list *ChainList, i, e int) bool {
	j := 1

	// finds the given index(i)
	for list != nil && j < i {
		list = list.next
		j++
	}

	// if end of list or out of the given index,
	// meaning no element at the given index.
	if list == nil || j > i {
		return false
	}

	// here,the list beeing seperated front part and after part by the given index,
	// the pointer area of new node pointing to pointer area of after part
	// the pointer area of front part pointing to the new node

	// it's important that the pointing order of pointer area,which can't be exchanged.
	node := new(ChainList) // new node
	node.data = e          // assigns the new node data
	node.next = list.next  // new pointer area pointing to old pointer area
	list.next = node       // old pointer area pointing to new node

	return true
}

// ListLength Gets the length of list
func ListLength(list *ChainList) int {
	var length int
	for list != nil {
		length++
		list = list.next
	}
	return length
}

// ListDelete Deletes the element form list
func ListDelete(list *ChainList, index int) *ChainList {
	j := 1
	tmpList := list // keeps the original list

	for list != nil && j < index { // finds the element at the given index
		list = list.next
		j++
	}

	if list == nil || j > index {
		return nil
	}

	tmpList.next = list.next //ignores the element at the given index

	return tmpList
}

// CreateListTail Creates list from tail
func CreateListTail(n int) *ChainList {
	list := new(ChainList)
	tail := list // the tail node(assumed)
	for i := 1; i <= n; i++ {
		l := new(ChainList)
		l.data = i
		tail.next = l // points to new node
		tail = l      // new node as tail
	}
	tail.next = nil // resets the next of tail pointing to nil
	return list
}

// CreateListHead Creates list from head
func CreateListHead(n int) *ChainList {
	list := new(ChainList) // the head node
	for i := 1; i <= n; i++ {
		tmpList := new(ChainList)
		tmpList.data = i
		tmpList.next = list.next
		list.next = tmpList
	}
	return list
}

// GetElement Gets all datas from list
func GetElement(list *ChainList) []int {
	data := make([]int, 0)
	for list != nil {
		data = append(data, list.data)
		list = list.next
	}
	return data
}

// ClearList Clears the list
func ClearList(list *ChainList) {
	var n *ChainList
	for list != nil {
		n = list        // temp variable for get the next element
		list.next = nil // destorys the reference
		list = n.next   // gets the next element
	}
}
