package traverse

import (
	"github.com/ikngtty/benchmark-go-dfs/tree"
)

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
	// HACK: I guess static size is more fast.
	pathStack := []*nodeWithParent{&pathRoot}

	for len(pathStack) > 0 {
		// pop
		pathItem := pathStack[len(pathStack)-1]
		pathStack = pathStack[:len(pathStack)-1]

		node := pathItem.Node
		if node.Value == destValue {
			return traceBackParents(pathItem)
		}

		children := node.Children
		if children == nil {
			continue
		}

		childrenLen := len(children)
		itemsToPush := make([]*nodeWithParent, childrenLen)
		for i, child := range children {
			item := nodeWithParent{child, pathItem}
			itemsToPush[childrenLen-1-i] = &item
		}

		// push
		pathStack = append(pathStack, itemsToPush...)
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
