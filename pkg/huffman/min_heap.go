package huffman

import "container/heap"

// Copied the code from https://pkg.go.dev/container/heap#example-package-PriorityQueue and change it to minheap

// An MinNode is something we manage in a priority queue.
type MinNode struct {
	value    rune     // The value of the item; arbitrary.
	priority int      // The priority of the item in the queue.
	left     *MinNode // Left child in the tree
	right    *MinNode // Right child in the tree
	index    int      // The index of the item in the heap.
}

// A MinHeap implements heap.Interface and holds Items.
type MinHeap []*MinNode

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap) Push(x any) {
	n := len(*pq)
	item := x.(*MinNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an MinNode in the queue.
func (pq *MinHeap) update(item *MinNode, value rune, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
