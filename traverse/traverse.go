package traverse

import (
	"github.com/ikngtty/benchmark-go-dfs/tree"
)

const maxStackLen = 1048575 // 2^20 - 1

func FindPathRec(
	node *tree.Node,
	destValue tree.NodeValue,
) []tree.NodeValue {

	reversedPath := findReversedPathRec(node, destValue)
	if reversedPath == nil {
		return nil
	}

	pathLen := len(reversedPath)
	path := make([]tree.NodeValue, pathLen)
	for i := 0; i < pathLen; i++ {
		path[i] = reversedPath[pathLen-1-i]
	}
	return path
}

func findReversedPathRec(
	node *tree.Node,
	destValue tree.NodeValue,
) []tree.NodeValue {

	if node.Value == destValue {
		return []tree.NodeValue{node.Value}
	}
	for _, child := range node.Children {
		path := findReversedPathRec(child, destValue)
		if path != nil {
			return append(path, node.Value)
		}
	}
	return nil
}

type nodeWithParent struct {
	Node   *tree.Node
	Parent *nodeWithParent
}

func FindPathLoop(
	node *tree.Node,
	destValue tree.NodeValue,
) []tree.NodeValue {

	pathRoot := nodeWithParent{node, nil}
	pathStack := make([]*nodeWithParent, maxStackLen)
	pathStack[0] = &pathRoot
	pathStackLen := 1

	for pathStackLen > 0 {
		// pop
		pathItem := pathStack[pathStackLen-1]
		pathStack[pathStackLen-1] = nil
		pathStackLen--

		node := pathItem.Node
		if node.Value == destValue {
			return traceBackParents(pathItem)
		}

		for _, child := range node.Children {
			item := nodeWithParent{child, pathItem}
			// push
			pathStack[pathStackLen] = &item
			pathStackLen++
		}
	}

	return nil
}

func traceBackParents(np *nodeWithParent) []tree.NodeValue {
	reversedPath := []tree.NodeValue{}
	cur := np
	for cur != nil {
		reversedPath = append(reversedPath, cur.Node.Value)
		cur = cur.Parent
	}

	pathLen := len(reversedPath)
	path := make([]tree.NodeValue, pathLen)
	for i := 0; i < pathLen; i++ {
		path[i] = reversedPath[pathLen-1-i]
	}
	return path
}
