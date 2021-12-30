package increment

import (
	"context"
	"sync"
)

type Node struct {
	number int64
	mutex  sync.Mutex
}

func NewNode() Node {
	return Node{
		number: -1,
	}
}

func (n *Node) Increment(ctx context.Context, void *VoidMessage) (*IncrementMessage, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	n.number++
	return &IncrementMessage{Number: n.number}, nil
}
