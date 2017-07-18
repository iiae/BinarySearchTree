/*
Author : Axoford12
本包用于实现二分搜索树的算法结构
 */
package binsertree

// 二分搜索树的分支
type Node struct {
	Left  *Node // 左孩子
	Right *Node // 右孩子
	Key   int   // 键
	Value int   // 值
}
// BST : BinarySearchTree
// 二分搜索树
type BST struct {
	Root  *Node // 根
	count int   // 节点数
}

func (b  *BST) Create(){
	b.Root = nil
	b.count = 0
}

func (b *BST) Count() int{
	return b.count
}

func (b *BST) IsEmpty() bool{
	return b.count == 0
}