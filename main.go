package main

import (
	"fmt"
	"math/rand"
)

//Node 是树状排序的节点数据结构
type Node struct {
	Father    *Node //储存父节点的引用
	Number    int   //储存数值大小
	Isdeleted bool  //选出较大数时进行是否摘除判断
}

func main() {
	fmt.Println("start")
	noorder := make([]Node, 11)
	father := make([][]Node, 0)
	father = append(father, noorder)
	for i, _ := range noorder {
		noorder[i].Number = RandInt(1, 100000)
		//fmt.Println(noorder[i].Number)
	}

	//进行排序输出
	sortNode(noorder, father)
	fmt.Println("pause")
	result := findMax(noorder)
	fmt.Println(result)

}

//RandInt 生产范围内的随机整数
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// return r.Intn(max-min) + min
	return rand.Intn(max-min) + min
}

//sortNode 用来实现二叉树排序
func sortNode(noorder []Node, father [][]Node) [][]Node {
	var f []Node
	if len(noorder)%2 == 0 {
		f = make([]Node, 0)
		for i := 0; i < len(noorder); i += 2 {
			f = append(f, max(noorder[i], noorder[i+1]))
			noorder[i].Father = &f[i/2]
			noorder[i+1].Father = &f[i/2]
		}
		father = append(father, f)
	} else {
		f = make([]Node, 0)
		for i := 0; i < len(noorder)-1; i += 2 {
			f = append(f, max(noorder[i], noorder[i+1]))
			noorder[i].Father = &f[i/2]
			noorder[i+1].Father = &f[i/2]
		}
		f = append(f, noorder[len(noorder)-1])
		noorder[len(noorder)-1].Father = &f[len(f)-1]
		father = append(father, f)
	}
	if len(noorder) == 1 {
		return sortNode(f, father)
	}

	return nil
}

//max 计算两数大小，返回新数据
func max(first, second Node) Node {
	var max Node
	if first.Number >= second.Number {
		max = first
	} else {
		max = second
	}
	return max
}

//findMax
func findMax(noorder []Node) int {
	node := noorder[0]
	for ; node.Father != nil; node = *node.Father {
	}
	return node.Number
}
