package tree

type NodeValue int

type Node struct {
	Value    NodeValue
	Children []*Node
}

type nodeWithDepth struct {
	Node  *Node
	Depth int
}

func GenerateATree(depth int) *Node {
	if depth < 1 {
		panic("depth should be 1 or more")
	}

	currentValue := NodeValue(1)
	rootNode := Node{currentValue, nil}
	firstQueueItem := nodeWithDepth{&rootNode, 1}
	queue := []nodeWithDepth{firstQueueItem}
	for len(queue) > 0 {
		// deque
		queueItem := queue[0]
		queue = queue[1:]
		if queueItem.Depth >= depth {
			break
		}

		currentValue++
		left := Node{currentValue, nil}
		currentValue++
		right := Node{currentValue, nil}
		queueItem.Node.Children = []*Node{&left, &right}

		// enque
		queue = append(queue,
			nodeWithDepth{&left, queueItem.Depth + 1},
			nodeWithDepth{&right, queueItem.Depth + 1})
	}
	return &rootNode
}
