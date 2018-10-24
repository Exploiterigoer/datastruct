package main

import (
	"data_struct_algorithm/algorithm/search"
	"data_struct_algorithm/algorithm/sort"
	"data_struct_algorithm/assistant"
	"data_struct_algorithm/datastruct/binarytree"
	"data_struct_algorithm/datastruct/stackqueue"
	"fmt"
)

var sgss *sortalgorithm.SelectionSort
var sgbb *sortalgorithm.Bubble
var sgis *sortalgorithm.InsertionSort
var sgms *sortalgorithm.MergerSort
var sgqs *sortalgorithm.QuickSort
var bssh *search.BinarySearch
var arrs []int
var stack *stackqueue.Stack
var queue *stackqueue.Queue
var btree *binarytree.BinTree
var maxHeap sortalgorithm.Heap

func init() {
	sgss = new(sortalgorithm.SelectionSort)
	sgbb = new(sortalgorithm.Bubble)
	sgis = new(sortalgorithm.InsertionSort)
	sgms = new(sortalgorithm.MergerSort)
	sgqs = new(sortalgorithm.QuickSort)
	bssh = new(search.BinarySearch)
	maxHeap = new(sortalgorithm.MaxHeap)
	arrs = assistant.ArrayGenerator(10, 100)
	stack = stackqueue.CreateStack(5)
	queue = stackqueue.CreateQueue(5)
	btree = binarytree.InitBinTree()
}

func main() {
	//fmt.Printf("本次共有%d个数组元素参加排序\n----------\n", arrs)
	//var cost int64
	/*
		// 注意：测试的数组元素个数不能达到十亿数量级别，否则机器内存被耗光
		o4sgss := arrs
		_, cost = assistant.CalcTime(sgss.SelectSort, &o4sgss)
		fmt.Printf("选择排序耗时%9.9f秒,%v\n----------\n", float64(cost)/1e9, o4sgss)

		// 二分查找需排好序
		bsarr := arrs
		fmt.Printf("本次二分查找的数组:%v\n----------\n", bsarr)
		result, cost := assistant.CalcTime(bssh.BinSearch2, &bsarr, 24)
		fmt.Printf("二分查找耗时%9.9f秒,所在下标:%d\n----------\n", float64(cost)/1e9, result)

		////////////////////////排序
		o4sgbb := arrs
		_, cost = assistant.CalcTime(sgbb.BubbleSort, &o4sgbb)
		fmt.Printf("冒泡排序耗时%9.9f秒,%v\n----------\n", float64(cost)/1e9, o4sgbb)

		o4sgis := arrs
		_, cost = assistant.CalcTime(sgis.InsertSort, &o4sgis)
		fmt.Printf("插入排序耗时%9.9f秒,%v\n----------\n", float64(cost)/1e9, o4sgis)
	*/

	//o4sgms := arrs
	//_, cost = assistant.CalcTime(sgms.MergerSortDown2Up, &o4sgms)
	//fmt.Printf("归并排序耗时:%9.9f秒\n----------\n", float64(cost)/1e9)
	//fmt.Println(o4sgms)

	//o4sgqs := arrs
	//_, cost = assistant.CalcTime(sgqs.QuicklySort, &o4sgqs)
	//fmt.Printf("随机快速排序耗时:%9.9f秒\n----------\n", float64(cost)/1e9)
	//_, cost = assistant.CalcTime(sgqs.QuicklySort2, &o4sgqs)
	//fmt.Printf("双路快速排序耗时:%9.9f秒\n----------\n", float64(cost)/1e9)
	//_, cost = assistant.CalcTime(sgqs.QuicklySort3, &o4sgqs)
	//fmt.Printf("三路快速排序耗时:%9.9f秒\n----------\n", float64(cost)/1e9)
	//fmt.Println(o4sgqs)
	/*
		// 栈操作
		for i := 1; i <= 5; i++ {
			stack.Push(i)
		}
		fmt.Printf("栈顶元素:%v\n----------\n", stack.Peek())
		fmt.Printf("弹出元素:%v\n----------\n", stack.Pop())
		fmt.Printf("栈元素个数:%d\n----------\n", stack.Size())
		for i := 1; i <= 4; i++ {
			fmt.Printf("弹出元素:%v\n----------\n", stack.Pop())
		}
		fmt.Printf("栈为空:%t\n----------\n", stack.IsEmpty())
		for i := 6; i <= 11; i++ {
			stack.Push(i)
		}
		fmt.Printf("栈元素个数:%d\n----------\n", stack.Size())
		fmt.Printf("栈顶元素:%v\n----------\n", stack.Peek())

		// 队列操作
		for i := 'A'; i-65 < 5; i++ {
			queue.Enqueue(i)
		}
		for !queue.IsEmpty() {
			fmt.Printf("出队元素:%c\n----------\n", queue.Dequeue())
		}
		fmt.Printf("队列是否空:%t\n----------\n", queue.IsEmpty())

		// 二叉树操作
		for i := 'A'; i-65 < 12; i++ {
			btree.CreateBinTree(byte(i))
		}
		var container []interface{}
		btree.LevelTraverse(&container)
		fmt.Printf("层序遍历二叉树(序列化结果):%s\n----------\n", container)
		container = nil
		btree.PreOrderTraverse(btree.Root, &container)
		fmt.Printf("先序遍历二叉树(序列化结果):%s\n----------\n", container)
		container = nil
		btree.InOrderTraverse(btree.Root, &container)
		fmt.Printf("中序遍历二叉树(序列化结果):%s\n----------\n", container)
		container = nil
		btree.LastOrderTraverse(btree.Root, &container)
		fmt.Printf("后序遍历二叉树(序列化结果):%s\n----------\n", container)
		fmt.Printf("计算二叉树节点数量:%d\n----------\n", btree.GetNodeCount())
		fmt.Printf("计算二叉树叶子节点数量:%d\n----------\n", btree.GetLeafNodeCount())
		fmt.Printf("通过公式计算二叉树的深度:%d\n----------\n", btree.GetDepth())
		fmt.Printf("递归方式计算二叉树的深度:%d\n----------\n", btree.GetDepthByRecursive(btree.Root))
		fmt.Printf("循环方式计算二叉树的深度:%d\n----------\n", btree.GetDepthByLoop(btree.Root))
	*/

	// 堆排序
	o4heap := arrs
	//maxHeap.Create(o4heap)
	//fmt.Println(maxHeap.Element(-1))
	//fmt.Println(maxHeap.Element(2)) // 父
	//fmt.Println(maxHeap.Element(6)) // 子
	//fmt.Println(maxHeap.Extract())                           // 将一个从小到大的数组heapify后
	//start := time.Now().UnixNano()
	//fmt.Println(maxHeap.HeapSort())
	//end := time.Now().UnixNano()
	//fmt.Printf("堆排序耗时%9.9f秒", float64(end-start)/1e9)
	//fmt.Println(maxHeap.Element(-1))
	//fmt.Println(maxHeap.Element(2)) // 父
	//fmt.Println(maxHeap.Element(6)) // 子
	//fmt.Println(maxHeap.(*sortalgorithm.MaxHeap).HeapSort())
	sortalgorithm.HeapSort2(&o4heap)
	fmt.Println(&o4heap)
}
