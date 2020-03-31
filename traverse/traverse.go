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

type nodePath struct {
	Parent       *nodePath
	Node         *tree.Node
	NextChildPos int
}

func FindPathLoop(
	node *tree.Node,
	destValue tree.NodeValue,
) []tree.NodeValue {

	rootPath := nodePath{nil, node, 0}
	curPath := &rootPath

	for {
		if curPath.Node.Value == destValue {
			return traceBackParents(curPath)
		}

		if curPath.NextChildPos >= len(curPath.Node.Children) {
			// pop
			curPath = curPath.Parent
			if curPath == nil {
				break
			}

			curPath.NextChildPos++
			continue
		}

		childNode := curPath.Node.Children[curPath.NextChildPos]
		// push
		childPath := nodePath{curPath, childNode, 0}
		curPath = &childPath
	}

	return nil
}

func traceBackParents(path *nodePath) []tree.NodeValue {
	reversedPath := []tree.NodeValue{}
	cur := path
	for cur != nil {
		reversedPath = append(reversedPath, cur.Node.Value)
		cur = cur.Parent
	}

	pathLen := len(reversedPath)
	values := make([]tree.NodeValue, pathLen)
	for i := 0; i < pathLen; i++ {
		values[i] = reversedPath[pathLen-1-i]
	}
	return values
}
