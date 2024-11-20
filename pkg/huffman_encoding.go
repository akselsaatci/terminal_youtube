package huffman

import (
	"container/heap"
)

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

func calculateFrequancy(input string) map[rune]int {

	res := map[rune]int{}
	for _, value := range input {
		if _, ok := res[value]; ok {
			res[value] += 1
		} else {
			res[value] = 1
		}
	}
	return res
}

func Encode(input string) string {
	freqTable := calculateFrequancy(input)
	pq := &MinHeap{}

	for key, val := range freqTable {
		heap.Push(pq, &MinNode{
			value:    key,
			priority: val,
		})
	}
	heap.Init(pq)

	var left, right *MinNode
	for pq.Len() > 1 {
		left = heap.Pop(pq).(*MinNode)
		right = heap.Pop(pq).(*MinNode)

		top := &MinNode{
			value:    '$', // Internal node (no character)
			priority: left.priority + right.priority,
			left:     left,
			right:    right,
		}

		heap.Push(pq, top)
	}

	root := heap.Pop(pq).(*MinNode)

	codes := make(map[rune]string)
	buildHuffmanCodes(root, "", codes)
	var encoded string
	for _, char := range input {
		encoded += codes[char]
	}

	return encoded
}

func buildHuffmanCodes(node *MinNode, currentCode string, codes map[rune]string) {
	if node == nil {
		return
	}
	// basicaly im checking if the current node is leaf if it is leaf node that means im interested in writing its code
	// if it is not it means that is just temp
	if node.left == nil && node.right == nil {
		codes[node.value] = currentCode
		return
	}

	buildHuffmanCodes(node.left, currentCode+"0", codes)
	buildHuffmanCodes(node.right, currentCode+"1", codes)
}

func Decode(input string) {}
