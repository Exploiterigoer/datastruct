package binarytree

import (
	"data_struct_algorithm/datastruct/stackqueue"
	"math"
)

// Node 节点结构体
type Node struct {
	data   interface{}
	lchild *Node
	rchild *Node
}

// BinTree 二叉树结构体
type BinTree struct {
	Root *Node
}

// InitBinTree 初始化二叉树
func InitBinTree() *BinTree {
	return new(BinTree)
}

// CreateBinTree 创建二叉树
func (bt *BinTree) CreateBinTree(data interface{}) {
	// 二叉树还没初始化
	if bt == nil {
		return
	}

	n := new(Node)
	n.data = data

	// bt.Root == nil 表明是一棵空二叉树，直接把新建的节点作为根节点并返回
	if bt.Root == nil {
		bt.Root = n
		return
	}

	// 创建一个容量为10的队列
	queue := stackqueue.CreateQueue(10)

	// 将二叉树的根节点加入队列
	queue.Enqueue(bt.Root)

	// 从队列中取出每一个节点
	var currentNode *Node
	for !queue.IsEmpty() {
		// 当前取出的节点
		currentNode = queue.Dequeue().(*Node)

		// 判断左子树尝试挂载新节点
		if currentNode.lchild == nil {
			currentNode.lchild = n
			break
		} else {
			queue.Enqueue(currentNode.lchild)
		}

		// 判断右子树尝试挂载新节点
		if currentNode.rchild == nil {
			currentNode.rchild = n
			break
		} else {
			queue.Enqueue(currentNode.rchild)
		}
	}
}

// PreOrderTraverse 先序遍历二叉树
func (bt *BinTree) PreOrderTraverse(node *Node, container *[]interface{}) {
	if node == nil {
		*container = append(*container, "#")
		return
	}
	*container = append(*container, string(node.data.(byte)))
	bt.PreOrderTraverse(node.lchild, container)
	bt.PreOrderTraverse(node.rchild, container)
}

// InOrderTraverse 中序遍历二叉树
func (bt *BinTree) InOrderTraverse(node *Node, container *[]interface{}) {
	if node == nil {
		*container = append(*container, "#")
		return
	}
	bt.InOrderTraverse(node.lchild, container)
	*container = append(*container, string(node.data.(byte)))
	bt.InOrderTraverse(node.rchild, container)
}

// LastOrderTraverse 后序遍历二叉树
func (bt *BinTree) LastOrderTraverse(node *Node, container *[]interface{}) {
	if node == nil {
		*container = append(*container, "#")
		return
	}
	bt.LastOrderTraverse(node.lchild, container)
	bt.LastOrderTraverse(node.rchild, container)
	*container = append(*container, string(node.data.(byte)))
}

// LevelTraverse 层序遍历二叉树，是序列化操作后的结果
func (bt *BinTree) LevelTraverse(container *[]interface{}) {
	queue := stackqueue.CreateQueue(10)
	queue.Enqueue(bt.Root)
	for !queue.IsEmpty() {
		currentNode := queue.Dequeue().(*Node)
		*container = append(*container, string(currentNode.data.(byte)))

		if currentNode.lchild != nil {
			queue.Enqueue(currentNode.lchild)
		} else {
			*container = append(*container, "#")
		}

		if currentNode.rchild != nil {
			queue.Enqueue(currentNode.rchild)
		} else {
			*container = append(*container, "#")
		}
	}
}

// GetNodeCount 获取二叉树的节点数，这里按照层序方式来计算节点数
func (bt *BinTree) GetNodeCount() int {
	if bt == nil || bt.Root == nil {
		return 0
	}
	var nodeCount int
	var currentNode *Node
	queue := stackqueue.CreateQueue(10)
	queue.Enqueue(bt.Root)

	for !queue.IsEmpty() {
		currentNode = queue.Dequeue().(*Node)

		// 算有左孩子的节点个数，并加入队列
		if currentNode.lchild != nil {
			nodeCount++
			queue.Enqueue(currentNode.lchild)
		}

		// 右孩子的节点个数不算，再算就重复了，但是加入队列
		if currentNode.rchild != nil {
			queue.Enqueue(currentNode.rchild)
		}

		// 算无左右孩子的节点个数，无需加入队列
		if currentNode.lchild == nil && currentNode.rchild == nil {
			nodeCount++
		}
	}
	return nodeCount
}

// GetLeafNodeCount 获取二叉树的叶子节点数，这里按照层序方式来计算叶子节点数
func (bt *BinTree) GetLeafNodeCount() int {
	if bt == nil || bt.Root == nil {
		return 0
	}

	var nodeCount int
	var currentNode *Node
	queue := stackqueue.CreateQueue(10)
	queue.Enqueue(bt.Root)

	for !queue.IsEmpty() {
		currentNode = queue.Dequeue().(*Node)
		// 所有非空节点都加入队列
		if currentNode.lchild != nil {
			queue.Enqueue(currentNode.lchild)
		}

		if currentNode.rchild != nil {
			queue.Enqueue(currentNode.rchild)
		}

		// 只计算同时没有左右孩子的节点个数
		if currentNode.rchild == nil && currentNode.lchild == nil {
			nodeCount++
		}
	}
	return nodeCount
}

// GetDepth 通过公式获取二叉树的深度
func (bt *BinTree) GetDepth() int {
	// 先取得总的节点数
	nc := bt.GetNodeCount()
	return int(math.Floor(math.Log2(float64(nc)))) + 1
}

// GetDepthByRecursive 通过递归调用获取二叉树的深度
func (bt *BinTree) GetDepthByRecursive(node *Node) int {
	if node == nil {
		return 0
	}
	// 这里的1表示当前层所在的层数
	ldepth := 1 + bt.GetDepthByRecursive(node.lchild)
	rdepth := 1 + bt.GetDepthByRecursive(node.rchild)
	if ldepth > rdepth {
		return ldepth
	}
	return rdepth
}

// GetDepthByLoop 通过循环获取二叉树的深度
func (bt *BinTree) GetDepthByLoop(node *Node) int {
	if bt == nil || bt.Root == nil {
		return 0
	}
	var depth int
	countFlag, finishFlag := 0, 1
	queue := stackqueue.CreateQueue(10)
	queue.Enqueue(bt.Root)
	var currentNode *Node

	for !queue.IsEmpty() {
		currentNode = queue.Dequeue().(*Node)
		countFlag++

		if currentNode.lchild != nil {
			queue.Enqueue(currentNode.lchild)
		}

		if currentNode.rchild != nil {
			queue.Enqueue(currentNode.rchild)
		}

		// countFlag == finishFlag 表示该层的节点已经取完
		// 下一轮循环将会取下一层的节点，这就是二叉树深度的记录时机
		if countFlag == finishFlag {
			countFlag = 0
			finishFlag = queue.Size()
			depth++
		}
	}
	return depth
}
